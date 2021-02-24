package unionFind_1

type UF struct {
	N int
	elements []int
	id []int
}

func NewUF(N int) *UF {
	uf := new(UF)
	uf.N = N
	for i := 0; i < N; i++ {
		uf.elements = append(uf.elements, i)
		uf.id = append(uf.id, i)
	}

	return uf
}

func (o UF) Union(p int, q int)  {
	// 0 1 2 3 4 5 6 7 8
	// 1 1 2 2 3 3 4 4 4
	// 3 and 4
	id1 := o.id[p]
	id2 := o.id[q]
	for _, elem := range o.elements {
		id := o.id[elem]
		if id == id2 {
			o.id[elem] = id1
		}
	}
}

func (o UF) Connected(p int, q int) bool {
	id1 := o.id[p]
	id2 := o.id[q]

	return (id1> 0) && (id2 > 0) && (id1 == id2)
}