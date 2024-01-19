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
