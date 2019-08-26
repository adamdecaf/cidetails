// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package cidetails

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	// currentCIProvider should be set to the name for whichever CI provider
	// we are running these tests in. This is manually set in the config (e.g. .travis.yml)
	// to match against what is discovered.
	currentCIProvider = os.Getenv("CURRENT_CI_PROVIDER")
)

func init() {
	fmt.Printf("CURRENT_CI_PROVIDER=%s\n", currentCIProvider)
}

func TestCIDetails__In(t *testing.T) {
	if currentCIProvider == "" {
		t.Skip("CURRENT_CI_PROVIDER is unset")
	}

	if !In(currentCIProvider) {
		t.Fatalf("unable to match against %s", currentCIProvider)
	}
}

func TestCIDetails__Name(t *testing.T) {
	if currentCIProvider == "" {
		t.Skip("CURRENT_CI_PROVIDER is unset")
	}

	name := Name()
	if !strings.EqualFold(name, currentCIProvider) {
		t.Errorf("Name: got %s expected %s", name, currentCIProvider)
	}
}

func TestCIDetails__InCI(t *testing.T) {
	if currentCIProvider == "" {
		t.Skip("CURRENT_CI_PROVIDER is unset")
	}

	if !InCI() {
		t.Errorf("unable to detect we're in CI with %s", currentCIProvider)
	}
}
