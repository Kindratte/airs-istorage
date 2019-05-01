/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// ToSlice converts chan Record to []Record
func ToSlice(records chan Record, perr *error) ([]Record, error) {
	var res []Record
	for r := range records {
		res = append(res, r)
	}
	return res, *perr
}

// Records returned by ToRecords
type Records []Record

// GetPart returns part by type or nil if part not found
func (rr Records) GetPart(partType int32) *Record {
	for _, r := range rr {
		if r.PartType == partType {
			return &r
		}
	}
	return nil
}

// RecordPartsFromRecords convert channels. Input channel must be sorted by ID
//func RecordPartsFromRecords(records chan Record, perr *error) (chan RecordParts, *error) {
/*

	session := getSession(ctx)

	iter := session.Query(getQueryText(workspaceID, recordType, useID, ID, partTypes)).Iter()
	res := make(chan istorage.RecordParts)
	var resError error

	go func() {
		var curRec *istorage.RecordParts
		var curPart = new(istorage.Part)
		var ID int64
		for iter.Scan(&ID, &curPart.PartType, &curPart.State, &curPart.Offset, &curPart.Value) {

			// Check if we has to exit
			if nil != ctx.Err() {
				resError = ctx.Err()
				close(res)
				break
			}

			isNewRec := nil == curRec
			if !isNewRec {
				isNewRec = curRec.ID != ID
			}

			// Send curRec
			if isNewRec {
				if curRec != nil {
					res <- *curRec
					curRec.ID = ID
				}
				// curRec = istorage.RecordParts{ID: dbrec.ID}
				curRec = new(istorage.RecordParts)
				curRec.ID = ID
			}
			curRec.Parts = append(curRec.Parts, istorage.Part{PartType: curPart.PartType, State: curPart.State, Offset: curPart.Offset, Value: curPart.Value})
			curPart = new(istorage.Part)
		}
		if curRec != nil {
			res <- *curRec
		}
		resError = iter.Close()
		close(res)
	}()

	return res, &resError

*/
//	return nil, perr
//}
