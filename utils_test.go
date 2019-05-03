/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetPart(t *testing.T) {

	var r1 Records
	r1 = append(r1, Record{ID: 0, PartType: 1})
	r1 = append(r1, Record{ID: 0, PartType: 4})

	var r2 Records
	r2 = append(r2, Record{ID: 1, PartType: 2})
	r2 = append(r2, Record{ID: 1, PartType: 3})

	assert.NotNil(t, r1.GetPart(1))
	assert.Nil(t, r1.GetPart(2))
	assert.Nil(t, r1.GetPart(3))
	assert.NotNil(t, r1.GetPart(4))

	assert.Nil(t, r2.GetPart(1))
	assert.NotNil(t, r2.GetPart(2))
	assert.NotNil(t, r2.GetPart(3))
	assert.Nil(t, r2.GetPart(4))

}

func Test_ToRecords(t *testing.T) {

	// Empty channel
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var actual []Records
		go func() {
			cRecords := ToRecords(cRecord)
			for records := range cRecords {
				actual = append(actual, records)

			}
			done <- "done"
		}()
		close(cRecord)
		<-done
		require.Equal(t, 0, len(actual))
	}

	// One record
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var actual []Records
		go func() {
			cRecords := ToRecords(cRecord)
			for records := range cRecords {
				actual = append(actual, records)

			}
			done <- "done"
		}()
		cRecord <- Record{RecordType: 1, ID: 0, PartType: 1}
		close(cRecord)
		<-done
		require.Equal(t, 1, len(actual))
	}

}

func Test_ToRecordsTwoRecords(t *testing.T) {
	// Two records, same ID
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var actual []Records
		go func() {
			cRecords := ToRecords(cRecord)
			for records := range cRecords {
				actual = append(actual, records)

			}
			done <- "done"
		}()
		cRecord <- Record{RecordType: 1, ID: 0, PartType: 1}
		cRecord <- Record{RecordType: 1, ID: 0, PartType: 2}
		close(cRecord)
		<-done
		require.Equal(t, 1, len(actual))
	}

	// Two records, dif ID
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var expected []Records
		expected = append(expected, Records{Record{RecordType: 1, ID: 0, PartType: 1}})
		expected = append(expected, Records{Record{RecordType: 1, ID: 1, PartType: 2}})
		var actual []Records
		go func() {
			cRecords := ToRecords(cRecord)
			for records := range cRecords {
				actual = append(actual, records)

			}
			done <- "done"
		}()
		cRecord <- Record{RecordType: 1, ID: 0, PartType: 1}
		cRecord <- Record{RecordType: 1, ID: 1, PartType: 2}
		close(cRecord)
		<-done
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}

}

func Test_ToRecordsSixRecords(t *testing.T) {
	// 6 records, 3 - 1 - 2
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var expected []Records
		expected = append(expected,
			Records{
				Record{RecordType: 1, ID: 0, PartType: 1},
				Record{RecordType: 1, ID: 0, PartType: 2},
				Record{RecordType: 1, ID: 0, PartType: 3},
			},
			Records{
				Record{RecordType: 2, ID: 0, PartType: 1},
			},
			Records{
				Record{RecordType: 2, ID: 1, PartType: 1},
				Record{RecordType: 2, ID: 1, PartType: 2},
			},
		)

		var actual []Records
		go func() {
			cRecords := ToRecords(cRecord)
			for records := range cRecords {
				actual = append(actual, records)

			}
			done <- "done"
		}()
		cRecord <- expected[0][0]
		cRecord <- expected[0][1]
		cRecord <- expected[0][2]
		cRecord <- expected[1][0]
		cRecord <- expected[2][0]
		cRecord <- expected[2][1]
		close(cRecord)
		<-done
		require.Equal(t, 3, len(actual))
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}
}

func Test_ToSlice(t *testing.T) {
	// Empty channel
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var actual []Record
		var errExpected error
		var errActual error

		go func() {
			actual, errActual = ToSlice(cRecord, &errExpected)
			done <- "done"
		}()
		close(cRecord)
		<-done
		require.Equal(t, 0, len(actual))
		require.Equal(t, errExpected, errActual)
	}

	// Empty channel with error
	{
		cRecord := make(chan Record)
		done := make(chan string)

		var actual []Record

		errExpected := errors.New("emit macho dwarf: elf header corrupted")
		var errActual error

		go func() {
			actual, errActual = ToSlice(cRecord, &errExpected)
			done <- "done"
		}()
		close(cRecord)
		<-done
		require.Equal(t, 0, len(actual))
		require.Equal(t, errExpected, errActual)
	}
}

func Test_ToSliceFew(t *testing.T) {
	// Few record in channel
	{
		cRecord := make(chan Record)
		done := make(chan string)
		var actual []Record
		var errExpected error
		var errActual error

		var expected = []Record{
			Record{RecordType: 1, ID: 1, PartType: 1},
			Record{RecordType: 1, ID: 1, PartType: 2},
			Record{RecordType: 1, ID: 0, PartType: 1},
		}

		go func() {
			actual, errActual = ToSlice(cRecord, &errExpected)
			done <- "done"
		}()
		for _, record := range expected {
			cRecord <- record
		}
		close(cRecord)
		<-done
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
		require.Equal(t, errExpected, errActual)

	}
}
