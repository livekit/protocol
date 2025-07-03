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

package livekit

import (
	"context"
	"fmt"
	"io"

	"buf.build/go/protoyaml"
	"github.com/dennwc/iters"
	proto "google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

type TrackID string
type ParticipantID string
type ParticipantIdentity string
type ParticipantName string
type RoomID string
type RoomName string
type ConnectionID string
type NodeID string
type RoomKey struct {
	ProjectID string
	RoomName  RoomName
}
type ParticipantKey struct {
	RoomKey
	Identity ParticipantIdentity
}
type JobID string
type DispatchID string
type AgentName string

func (s TrackID) String() string             { return string(s) }
func (s ParticipantID) String() string       { return string(s) }
func (s ParticipantIdentity) String() string { return string(s) }
func (s ParticipantName) String() string     { return string(s) }
func (s RoomID) String() string              { return string(s) }
func (s RoomName) String() string            { return string(s) }
func (s ConnectionID) String() string        { return string(s) }
func (s NodeID) String() string              { return string(s) }
func (s JobID) String() string               { return string(s) }
func (s DispatchID) String() string          { return string(s) }
func (s AgentName) String() string           { return string(s) }

type stringTypes interface {
	ParticipantID | RoomID | TrackID | ParticipantIdentity | ParticipantName | RoomName | ConnectionID | NodeID
}

func IDsAsStrings[T stringTypes](ids []T) []string {
	strs := make([]string, 0, len(ids))
	for _, id := range ids {
		strs = append(strs, string(id))
	}
	return strs
}

func StringsAsIDs[T stringTypes](ids []string) []T {
	asID := make([]T, 0, len(ids))
	for _, id := range ids {
		asID = append(asID, T(id))
	}

	return asID
}

type Guid interface {
	TrackID | ParticipantID | RoomID
}

type GuidBlock [9]byte

func (r *RoomConfiguration) UnmarshalYAML(value *yaml.Node) error {
	// Marshall the Node back to yaml to pass it to the protobuf specific unmarshaller
	str, err := yaml.Marshal(value)
	if err != nil {
		return err
	}

	return protoyaml.Unmarshal(str, r)
}

func (r *RoomConfiguration) MarshalYAML() (interface{}, error) {
	return marshalProto(r)
}

func (r *RoomEgress) UnmarshalYAML(value *yaml.Node) error {
	// Marshall the Node back to yaml to pass it to the protobuf specific unmarshaller
	str, err := yaml.Marshal(value)
	if err != nil {
		return err
	}

	return protoyaml.Unmarshal(str, r)
}

func (r *RoomEgress) MarshalYAML() (interface{}, error) {
	return marshalProto(r)
}

func (r *RoomAgent) UnmarshalYAML(value *yaml.Node) error {
	// Marshall the Node back to yaml to pass it to the protobuf specific unmarshaller
	str, err := yaml.Marshal(value)
	if err != nil {
		return err
	}

	return protoyaml.Unmarshal(str, r)
}

func (r *RoomAgent) MarshalYAML() (interface{}, error) {
	return marshalProto(r)
}

func marshalProto(o proto.Message) (map[string]interface{}, error) {
	// Marshall the Node to yaml using the protobuf specific marshaller to ensure the proper field names are used
	str, err := protoyaml.MarshalOptions{UseProtoNames: true}.Marshal(o)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})

	err = yaml.Unmarshal(str, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func cloneProto[T proto.Message](m T) T {
	return proto.Clone(m).(T)
}

func IsJobType(jobType JobType) bool {
	_, ok := JobType_name[int32(jobType)]
	return ok
}

func (p *Pagination) Filter(v PageItem) bool {
	if p == nil {
		return true
	}
	if p.AfterId != "" {
		if id := v.ID(); id != "" && id <= p.AfterId {
			return false
		}
	}
	return true
}

type pageIterReq[T any] interface {
	GetPage() *Pagination
	Filter(v T) bool
}

type pageIterResp[T any] interface {
	GetItems() []T
}

type PageItem interface {
	ID() string
}

// FilterStats tracks statistics about the filtering process
type FilterStats struct {
	TotalItemsBeforeFilter int
	TotalItemsAfterFilter  int
	ItemsInCurrentPage     int
	FilteredInCurrentPage  int
}

type StatsCallback func(stats FilterStats)

// ListPageIterWithStats creates a paginated iterator that tracks and emits filtering statistics
func ListPageIterWithStats[T PageItem, Req pageIterReq[T], Resp pageIterResp[T]](
	fnc func(ctx context.Context, req Req) (Resp, error),
	req Req,
	statsCallback StatsCallback,
	applyFilter bool,
) iters.PageIter[T] {
	it := &listPageIterWithStats[T, Req, Resp]{
		fnc:           fnc,
		req:           req,
		statsCallback: statsCallback,
		stats:         FilterStats{},
		applyFilter:   applyFilter,
	}
	if applyFilter {
		return iters.FilterPage(it, func(v T) bool {
			return req.Filter(v)
		})
	}
	return it
}

func ListPageIter[T PageItem, Req pageIterReq[T], Resp pageIterResp[T]](fnc func(ctx context.Context, req Req) (Resp, error), req Req) iters.PageIter[T] {
	// Create a callback that applies the same filtering logic as iters.FilterPage
	statsCallback := func(stats FilterStats) {
		// Optional: You can add logging here if needed
		// fmt.Printf("Filter Stats: Total Before: %d, Total After: %d, Current Page: %d, Filtered This Page: %d\n",
		// 	stats.TotalItemsBeforeFilter,
		// 	stats.TotalItemsAfterFilter,
		// 	stats.ItemsInCurrentPage,
		// 	stats.FilteredInCurrentPage)
	}

	// Use ListPageIterWithStats with applyFilter=false to get unfiltered items, then apply filtering
	it := ListPageIterWithStats(fnc, req, statsCallback, false)
	return iters.FilterPage(it, func(v T) bool {
		return req.Filter(v)
	})
}

// ListPageIterSilent creates a paginated iterator without statistics logging
func ListPageIterSilent[T PageItem, Req pageIterReq[T], Resp pageIterResp[T]](fnc func(ctx context.Context, req Req) (Resp, error), req Req) iters.PageIter[T] {
	return ListPageIterWithStats(fnc, req, nil, true)
}

type listPageIter[T PageItem, Req pageIterReq[T], Resp pageIterResp[T]] struct {
	fnc  func(ctx context.Context, opts Req) (Resp, error)
	req  Req
	done bool
}

type listPageIterWithStats[T PageItem, Req pageIterReq[T], Resp pageIterResp[T]] struct {
	fnc           func(ctx context.Context, opts Req) (Resp, error)
	req           Req
	done          bool
	stats         FilterStats
	statsCallback StatsCallback
	applyFilter   bool
}

func (it *listPageIter[T, Req, Resp]) NextPage(ctx context.Context) ([]T, error) {
	if it.done {
		return nil, io.EOF
	}
	opts := it.req.GetPage()
	resp, err := it.fnc(ctx, it.req)
	page := resp.GetItems()
	if opts == nil {
		// No pagination set - returns all items.
		// We have to do this to support legacy implementations.
		it.done = true
		return page, err
	}
	// Advance pagination cursor.
	hasID := false
	for i := len(page) - 1; i >= 0; i-- {
		if id := page[i].ID(); id > opts.AfterId {
			opts.AfterId = id
			hasID = true
		}
	}
	if err == nil && (len(page) == 0 || !hasID) {
		err = io.EOF
		it.done = true
	}
	return page, err
}

func (it *listPageIter[_, _, _]) Close() {
	it.done = true
}

func (it *listPageIterWithStats[T, Req, Resp]) NextPage(ctx context.Context) ([]T, error) {
	if it.done {
		return nil, io.EOF
	}

	opts := it.req.GetPage()
	resp, err := it.fnc(ctx, it.req)
	page := resp.GetItems()

	// Track total items before filtering
	it.stats.TotalItemsBeforeFilter += len(page)
	it.stats.ItemsInCurrentPage = len(page)

	if it.applyFilter {
		// Apply filtering and track statistics
		filteredItems := make([]T, 0, len(page))
		for _, item := range page {
			if it.req.Filter(item) {
				filteredItems = append(filteredItems, item)
			}
		}

		// Update statistics
		it.stats.TotalItemsAfterFilter += len(filteredItems)
		it.stats.FilteredInCurrentPage = len(page) - len(filteredItems)

		// Emit statistics via callback
		if it.statsCallback != nil {
			it.statsCallback(it.stats)
		}

		if opts == nil {
			it.done = true
			return filteredItems, err
		}

		// Advance pagination cursor based on filtered items
		hasID := false
		for i := len(filteredItems) - 1; i >= 0; i-- {
			if id := filteredItems[i].ID(); id > opts.AfterId {
				opts.AfterId = id
				hasID = true
			}
		}

		if err == nil && (len(filteredItems) == 0 || !hasID) {
			err = io.EOF
			it.done = true
		}

		return filteredItems, err
	} else {
		// Don't apply filtering, just track statistics
		it.stats.TotalItemsAfterFilter += len(page) // All items pass through
		it.stats.FilteredInCurrentPage = 0

		// Emit statistics via callback
		if it.statsCallback != nil {
			it.statsCallback(it.stats)
		}

		if opts == nil {
			it.done = true
			return page, err
		}

		// Advance pagination cursor based on all items
		hasID := false
		for i := len(page) - 1; i >= 0; i-- {
			if id := page[i].ID(); id > opts.AfterId {
				opts.AfterId = id
				hasID = true
			}
		}

		if err == nil && (len(page) == 0 || !hasID) {
			err = io.EOF
			it.done = true
		}

		return page, err
	}
}

func (it *listPageIterWithStats[_, _, _]) Close() {
	it.done = true
}

func (p *ListUpdate) Validate() error {
	if p == nil {
		return nil
	}
	for _, v := range p.Set {
		if v == "" {
			return fmt.Errorf("empty element in the list")
		}
	}
	return nil
}

func applyUpdate[T any](dst *T, set *T) {
	if set != nil {
		*dst = *set
	}
}

func applyUpdatePtr[T any](dst **T, set *T) {
	if set != nil {
		*dst = set
	}
}

func applyListUpdate[T ~string](dst *[]T, u *ListUpdate) {
	if u == nil {
		return
	}
	arr := make([]T, 0, len(u.Set))
	for _, v := range u.Set {
		arr = append(arr, T(v))
	}
	*dst = arr
}

func applyMapDiff(dst *map[string]string, diff map[string]string) {
	m := *dst
	for k, v := range diff {
		if v != "" {
			if m == nil {
				m = make(map[string]string)
			}
			m[k] = v
		} else {
			delete(m, k)
		}
	}
	*dst = m
}
