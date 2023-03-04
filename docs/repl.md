# REPL

## What is REPL

REPL stands for "Read Eval Print Loop". Many mainstream programming languages such as Python, Ruby, and JavaScript, have their own REPLs. So does our charrylang. IMO, REPL is the same thing as "console", or "interactive mode".

## The architecture of REPL

The implementation of our REPL is very easy. It just has one function called `Start(in io.Reader, out io.Writer)`. We will use it to read inputs from `os.Stdin`, and write outputs to `os.Stdout`. Simply speaking, it reads from our terminal, then prints our results to our terminal, our screen!