/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import "context"

// PutBatch puts records in transaction
// WorkspaceID from records are ignored, one from parameters are used
var PutBatch func(ctx context.Context, workspaceID int64, batch []Record) error

// Get returns records with given workspaceID and recordType
// If useID is true result is filtered by ID
// If partType is not empty result is filtered by partTypes
// GetAll must analyze ctx.Done
// *error will be valid when chan is closed
var Get func(ctx context.Context, workspaceID int64, recordType int32, useID bool, ID int64, partTypes []int32) (chan RecordParts, *error)

/*

Channel: Pointer or Value

https://groups.google.com/forum/#!topic/golang-nuts/eM_a09l8yU0
You may be surprised how large a struct can get before passing it has a noticeable performance impact compared to passing a pointer to that
struct (not to mention accessing data behind a pointer involves an indirection, and especially when the data is shared across processors,
	there can be additional cost.

BenchmarkChanPointer-4   	 3000000	       478 ns/op	      80 B/op	       2 allocs/op
BenchmarkChanValue-4   	 3000000	       401 ns/op	      48 B/op	       1 allocs/op

*/
