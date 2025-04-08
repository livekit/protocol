// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pion/sdp/v3"
	"github.com/pion/webrtc/v4"
)

func GetMidValue(media *sdp.MediaDescription) string {
	for _, attr := range media.Attributes {
		if attr.Key == sdp.AttrKeyMID {
			return attr.Value
		}
	}
	return ""
}

func ExtractFingerprint(desc *sdp.SessionDescription) (string, string, error) {
	fingerprints := make([]string, 0)

	if fingerprint, haveFingerprint := desc.Attribute("fingerprint"); haveFingerprint {
		fingerprints = append(fingerprints, fingerprint)
	}

	for _, m := range desc.MediaDescriptions {
		if fingerprint, haveFingerprint := m.Attribute("fingerprint"); haveFingerprint {
			fingerprints = append(fingerprints, fingerprint)
		}
	}

	if len(fingerprints) < 1 {
		return "", "", webrtc.ErrSessionDescriptionNoFingerprint
	}

	for _, m := range fingerprints {
		if m != fingerprints[0] {
			return "", "", webrtc.ErrSessionDescriptionConflictingFingerprints
		}
	}

	parts := strings.Split(fingerprints[0], " ")
	if len(parts) != 2 {
		return "", "", webrtc.ErrSessionDescriptionInvalidFingerprint
	}
	return parts[1], parts[0], nil
}

func ExtractDTLSRole(desc *sdp.SessionDescription) webrtc.DTLSRole {
	for _, md := range desc.MediaDescriptions {
		setup, ok := md.Attribute(sdp.AttrKeyConnectionSetup)
		if !ok {
			continue
		}

		if setup == sdp.ConnectionRoleActive.String() {
			return webrtc.DTLSRoleClient
		}

		if setup == sdp.ConnectionRolePassive.String() {
			return webrtc.DTLSRoleServer
		}
	}

	//
	// If 'setup' attribute is not available, use client role
	// as that is the default behaviour of answerers
	//
	// There seems to be some differences in how role is decided.
	// libwebrtc (Chrome) code - (https://source.chromium.org/chromium/chromium/src/+/main:third_party/webrtc/pc/jsep_transport.cc;l=592;drc=369fb686729e7eb20d2bd09717cec14269a399d7)
	// does not mention anything about ICE role when determining
	// DTLS Role.
	//
	// But, ORTC has this - https://github.com/w3c/ortc/issues/167#issuecomment-69409953
	// and pion/webrtc follows that (https://github.com/pion/webrtc/blob/e071a4eded1efd5d9b401bcfc4efacb3a2a5a53c/dtlstransport.go#L269)
	//
	// So if remote is ice-lite, pion will use DTLSRoleServer when answering
	// while browsers pick DTLSRoleClient.
	//
	return webrtc.DTLSRoleClient
}

func ExtractICECredential(desc *sdp.SessionDescription) (string, string, error) {
	pwds := []string{}
	ufrags := []string{}

	if ufrag, haveUfrag := desc.Attribute("ice-ufrag"); haveUfrag {
		ufrags = append(ufrags, ufrag)
	}
	if pwd, havePwd := desc.Attribute("ice-pwd"); havePwd {
		pwds = append(pwds, pwd)
	}

	for _, m := range desc.MediaDescriptions {
		if ufrag, haveUfrag := m.Attribute("ice-ufrag"); haveUfrag {
			ufrags = append(ufrags, ufrag)
		}
		if pwd, havePwd := m.Attribute("ice-pwd"); havePwd {
			pwds = append(pwds, pwd)
		}
	}

	if len(ufrags) == 0 {
		return "", "", webrtc.ErrSessionDescriptionMissingIceUfrag
	} else if len(pwds) == 0 {
		return "", "", webrtc.ErrSessionDescriptionMissingIcePwd
	}

	for _, m := range ufrags {
		if m != ufrags[0] {
			return "", "", webrtc.ErrSessionDescriptionConflictingIceUfrag
		}
	}

	for _, m := range pwds {
		if m != pwds[0] {
			return "", "", webrtc.ErrSessionDescriptionConflictingIcePwd
		}
	}

	return ufrags[0], pwds[0], nil
}

