// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

// +build !integration

package inventory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemCopy(t *testing.T) {
	sys1 := &System{
		Data: map[string]interface{}{
			"foo": "bar",
		},
	}

	sys2 := sys1.Copy()
	assert.Equal(t, *sys1, *sys2, "copied system should be identical")
	sys2.Data["foo2"] = "bar2"
	assert.NotEqual(t, *sys1, *sys2, "copied system change should not affect origin")
}
