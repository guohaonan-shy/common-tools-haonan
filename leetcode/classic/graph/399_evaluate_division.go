package graph

type Neibor struct {
	Div string
	Val float64
}

var Graph = map[string][]*Neibor{}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	func() {
		for i, equation := range equations {
			neiborRaw0 := &Neibor{
				Div: equation[1],
				Val: values[i],
			}
			if _, ok := Graph[equation[0]]; ok {
				Graph[equation[0]] = append(Graph[equation[0]], neiborRaw0)
			} else {
				Graph[equation[0]] = []*Neibor{neiborRaw0}
			}

			neiborRaw1 := &Neibor{
				Div: equation[0],
				Val: 1 / values[i],
			}
			if _, ok := Graph[equation[1]]; ok {
				Graph[equation[1]] = append(Graph[equation[1]], neiborRaw1)
			} else {
				Graph[equation[1]] = []*Neibor{neiborRaw1}
			}
		}
	}()

	// query
	res := make([]float64, 0)

	for _, query := range queries {

		reached := make(map[string]bool, 0)
		var dfs func(string, string) float64
		dfs = func(a, b string) float64 {

			var (
				ok1, ok2 bool
				neibors  = []*Neibor{}
			)

			neibors, ok1 = Graph[a]
			_, ok2 = Graph[b]

			if !ok1 || !ok2 {
				return -1
			}

			reached[a] = true
			for _, neibor := range neibors {
				if neibor.Div == b {
					return neibor.Val
				}
			}

			for _, neibor := range neibors {
				if !reached[neibor.Div] {
					init := neibor.Val
					val := dfs(neibor.Div, b)
					if val != -1 {
						return init * val
					}
				}
			}

			return -1
		}
		re := dfs(query[0], query[1])
		res = append(res, re)
	}
	return res
}

//	func simplify(a, b string) (string, string) {
//		idxs := []int{}
//		for j := range a {
//			target := a[j]
//			idx := strings.IndexByte(b, target)
//			if idx == -1 {
//				continue
//			}
//
//			idxs = append(idxs, j)
//			b = b[:idx] + b[idx+1:]
//		}
//
//		for i := len(idxs) - 1; i >= 0; i-- {
//			a = a[:idxs[i]] + a[idxs[i]+1:]
//		}
//
//		return a, b
//	}
type CommonStruct struct {
	start string
	end   string
	value float64
}

func calcEquationV2(equations [][]string, values []float64, queries [][]string) []float64 {
	adjacenyMap := make(map[string][]*CommonStruct, 0)
	statusMap := make(map[[2]string]bool, 0)

	for i, equation := range equations {
		start, end := equation[0], equation[1]

		if _, ok := adjacenyMap[start]; ok {
			adjacenyMap[start] = append(adjacenyMap[start], &CommonStruct{
				start: start,
				end:   end,
				value: values[i],
			})
		} else {
			adjacenyMap[start] = []*CommonStruct{
				{
					start: start,
					end:   end,
					value: values[i],
				},
			}
		}

		statusMap[[2]string{start, end}] = false

		// know the a / b, we can also know b / a
		if _, ok := adjacenyMap[end]; ok {
			adjacenyMap[end] = append(adjacenyMap[end], &CommonStruct{
				start: end,
				end:   start,
				value: 1.0 / values[i],
			})
		} else {
			adjacenyMap[end] = []*CommonStruct{
				{
					start: end,
					end:   start,
					value: 1.0 / values[i],
				},
			}
		}

		statusMap[[2]string{end, start}] = false

	}

	var dfs func(cur *CommonStruct, target string, pre float64) float64
	dfs = func(cur *CommonStruct, target string, pre float64) float64 {

		if cur.end == target {
			return pre * cur.value
		}

		statusMap[[2]string{cur.start, cur.end}] = true
		val := -1.0
		for _, next := range adjacenyMap[cur.end] {

			if statusMap[[2]string{next.start, next.end}] {
				continue
			}
			val = dfs(next, target, pre*cur.value)
			if val != -1.0 {
				break
			}
		}
		statusMap[[2]string{cur.start, cur.end}] = false
		return val
	}

	res := make([]float64, 0, len(queries))
	for _, query := range queries {

		start, end := query[0], query[1]

		// corner case: like 'aa' vs 'aa'
		val := -1.0
		adjacenyList, ok := adjacenyMap[start]
		_, ok2 := adjacenyMap[end]

		if !ok || !ok2 {
			res = append(res, val)
			continue
		}

		if start == end && ok {
			val = 1.0
		} else {
			for _, next := range adjacenyList {
				val = dfs(next, end, 1.0)
				if val != -1.0 {
					break
				}
			}
		}
		res = append(res, val)
	}
	return res
}
