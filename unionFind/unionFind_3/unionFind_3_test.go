package unionFind_3

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T)  {
	uf := NewUF(13)
	uf.Union(6, 8)
	uf.Union(10, 8)
	uf.Union(8, 0)

	uf.Union(11,12)
	uf.Union(11,9)
	uf.Union(0, 9)

	fmt.Println(uf.Connected(12,10))

}
