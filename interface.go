/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import "context"

// PutBatch puts records in transaction
var PutBatch func(ctx context.Context, workspaceID int64, batch []Record) error

// Get returns records with given workspaceID and recordType
// If useID is true result is filtered by ID
// If partType is not empty result is filtered by partTypes
// GetAll must analyze ctx.Done
// *error will be valid when chan is closed
var Get func(ctx context.Context, workspaceID int64, recordType int32, useID bool, ID int64, partTypes []int32) (chan RecordParts, *error)
