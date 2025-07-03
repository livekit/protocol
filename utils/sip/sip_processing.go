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

package sip

import (
	"regexp"
	"strings"
)

var (
	reNumber     = regexp.MustCompile(`^\+?[\d\- ()]+$`)
	reNumberRepl = strings.NewReplacer(
		" ", "",
		"-", "",
		"(", "",
		")", "",
	)
)

// NormalizeNumber normalizes a phone number by removing formatting characters and ensuring it starts with a "+".
// If the input is empty, it returns an empty string.
// If the input doesn't match the expected number pattern, it returns the original input unchanged.
func NormalizeNumber(num string) string {
	if num == "" {
		return ""
	}
	if !reNumber.MatchString(num) {
		return num
	}
	num = reNumberRepl.Replace(num)
	if !strings.HasPrefix(num, "+") {
		return "+" + num
	}
	return num
}
