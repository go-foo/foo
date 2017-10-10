// Copyright (AT) 2017 Bill
// created 10 Oct, 2017
package foo

import (
    "fmt"
)

type Foo struct {}
func (f *Foo) Bar() {
    fmt.Println("bar!")
}

func NewFoo() *Foo {
    return &Foo{}
}
