// Copyright 2026 LiveKit, Inc.
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

package decodebudget

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// JSON limits JSON structure before a downstream decoder constructs
// collections. Callers must separately limit the serialized byte size. Both
// limits must be positive; the zero value does not disable validation.
type JSON struct {
	// MaxTokens includes scalar values, object keys, and opening/closing delimiters.
	MaxTokens int
	// MaxDepth counts nested arrays/objects, including a root container at depth one.
	MaxDepth int
}

// Validate accepts exactly one complete JSON value, including scalar roots.
func (b JSON) Validate(data []byte) error {
	if b.MaxTokens <= 0 || b.MaxDepth <= 0 {
		return errors.New("JSON budget limits must be positive")
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	depth, tokens := 0, 0
	seen, complete := false, false
	for {
		token, err := dec.Token()
		if err == io.EOF {
			if !seen || depth != 0 {
				return io.ErrUnexpectedEOF
			}
			return nil
		}
		if err != nil {
			return err
		}
		if complete {
			return errors.New("unexpected value after JSON document")
		}
		tokens++
		if tokens > b.MaxTokens {
			return ErrMessageComplexity
		}
		seen = true
		if d, ok := token.(json.Delim); ok {
			switch d {
			case '{', '[':
				depth++
			case '}', ']':
				depth--
			}
			if depth > b.MaxDepth {
				return ErrMessageComplexity
			}
		}
		if depth == 0 {
			complete = true
		}
	}
}
