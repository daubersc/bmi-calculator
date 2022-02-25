package models

import (
	"fmt"
	"math"
)

const (
	Weight float64 = 2.205
	Height         = 39.37
)

/*
BmiCalculator structure holds all relevant data required to calculate the Body
Mass Index (BMI), which is Weight, Height and the units of both. Note that we
use metric units for the calculation, therefore getters and setters have to
consider conversions.
*/
type BmiCalculator struct {
	Weight float64
	Height float64
	Unit   Unit
}

/*
New constructs a blank BmiCalculator structure with metrical units.
*/
func New() BmiCalculator {
	return BmiCalculator{
		Weight: 0,
		Height: 0,
		Unit:   Metrical,
	}
}

/*
SetWeight sets the Weight of a BmiCalculator while considering the
BmiCalculator's Unit. The Unit defaults to Metrical, so only diverging units
need to be converted.
*/
func (c *BmiCalculator) SetWeight(weight float64) {
	units := c.Unit.Weight()

	// Illegal Weight.
	if weight <= 0 {
		fmt.Printf("Illegal Weight: %.2f %s.", weight, units)
		return
	}

	// needs to be advanced to avoid temporary variables.
	fmt.Printf("Set Weight to %.2f %s.", weight, units)

	// Convert to metrical.
	if c.Unit == Imperial {
		weight /= Weight
	}

	c.Weight = weight
}

/*
SetHeight sets the BmiCalculator's Height while considering the BmiCalculator's
Unit. The Unit defaults to Metrical, so only diverging units need to be
converted.
*/
func (c *BmiCalculator) SetHeight(height float64) {
	units := c.Unit.Height()

	// Illegal Height
	if height <= 0 {
		fmt.Printf("Illegal Height %.2f %s.\n", height, units)
		return
	}

	// Convert and inform
	if c.Unit == Imperial {
		ft, in := convert(height)
		height /= Height
		fmt.Printf(`Height set to %d' %2.f".`, ft, in)
	} else {
		fmt.Printf(`Height set to %.2f m.`, height)
	}

	c.Height = height
}

/**
convert inches to ft and inch and return it.
*/
func convert(inch float64) (int, float64) {
	ft := math.Floor(inch / 12)
	in := inch - ft*12
	return int(ft), in
}

/*
SwitchUnit Switches the Unit of a person from Imperial to Metrical and vice
versa.
*/
func (c *BmiCalculator) SwitchUnit() {
	if c.Unit == Imperial {
		c.Unit = Metrical
	} else {
		c.Unit = Imperial
	}

	fmt.Printf("Selected %s units.", c.Unit)
}
