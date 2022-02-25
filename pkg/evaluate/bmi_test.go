package evaluate

import (
	"bmi-calculator/pkg/storage/models"
	"math"
	"regexp"
	"testing"
)

func TestEvaluate(t *testing.T) {
	c := models.New()

	// empty
	expected := "Your BMI is %.2f. Please ensure both height and weight are set."
	want := regexp.MustCompile(expected)
	bmi, resp := Evaluate(c)

	if bmi != -1 && !want.MatchString(resp) {
		t.Fatalf("Failed Evaluation. Got %s, expected %s", resp, expected)
	}

	// Setup for the actual tests
	// Tolerance for floating points.
	tolerance := 0.00001
	c.SetHeight(1.8)
	limits := [7]float64{9, 11, 15, 18, 24, 29, 31}
	response := `Your BMI is %.2f. `
	responses := [7]string{
		response + `That is considered incompatible with life.`,
		response + `This is critically underweight.`,
		response + `That is highly underweight. Please consult a hospital.`,
		response + `That is considered underweight.`,
		response + `That is considered normal weight.`,
		response + `That is considered pre-obese.`,
		response + `That is considered obese. Please consult a doctor.`,
	}
	// Calculate the weights for the bmi limits
	for i := 0; i < len(limits); i++ {
		c.SetWeight(getWeight(limits[i]))

		expected = responses[i]
		want = regexp.MustCompile(responses[i])
		bmi, resp = Evaluate(c)

		if math.Abs(limits[i]-bmi) > tolerance {
			t.Fatalf("Mismatching BMI: %.2f (actual) %.2f (expected)", bmi, limits[i])
		}
		if !want.MatchString(resp) {
			t.Fatalf("Failed Evaluation. Got %s, expected %s", resp, expected)
		}
	}

}

/*
getWeight is a helper method to get the Weight for a specific height and BMI
value. bmi = w / h² -> w = h² * bmi
*/
func getWeight(bmi float64) float64 {
	h := 1.8
	return math.Pow(h, 2) * bmi
}
