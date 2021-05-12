# Monkey - Go

This repo contains my annotations of [Writing An Interpreter In Go](https://interpreterbook.com) book!

## 1. Lexing

### 1.1 Lexical Analysis

To transform an amount of text (source code) in something meaningful to our language
we need to transform it, the transformations are described below:

`Source Code` -> `Tokens` -> `Abstract Syntax Tree (AST)`

The first transformation (Source Code to Token) is called _Lexical Analysis_ (lexing),
that job is done by a __lexer__. We can have some differences inside the lexer category: `Tokenizer` and `Scanner`

Tokens are small and categorizable data structs! See the example below:
```text
let x = 5 + 5;
```

The lexer output can be:
```text
[LET, IDENTIFIER("X"), EQUAL_SIGN, INTEGER(5), PLUS_SIGN, INTEGER(5), SEMICOLON]
```

All the tokens are a representation of the original source code, and they can be different
depending on the implementation!

> A well implemented lexer might store useful information, like, line number, column number,
> and file name.
> It can store for many reasons and one of them is to output a more useful error message!

The second transformation (Tokens to AST) is done by the __parser__.

### 1.2 Defining Our Tokens

The first thing to do is to define the tokens of our language! We need to think
in an initial structure we want our code to be and by them think/create the tokens!

The token structure is defined [token/token.go](token/token.go).

### 1.3 The Lexer

The Lexer can be found on [lexer/lexer.go](lexer/lexer.go)

TODO: Supports Unicode characters!

### 1.4 Extending our Token Set and Lexer

Just extending the implementation!

### 1.5 Start of a REPL

The REPL can be found on [repl/repl.go](repl/repl.go).
