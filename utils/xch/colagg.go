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

package xch

import (
	"fmt"
	"strings"

	"github.com/ClickHouse/ch-go/proto"
)

// ColAggFunc represents AggregateFunction columns.
type ColAggFunc[T interface {
	proto.Decoder
	proto.Encoder
}] struct {
	Values []T
	typ    proto.ColumnType
	ctor   func() T
}

type ColAggFuncUniqStr = ColAggFunc[*UniquesHashSet]
type ColAggFuncUniqStrVar = ColAggFunc[*UniquesHashSetVar]

func MakeColAggFuncUniqStr() ColAggFuncUniqStr {
	return ColAggFuncUniqStr{
		typ:  proto.ColumnType("AggregateFunction(uniq, String)"),
		ctor: NewUniquesHashSet,
	}
}

func MakeColAggFuncUniqStrVar(n int) ColAggFuncUniqStrVar {
	return ColAggFuncUniqStrVar{
		typ:  proto.ColumnType(fmt.Sprintf("AggregateFunction(uniq, String%s)", strings.Repeat(", String", n-1))),
		ctor: NewUniquesHashSetVar,
	}
}

// Append aggregate to column.
func (c *ColAggFunc[T]) Append(v T) {
	c.Values = append(c.Values, v)
}

func (c *ColAggFunc[T]) AppendArr(v []T) {
	for _, e := range v {
		c.Append(e)
	}
}

// Type returns ColumnType of ColAggFunc.
func (c ColAggFunc[T]) Type() proto.ColumnType {
	return c.typ
}

// Rows returns count of rows in column.
func (c ColAggFunc[T]) Rows() int {
	return len(c.Values)
}

// First returns first row of column.
func (c ColAggFunc[T]) First() T {
	return c.Row(0)
}

// Row returns row with number i.
func (c ColAggFunc[T]) Row(i int) T {
	return c.Values[i]
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColAggFunc[T]) Reset() {
	clear(c.Values)
	c.Values = c.Values[:0]
}

// EncodeColumn encodes ColAggFunc rows to *Buffer.
func (c ColAggFunc[T]) EncodeColumn(b *proto.Buffer) {
	for _, s := range c.Values {
		s.Encode(b)
	}
}

// WriteColumn writes ColAggFunc rows to *Writer.
func (c ColAggFunc[T]) WriteColumn(w *proto.Writer) {
	for _, s := range c.Values {
		w.ChainBuffer(s.Encode)
	}
}

// DecodeColumn decodes ColAggFunc rows from *Reader.
func (c *ColAggFunc[T]) DecodeColumn(r *proto.Reader, rows int) error {
	c.Values = append(c.Values[:0], make([]T, rows)...)
	for i := range rows {
		c.Values[i] = c.ctor()
		if err := c.Values[i].Decode(r); err != nil {
			return err
		}
	}
	return nil
}
