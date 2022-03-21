// Copyright (AT) 2017 Bill
// created 10 Oct, 2017
package foo

import (
	"gopkg.in/foo/internal/log"
)

const VERSION = "2.0.0"

// Foo is a name, just for fun.
type Foo struct{}

func (f *Foo) Bar() {
	log.Println("bar!", VERSION)
}

// NewFoo returns an instance of Foo. For example:
//
//	foo := NewFoo()
//	foo.Bar()
//
// the above foo will ouput some text to stdout.
func NewFoo() *Foo {
	return &Foo{}
}
