/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import "context"

// Put puts records in transaction
var Put func(ctx context.Context, workspaceID int64, batch []Record) error

/*Get returns channel for records with given workspaceID and recordType
- res buffer length must be zero
- res must be sorted by ID, PartType
- If IDs is not empty result is filtered by IDs
- If IDs is not empty and partTypes is not empty result is filtered by IDs and partTypes
- Get must analyze ctx.Err AFTER each write to channel
- *perr will be valid when chan is closed
*/
var Get func(ctx context.Context, workspaceID int64, recordType int32, IDs []int64, partTypes []int32) (res <-chan Record, perr *error)
