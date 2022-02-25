# BMI calculator

This is a simple Command Line Interface (CLI) Application to evaluate your BMI
written in Go. It uses the Hexagonal Architecture and Domain Driven Design for
extensibility.

It includes the basic functionality of Go in terms of loops, conditionals, types
imports, pointers, functions, structures and methods.

## Structure

- cmd:
    - cli: holds the Command Line Interface main function.
- pkg:
    - calculate: includes the calculations and evaluations of the BMI.
    - storage: includes the models.

### Execute

There is an .exe in the build. To build a different executable head execute the
following command: 

`go build -o build/bmi-calculator github.com/daubersc/bmi-calculator/cmd/cli` 