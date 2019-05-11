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
- res must be sorted by workspaceID, recordType, ID, PartType
- If ID is not 0 result is filtered by ID
  - Thus 0 is prohibited as ID value
- If ID is not 0 and partTypes is not empty result is filtered by ID and partTypes
- Get must analyze ctx.Err AFTER each write to channel
- *perr will be valid when chan is closed
*/
var Get func(ctx context.Context, workspaceID int64, recordType int32, ID int64, partTypes []int32) (res chan Record, perr *error)
