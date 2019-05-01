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

func BenchmarkChanInterface(b *testing.B) {
	c := make(chan interface{})
	done := make(chan string)

	var sum int64
	go func() {
		var rp Record
		for rpr := range c {
			rp = rpr.(Record)
			sum += rp.ID
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := Record{ID: 0, PartType: 0, State: 1, Offset: 2}
		c <- r
	}
	close(c)
	<-done
}

func BenchmarkChanValue(b *testing.B) {
	c := make(chan Record)
	done := make(chan string)

	var sum int64
	go func() {
		for rp := range c {
			sum += rp.ID
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := Record{ID: 0, PartType: 0, State: 1, Offset: 2}
		c <- r
	}
	close(c)
	<-done
}

func BenchmarkChanValueFromPointer(b *testing.B) {
	c := make(chan Record)
	done := make(chan string)
	var sum int64

	go func() {
		for rp := range c {
			sum += rp.ID
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := &Record{ID: 0, PartType: 0, State: 1, Offset: 2}
		c <- *r
	}
	close(c)
	<-done
}

func BenchmarkChanPointer(b *testing.B) {
	c := make(chan *Record)
	done := make(chan string)
	var sum int64

	go func() {
		for rp := range c {
			sum += rp.ID
		}
		done <- "bye"
	}()

	for n := 0; n < b.N; n++ {
		r := &Record{ID: 0, PartType: 0, State: 1, Offset: 2}
		c <- r
	}
	close(c)
	<-done
}
