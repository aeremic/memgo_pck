package memgo

import (
	"os"
	"testing"
)

var memgo Memgo

func setUp(m *testing.M, t *testing.T) *Memgo {
	memgo, err := NewMemgo("127.0.0.1", "1234")
	if err != nil {
		t.Fatalf("error on init %s", err.Error())
	}

	return memgo
}

func dispose(m *testing.M, t *testing.T) {
	stop := memgo.Dispose()
	if stop != true {
		t.Fatalf("stop result invalid")
	}
}

func TestMain(m *testing.M, t *testing.T) {
	setUp(m, t)

	code := m.Run()

	dispose(m, t)
	os.Exit(code)
}

func TestProxyBasic(m *testing.M, t *testing.T) {
	// t.Skip("Skipping this test..")

	set := memgo.Set("keyTest", "valueTest")
	if set != true {
		t.Fatalf("set result invalid")
	}
}

func TestGetAll(m *testing.M, t *testing.T) {
	// t.Skip("Skipping this test..")

	set := memgo.Set("keyTest", "valueTest")
	if set != true {
		t.Fatalf("set result invalid")
	}

	getall := memgo.GetAll()
	if getall == "" {
		t.Fatalf("getall result invalid")
	}
}
