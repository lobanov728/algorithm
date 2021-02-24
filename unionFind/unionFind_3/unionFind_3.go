package unionFind_3

type UF struct {
	N int
	elements []int
	id []int
	sz []int
}

func NewUF(N int) *UF {
	uf := new(UF)
	uf.N = N
	for i := 0; i < N; i++ {
		uf.elements = append(uf.elements, i)
		uf.id = append(uf.id, i)
		uf.sz = append(uf.sz, 1)
	}

	return uf
}

func (o UF) root(i int) int {
	// 0 1 2 3 4 5 6 7 8 9
	// 0 1 2 3 3 4 6 7 8 9
	// 3 and 4
	// 4 and 5
	for ; o.elements[i] != o.id[i]; {
		i = o.id[i]
	}
	return i
}

func (o UF) Union(p int, q int)  {
	// 0 1 2 3 4 5 6 7 8 9
	// 0 1 2 3 4 5 6 7 8 9
	// 3 and 4
	o.id[q] = o.root(p)
}

func (o UF) Connected(p int, q int) bool {
	root1 := o.root(p)
	root2 := o.root(q)

	return root1 == root2
}