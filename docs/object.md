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

| Object      | Member | Description                                |
| ----------- | ------ | ------------------------------------------ |
| Integer     |        | An integer like -1, 0, 2, ...              |
| Boolean     |        | Boolean, only true or false                |
| Null        |        | Null, its nil in host language             |
| ReturnType  |        | A return value                             |
| Error       |        | An error                                   |
| Environment |        | An environment, used for binding variables |
| Function    |        | A function                                 |

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

