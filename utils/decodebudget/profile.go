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

// Profile holds the tunable constants of the allocation cost model. The
// defaults are calibrated against production OTLP payloads; start from
// DefaultProfile and change only the fields you need, since a zero field
// charges nothing.
type Profile struct {
	// MarginPercent covers allocations the model does not itemize; the
	// calibration tests measured them at under 3%.
	MarginPercent int
	// MapGrowthFactor scales a map entry's key and value slots: tables grow by
	// doubling and are at most 7/8 full.
	MapGrowthFactor  int
	MapEntryOverhead int
	ExtensionCost    int

	// protojson allocates per token and per object on top of the decoded
	// message, and its reflection-based list and map setters box each scalar
	// element and map entry.
	JSONTokenCost       int
	JSONListScalarCost  int
	JSONMapEntryBoxCost int
	JSONObjectCost      int
}

var defaultProfile = Profile{
	MarginPercent:       5,
	MapGrowthFactor:     4,
	MapEntryOverhead:    16,
	ExtensionCost:       64,
	JSONTokenCost:       8,
	JSONListScalarCost:  24,
	JSONMapEntryBoxCost: 32,
	JSONObjectCost:      8,
}

// DefaultProfile returns the profile a Proto with a nil Profile uses.
func DefaultProfile() Profile {
	return defaultProfile
}

func (p *Profile) withMargin(n int) int {
	return n + n*p.MarginPercent/100
}