func ExtractStreamID(media *sdp.MediaDescription) (string, bool) {
	var streamID string
	msid, ok := media.Attribute(sdp.AttrKeyMsid)
	if !ok {
		return "", false
	}
	ids := strings.Split(msid, " ")
	if len(ids) < 2 {
		streamID = msid
	} else {
		streamID = ids[1]
	}
	return streamID, true
}

func GetTrackIDFromMediaDescription(m *sdp.MediaDescription) string {
	trackId := ""
	msid, ok := m.Attribute(sdp.AttrKeyMsid)
	if ok {
		if split := strings.Split(msid, " "); len(split) >= 2 {
			trackId = split[1]
		}
	}
	return trackId
}

func IsMediaDescriptionSimulcast(m *sdp.MediaDescription) bool {
	_, ok := m.Attribute("simulcast")
	return ok
}

func CodecsFromMediaDescription(m *sdp.MediaDescription) (out []sdp.Codec, err error) {
	s := &sdp.SessionDescription{
		MediaDescriptions: []*sdp.MediaDescription{m},
	}

	for _, payloadStr := range m.MediaName.Formats {
		payloadType, err := strconv.ParseUint(payloadStr, 10, 8)
		if err != nil {
			return nil, err
		}

		codec, err := s.GetCodecForPayloadType(uint8(payloadType))
		if err != nil {
			if payloadType == 0 {
				continue
			}
			return nil, err
		}

		out = append(out, codec)
	}

	return out, nil
}

func GetBundleMid(parsed *sdp.SessionDescription) (string, bool) {
	if groupAttribute, found := parsed.Attribute(sdp.AttrKeyGroup); found {
		bundleIDs := strings.Split(groupAttribute, " ")
		if len(bundleIDs) > 1 && strings.EqualFold(bundleIDs[0], "BUNDLE") {
			return bundleIDs[1], true
		}
	}

	return "", false
}

type sdpFragmentICE struct {
	ufrag   string
	pwd     string
	lite    *bool
	options string
}

func (i *sdpFragmentICE) Unmarshal(attributes []sdp.Attribute) error {
	getAttr := func(key string) (string, bool) {
		for _, a := range attributes {
			if a.Key == key {
				return a.Value, true
			}
		}

		return "", false
	}

	iceUfrag, found := getAttr("ice-ufrag")
	if found {
		i.ufrag = iceUfrag
	}

	icePwd, found := getAttr("ice-pwd")
	if found {
		i.pwd = icePwd
	}

	_, found = getAttr(sdp.AttrKeyICELite)
	if found {
		lite := true
		i.lite = &lite
	}

	iceOptions, found := getAttr("ice-options")
	if found {
		i.options = iceOptions
	}

	return nil
}

func (i *sdpFragmentICE) Marshal() (string, error) {
	iceFragment := []byte{}
	addKeyValue := func(key string, value string) {
		iceFragment = append(iceFragment, key...)
		if value != "" {
			iceFragment = append(iceFragment, value...)
		}
		iceFragment = append(iceFragment, "\r\n"...)
	}

	if i.ufrag != "" {
		addKeyValue("a=ice-ufrag:", i.ufrag)
	}
	if i.pwd != "" {
		addKeyValue("a=ice-pwd:", i.pwd)
	}
	if i.lite != nil && *i.lite {
		addKeyValue("a=ice-lite", "")
	}
	if i.options != "" {
		addKeyValue("a=ice-options:", i.options)
	}

	return string(iceFragment), nil
}

type sdpFragmentMedia struct {
	info            string
	mid             string
	ice             *sdpFragmentICE
	candidates      []string
	endOfCandidates *bool
}

