package command

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// TODO(yoshi): map of func name to func or map of struct name to struct

func ParseCommands(args []string) error {
	if len(args) == 1 {
		arg := args[0]

		switch {
		case strings.Contains(arg, "db"):
			s := strings.Split(arg, ":")
			fmt.Println(s)
			// TODO(yoshi): handle "db:foo:bar"
			if len(s) == 2 {
				st := reflect.TypeOf(DB{})
				_, ok := st.MethodByName(s[1])
				// If method exists
				if ok {
					reflect.ValueOf(DB{}).MethodByName(s[1]).Call([]reflect.Value{})
				}
			}

		case strings.Contains(arg, "debug"):
			fmt.Println("Arg contains debug!")
		default:
			fmt.Println("Default brrr")
		}

		os.Exit(1)

	} else if len(args) > 1 {
		fmt.Println("Too many arguments!")
		os.Exit(1)
	}

	// TODO(yoshi): fix
	return nil
}
