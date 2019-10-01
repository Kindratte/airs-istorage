/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// Record used by Put
type Record struct {

	// Record itself

	RecordType int32
	ID         int64

	// Document

	DocType int32
	DocID   int64

	// Parent

	ParentType int32
	ParentID   int64

	// Part

	PartType int32
	PartID   int64

	// Offset of the event which affected the value
	Offset int64
	// 0-deleted, 1-active, 2-closed
	State int32
	// Data
	Value []byte
}