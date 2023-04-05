# REPL

## What is REPL

REPL stands for "Read Eval Print Loop". Many mainstream programming languages such as Python, Ruby, and JavaScript, have their own REPLs. So does our charrylang. IMO, REPL is the same thing as "console", or "interactive mode". At first, we create a scanner to receive input from our keyboard and an environment whose data type is `*Object.Environment` as the root environment for this REPL instance. Second, we create lexer, parser and evaluator to parse and evaluate our input. Finally, we print the result to our terminal. In order to make our REPL more like a console, I create an infinite loop for getting any loop except 'exit', which I use for exiting the REPL. 'exit' is not necessary condition to jump out the infinite loop. We can use `ctrl/cmd+C` to termnite the REPL because REPL is in fact only a go rountine. However, it is not elegent. Nobody likes seeing red 'X' every time they forcely terminate a process.

## The architecture of REPL

The implementation of our REPL is very easy. It just has one function called `Start(in io.Reader, out io.Writer)`. We will use it to read inputs from `os.Stdin`, and write outputs to `os.Stdout`. Simply speaking, it reads from our terminal, then prints our results to our terminal, our screen!