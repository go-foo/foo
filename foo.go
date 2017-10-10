// Copyright (AT) 2017 Bill
// created 10 Oct, 2017
package foo

import (
    "fmt"
)

const VERSION = "1.0.0"

type Foo struct {}
func (f *Foo) Bar() {
    fmt.Println("bar! " + VERSION)
}

func NewFoo() *Foo {
    return &Foo{}
}
