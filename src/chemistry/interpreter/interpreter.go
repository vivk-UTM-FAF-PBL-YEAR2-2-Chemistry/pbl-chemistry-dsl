package interpreter

import (
	"chemistry/config"
	"chemistry/evaluator"
	"chemistry/lexer"
	"chemistry/object"
	"chemistry/parser"
	"fmt"
	"io"
	"os"
	"strings"
)

func Start(lines []string) {
	out := os.Stdout
	outErr := os.Stderr

	env := object.NewEnvironment()

	for index, line := range lines {
		line = strings.Trim(line, "\r")
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(index+1, outErr, p.Errors())
			os.Exit(config.PARSE_ERR)
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Type() != object.NONE_OBJ {
			var output io.Writer = nil
			var exit bool = false
			if evaluated.Type() == object.ERROR_OBJ {
				output = outErr
				exit = true
			} else {
				output = out
				exit = false
			}

			if exit {
				fmt.Fprintf(output, "%d : ", index+1)
			}

			io.WriteString(output, evaluated.Inspect())
			io.WriteString(output, "\n")

			if exit {
				os.Exit(config.EVAL_ERR)
			}
		}
	}
}

func printParserErrors(index int, out io.Writer, errors []string) {
	fmt.Fprintf(out, "%d :", index)
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
