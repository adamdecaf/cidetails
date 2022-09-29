// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package cidetails

import (
	"strings"
)

// In returns true if any of the names match CI server names.
func In(names ...string) bool {
	for i := range names {
		return in(names[i]) != nil
	}
	return false
}

func in(name string) *vendor {
	for i := range vendors {
		for j := range vendors[i].names {
			if strings.EqualFold(vendors[i].names[j], name) && vendors[i].envVar() {
				return &vendors[i]
			}
		}
	}
	return nil
}

// Name returns a string containing name of the CI server the code is running on.
// If CI server is not detected, it returns an empty string.
//
// Don't depend on the value of this string, use In.
func Name() string {
	if v := find(); v != nil {
		return v.names[0] // first name is most recent
	}
	return ""
}

// InCI returns true if we detected ourselves in a CI server.
func InCI() bool {
	return find() != nil
}

// IsPR returns true if the CI server has Pull Request detection and
// if we are testing a pull request.
func IsPR() bool {
	if v := find(); v != nil {
		return v.pr()
	}
	return false
}

func find() *vendor {
	for i := range vendors {
		for j := range vendors[i].names {
			if v := in(vendors[i].names[j]); v != nil {
				return &vendors[i]
			}
		}
	}
	return nil
}