func (m *sdpFragmentMedia) Unmarshal(md *sdp.MediaDescription) error {
	// MediaName conversion to string taken from github.com/pion/sdp
	var info []byte
	appendList := func(list []string, sep byte) {
		for i, p := range list {
			if i != 0 && i != len(list) {
				info = append(info, sep)
			}
			info = append(info, p...)
		}
	}

	info = append(append(info, md.MediaName.Media...), ' ')

	info = strconv.AppendInt(info, int64(md.MediaName.Port.Value), 10)
	if md.MediaName.Port.Range != nil {
		info = append(info, '/')
		info = strconv.AppendInt(info, int64(*md.MediaName.Port.Range), 10)
	}

	appendList(md.MediaName.Protos, '/')
	info = append(info, ' ')
	appendList(md.MediaName.Formats, ' ')
	m.info = string(info)

	mid, found := md.Attribute(sdp.AttrKeyMID)
	if found {
		m.mid = mid
	}

	m.ice = &sdpFragmentICE{}
	if err := m.ice.Unmarshal(md.Attributes); err != nil {
		return err
	}

	for _, a := range md.Attributes {
		if a.IsICECandidate() {
			m.candidates = append(m.candidates, a.Value)
		}
	}

	_, found = md.Attribute(sdp.AttrKeyEndOfCandidates)
	if found {
		endOfCandidates := true
		m.endOfCandidates = &endOfCandidates
	}
	return nil
}

func (m *sdpFragmentMedia) Marshal() (string, error) {
	mediaFragment := []byte{}
	addKeyValue := func(key string, value string) {
		mediaFragment = append(mediaFragment, key...)
		if value != "" {
			mediaFragment = append(mediaFragment, value...)
		}
		mediaFragment = append(mediaFragment, "\r\n"...)
	}

	if m.info != "" {
		addKeyValue("m=", m.info)
	}

	if m.mid != "" {
		addKeyValue("a=mid:", m.mid)
	}

	if m.ice != nil {
		iceFragment, err := m.ice.Marshal()
		if err != nil {
			return "", err
		}
		mediaFragment = append(mediaFragment, iceFragment...)
	}

	for _, c := range m.candidates {
		addKeyValue("a=candidate:", c)
	}
	if m.endOfCandidates != nil && *m.endOfCandidates {
		addKeyValue("a=end-of-candidates", "")
	}

	return string(mediaFragment), nil
}

type SDPFragment struct {
	group string
	ice   *sdpFragmentICE
	media *sdpFragmentMedia
}

// primarily for use with WHIP Trickle ICE - https://www.rfc-editor.org/rfc/rfc9725.html#name-trickle-ice
func (s *SDPFragment) Unmarshal(frag string) error {
	lines := strings.Split(frag, "\n")
	for _, line := range lines {
		line = strings.TrimRight(line, " \r")
		if len(line) == 0 {
			continue
		}

		if line[0] == 'm' {
			if s.media != nil {
				return errors.New("too many media sections")
			}

			s.media = &sdpFragmentMedia{}
			s.media.ice = &sdpFragmentICE{}
			s.media.info = line[2:]
		}

		if line[0] != 'a' {
			// not an attribute, skip
			continue
		}

		if line[1] != '=' {
			return errors.New("invalid attribute")
		}

		attrParts := strings.Split(line[2:], ":")
		if len(attrParts) != 2 {
			return errors.New("invalid attribute")
		}

		if s.ice == nil {
			s.ice = &sdpFragmentICE{}
		}

		switch attrParts[0] {
		case sdp.AttrKeyGroup:
			s.group = attrParts[1]

		case "ice-ufrag":
			if s.media != nil {
				s.media.ice.ufrag = attrParts[1]
			} else {
				s.ice.ufrag = attrParts[1]
			}

		case "ice-pwd":
			if s.media != nil {
				s.media.ice.pwd = attrParts[1]
			} else {
				s.ice.pwd = attrParts[1]
			}

		case sdp.AttrKeyICELite:
			lite := true
			if s.media != nil {
				s.media.ice.lite = &lite
			} else {
				s.ice.lite = &lite
			}

		case "ice-options":
			if s.media != nil {
				s.media.ice.options = attrParts[1]
			} else {
				s.ice.options = attrParts[1]
			}

		case sdp.AttrKeyMID:
			if s.media != nil {
				s.media.mid = attrParts[1]
			}

		case sdp.AttrKeyCandidate:
			if s.media != nil {
				s.media.candidates = append(s.media.candidates, attrParts[1])
			}

		case sdp.AttrKeyEndOfCandidates:
			endOfCandidates := true
			if s.media != nil {
				s.media.endOfCandidates = &endOfCandidates
			}
		}
	}

	if s.media == nil {
		return errors.New("missing media section")
	}

	if s.group != "" {
		bundleIDs := strings.Split(s.group, " ")
		if len(bundleIDs) > 1 && strings.EqualFold(bundleIDs[0], "BUNDLE") {
			if s.media.mid != bundleIDs[1] {
				return fmt.Errorf("bundle media mismatch, expected: %s, got: %s", bundleIDs[1], s.media.mid)
			}
		}
	}

	return nil
}

