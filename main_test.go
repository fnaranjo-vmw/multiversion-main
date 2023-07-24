package main

import (
	"github.com/fnaranjo-vmw/multiversion-main/api"
	"testing"
)

func TestDefaultFoo(t *testing.T) {
	foo := api.DefaultFoo()
	if foo.Data != 1002 {
		t.Fail()
	}
	if foo.OtherData != "multiversion-main" {
		t.Fail()
	}
}

func TestNewFoo(t *testing.T) {
	foo := api.NewFoo(987)
	if foo.Data != 997 {
		t.Fail()
	}
	if foo.OtherData != "multiversion-main" {
		t.Fail()
	}
}
