package graph

import "container/list"

type vertex struct {
	word     string
	edges    []*vertex
	distance int
}

var graph = make(map[string]*vertex)

// WordLadder returns the minimum number of transformations from start to end in a dictionary
// where words are all equal in length, and a transformation can only happen if the difference
// between two words is only in one letter. Zero is returned if no such transformations can occur.
func WordLadder(start, end string, dic []string) int {
	graph = make(map[string]*vertex)

	graph[start] = &vertex{word: start, edges: nil, distance: 0}
	for _, w := range dic {
		graph[w] = &vertex{word: w}
	}

	for w1 := range graph {
		for _, w2 := range dic {
			if isDifferentByOneLetter(w1, w2) {
				graph[w1].edges = append(graph[w1].edges, graph[w2])
			}
		}
	}
	return bfsMinTransformation(start, end)
}

func bfsMinTransformation(beginWord string, endWord string) int {
	min := 0
	distance := 0
	source := graph[beginWord]
	seen := make(map[*vertex]struct{})
	queue := list.New()

	queue.PushBack(source)
	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*vertex)
		queue.Remove(queue.Front())
		distance++
		for _, v := range tmp.edges {
			if _, ok := seen[v]; ok {
				continue
			}
			v.distance = distance
			seen[v] = struct{}{}
			if v.word == endWord && (v.distance < min || min == 0) {
				min = v.distance
			}
			queue.PushBack(v)
		}
	}
	return min
}

func isDifferentByOneLetter(a, b string) bool {
	oneDiff := false
	for i := range a {
		if a[i] != b[i] {
			if oneDiff {
				return false
			}
			oneDiff = true
		}
	}
	return oneDiff
}
