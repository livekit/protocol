package rpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils/prototest"
)

func TestNewCreateSIPParticipantRequest(t *testing.T) {
	r := &livekit.CreateSIPParticipantRequest{
		SipTrunkId:          "trunk",
		SipCallTo:           "+3333",
		RoomName:            "room",
		ParticipantIdentity: "",
		ParticipantName:     "",
		ParticipantMetadata: "meta",
		ParticipantAttributes: map[string]string{
			"extra": "1",
		},
		Headers: map[string]string{
			"X-B": "B2",
			"X-C": "C",
		},
		Dtmf:              "1234#",
		PlayDialtone:      true,
		WaitUntilAnswered: true,
		MediaEncryption:   livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE,
	}
	tr := &livekit.SIPOutboundTrunkInfo{
		SipTrunkId:         "trunk",
		Address:            "sip.example.com",
		Numbers:            []string{"+1111"},
		DestinationCountry: "us",
		AuthUsername:       "user",
		AuthPassword:       "pass",
		Headers: map[string]string{
			"X-A": "A",
			"X-B": "B1",
		},
	}
	expAttrs1 := map[string]string{
		"extra":                    "1",
		livekit.AttrSIPCallID:      "call-id",
		livekit.AttrSIPTrunkID:     "trunk",
		livekit.AttrSIPTrunkNumber: "+1111",
		livekit.AttrSIPPhoneNumber: "+3333",
		livekit.AttrSIPHostName:    "sip.example.com",
	}
	exp := &InternalCreateSIPParticipantRequest{
		ProjectId:             "p_123",
		SipCallId:             "call-id",
		SipTrunkId:            "trunk",
		Address:               "sip.example.com",
		Hostname:              "xyz.sip.livekit.cloud",
		DestinationCountry:    "us",
		Number:                "+1111",
		CallTo:                "+3333",
		Username:              "user",
		Password:              "pass",
		RoomName:              "room",
		ParticipantIdentity:   "sip_+3333",
		ParticipantMetadata:   "meta",
		Token:                 "token",
		WsUrl:                 "url",
		Dtmf:                  "1234#",
		PlayDialtone:          true,
		ParticipantAttributes: expAttrs1,
		Headers: map[string]string{
			"X-A": "A",
			"X-B": "B2",
			"X-C": "C",
		},
		WaitUntilAnswered: true,
		MediaEncryption:   livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE,
		Media: &livekit.SIPMediaConfig{
			Encryption: new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		},
	}
	res, errs := NewCreateSIPParticipantRequest("p_123", "call-id", "xyz.sip.livekit.cloud", "url", "token", r, tr)
	require.Empty(t, errs)
	require.True(t, proto.Equal(exp, res), "%v\nvs\n%v", exp, res)

	r.HidePhoneNumber = true
	r.MediaEncryption = 0
	r.Media = &livekit.SIPMediaConfig{
		Encryption: new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
	}
	res, errs = NewCreateSIPParticipantRequest("p_123", "call-id", "xyz.sip.livekit.cloud", "url", "token", r, tr)
	require.Empty(t, errs)
	exp = &InternalCreateSIPParticipantRequest{
		ProjectId:           "p_123",
		SipCallId:           "call-id",
		SipTrunkId:          "trunk",
		Address:             "sip.example.com",
		Hostname:            "xyz.sip.livekit.cloud",
		DestinationCountry:  "us",
		Number:              "+1111",
		CallTo:              "+3333",
		Username:            "user",
		Password:            "pass",
		RoomName:            "room",
		Token:               "token",
		WsUrl:               "url",
		Dtmf:                "1234#",
		PlayDialtone:        true,
		ParticipantIdentity: "sip_+3333",
		ParticipantAttributes: map[string]string{
			"extra":                "1",
			livekit.AttrSIPCallID:  "call-id",
			livekit.AttrSIPTrunkID: "trunk",
		},
		ParticipantMetadata: "meta",
		Headers: map[string]string{
			"X-A": "A",
			"X-B": "B2",
			"X-C": "C",
		},
		WaitUntilAnswered: true,
		MediaEncryption:   livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW,
		Media: &livekit.SIPMediaConfig{
			Encryption: new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		},
	}
	require.True(t, proto.Equal(exp, res), "%v\nvs\n%v", exp, res)

	r.HidePhoneNumber = false
	r.SipNumber = tr.Numbers[0]
	r.Trunk = &livekit.SIPOutboundConfig{
		Hostname:            tr.Address,
		Transport:           tr.Transport,
		DestinationCountry:  "us",
		AuthUsername:        tr.AuthUsername,
		AuthPassword:        tr.AuthPassword,
		HeadersToAttributes: tr.HeadersToAttributes,
		AttributesToHeaders: tr.AttributesToHeaders,
	}
	r.SipTrunkId = ""
	exp.SipTrunkId = ""
	for k, v := range tr.Headers {
		if _, ok := r.Headers[k]; !ok {
			r.Headers[k] = v
		}
	}
	exp.ParticipantAttributes = expAttrs1
	exp.ParticipantAttributes[livekit.AttrSIPTrunkID] = ""
	res, errs = NewCreateSIPParticipantRequest("p_123", "call-id", "xyz.sip.livekit.cloud", "url", "token", r, nil)
	require.Empty(t, errs)
	require.True(t, proto.Equal(exp, res), "%v\nvs\n%v", exp, res)
}

