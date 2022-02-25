package evaluate

import (
	"bmi-calculator/pkg/storage/models"
	"fmt"
	"math"
)

/*
calculateBMI calculates the BMI for a BmiCalculator and stores it. The BMI uses
the formula weight / height ². The evaluation happens in the evaluation package,
as for the hexagonal architecture.
*/
func calculateBMI(c models.BmiCalculator) float64 {

	// Skip if height or weight is not set:
	if c.Weight <= 0 || c.Height <= 0 {
		return -1
	} else {
		return c.Weight / math.Pow(c.Height, 2)
	}
}

/*
Evaluate calculates the BMI and evaluates it according to the World Health
Organization's classification.
*/
func Evaluate(c models.BmiCalculator) (float64, string) {

	// Calculate.
	bmi := calculateBMI(c)
	response := `Your BMI is %.2f. `

	// Evaluate
	if bmi < 0 {
		response = `Please ensure both height and weight are set.`
	} else if bmi < 10 {
		response += `That is considered incompatible with life.`
	} else if bmi < 12 {
		response += `This is critically underweight.`
	} else if bmi < 16 {
		response += `That is highly underweight. Please consult a hospital.`
	} else if bmi < 18.5 {
		response += `That is considered underweight.`
	} else if bmi < 25 {
		response += `That is considered normal weight.`
	} else if bmi < 30 {
		response += `That is considered pre-obese.`
	} else {
		response += `That is considered obese. Please consult a doctor.`
	}

	fmt.Printf(response, bmi)
	return bmi, response
}
