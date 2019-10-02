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
	"github.com/untillpro/godif"
)

// DeclareForTest2 s.e.
func DeclareForTest2() {
	DeclareForTest()
	godif.Require(&GetEvents)
	godif.Require(&PutEvents)
}

// TestImpl2 s.e.
// Storage must be empty before testing
func TestImpl2(actx context.Context, t *testing.T) {
	TestImpl(actx, t)
	t.Run("testBasicUsage2", testBasicUsage2)
	t.Run("testFilteringWSByGetEvents", testFilteringWSByGetEvents)
}

func testBasicUsage2(t *testing.T) {
	WSID := newWsID()

	var source = []Event{
		{Offset1: 0, Offset2: 1, Domain: 2, Type: 3, Data: []byte{4, 5}, MetaData: []byte{6, 7}},
		{Offset1: 1, Offset2: 2, Domain: 3, Type: 4, Data: []byte{5, 6}, MetaData: []byte{7}},
		{Offset1: 2, Offset2: 3, Domain: 4, Type: 5, Data: []byte{6}, MetaData: []byte{7, 8, 9}},
	}

	// Empty log
	{
		actual, err := ToEventSlice(GetEvents(ctx, WSID, 0, 0))
		require.Nil(t, err)
		assert.True(t, len(actual) == 0)
	}

	// Insert one event
	{
		err := PutEvents(ctx, WSID, source[0:1])
		require.Nil(t, err)
		actual, err := ToEventSlice(GetEvents(ctx, WSID, 0, 0))
		expected := source[0:1]
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}

	// Insert two events in batch
	{
		err := PutEvents(ctx, WSID, source[1:3])
		require.Nil(t, err)
		// First event must be the same
		{
			actual, err := ToEventSlice(GetEvents(ctx, WSID, 0, 0))
			require.Nil(t, err)
			expected := source[0:1]
			assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
		}
		// Events [1:2]
		{
			actual, err := ToEventSlice(GetEvents(ctx, WSID, 1, 2))
			require.Nil(t, err)
			expected := source[1:3]
			assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
		}
		// Events [0:2]
		{
			actual, err := ToEventSlice(GetEvents(ctx, WSID, 0, 2))
			require.Nil(t, err)
			expected := source[0:3]
			assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)

		}
	}

}

func testFilteringWSByGetEvents(t *testing.T) {
	WSID1 := newWsID()
	WSID2 := newWsID()
	var source = []Event{
		{Offset1: 0, Offset2: 1, Domain: 2, Type: 3, Data: []byte{4, 5}, MetaData: []byte{6, 7}},
		{Offset1: 0, Offset2: 2, Domain: 3, Type: 4, Data: []byte{5, 6}, MetaData: []byte{7}},
	}
	// Empty WS1
	{
		actual, err := ToEventSlice(GetEvents(ctx, WSID1, 0, 0))
		require.Nil(t, err)
		assert.True(t, len(actual) == 0)
	}
	// Empty WS2
	{
		actual, err := ToEventSlice(GetEvents(ctx, WSID2, 0, 0))
		require.Nil(t, err)
		assert.True(t, len(actual) == 0)
	}

	// Insert one event to both WS
	{
		err := PutEvents(ctx, WSID1, source[0:1])
		require.Nil(t, err)
		err = PutEvents(ctx, WSID2, source[1:2])
		require.Nil(t, err)

		actual, err := ToEventSlice(GetEvents(ctx, WSID1, 0, 0))
		expected := source[0:1]
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)

		actual, err = ToEventSlice(GetEvents(ctx, WSID2, 0, 0))
		expected = source[1:2]
		assert.True(t, reflect.DeepEqual(expected, actual), "Expected %#v actual %#v", expected, actual)
	}

}
