package utils

import (
	"fmt"
	"strings"
)

func ResolveArgs(args []string) map[string]string {
	m := make(map[string]string)
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) < 2 {
			panic(fmt.Sprintf("Invalid argument (format: <key>=<value>): %s", arg))
		}
		m[parts[0]] = parts[1]
	}
	return m
}
