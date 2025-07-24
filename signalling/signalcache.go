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

package signalling

import (
	"math/rand"
	"sync"

	"github.com/gammazero/deque"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
)

type SignalCacheParams struct {
	Logger         logger.Logger
	FirstMessageId uint32 // should be used for testing only
}

type element struct {
	messageId uint32
	value     proto.Message
}

type SignalCache struct {
	params SignalCacheParams

	lock                         sync.Mutex
	nextMessageId                uint32
	lastProcessedRemoteMessageId uint32
	messages                     deque.Deque[element]
}

func NewSignalCache(params SignalCacheParams) *SignalCache {
	s := &SignalCache{
		params:        params,
		nextMessageId: params.FirstMessageId,
	}
	if s.nextMessageId == 0 {
		s.nextMessageId = uint32(rand.Intn(1<<8) + 1)
	}
	s.messages.SetBaseCap(16)
	return s
}

func (s *SignalCache) NextMessageId() uint32 {
	s.lock.Lock()
	defer s.lock.Unlock()

	nextMessageId := s.nextMessageId
	s.nextMessageId++
	return nextMessageId
}

func (s *SignalCache) SetLastProcessedRemoteMessageId(lastProcessedRemoteMessageId uint32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.lastProcessedRemoteMessageId = lastProcessedRemoteMessageId
}

func (s *SignalCache) Add(msg proto.Message, messageId uint32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	elem := element{messageId, msg}

	rindex := -1
	if s.messages.Len() != 0 {
		rindex = s.messages.RIndex(func(e element) bool {
			return e.messageId < messageId
		})
	}
	s.messages.Insert(rindex+1, elem)
}

func (s *SignalCache) Clear(till uint32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.clearLocked(till)
}

func (s *SignalCache) clearLocked(till uint32) {
	for s.messages.Len() != 0 {
		front := s.messages.Front()
		if front.messageId > till {
			break
		}
		s.messages.PopFront()
	}
}

func (s *SignalCache) GetFromFront() ([]proto.Message, uint32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.getFromFrontLocked()
}

func (s *SignalCache) getFromFrontLocked() ([]proto.Message, uint32) {
	var msgs []proto.Message
	for elem := range s.messages.Iter() {
		msgs = append(msgs, utils.CloneProto(elem.value))
	}

	return msgs, s.lastProcessedRemoteMessageId
}

func (s *SignalCache) ClearAndGetFrom(from uint32) ([]proto.Message, uint32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.clearLocked(from - 1)
	return s.getFromFrontLocked()
}
