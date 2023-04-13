# Parse

## Structure

### 0. Members

```go
type Parser struct {
	l *lexer.Lexer		// a pointer to lexer instance

	curToken token.Token	// current token
	peekToken token.Token	// next token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
	errors []string
}
```

### 1. Public Methods

| Method Name                   | Description                                                  |
| ----------------------------- | ------------------------------------------------------------ |
| `New(l *lexer.Lexer) *Parser` | Create a parser and register prefix parsing functions and infix parsing functions |
| `Errors() []string`           | Return errors in parser when it parsing a program.           |
| `ParseProgram() *ast.Program` | Parse the sequence of tokens into an AST(Abstract Syntax Tree) |

### 2. Private Methods

### 3. Tool Methods

| Method Name                                                  | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `nextToken()`                                                | Update current token of parser to its peek token, and its peek token to the peek token of its peek token |
| `registerPrefix(tokenType token.TokenType, fn prefixParseFn)` | Register the prefix parse function to the map of prefix parse functions |
| `registerInfix(tokenType token.TokenType, fn infixParseFn)`  | Register the infix parse function to the map of infix parse functions |
| `curTokenIs(t token.TokenType) bool`                         | Return if the type of current token is token type t          |
| `peekTokenIs(t token.TokenType) bool`                        | Return if the type of peek token is token type t             |
| `expectedPeek(t token.TokenType) bool`                       | An assert function to handle the peek token                  |
| `curPrecedence() int`                                        | Return the number of precedence of current operator          |
| `peekPrecedence() int`                                       | Return the number of precedence of peek operator             |
| `peekError(t token.TokenType)`                               | Add error message if peek token's type is not what it should be |
| `noPrefixParseFnError(t token.TokenType)`                    | Add error message if there is no prefix parse function for token type t is found |

### 4. Assist Constants

#### 4.1 Precedence

```go
const (
	_ int = iota
	LOWEST
	EQUALS	// ==
	LESSGREATER	// > or <
	SUM	// + or -
	PRODUCT	// * or /
	PREFIX	// -X or !X
	CALL	// myFunction(X)
	INDEX
)
```

#### 4.2 Precedence Map

```go
var precedences = map[token.TokenType]int {
	token.EQ: EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT: LESSGREATER,
	token.GT: LESSGREATER,
	token.PLUS: SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN: CALL,
	token.LBRACKET: INDEX,
}
```

## Overview
The core of Charrylang is this parser! We can get tokens from lexer in sequence. However, it is meaningless if tokens cannot be organized as a good structure. The parser's job is this! It is a top-down precedence parser that can make our tokens into a well-organized abstract syntax tree(AST). After we get this AST, we can handle each node through the tree and get our result!

## Concepts

### 1. How to implement top-down precedence parsing?

### 2. How to know each expression's precedence?