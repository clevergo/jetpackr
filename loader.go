// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package jetpackr

import (
	"io"

	"github.com/CloudyKit/jet/v4"
	"github.com/gobuffalo/packr/v2"
)

var _ jet.Loader = (*Loader)(nil)

// Loader a file system loader.
type Loader struct {
	box *packr.Box
}

// New returns a loader associate with the given packr box.
func New(box *packr.Box) *Loader {
	return &Loader{box: box}
}

// Open implements interface loader's open function.
func (l *Loader) Open(name string) (io.ReadCloser, error) {
	return l.box.Open(name)
}

// Exists implements interface loader's exists function.
func (l *Loader) Exists(name string) (string, bool) {
	return name, l.box.Has(name)
}
