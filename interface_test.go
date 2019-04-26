/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import (
	"testing"
)

func BenchmarkChanValue(b *testing.B) {
	c := make(chan RecordParts)
	done := make(chan string)

	go func() {
		for range c {
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := RecordParts{ID: 0}
		r.Parts = append(r.Parts, Part{PartType: 0, State: 1, Offset: 2})
		c <- r
	}
	close(c)
	<-done
}

func BenchmarkChanValueFromPointer(b *testing.B) {
	c := make(chan RecordParts)
	done := make(chan string)

	go func() {
		for range c {
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := new(RecordParts)
		r.ID = 0
		r.Parts = append(r.Parts, Part{PartType: 0, State: 1, Offset: 2})
		c <- *r
	}
	close(c)
	<-done
}

func BenchmarkChanPointer(b *testing.B) {
	c := make(chan *RecordParts)
	done := make(chan string)

	go func() {
		for range c {
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := new(RecordParts)
		r.ID = 0
		r.Parts = append(r.Parts, Part{PartType: 0, State: 1, Offset: 2})
		c <- r
	}
	close(c)
	<-done
}
