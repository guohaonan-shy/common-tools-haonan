package graph

import "strings"

type Neibor struct {
	Div string
	Val float64
}

var Graph = map[string][]*Neibor{}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	func() {
		for i, equation := range equations {
			// 原版
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
				Val: float64(1) / values[i],
			}
			if _, ok := Graph[equation[1]]; ok {
				Graph[equation[1]] = append(Graph[equation[1]], neiborRaw1)
			} else {
				Graph[equation[1]] = []*Neibor{neiborRaw1}
			}
			// 最简因式
			var1, var2 := simplify(equation[0], equation[1])
			neibor0 := &Neibor{
				Div: var2,
				Val: values[i],
			}
			if _, ok := Graph[var1]; ok {
				Graph[var1] = append(Graph[var1], neibor0)
			} else {
				Graph[var1] = []*Neibor{neibor0}
			}

			neibor1 := &Neibor{
				Div: var1,
				Val: float64(1) / values[i],
			}
			if _, ok := Graph[var2]; ok {
				Graph[var2] = append(Graph[var2], neibor1)
			} else {
				Graph[var2] = []*Neibor{neibor1}
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
					temp := neibor.Val
					temp *= dfs(neibor.Div, b)
					if temp != -1 {
						return temp
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

func simplify(a, b string) (string, string) {
	idxs := []int{}
	for j := range a {
		target := a[j]
		idx := strings.IndexByte(b, target)
		if idx == -1 {
			continue
		}

		idxs = append(idxs, j)
		b = b[:idx] + b[idx+1:]
	}

	for i := len(idxs) - 1; i >= 0; i-- {
		a = a[:idxs[i]] + a[idxs[i]+1:]
	}

	return a, b
}
