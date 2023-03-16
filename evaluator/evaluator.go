package evaluator

import (
	"charrylang/ast"
	"charrylang/object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}

	return nil
}

func evalStatements(stmt []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmt {
		result = Eval(statement)
	}

	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
}

// !false -> true; !true -> false
func evalBangOperatorExpression(right object.Object) object.Object {
	// right is switched by real value but not reference
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// switch realRight := right.(type) {
// case *object.Boolean:
// 	if realRight.Value == true {
// 		return FALSE
// 	} else {
// 		return TRUE
// 	}
// case *object.Integer:
// 	if realRight.Value == 0 {
// 		return TRUE
// 	} else {
// 		return FALSE
// 	}
// case *object.Null:
// 	return TRUE
// default:
// 	return FALSE
// }