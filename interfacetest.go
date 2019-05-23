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

// Same record type, two IDs, second has two parts
var source1 = []Record{
	{RecordType: 2, ID: 3, PartType: 4, State: 5, Offset: 6, Value: []byte{7, 8, 9}},
	{RecordType: 2, ID: 13, PartType: 4, State: 15, Offset: 16, Value: []byte{17, 18, 19}},
	{RecordType: 2, ID: 13, PartType: 24, State: 25, Offset: 26, Value: []byte{27, 28, 29}},
}

// Same as source1
var source2 = []Record{
	{RecordType: 102, ID: 103, PartType: 104, State: 105, Offset: 106, Value: []byte{107, 108, 109}},
	{RecordType: 102, ID: 113, PartType: 104, State: 115, Offset: 116, Value: []byte{117, 118, 119}},
	{RecordType: 102, ID: 113, PartType: 124, State: 125, Offset: 126, Value: []byte{127, 128, 129}},
}

// Different record types
var source3 = []Record{
	{RecordType: 202, ID: 203, PartType: 204, State: 205, Offset: 206, Value: []byte{207, 208, 209}},
	{RecordType: 201, ID: 213, PartType: 204, State: 215, Offset: 216, Value: []byte{217, 118, 219}},
	{RecordType: 201, ID: 213, PartType: 224, State: 225, Offset: 226, Value: []byte{227, 128, 229}},
}

// Different PartIDs
var source4 = []Record{

	{RecordType: 201, ID: 1203, PartType: 3204, PartID: 0, State: 205, Offset: 206, Value: []byte{207, 208, 209}},
	{RecordType: 201, ID: 1203, PartType: 3204, PartID: 1, State: 215, Offset: 216, Value: []byte{217, 118, 219}},
	{RecordType: 201, ID: 1203, PartType: 3205, PartID: 0, State: 205, Offset: 206, Value: []byte{207, 208, 209}},
	{RecordType: 201, ID: 1203, PartType: 3205, PartID: 1, State: 215, Offset: 216, Value: []byte{217, 118, 219}},

	{RecordType: 201, ID: 1204, PartType: 2204, PartID: 10, State: 205, Offset: 206, Value: []byte{207, 208, 209}},
	{RecordType: 201, ID: 1204, PartType: 2204, PartID: 11, State: 215, Offset: 216, Value: []byte{217, 118, 219}},
	{RecordType: 201, ID: 1204, PartType: 2204, PartID: 12, State: 225, Offset: 226, Value: []byte{217, 118, 219}},
	{RecordType: 201, ID: 1204, PartType: 2205, PartID: 0, State: 205, Offset: 206, Value: []byte{207, 208, 209}},
	{RecordType: 201, ID: 1204, PartType: 2205, PartID: 1, State: 215, Offset: 216, Value: []byte{217, 118, 219}},
	{RecordType: 201, ID: 1204, PartType: 2205, PartID: 2, State: 225, Offset: 226, Value: []byte{217, 118, 219}},
}

// TestImpl s.e.
// Storage must be empty before testing
func TestImpl(actx context.Context, t *testing.T) {

	ctx = actx

	t.Run("testBasicUsage", testBasicUsage)
	t.Run("TestOrder", testOrder)
	t.Run("testOrderPartID", testOrderPartID)
	t.Run("testFilteringRecordType", testFilteringRecordType)
	t.Run("testFilteringWs", testFilteringWs)
	t.Run("testCancelByErr", testCancelByErr)

}

func ids(ID int64) []int64 {
	return []int64{ID}
}

func ids2(ID int64) ([]int64, []int32) {
	return []int64{ID}, nil
}

func pid(ID int64) *int64 {
	return &ID
}

var wsID = int64(0)

func newWsID() int64 {
	wsID++
	return wsID
}

func testBasicUsage(t *testing.T) {

	WsID := newWsID()

	require.Nil(t, Put(ctx, WsID, source1[0:1]))

	// Record with two parts, put them as a two-records batch

	require.Nil(t, Put(ctx, WsID, source1[1:3]))

	// Fetch them all

	actual, err := ToSlice(Get(ctx, WsID, 2, nil, nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(source1, actual), "Expected %#v actual %#v", source1, actual)

	// Fetch by ID and PartType

	actual, err = ToSlice(Get(ctx, WsID, 2, ids(3), nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(source1[0:1], actual), "Expected %#v actual %#v", source1[0:1], actual)

	actual, err = ToSlice(Get(ctx, WsID, 2, ids(13), nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(source1[1:3], actual), "Expected %#v actual %#v", source1[1:3], actual)

}

func testOrder(t *testing.T) {

	// Result must be sorted by workspaceID, recordType, ID, PartType

	WsID := newWsID()

	// Insert in reverse order

	err := Put(ctx, WsID, source2[2:3])
	require.Nil(t, err)

	err = Put(ctx, WsID, source2[1:2])
	require.Nil(t, err)

	err = Put(ctx, WsID, source2[0:1])
	require.Nil(t, err)

	// Test

	actual, err := ToSlice(Get(ctx, WsID, 102, nil, nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(source2, actual), "Expected %#v actual %#v", source2, actual)

}

func testOrderPartID(t *testing.T) {
	// Result must be sorted by ID, PartType, PartID

	WsID := newWsID()

	for i := len(source4) - 1; i >= 0; i-- {
		err := Put(ctx, WsID, source4[i:i+1])
		require.Nil(t, err)
	}

	actual, err := ToSlice(Get(ctx, WsID, 201, nil, nil))
	require.Nil(t, err, "Get error")
	assert.True(t, reflect.DeepEqual(source4, actual), "Expected %#v Actual %#v", source4, actual)
}

func testFilteringRecordType(t *testing.T) {
	WsID := newWsID()

	require.Nil(t, Put(ctx, WsID, source3))
	{
		actual, err := ToSlice(Get(ctx, WsID, 202, nil, nil))
		require.Nil(t, err, "Get error")
		expected := source3[0:1]
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}
	{
		actual, err := ToSlice(Get(ctx, WsID, 201, nil, nil))
		require.Nil(t, err, "Get error")
		expected := source3[1:3]
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}
}

func testFilteringWs(t *testing.T) {
	WsID := newWsID()

	require.Nil(t, Put(ctx, WsID, source1))
	require.Nil(t, Put(ctx, WsID+1, source2))
	require.Nil(t, Put(ctx, WsID+2, source3))

	{
		actual, err := ToSlice(Get(ctx, WsID, 2, nil, nil))
		require.Nil(t, err, "Get error")
		expected := source1
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}
	{
		actual, err := ToSlice(Get(ctx, WsID+1, 102, nil, nil))
		require.Nil(t, err, "Get error")
		expected := source2
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}
	{
		actual, err := ToSlice(Get(ctx, WsID+2, 201, nil, nil))
		require.Nil(t, err, "Get error")
		expected := source3[1:3]
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}

}

// Get must analyze ctx.Err AFTER each write to channel
func testCancelByErr(t *testing.T) {

	WsID := newWsID()

	ctxCancel, cancel := context.WithCancel(ctx)

	require.Nil(t, Put(ctxCancel, WsID, source1))

	c, perr := Get(ctxCancel, WsID, 2, nil, nil)

	var actual []Record
	actual = append(actual, <-c)
	cancel()

	for r := range c {
		actual = append(actual, r)
	}
	require.Nil(t, *perr)
	expected := source1[0:2]
	assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)

}
