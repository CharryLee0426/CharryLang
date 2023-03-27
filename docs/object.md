# Object

## Structure

### 1. Defined Type

| Type Name  | Type   | Description                                |
| ---------- | ------ | ------------------------------------------ |
| ObjectType | string | just a mark that represents an object type |

### 2. Interface

| Interface | Method  | Description                     |
| --------- | ------- | ------------------------------- |
| Object    | Type    | return its object type          |
|           | Inspect | return its detailed information |

### 3. Object

| Object      | Member/Method                  | Description                                                  |
| ----------- | ------------------------------ | ------------------------------------------------------------ |
| Integer     | Value int 64                   | An integer like -1, 0, 2, ...                                |
| Boolean     | Value bool                     | Boolean, only true or false                                  |
| Null        | N/A                            | Null, its nil in host language                               |
| ReturnType  | Value Object                   | A return value                                               |
| Error       | Message string                 | An error                                                     |
| Environment | store map\[string\]\[Object\]  | An environment, used for binding variables                   |
|             | outer *Environment             | store is a map which is useful for variable binding and context understanding; outer is a pointer to its father environment |
| Function    | Parameters \[\]*ast.Identifier | A function                                                   |
|             | Body *ast.BlockStatement       |                                                              |
|             | Env *Environment               | a function has an environment.                               |

### 4. Assist Constants

Assist constants in `object.go` are used to find and generate object type infomation easily.

| Assistant constant | Description          |
| ------------------ | -------------------- |
| INTEGER_OBJ        | An integer object    |
| BOOLEAN_OBJ        | A boolean object     |
| NULL_OBJ           | A null object        |
| RETURN_VALUE_OBJ   | A return type object |
| ERROR_OBJ          | An error object      |
| FUNCTION_OBJ       | A function object    |

## Overview
In `object.go`, we defined many objects we used during the process of evaluating. For short, object contains real values in abstract syntax tree nodes. Our host programming language, Go, can handle objects and follow the guildence of parser.

## Important Concepts

### 1. Environment

Environment is our solution to **variable binding**. We need variable binding because our program should keep our variables defined before and reuses them in order to complete our tasks. For example, if we defines three variables a, b, and c in our program. Then we should use them after we defined them like this.

```
>>> let a = 1;
>>> let b = 2;
>>> let c = 3;
>>> a
>>> 1
```

So how can we let our programs remember variables? It's pretty easy! We can use a map to store our variables. For performance, we can use a hash map to store our variables so that we can retrive our variables. In Charrylang, we use environment to store our variables.

When we create a problem, we'll create an environment at the initialization stage. This environment is called the **root environment**. After that, we will use it as our evaluation argument in order to put variables into our environment. For the above program, our environment will be like this:

```
Environment
|
|--store
		|
		|- "a" : &Integer{Value: 1}
		|- "b" : &Integer{Value: 2}
		|- "c" : &Integer{Value: 3}
```

When the fourth statment is executed, it will retrive the object by using the identifier a as the key.

### 2. Function

The syntax of function in Charrylang is defined like this:

```
fn (param1, param2, ...) {...}
```

It can be assigned to an identifier in Charrylang. In other words, funciton in Charrylang is a type of variable, like Javascript and some other modern programming languages. It is also the foundation of closure feature. For example, this program is correct in Charrylang.

```
>>> let as = fn(x, y) {return x+y;}
>>> as(2, 3)
5
>>>
```

The really important problem for evaluating function is that if we copying parameters by values or by references? For most programming languages, they choose the first implementation for the reason that it can prevent from unnecessary modifying values. So does charrylang. So how we can implement it?

The answer is **enclosed environment**. When we define a function, we add the caller's environment to function literal. We call it `outer` environment, or father environment. When the program calls this function, the function creates another environment to store all parameters, temp variables, and result variables. When the function gets the result and return it to the caller, the program goes back to the outer environment.

The logic of environment and outer environment is like this:

```
Root Environment
|
|
|---function defined outer : root
|
|
|---calling function
    |
    |--- inner environment for the function created
    |--- evaluating the function
    |--- return the result
----| goes back to the outer environment
|
|--- continue...
```

