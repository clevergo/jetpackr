// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package packrloader

import (
	"bytes"
	"io/ioutil"
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

func TestLoaderExists(t *testing.T) {
	box := packr.New("views", "./example/views")
	l := New(box)
	views := []string{"index.tmpl", "non-existence.tmpl"}
	for _, name := range views {
		actualName, actualExists := l.Exists(name)
		expectedExists := box.Has(name)
		if expectedExists != actualExists {
			t.Errorf("expected exists %t, got %t", expectedExists, actualExists)
		}
		if actualName != name {
			t.Errorf("expected name %s, got %s", name, actualName)
		}
	}
}

func TestLoaderOpen(t *testing.T) {
	box := packr.New("views", "./example/views")
	l := New(box)
	views := []string{"index.tmpl", "non-existence.tmpl"}
	for _, name := range views {
		file, err := l.Open(name)
		expectedFile, expectedErr := box.Open(name)
		if expectedErr != nil {
			if err == nil || err.Error() != expectedErr.Error() {
				t.Errorf("expected error %s, got %s", expectedErr, err)
			}
			continue
		}
		content, _ := ioutil.ReadAll(file)
		expectedContent, _ := ioutil.ReadAll(expectedFile)
		if !bytes.Equal(content, expectedContent) {
			t.Errorf("expected file content %s, got %s", content, expectedContent)
		}
	}
}
