package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File to read expression from")
	outputFile = flag.String("o", "", "File to write the results")
)

func getInputType(inputExpression, inputFile string) (io.Reader, error) {
	switch {
	case inputExpression != "":
		input := strings.NewReader(inputExpression)
		return input, nil
	case inputFile != "":
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, fmt.Errorf("Вказаний файль для зчитування не міг бути знайденим!")
		}
		return file, nil
	default:
		return nil, fmt.Errorf("Не було надано виразу для обчислення, або файл для зчитування виразу!")
	}
}

func getOutputType(outputFile string) io.Writer {
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Вказаний файл для виводу даний не вдалося створити! Вивід результату буде відбуватись в stdout!")
			output := os.Stdout
			return output
		}
		return file
	} else {
		output := os.Stdout
		return output
	}
}


func main() {
	flag.Parse()
	input, err := getInputType(*inputExpression, *inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	output := getOutputType(*outputFile)

	handler := &lab2.ComputeHandler{ 
		Input:  input, 
		Output: output, 
	}
	err = handler.Compute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
