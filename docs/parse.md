# Parse

## Structure
### 1. Public Methods

### 2. Private Methods

### 3. Tool Methods

### 4. Assist Constants

## Overview
The core of Charrylang is this parser! We can get tokens from lexer in sequence. However, it is meaningless if tokens cannot be organized as a good structure. The parser's job is this! It is a top-down precedence parser that can make our tokens into a well-organized abstract syntax tree(AST). After we get this AST, we can handle each node through the tree and get our result!

## Concepts

### 1. How to implement top-down precedence parsing?

### 2. How to know each expression's precedence?