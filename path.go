package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
import "C"

// Path is the type defined in mecab_path_t
type Path _Ctype_struct_mecab_path_t

func (l *Path) unalias() *C.mecab_path_t {
	return (*C.mecab_path_t)(l)
}

// Rnode is a method to return a pointer to the right node
func (l *Path) Rnode() *Node {
	return (*Node)(l.unalias().rnode)
}

// Rnext is a method to return a pointer to the next right path
func (l *Path) Rnext() *Path {
	return (*Path)(l.unalias().rnext)
}

// Lnode is a method to return a pointer to the left node
func (l *Path) Lnode() *Node {
	return (*Node)(l.unalias().lnode)
}

// Lnext is a method to return a pointer to the next left path
func (l *Path) Lnext() *Path {
	return (*Path)(l.unalias().lnext)
}

// Cost is a method to return a local cost
func (l *Path) Cost() int {
	return (int)(l.unalias().cost)
}

// Prob is a method to return a marginal probability
func (l *Path) Prob() float32 {
	return (float32)(l.unalias().prob)
}
