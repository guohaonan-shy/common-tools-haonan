package graph

type Genetics struct {
	Str string
	Cnt int
}

func minMutation(startGene string, endGene string, bank []string) int {
	bank = append(bank, startGene)

	graph := make(map[string][]string, len(bank))

	init := func() {
		for i := range bank {
			gene := bank[i]
			for j := 0; j < len(bank); j++ {
				if i == j {
					continue
				}

				if diffOnlyOne(gene, bank[j]) {
					if graph[gene] == nil {
						graph[gene] = []string{bank[j]}
					} else {
						graph[gene] = append(graph[gene], bank[j])
					}
				}
			}
		}
		return
	}

	init()

	queue := []*Genetics{
		{
			Str: startGene,
			Cnt: 0,
		},
	}
	reached := map[string]bool{startGene: true}

	for len(queue) > 0 {
		gene := queue[0]
		queue = queue[1:]

		for _, next := range graph[gene.Str] {

			if next == endGene {
				return gene.Cnt + 1
			}

			if !reached[next] {
				reached[next] = true
				queue = append(queue, &Genetics{Str: next, Cnt: gene.Cnt + 1})
			}
		}
	}
	return -1
}

func diffOnlyOne(a, b string) bool {
	cnt := 0
	for i := range a {
		if a[i] != b[i] {
			cnt++
		}

		if cnt > 1 {
			return false
		}
	}

	return cnt == 1
}

func minMutationV2(startGene string, endGene string, bank []string) int {
	// because the startGene is valid, we need to add it into bank and analysis the connectivity
	bank = append(bank, startGene)
	var (
		initGraph   func()
		adjacentMap = make(map[string][]string, 0)
	)
	initGraph = func() {

		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				if i == j {
					continue
				}

				if !diffOnlyOneV2(bank[i], bank[j]) {
					continue
				}

				// forward
				if _, ok := adjacentMap[bank[i]]; ok {
					adjacentMap[bank[i]] = append(adjacentMap[bank[i]], bank[j])
				} else {
					adjacentMap[bank[i]] = []string{bank[j]}
				}

				// reverse
				if _, ok := adjacentMap[bank[j]]; ok {
					adjacentMap[bank[j]] = append(adjacentMap[bank[j]], bank[i])
				} else {
					adjacentMap[bank[j]] = []string{bank[i]}
				}
			}
		}
		return
	}

	initGraph()

	queue := make([]*Genetics, 0)
	queue = append(queue, &Genetics{
		Str: startGene,
		Cnt: 0,
	})
	reachedMap := make(map[string]bool, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, next := range adjacentMap[cur.Str] {

			if next == endGene {
				return cur.Cnt + 1
			}

			if !reachedMap[next] {
				reachedMap[next] = true
				queue = append(queue, &Genetics{
					Str: next,
					Cnt: cur.Cnt + 1,
				})
			}

		}
	}
	return -1
}

func diffOnlyOneV2(a, b string) bool {
	diff := 0
	for i := 0; i < 8; i++ {
		if a[i] != b[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}
