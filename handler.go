package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	for scanner.Scan() {
		data := scanner.Text()
		res, err := CountPostfix(data)
		if err != nil {
			return err
		}
		_, err = ch.Output.Write([]byte(res))
		if err != nil {
			return fmt.Errorf("Відбулася помилка при записі результату!")
		}
		ch.Output.Write([]byte("\n"))
	}
	return nil
}
