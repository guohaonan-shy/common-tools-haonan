package every_day

import "math"

func wallsAndGates(rooms [][]int) {
	for i, row := range rooms {
		for j := range row {
			if row[j] == math.MaxInt32 {
				bfsRooms(i, j, rooms)
			}
		}
	}
}

type bfsNode struct {
	row, column int
	depth       int
}

var bfsDirections = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func bfsRooms(i, j int, rooms [][]int) {
	queue := make([]*bfsNode, 0)
	queue = append(queue, &bfsNode{
		row:    i,
		column: j,
		depth:  0,
	})
	minimum := math.MaxInt32
	reached := make(map[[2]int]bool, 0) // check whether available nodes is in the queue but still not check
	reached[[2]int{i, j}] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, direction := range bfsDirections {
			nextRow := cur.row + direction[0]
			nextCol := cur.column + direction[1]

			if nextRow < 0 || nextRow >= len(rooms) || nextCol < 0 || nextCol >= len(rooms[0]) {
				continue
			}

			if reached[[2]int{nextRow, nextCol}] {
				continue
			}

			if rooms[nextRow][nextCol] == -1 {
				continue
			}

			if rooms[nextRow][nextCol] == 0 {
				minimum = min(minimum, cur.depth+1)
				continue
			}

			//if rooms[nextRow][nextCol] != math.MaxInt32 && cur.depth+rooms[nextRow][nextCol] > minimum {
			//	continue
			//}

			/*
				if some node has completed the iteration, it means the value of this position is not inf
				we don't continue to iterate this type of node, we can just calculate the minimum by cur.depth + 1 + rooms[nextRow][nextCol]
				cut branch firstly can be beneficial for time complexity
			*/
			if rooms[nextRow][nextCol] != math.MaxInt32 {
				minimum = min(minimum, cur.depth+rooms[nextRow][nextCol]+1)
				continue
			}
			queue = append(queue, &bfsNode{
				row:    nextRow,
				column: nextCol,
				depth:  cur.depth + 1,
			})
			reached[[2]int{nextRow, nextCol}] = true
		}
	}
	rooms[i][j] = minimum
}

/*
	A general thought about BFS:
	1. confirm the start point of bfs
	2. get the head of node in queue as the current node
	3. iterate four directions:
		3.1. check if the four direction of current node is valid or not: out of range => filter
		3.2. check if this node is reached or has already added in to queue
		3.3. specified conditions in practice => filter
		3.4. target
		3.5. if need to further search, add next node into queue and tag it as reached
		3.6. go to step 2
*/
