/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

// Event s.e.
type Event struct {
	Offset1  int64
	Offset2  int64
	Domain   int32
	Type     int32
	Data     []byte
	MetaData []byte
}
