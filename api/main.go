package api

import (
	"github.com/fnaranjo-vmw/multiversion-companion/a"
	"github.com/fnaranjo-vmw/multiversion-companion/b"
)

func DefaultFoo() a.Foo {
	foo := b.DefaultFoo()
	foo.Data += 10
	foo.OtherData = "multiversion-main"
	return foo
}

func NewFoo(data int) a.Foo {
	foo := b.NewFoo(data)
	foo.Data += 10
	foo.OtherData = "multiversion-main"
	return foo
}
