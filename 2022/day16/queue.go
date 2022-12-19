package main

type Item struct {
	valve    string
	distance int
	index    int
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p)
	item := x.(*Item)
	item.index = n
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*p = old[0 : n-1]
	return item
}
