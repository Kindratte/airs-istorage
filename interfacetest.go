/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx context.Context

// TestImpl s.e.
// Storage must be empty before testing
func TestImpl(actx context.Context, t *testing.T) {

	ctx = actx

	t.Run("TestBasicUsage", testBasicUsage)
	// t.Run("TestFailedInit", TestFailedInit)
	// t.Run("TestFailedStart", TestFailedStart)
}

func pid(ID int64) *int64 {
	return &ID
}

func testBasicUsage(t *testing.T) {

	// Note: WsID = 1

	// Record with one part, one-record batch

	expected := []Record{
		{RecordType: 2, PartType: 4, ID: 3, State: 5, Offset: 6, Value: []byte{7, 8, 9}},
		{RecordType: 2, PartType: 4, ID: 13, State: 15, Offset: 16, Value: []byte{17, 18, 19}},
		{RecordType: 2, PartType: 24, ID: 13, State: 25, Offset: 26, Value: []byte{27, 28, 29}},
	}

	err := Put(ctx, 1, expected[0:1])

	require.Nil(t, err, "Error inserting record")

	// Record with two parts, put them as a two-records batch

	err = Put(ctx, 1, expected[1:3])
	require.Nil(t, err, "Error in batch")

	// Fetch them all

	var actual []Record
	actual, err = ToSlice(Get(ctx, 1, 2, nil, nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)

	// Fetch by PartType, ID

	actual, err = ToSlice(Get(ctx, 1, 2, []int32{4}, nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(expected[0:2], actual), "Expected %#v actual %#v", expected[0:2], actual)

	actual, err = ToSlice(Get(ctx, 1, 2, []int32{4}, pid(3)))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(expected[0:1], actual), "Expected %#v actual %#v", expected[0:1], actual)

	actual, err = ToSlice(Get(ctx, 1, 2, []int32{4, 24}, pid(13)))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(expected[1:3], actual), "Expected %#v actual %#v", expected[1:3], actual)

}
