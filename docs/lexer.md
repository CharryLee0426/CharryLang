# Lexer

## UML
![lexer](/Users/lichen/Desktop/CSKnowledge/Charrylang/docs/diagram/out/Lexer.png)

## Types

It contains only one type called `Lexer`, which has four attributes.

| Attribute    | Description                      |
| ------------ | -------------------------------- |
| input        | Our program in plain text format |
| position     | current position                 |
| readPosition | next position                    |
| ch           | current character                |



> why do we need position and readPosition?
>
> The readPosition attribute, which points to the next position forever, is very important to peek the content after current character.

## Methods

### 1. Foundation Methods

`readChar()`: read next character into `l.ch`, then put forward `position` and `readPosition` pointers. When it comes to the end of input, it will set `l.ch` to zero, which means `EOF` in ASCII.

`readIdentifier()`: it will read letters to create a string until the current character isn't a letter, then return the string. It is one of our core foundations. It will help us to get literals and analyse tokens in our programs.

`readNumber()`: like `readIdentifier()`, it will read numbers and return a string that represents numbers it reads.

### 2. Justifying Methods

`peekChar()`: see the next character, equivalent to the character at `readPosition`.

> It is an assert function that many mainstream interpreters have such function. It is helpful to handle some double-char operators such as `==` and `!=`.

### 3. Public Methods

`NextToken() token.Token`: get next token in our programs. It contains a switch blocks in order to make sure what the type is.

`New(input string) *Lexer`: create a lexer based on our input, which is our programs in string format. It is the constructer of Lexer in other way.
