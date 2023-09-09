// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build darwin || (linux && !android)

package openssl

import (
	"testing"
)

func TestRand(t *testing.T) {
	_, err := RandReader.Read(make([]byte, 5))
	if err != nil {
		t.Fatal(err)
	}
}
