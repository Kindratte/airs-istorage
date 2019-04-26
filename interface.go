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

/*Get returns records with given workspaceID and recordType
- Result must be sorted by workspaceID, recordType, PartType, ID
- If partType is not empty result is filtered by partTypes
- If ID is not nil result is filtered by ID
- Get must analyze ctx.Done
- *error will be valid when chan is closed
*/
var Get func(ctx context.Context, workspaceID int64, recordType int32, partTypes []int32, ID *int64) (chan Record, *error)
