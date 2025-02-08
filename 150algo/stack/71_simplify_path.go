package stack

import "strings"

func simplifyPath(path string) string {

	dirs := strings.Split(path, "/")

	var stack []string
	var skipNext int

	for i := len(dirs) - 1; i >= 0; i-- {
		if dirs[i] == ".." {
			skipNext++
			continue
		}

		if dirs[i] == "" || dirs[i] == "." {
			continue
		}

		if skipNext > 0 {
			skipNext--
			continue
		}

		stack = append(stack, dirs[i])
	}

	if len(stack) == 0 {
		return "/"
	}

	var result string
	for i := len(stack) - 1; i >= 0; i-- {
		result = result + "/" + stack[i]
	}
	return result
}
