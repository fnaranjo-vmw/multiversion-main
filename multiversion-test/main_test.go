package main

import (
	"fmt"
	"testing"
	c001 "github.com/fnaranjo-vmw/multiversion-main/001/api"
	c002 "github.com/fnaranjo-vmw/multiversion-main/002/api"
	cdev "github.com/fnaranjo-vmw/multiversion-main/dev/api"
)

func TestExpectedDefaultFoo001(t *testing.T) {
	if c001.DefaultFoo().Data != 972 {
		t.Log("We were expecting Data to be 972 but it was " + fmt.Sprint(c001.DefaultFoo().Data))
		t.Fail()
	}
}

func TestActualDefaultFoo001(t *testing.T) {
	if c001.DefaultFoo().Data != 993 {
		t.Fail()
	}
	if c001.DefaultFoo().OtherData == c001.DefaultFoo().OtherData {
		t.Log("OtherData didn't exist in this version")
		t.Fail()
	}
}

func TestDefaultFoo002(t *testing.T) {
	if c002.DefaultFoo().Data != 997 {
		t.Log("We were expecting Data to be 997 but it was " + fmt.Sprint(c002.DefaultFoo().Data))
		t.Fail()
	}
	if c002.DefaultFoo().OtherData != c002.DefaultFoo().OtherData {
		t.Fail()
	}
}

func TestDefaultFooDev(t *testing.T) {
	if cdev.DefaultFoo().Data != 1002 {
		t.Fail()
	}
	if c002.DefaultFoo().OtherData != "multiversion-main" {
		t.Fail()
	}
}
