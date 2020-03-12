// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package packrloader

import (
	"reflect"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func TestNew(t *testing.T) {
	box := packr.New("test", "test")
	l := New(box)
	if !reflect.DeepEqual(box, l.box) {
		t.Errorf("expected box %v, got %v", box, l.box)
	}
}
