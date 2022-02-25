package main

import (
	"bmi-calculator/pkg/evaluate"
	"bmi-calculator/pkg/storage/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
parseArgs takes the user input and slices it up. It then evaluates the command
and the parameter(s). Each command should have upmost 1 parameter, but for
imperial height the param is presented in ft and in, which will be converted.

Returns the command and the parameter with the correct and expected data types:
commands are string and param is float64.
*/
func parseArgs(stdin string) (string, float64) {
	var param float64 = -1
	var cmd = ""
	var err error

	args := strings.Split(stdin, " ")

	// There should be at least 1 argument, the command.
	// There are up to 2 optional parameters, e.g. height <ft> <in>
	if 1 <= len(args) && len(args) <= 3 {
		cmd = strings.ToLower(args[0])
		cmd = strings.Replace(cmd, "\n", "", -1)
		cmd = strings.Replace(cmd, "\r", "", -1)

		// Remove eventual units in params
		for i := 1; i < len(args); i++ {
			reg, err := regexp.Compile("[^0-9.]+")
			logs(err)
			args[i] = reg.ReplaceAllString(args[i], "")
		}

		// exactly one param given -> simply parse to float.
		if len(args) == 2 {
			param, err = strconv.ParseFloat(args[1], 64)
			logs(err)

			// if two params given and the command is height:
			// assume its feet and inches and convert to just inch.
		} else if len(args) == 3 && cmd == "height" {
			ft, err := strconv.ParseFloat(args[1], 64)
			logs(err)
			in, err := strconv.ParseFloat(args[2], 64)
			logs(err)
			param = ft*12 + in
		}
	}
	return cmd, param
}

/*
logs eventual fatal errors and quits the application.
*/
func logs(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
help prints the help text to get an overview over supported commands. The
supported commands are:

- help (see this method)

- units (see models.BmiCalculator)

- height (see models.BmiCalculator)

- weight (see models.BmiCalculator)

- eval (see calculate.Evaluate())

- quit (see os.Exit)
*/
func help() {
	stdout := "The BMI Calculator supports the following commands:\n" +
		"- help: Prints the help text.\n" +
		"- units: Switches from metrical units to imperial units.\n" +
		"- height <value>: Sets your body height. Use <m> for metrical " +
		"units or <ft in> or <in> for imperial units.\n" +
		"- weight <value>: Sets your weight. Use <kg> for metrical units or " +
		"<lbs> for imperial units.\n" +
		"- eval: Evaluates your BMI after you have set your height and " +
		"weight.\n" +
		"- quit: stops the application."
	fmt.Println(stdout)
}

/*
status prints an info text to notify the user about settings and such.
*/
func status(c models.BmiCalculator) {
	stdout := "BMI Calculator\n" +
		"--------------\n" +
		"Calculates and evaluates your Body Mass Index (BMI). Units are " +
		c.Unit + ". Please use help for further instructions on usage."

	fmt.Println(stdout)
}

/*
main starts the CLI and keeps it running until the user wants to quit.

DevNote: Using fmt.Scanf() instead of Reader would cause errors if users enter
the units along so for defensive programming, we parse the arguments manually
(see parseArgs).
*/
func main() {

	// Functionality revolving around BMI.
	calc := models.New()

	// User Input.
	r := bufio.NewReader(os.Stdin)

	// Prompts the input text.
	status(calc)

	for true {
		fmt.Print("\n>")

		stdin, _ := r.ReadString('\n')

		cmd, param := parseArgs(stdin)

		switch cmd {
		case "quit":
			os.Exit(0)
		case "units":
			calc.SwitchUnit()
		case "eval":
			evaluate.Evaluate(calc)
		case "height":
			calc.SetHeight(param)
		case "weight":
			calc.SetWeight(param)
		default:
			help()
		}
	}
}
