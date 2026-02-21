# mini-lang

A minimal programming language interpreter written in Go, exploring parsing, evaluation, and language design from first principles.

## Language Features

- **Integers** - Whole number arithmetic
- **Booleans** - `true` and `false` with logical operators
- **Strings** - Double-quoted text
- **Arrays** - Ordered collections with index access
- **Hashes** - Key-value maps with string keys
- **Prefix, Infix, and Index Operators** - `-5`, `a + b`, `arr[0]`, `hash["key"]`
- **Conditionals** - `if/else` expressions
- **Global and Local Bindings** - `let` for variable declaration
- **First-Class Functions** - Functions as values, higher-order functions
- **Return Statements** - Explicit returns from functions
- **Closures** - Functions that capture their environment

## Sample Programs

### Variables and Data Types

```javascript
let recipe = "Pancakes";
let ingredients = ["flour", "milk", "eggs", "butter"];
let pancake = {
  "servings": 4,
  "cookTime": 15,
  "difficulty": "easy"
};
```

### Functions and Closures

```javascript
let cook = fn(recipe) {
    let name = recipe["name"];
    let time = recipe["cookTime"];
    puts("Cooking " + name + " for " + time + " minutes");
};

cook(pancake);
// => prints: "Cooking Pancakes for 15 minutes"
```

### Recursion

```javascript
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
```

### Higher-Order Functions

```javascript
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };

  iter(arr, []);
};

let numbers = [1, 1 + 1, 4 - 1, 2 * 2, 2 + 3, 12 / 2];
map(numbers, fibonacci);
// => returns: [1, 1, 2, 3, 5, 8]
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
```

The REPL supports debugging modes: type `/tokens` to see lexer output, `/parse` to see the AST, or `/eval` to evaluate code.

```bash
# Run tests
go test ./...
```

## Limitations

- ASCII characters only
- Integers only (no floats or hex)
- No error line numbers yet

## Related Projects

This is part of the [mini-systems](../) collection - small implementations of foundational software systems built for learning and exploration.

## Resources

- [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball
- [Pratt Parsing](https://en.wikipedia.org/wiki/Pratt_parser) by Wikipedia