// primarily for use with WHIP ICE Restart - https://www.rfc-editor.org/rfc/rfc9725.html#name-ice-restarts
func (s *SDPFragment) Marshal() (string, error) {
	sdpFragment := []byte{}
	addKeyValue := func(key string, value string) {
		sdpFragment = append(sdpFragment, key...)
		if value != "" {
			sdpFragment = append(sdpFragment, value...)
		}
		sdpFragment = append(sdpFragment, "\r\n"...)
	}

	if s.group != "" {
		addKeyValue("a=group:", s.group)
	}

	if s.ice != nil {
		iceFragment, err := s.ice.Marshal()
		if err != nil {
			return "", err
		}
		sdpFragment = append(sdpFragment, iceFragment...)
	}

	if s.media != nil {
		mediaFragment, err := s.media.Marshal()
		if err != nil {
			return "", err
		}
		sdpFragment = append(sdpFragment, mediaFragment...)
	}

	return string(sdpFragment), nil
}

func (s *SDPFragment) Mid() string {
	if s.media != nil {
		return s.media.mid
	}

	return ""
}

func (s *SDPFragment) Candidates() []string {
	if s.media != nil {
		return s.media.candidates
	}

	return nil
}

func (s *SDPFragment) ExtractICECredential() (string, string, error) {
	pwds := []string{}
	ufrags := []string{}

	if s.ice != nil {
		if s.ice.ufrag != "" {
			ufrags = append(ufrags, s.ice.ufrag)
		}
		if s.ice.pwd != "" {
			pwds = append(pwds, s.ice.pwd)
		}
	}

	if s.media != nil {
		if s.media.ice.ufrag != "" {
			ufrags = append(ufrags, s.media.ice.ufrag)
		}
		if s.media.ice.pwd != "" {
			pwds = append(pwds, s.media.ice.pwd)
		}
	}

	if len(ufrags) == 0 {
		return "", "", webrtc.ErrSessionDescriptionMissingIceUfrag
	} else if len(pwds) == 0 {
		return "", "", webrtc.ErrSessionDescriptionMissingIcePwd
	}

	for _, m := range ufrags {
		if m != ufrags[0] {
			return "", "", webrtc.ErrSessionDescriptionConflictingIceUfrag
		}
	}

	for _, m := range pwds {
		if m != pwds[0] {
			return "", "", webrtc.ErrSessionDescriptionConflictingIcePwd
		}
	}

	return ufrags[0], pwds[0], nil
}

