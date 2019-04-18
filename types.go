/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// RecordType defines record type
type RecordType int32

// Record s.e.
type Record struct {
	ID    int64
	Parts []*Part
}

// PartType defines type of the record part
type PartType int32

// Part of the record
type Part struct {
	Type PartType

	// 0-active, 1-closed, 2-deleted/moved
	State int32

	Offset int64
	// Default value
	Value []byte
}

// GetPart part by type of nil if part not found
// For speed consideration ref. https://www.darkcoding.net/software/go-slice-search-vs-map-lookup/
func (r *Record) GetPart(partType PartType) *Part {
	for _, part := range r.Parts {
		if part.Type == partType {
			return part
		}
	}
	return nil
}
