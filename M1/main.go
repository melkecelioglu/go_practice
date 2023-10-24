package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// InputData represents the data for the input form.
type InputData struct {
	Number1  int
	Number2  int
	Operator string
	Result   int
	Error    string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		inputData := InputData{}
		tmpl := template.Must(template.ParseFiles("input.html"))

		if r.Method == http.MethodPost {
			inputData.Number1, _ = strconv.Atoi(r.FormValue("number1"))
			inputData.Number2, _ = strconv.Atoi(r.FormValue("number2"))
			inputData.Operator = r.FormValue("operator")
			inputData.Result, inputData.Error = calculate(inputData.Number1, inputData.Number2, inputData.Operator)
		}

		tmpl.Execute(w, inputData)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		inputData := InputData{}
		tmpl := template.Must(template.ParseFiles("output.html"))

		if r.Method == http.MethodPost {
			inputData.Number1, _ = strconv.Atoi(r.FormValue("number1"))
			inputData.Number2, _ = strconv.Atoi(r.FormValue("number2"))
			inputData.Operator = r.FormValue("operator")
			inputData.Result, inputData.Error = calculate(inputData.Number1, inputData.Number2, inputData.Operator)
		}

		tmpl.Execute(w, inputData)
	})
	http.ListenAndServe(":8080", nil)

}

func calculate(number1, number2 int, operator string) (int, string) {
	output := 0
	var err string

	switch operator {
	case "+":
		output = number1 + number2
	case "-":
		output = number1 - number2
	case "*":
		output = number1 * number2
	case "/":
		if number2 == 0 {
			err = "Division by zero"
		} else {
			output = number1 / number2
		}
	case "%":
		if number2 == 0 {
			err = "Modulo by zero"
		} else {
			output = number1 % number2
		}
	default:
		err = "Invalid Operation"
	}

	return output, err
}