// primarily for use with WHIP ICE Restart - https://www.rfc-editor.org/rfc/rfc9725.html#name-ice-restarts
func (s *SDPFragment) PatchICECredentialIntoSDP(parsed *sdp.SessionDescription) error {
	// ice-options and ice-lite should match
	if s.ice != nil && (s.ice.lite != nil || s.ice.options != "") {
		for _, a := range parsed.Attributes {
			switch a.Key {
			case "ice-lite":
				if s.ice.lite == nil || !*s.ice.lite {
					return errors.New("ice lite mismatch")
				}
			case "ice-pwd":
				if s.ice.options != "" && s.ice.options != a.Value {
					return errors.New("ice options mismatch")
				}
			}
		}
	}
	if s.media != nil && s.media.mid != "" && s.media.ice != nil && (s.media.ice.lite != nil || s.media.ice.options != "") {
		for _, md := range parsed.MediaDescriptions {
			mid, found := md.Attribute(sdp.AttrKeyMID)
			if !found || mid != s.media.mid {
				continue
			}

			for _, a := range md.Attributes {
				switch a.Key {
				case "ice-lite":
					if s.media.ice.lite == nil || !*s.media.ice.lite {
						return errors.New("ice lite mismatch")
					}
				case "ice-options":
					if s.media.ice.options != "" && s.media.ice.options != a.Value {
						return errors.New("ice options mismatch")
					}
				}
			}
		}
	}

	if s.ice != nil && s.ice.ufrag != "" && s.ice.pwd != "" {
		for _, a := range parsed.Attributes {
			switch a.Key {
			case "ice-ufrag":
				a.Value = s.ice.ufrag
			case "ice-pwd":
				a.Value = s.ice.pwd
			}
		}
	}

	if s.media != nil && s.media.mid != "" {
		for _, md := range parsed.MediaDescriptions {
			mid, found := md.Attribute(sdp.AttrKeyMID)
			if !found || mid != s.media.mid {
				continue
			}

			for _, a := range md.Attributes {
				switch a.Key {
				case "ice-ufrag":
					if s.media.ice.ufrag != "" {
						a.Value = s.media.ice.ufrag
					}
				case "ice-pwd":
					if s.media.ice.pwd != "" {
						a.Value = s.media.ice.pwd
					}
				}
			}

			// clean out existing candidates and patch in new ones
			for idx, a := range md.Attributes {
				if a.IsICECandidate() || a.Key == sdp.AttrKeyEndOfCandidates {
					md.Attributes = append(md.Attributes[:idx], md.Attributes[idx+1:]...)
				}
			}

			for _, ic := range s.media.candidates {
				md.Attributes = append(
					md.Attributes,
					sdp.Attribute{
						Key:   sdp.AttrKeyCandidate,
						Value: ic,
					},
				)
			}

			if s.media.endOfCandidates != nil && *s.media.endOfCandidates {
				md.Attributes = append(
					md.Attributes,
					sdp.Attribute{Key: sdp.AttrKeyEndOfCandidates},
				)
			}
		}
	}
	return nil
}

// primarily for use with WHIP ICE Restart - https://www.rfc-editor.org/rfc/rfc9725.html#name-ice-restarts
func ExtractSDPFragment(parsed *sdp.SessionDescription) (*SDPFragment, error) {
	bundleMid, found := GetBundleMid(parsed)
	if !found {
		return nil, errors.New("could not get bundle mid")
	}

	s := &SDPFragment{}
	s.ice = &sdpFragmentICE{}
	if err := s.ice.Unmarshal(parsed.Attributes); err != nil {
		return nil, err
	}

	foundBundleMedia := false
	for _, md := range parsed.MediaDescriptions {
		mid, found := md.Attribute(sdp.AttrKeyMID)
		if !found || mid != bundleMid {
			continue
		}

		foundBundleMedia = true

		s.media = &sdpFragmentMedia{}
		if err := s.media.Unmarshal(md); err != nil {
			return nil, err
		}
		break
	}

	if !foundBundleMedia {
		return nil, fmt.Errorf("could not find bundle media: %s", bundleMid)
	}

	return s, nil
}
