package memgo

import (
	"testing"
)

func TestProxyBasic(t *testing.T) {
	// t.Skip("Skipping this test..")

	memgo, err := NewMemgo("127.0.0.1", "1234")
	if err != nil {
		t.Fatalf("error on init %s", err.Error())
	}

	set := memgo.Set("keyTest", "valueTest")
	if set != true {
		t.Fatalf("set result invalid")
	}

	stop := memgo.Dispose()
	if stop != true {
		t.Fatalf("stop result invalid")
	}
}
