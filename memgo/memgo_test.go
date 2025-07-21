package memgo

import (
	"testing"
)

func setUp(t *testing.T) *Memgo {
	m, err := NewMemgo("127.0.0.1", "1234")
	if err != nil {
		t.Fatalf("error on init %s", err.Error())
	}

	return m
}

func dispose(t *testing.T, memgo *Memgo) {
	stop := memgo.Dispose()
	if stop != true {
		t.Fatalf("stop result invalid")
	}
}

func TestMemgoBasic(t *testing.T) {
	// t.Skip("Skipping this test..")

	memgo := setUp(t)
	t.Run("TestProxyBasic", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}
	})
	t.Run("TestGetAll", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}

		getall := memgo.GetAll()
		if getall == "" {
			t.Fatalf("getall result invalid")
		}
	})

	dispose(t, memgo)
}
