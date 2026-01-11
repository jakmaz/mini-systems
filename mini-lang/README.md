# mini-lang

A minimal programming language interpreter written in Go, exploring parsing, evaluation, and language design from first principles.

## Language Features

- C-like syntax with variable bindings
- Data types: integers, booleans, strings, arrays, hash maps
- Arithmetic and boolean expressions
- First-class and higher-order functions with closures
- Built-in functions

## Sample Programs

```javascript
// Variable bindings and arithmetic
let x = 5;
let y = x * 2 + 1;

// Functions and closures
let add = fn(a, b) { a + b };
let makeAdder = fn(x) { fn(y) { x + y } };
let addTwo = makeAdder(2);
addTwo(3); // returns 5

// Arrays and hash maps
let arr = [1, 2, 3, 4];
let person = {"name": "John", "age": 30};
```

## Architecture

**Lexer** → **Parser** → **AST** → **Evaluator**

- **Lexer**: Tokenizes source code
- **Parser**: Recursive descent parser using Pratt parsing for expressions
- **AST**: Abstract syntax tree representation
- **Evaluator**: Tree-walking interpreter

## Usage

```bash
# Run REPL
go run main.go

# Run tests
go test ./...
```

## Limitations

- ASCII characters only
- Integers only (no floats or hex)
- No error line numbers yet

## Related Projects

This is part of the [mini-systems](../) collection - small implementations of foundational software systems built for learning and exploration.
