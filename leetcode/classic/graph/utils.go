package graph

func BuildGraph(str string) [][]byte {
	str = str[1 : len(str)-1]

	graph := make([][]byte, 0)

	start := false
	subStr := []byte{}
	for i := range str {
		if str[i] == '[' {
			start = true
			continue
		}

		if str[i] == ']' {
			start = false
			graph = append(graph, subStr)
			subStr = []byte{}
			continue
		}

		if start == false {
			continue
		}

		if (str[i] >= '0' && str[i] <= '9') || (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			subStr = append(subStr, str[i])
		}
	}
	return graph
}
