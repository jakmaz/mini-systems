# mini-lang

A minimal programming language interpreter written in Go.

## Features

- C-like syntax
- Variable bindings
- Integers and booleans
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- String data structure
- Array data structure
- Hash data structure

## Assumptions

- Only ASCII characters are supported
- Only integers as numbers (no float, hex etc)

## Usage

```bash
go run main.go
go test ./...
```

## Implementation

### Parser

Mini-lang implements recursive descent parser (Pratt Parsing)

## Todo

- Add line number to error messages

## Acknowledgements

- [Writng an Interpreter in Go](https://interpreterbook.com/)
