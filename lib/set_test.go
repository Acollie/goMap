package lib

import "testing"

func TestSet_New(t *testing.T) {
	mapTest := new(Set[string]).New()

	if mapTest == nil {
		t.Error("Expected mapTest to not be nil")
	} else {
		t.Log("mapTest is not nil")
	}
}

func TestSet_Add(t *testing.T) {
	mapTest := new(Set[string]).New()
	mapTest.Add("foo bar")

	if mapTest.Len() != 1 {
		t.Error("Expected mapTest to have length 1")
	} else {
		t.Log("mapTest has length 1")
	}
}

func TestSet_Remove(t *testing.T) {
	mapTest := new(Set[string]).New()
	mapTest.Add("foo bar")
	mapTest.Remove("foo bar")

	if mapTest.Len() != 0 {
		t.Error(mapTest.Len())
		t.Error("Expected mapTest to have length 0")
	} else {
		t.Log("mapTest has length 0")
	}
}

func TestSetAddRemove(t *testing.T) {
	mapTest := new(Set[string]).New()
	mapTest.Add("foo bar")
	mapTest.Remove("foo bar")
	mapTest.Add("foo bar")

	if mapTest.Len() != 1 {
		t.Error("Expected mapTest to have length 1")
	} else {
		t.Log("mapTest has length 1")
	}
}

func TestOtherTypes(t *testing.T) {
	mapTest := new(Set[int]).New()
	mapTest.Add(1)
	mapTest.Add(2)
	mapTest.Add(3)
	mapTest.Remove(2)

	if mapTest.Len() != 2 {
		t.Error("Expected mapTest to have length 2")
	} else {
		t.Log("mapTest has length 2")
	}

	if mapTest.In(1) != true {
		t.Error("Expected mapTest to contain 1")
	} else {
		t.Log("mapTest contains 1")
	}

}

func TestSet_In(t *testing.T) {
	mapTest := new(Set[string]).New()
	mapTest.Add("foo bar")

	if mapTest.In("foo bar") != true {
		t.Error("Expected mapTest to contain foo bar")
	} else {
		t.Log("mapTest contains foo bar")
	}

	if mapTest.In("bar") != false {
		t.Error("Expected mapTest to not contain bar")
	} else {
		t.Log("mapTest does not contain bar")
	}
}
