/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import "context"

// PutBatch puts record in transaction
// WorkspaceID from records are ignored, one from parameters are used
var PutBatch func(ctx context.Context, workspaceID int64, batch []*Record) error

// Get returns a record, nil if not found
// If partTypes is empty all parts are returned
var Get func(ctx context.Context, workspaceID int64, recordType int32, ID int64, partTypes []int32) (*Record, error)

// GetAll returns all records of the given type
// GetAll must analyze ctx.Done
// If partType is empty all parts are returned
var GetAll func(ctx context.Context, wsID int64, recordType int32, partTypes []int32) (chan *Record, error)

// WipeWorkspace removes all records in workspace
var WipeWorkspace func(ctx context.Context, wsID int64)
