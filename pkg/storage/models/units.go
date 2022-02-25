package models

/*
Unit represents the units used for IO while interacting with the user as people
who are used to Imperial can't fully grasp Metrical and vice versa.
*/
type Unit string

const (
	// Metrical The metrical name.
	Metrical Unit = "metrical"

	// Imperial The imperial name.
	Imperial = "imperial"
)

/*
Height returns the units for Height as a string
*/
func (u Unit) Height() string {
	if u == Imperial {
		return "in"
	} else {
		return "m"
	}
}

/*
Weight returns the units for the Weight as a string.
*/
func (u Unit) Weight() string {
	if u == Imperial {
		return "lbs"
	} else {
		return "kg"
	}
}
