/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package istorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPart(t *testing.T) {

	r1 := &RecordParts{ID: 0}
	r1.Parts = append(r1.Parts, Part{PartType: 1})
	r1.Parts = append(r1.Parts, Part{PartType: 4})

	r2 := RecordParts{ID: 1}
	r2.Parts = append(r2.Parts, Part{PartType: 2})
	r2.Parts = append(r2.Parts, Part{PartType: 3})

	assert.NotNil(t, r1.GetPart(1))
	assert.Nil(t, r1.GetPart(2))
	assert.Nil(t, r1.GetPart(3))
	assert.NotNil(t, r1.GetPart(4))

	assert.Nil(t, r2.GetPart(1))
	assert.NotNil(t, r2.GetPart(2))
	assert.NotNil(t, r2.GetPart(3))
	assert.Nil(t, r2.GetPart(4))

}
