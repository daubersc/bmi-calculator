package models

import (
	"regexp"
	"testing"
)

/*
TestUnit_Height checks whether the correct unit is returned based on the Unit
enum. Granted there cannot be other units, we can't do negative tests.
*/
func TestUnit_Height(t *testing.T) {
	// Metrical test
	var u Unit = Metrical
	unit := "m"
	want := regexp.MustCompile(unit)
	actual := u.Height()
	if !want.MatchString(actual) {
		t.Fatal("wrong unit")
	}

	// Imperial test
	u = Imperial
	unit = "in"
	want = regexp.MustCompile(unit)
	actual = u.Height()

	if !want.MatchString(actual) {
		t.Fatal("wrong unit")
	}
}

/*
TestUnit_Weight checks whether the correct unit is returned based on the Unit
enum. Granted there cannot be other units, we can't do negative tests.
*/
func TestUnit_Weight(t *testing.T) {
	// Metrical test
	var u Unit = Metrical
	unit := "kg"
	want := regexp.MustCompile(unit)
	actual := u.Weight()
	if !want.MatchString(actual) {
		t.Fatal("wrong unit")
	}

	// Imperial test
	u = Imperial
	unit = "lbs"
	want = regexp.MustCompile(unit)
	actual = u.Weight()

	if !want.MatchString(actual) {
		t.Fatal("wrong unit")
	}
}
