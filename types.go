/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// Record s.e.
type Record struct {
	WorkspaceD int64
	ID         int64
	Type       int32
	Parts      []*Part
}

// Part of the record
type Part struct {
	Type int32

	// 0-active, 1-closed, 2-deleted/moved
	State int32

	Offset int64
	// Default value
	Value []byte
}

// GetPart returns part by type or nil if part not found
// For speed consideration ref. https://www.darkcoding.net/software/go-slice-search-vs-map-lookup/
func (r *Record) GetPart(partType int32) *Part {
	for _, part := range r.Parts {
		if part.Type == partType {
			return part
		}
	}
	return nil
}
