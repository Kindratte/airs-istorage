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
var PutBatch func(ctx context.Context, workspaceID int64, batch []*Record) error

// GetByID returns a record, nil if not found
// If partTypes is empty all parts are returned
var GetByID func(ctx context.Context, workspaceID int64, recordType int32, ID int64, partTypes []int32) (*Record, error)

// Get returns records with given workspaceID and recordType
// If useID is true result is filtered by ID
// If partType is not empty result is filtered by partTypes
// GetAll must analyze ctx.Done
// *error will be valid when chan is closed
var Get func(ctx context.Context, workspaceID int64, recordType int32, useID bool, ID int64, partTypes []int32) (chan *Record, *error)
