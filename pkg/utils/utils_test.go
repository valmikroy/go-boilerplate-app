package utils

import "testing"

func TestHello(t *testing.T) {
	str := String_one()

	if str != "Hello World +" {
		t.Errorf("Unexpected string returned")
	}

	str = String_two()

	if str != "Hello World ++" {
		t.Errorf("Unexpected string returned")
	}
}
