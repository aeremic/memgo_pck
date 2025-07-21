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
	t.Run("TestSet", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}
	})
	t.Run("TestGet", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}

		get := memgo.Get("keyTest")
		if get == "NOT_FOUND" {
			t.Fatalf("get result invalid")
		}
	})
	t.Run("Delete", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}

		get := memgo.Delete("keyTest")
		if get != true {
			t.Fatalf("delete result invalid")
		}
	})
	t.Run("DeleteAll", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}

		get := memgo.DeleteAll()
		if get != true {
			t.Fatalf("deleteall result invalid")
		}
	})
	t.Run("TestGetAll", func(t *testing.T) {
		set := memgo.Set("keyTest", "valueTest")
		if set != true {
			t.Fatalf("set result invalid")
		}

		getall := memgo.GetAll()
		if getall == "NOT_FOUND" {
			t.Fatalf("getall result invalid")
		}
	})

	dispose(t, memgo)
}
