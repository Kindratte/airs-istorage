/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// Record used by Put
type Record struct {
	RecordType int32
	ID         int64
	PartType   int32

	// 0-active, 1-closed, 2-deleted/moved
	State int32

	Offset int64
	// Default value
	Value []byte
}

// Part of the record
type Part struct {
	PartType int32

	// 0-active, 1-closed, 2-deleted/moved
	State int32

	Offset int64
	// Default value
	Value []byte
}

