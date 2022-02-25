package models

import (
	"math"
	"testing"
)

// TestNew tests the construction of the structure with the correct values.
func TestNew(t *testing.T) {
	// expected values
	var w float64 = 0
	var h float64 = 0
	u := Metrical

	// actual values
	c := New()

	if u != c.Unit {
		t.Fatalf("Wrong unit system. Expected %s, got %s", u, c.Unit)
	} else if w != c.Weight {
		t.Fatalf("Wrong w. Expected %.0f, got %.0f", w, c.Weight)
	} else if h != c.Height {
		t.Fatalf("Wrong w. Expected %.0f, got %.0f", h, c.Height)
	}
}

var calculator BmiCalculator = New()

/*
TestBmiCalculator_SwitchUnit checks whether the unit is actually switched.
*/
func TestBmiCalculator_SwitchUnit(t *testing.T) {
	// pre-existing
	unit := Metrical
	if calculator.Unit != unit {
		t.Fatalf("Wrong unit system. Expected %s, got %s", unit, calculator.Unit)
	}

	// Perform action and check
	calculator.SwitchUnit()
	unit = Imperial
	if calculator.Unit != unit {
		t.Fatalf("Wrong unit system. Expected %s, got %s", unit, calculator.Unit)
	}
}

/*
TestBmiCalculator_SetHeight checks whether the height is correctly set. Since
the setting of metrical is trivial, we check the conversion of height from
imperial to meter that is included.
*/
func TestBmiCalculator_SetHeight(t *testing.T) {
	// parameter
	var inch float64 = 70
	var meter = 1.778

	calculator.SetHeight(-1)
	if math.Abs(calculator.Height-0) > 0.00001 {
		t.Fatalf("Wrong height. Expected %d, got %f", 0, calculator.Height)
	}

	// actual, since you can't compare floats natively. Solved using tolerance.
	calculator.SetHeight(inch)
	if math.Abs(calculator.Height-meter) > 0.00001 {
		t.Fatalf("Wrong height. Expected %f, got %f", meter, calculator.Height)
	}

}

/*
TestBmiCalculator_SetWeight checks whether the weight is correctly set. Since
the setting of metrical is trivial, we check the conversion lbs to kg that is
included.
*/
func TestBmiCalculator_SetWeight(t *testing.T) {
	// parameter
	var lbs float64 = 150
	var kg = 68.0272108

	//
	calculator.SetWeight(-1)
	if math.Abs(calculator.Weight-0) > 0.00001 {
		t.Fatalf("Wrong weight. Expected %d, got %f", 0, calculator.Weight)
	}

	// actual, since you can't compare floats natively. Solved using tolerance.
	calculator.SetWeight(lbs)
	if math.Abs(calculator.Weight-kg) > 0.00001 {
		t.Fatalf("Wrong weight. Expected %f, got %f", kg, calculator.Weight)
	}
}
