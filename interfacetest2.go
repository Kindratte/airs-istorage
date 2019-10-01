/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import (
	"context"
	"testing"

	"github.com/untillpro/godif"
)

// DeclareForTest2 s.e.
func DeclareForTest2() {
	godif.Require(&GetEvents)
	godif.Require(&PutEvents)
	godif.Require(&GetMaxOffset1)
}

func testBasicUsage2(t *testing.T) {
	// WsID := newWsID()

	// assert.Equal(t, PutEvents(ctx, WsID, []Event{e}))

	// e := Event{Offset1: 0, Offset2: 1, Domain: 2, Type: 3, Data: []byte{4, 5}, MetaData: []byte{6, 7}}

	// assert.Nil(t, PutEvents(ctx, WsID, []Event{e}))
}

// TestImpl2 s.e.
// Storage must be empty before testing
func TestImpl2(actx context.Context, t *testing.T) {
	ctx = actx
	t.Run("testBasicUsage2", testBasicUsage2)
}
