package tasks

import "testing"

func TestHello(t *testing.T) {
	str := String_one("blah1")

	if str != "Hello blah1" {
		t.Errorf("Unexpected string returned")
	}

	str = String_two("blah2")

	if str != "Hello World blah2" {
		t.Errorf("Unexpected string returned")
	}
}