func TestNewCreateSIPParticipantRequestMediaConfig(t *testing.T) {
	newReq := func() *livekit.CreateSIPParticipantRequest {
		return &livekit.CreateSIPParticipantRequest{
			SipTrunkId: "trunk",
			SipCallTo:  "+3333",
			RoomName:   "room",
		}
	}
	newTrunk := func() *livekit.SIPOutboundTrunkInfo {
		return &livekit.SIPOutboundTrunkInfo{
			SipTrunkId: "trunk",
			Address:    "sip.example.com",
			Numbers:    []string{"+1111"},
		}
	}
	check := func(
		t testing.TB,
		r *livekit.CreateSIPParticipantRequest, tr *livekit.SIPOutboundTrunkInfo,
		media *livekit.SIPMediaConfig,
	) {
		res, errs := NewCreateSIPParticipantRequest("p_123", "call-id", "xyz.sip.livekit.cloud", "url", "token", r, tr)
		require.Empty(t, errs)
		require.Equal(t, *media.Encryption, res.MediaEncryption)
		require.NotNil(t, res.Media)
		prototest.Equals(t, media, res.Media)
	}
	t.Run("trunk only legacy", func(t *testing.T) {
		// Regression: trunk-level MediaEncryption must be honored when the request specifies
		// neither MediaEncryption nor Media. A prior version called req.Upgrade() at the top of
		// NewCreateSIPParticipantRequest, which pinned req.Media.Encryption to req.MediaEncryption (0)
		// before the trunk was consulted, causing outbound INVITEs to omit SRTP and upstream
		// providers (e.g. Twilio) to reject with 488 / 32208.
		r := newReq()

		tr := newTrunk()
		tr.MediaEncryption = livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				Encryption: new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
			},
		)
	})
	t.Run("req only legacy", func(t *testing.T) {
		r := newReq()
		r.MediaEncryption = livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE

		tr := newTrunk()

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				Encryption: new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
			},
		)
	})
	t.Run("trunk only media", func(t *testing.T) {
		r := newReq()

		tr := newTrunk()
		tr.Media = &livekit.SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		}

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				MediaTimeout: durationpb.New(10 * time.Second),
				Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
			},
		)
	})
	t.Run("req only media", func(t *testing.T) {
		r := newReq()
		r.Media = &livekit.SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		}

		tr := newTrunk()

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				MediaTimeout: durationpb.New(10 * time.Second),
				Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
			},
		)
	})
	t.Run("both legacy", func(t *testing.T) {
		r := newReq()
		r.MediaEncryption = livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW

		tr := newTrunk()
		tr.MediaEncryption = livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				Encryption: new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
			},
		)
	})
	t.Run("both media", func(t *testing.T) {
		r := newReq()
		r.Media = &livekit.SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		}

		tr := newTrunk()
		tr.Media = &livekit.SIPMediaConfig{
			MediaTimeout: durationpb.New(15 * time.Second),
			Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		}

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				MediaTimeout: durationpb.New(10 * time.Second),
				Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
			},
		)
	})
	t.Run("trunk legacy and req media", func(t *testing.T) {
		r := newReq()
		r.Media = &livekit.SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
		}

		tr := newTrunk()
		tr.MediaEncryption = livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				MediaTimeout: durationpb.New(10 * time.Second),
				Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
			},
		)
	})
	t.Run("trunk media and req legacy", func(t *testing.T) {
		r := newReq()
		r.MediaEncryption = livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW

		tr := newTrunk()
		tr.Media = &livekit.SIPMediaConfig{
			MediaTimeout: durationpb.New(10 * time.Second),
			Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_REQUIRE),
		}

		check(
			t, r, tr,
			&livekit.SIPMediaConfig{
				MediaTimeout: durationpb.New(10 * time.Second),
				Encryption:   new(livekit.SIPMediaEncryption_SIP_MEDIA_ENCRYPT_ALLOW),
			},
		)
	})
}
