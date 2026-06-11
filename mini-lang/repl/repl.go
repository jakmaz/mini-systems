package repl

import (
	"bufio"
	"fmt"
	"io"

	"minilang/ast"
	"minilang/compiler"
	"minilang/evaluator"
	"minilang/lexer"
	"minilang/object"
	"minilang/parser"
	"minilang/token"
	"minilang/vm"
)

const prompt = "> "

type mode string

const (
	modeTokens  mode = "tokens"
	modeParse   mode = "parse"
	modeEval    mode = "eval"
	modeCompile mode = "compile"
	modeRun     mode = "run"
)

var commandHelp = map[string]string{
	"/tokens":  "Show tokens produced by the lexer",
	"/parse":   "Show AST produced by the parser",
	"/eval":    "Evaluate input using the tree-walking evaluator",
	"/compile": "Show bytecode produced by the compiler",
	"/run":     "Compile and run input on the VM",
	"/help":    "Show this help message",
	"/exit":    "Exit the REPL",
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	m := modeEval

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	for {
		fmt.Fprint(out, prompt)
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		// Handle commands
		if handled, newMode := handleCommand(line, m, out); handled {
			m = newMode
			continue
		}

		// Process based on mode
		switch m {
		case modeTokens:
			printTokens(line, out)
		case modeParse:
			printAST(line, out)
		case modeEval:
			printEval(line, out, env)
		case modeCompile:
			printCompile(line, out, symbolTable, &constants)
		case modeRun:
			printRun(line, out, symbolTable, &constants, globals)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(out, "scanner error: %s\n", err)
	}
}

func handleCommand(line string, current mode, out io.Writer) (bool, mode) {
	switch line {
	case "/tokens":
		fmt.Fprintln(out, "Mode: tokens")
		return true, modeTokens
	case "/parse":
		fmt.Fprintln(out, "Mode: parse")
		return true, modeParse
	case "/eval":
		fmt.Fprintln(out, "Mode: eval")
		return true, modeEval
	case "/compile":
		fmt.Fprintln(out, "Mode: compile")
		return true, modeCompile
	case "/run":
		fmt.Fprintln(out, "Mode: run")
		return true, modeRun
	case "/help":
		fmt.Fprintln(out, "Available commands:")
		for cmd, desc := range commandHelp {
			fmt.Fprintf(out, "  %-10s %s\n", cmd, desc)
		}
		return true, current
	case "/exit", "/quit":
		fmt.Fprintln(out, "Goodbye!")
		return true, current
	}
	return false, current
}

func parseLine(line string, out io.Writer) (*ast.Program, bool) {
	l := lexer.New(line)
	p := parser.New(l)
	program := p.ParseProgram()

	errors := p.Errors()
	if len(errors) != 0 {
		fmt.Fprintf(out, "parser has %d errors\n", len(errors))
		for _, msg := range errors {
			fmt.Fprintf(out, "parser error: %s\n", msg)
		}
		return nil, false
	}

	return program, true
}

func printTokens(line string, out io.Writer) {
	l := lexer.New(line)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Fprintf(out, "%+v\n", tok)
	}
}

func printAST(line string, out io.Writer) {
	program, ok := parseLine(line, out)
	if !ok {
		return
	}

	fmt.Fprintln(out, program.String())
}

func printEval(line string, out io.Writer, env *object.Environment) {
	program, ok := parseLine(line, out)
	if !ok {
		return
	}

	eval := evaluator.Eval(program, env)
	if eval != nil {
		fmt.Fprintln(out, eval.Inspect())
	}
}

func printCompile(line string, out io.Writer, symbolTable *compiler.SymbolTable, constants *[]object.Object) {
	program, ok := parseLine(line, out)
	if !ok {
		return
	}

	c := compiler.NewWithState(symbolTable, *constants)
	err := c.Compile(program)
	if err != nil {
		fmt.Fprintln(out, err)
		return
	}

	*constants = c.Bytecode().Constants
	fmt.Fprintln(out, c.Bytecode().Instructions)
}

func printRun(line string, out io.Writer, symbolTable *compiler.SymbolTable, constants *[]object.Object, globals []object.Object) {
	program, ok := parseLine(line, out)
	if !ok {
		return
	}

	c := compiler.NewWithState(symbolTable, *constants)
	err := c.Compile(program)
	if err != nil {
		fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
		return
	}

	*constants = c.Bytecode().Constants
	v := vm.NewWithGlobalsStore(c.Bytecode(), globals)
	err = v.Run()
	if err != nil {
		fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		return
	}

	result := v.LastPoppedStackElem()
	if result != nil {
		fmt.Fprintln(out, result.Inspect())
	}
}
