package repl

import (
	"bufio"
	"fmt"
	"io"

	"minilang/lexer"
	"minilang/parser"
	"minilang/token"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	mode := "parse" // simple string flag

	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()

		// Handle commands
		switch line {
		case "/tokens":
			mode = "tokens"
			fmt.Fprintln(out, "Mode: tokens")
			continue
		case "/parse":
			mode = "parse"
			fmt.Fprintln(out, "Mode: parse")
			continue
		}

		// Process based on mode
		switch mode {
		case "tokens":
			printTokens(line, out)
		case "parse":
			printAST(line, out)
		case "eval":
			fmt.Println("Not Implemented yet!")
			// eval(line, out)
		}
	}
}

func printTokens(line string, out io.Writer) {
	l := lexer.New(line)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Fprintf(out, "%+v\n", tok)
	}
}

func printAST(line string, out io.Writer) {
	l := lexer.New(line)
	p := parser.New(l)
	program := p.ParseProgram()

	errors := p.Errors()
	if len(errors) != 0 {
		fmt.Fprintf(out, "parser has %d errors\n", len(errors))
		for _, msg := range errors {
			fmt.Fprintf(out, "parser error: %s\n", msg)
		}
		return
	}

	fmt.Fprintln(out, program.String())
}
