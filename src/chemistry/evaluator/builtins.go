package evaluator

import (
	"chemistry/balancer"
	"chemistry/object"
	"fmt"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.String{Value: strconv.Itoa(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NONE
		},
	},
	"balance": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				result, err := balancer.Balance(arg.Inspect())
				if err == nil {
					return &object.String{Value: result}
				}
				return newError("argument to `balance` not supported, got error `%s`",
					err)
			default:
				return newError("argument to `balance` not supported, got %s",
					args[0].Type())
			}
		},
	},
}
