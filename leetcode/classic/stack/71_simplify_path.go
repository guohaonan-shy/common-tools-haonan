package stack

import "strings"

func simplifyPath(path string) string {
	name := ""
	stack := make([]string, 0)
	for i := range path {
		if path[i] == '/' {
			if name != "" {
				if name == ".." {
					if len(stack) > 0 {
						stack = stack[:len(stack)-1]
					}
				} else if name == "." {
				} else {
					stack = append(stack, name)
				}
				name = ""
			} else {
				continue
			}
		} else {
			name += string(path[i])
		}
	}

	if name != "" {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name == "." {
		} else {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}
