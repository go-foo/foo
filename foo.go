// Copyright (AT) 2017 Bill
// created 10 Oct, 2017
package foo

import (
    "gopkg.in/foo.v1/internal/log"
)

const VERSION = "1.0.2"

type Foo struct {}
func (f *Foo) Bar() {
    log.Println("bar!", VERSION)
}

func NewFoo() *Foo {
    return &Foo{}
}
