/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// ToSlice converts chan Record to []Record
func ToSlice(records <-chan Record, perr *error) ([]Record, error) {
	var res []Record
	for r := range records {
		res = append(res, r)
	}
	return res, *perr
}

// Records returned by ToRecords
type Records []Record

// GetPart returns first part by type or nil if part not found
func (rr Records) GetPart(partType int32) *Record {
	for _, r := range rr {
		if r.PartType == partType {
			return &r
		}
	}
	return nil
}

// GetParts returns all parts by type or nil if part not found
func (rr Records) GetParts(partType int32) []Record {
	var res []Record
	for _, r := range rr {
		if r.PartType == partType {
			res = append(res, r)
		}
	}
	return nil
}

// ToRecords s.e.
func ToRecords(records chan Record) chan Records {

	res := make(chan Records)

	go func() {
		var curSlice Records
		for r := range records {

			isNewRec := nil == curSlice
			if !isNewRec {
				isNewRec = !(curSlice[0].ID == r.ID && curSlice[0].RecordType == r.RecordType)
			}

			// Send curSlice
			if isNewRec {
				if curSlice != nil {
					res <- curSlice
					curSlice = nil
				}
			}
			curSlice = append(curSlice, r)
		}
		if curSlice != nil {
			res <- curSlice
		}
		close(res)
	}()
	return res
}
