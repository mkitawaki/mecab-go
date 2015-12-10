package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
import "C"
import "errors"

// Path is the type defined in mecab_path_t
type Path struct {
	ptr *_Ctype_struct_mecab_path_t
}

var (
	// ErrNotFoundPath is the error that is returned when a path is not found
	ErrNotFoundPath = errors.New("Path: path not found")
)

func newPath(p *_Ctype_struct_mecab_path_t) (*Path, error) {
	if p == nil {
		return nil, ErrNotFoundPath
	}
	return &Path{ptr: p}, nil
}

func (l *Path) toMecabPathT() *C.mecab_path_t {
	return (*C.mecab_path_t)(l.ptr)
}

// Rnode is a method to return a pointer to the right node
func (l *Path) Rnode() (*Node, error) {
	return newNode(l.toMecabPathT().rnode)
}

// Rnext is a method to return a pointer to the next right path
func (l *Path) Rnext() (*Path, error) {
	return newPath(l.toMecabPathT().rnext)
}

// Lnode is a method to return a pointer to the left node
func (l *Path) Lnode() (*Node, error) {
	return newNode(l.toMecabPathT().lnode)
}

// Lnext is a method to return a pointer to the next left path
func (l *Path) Lnext() (*Path, error) {
	return newPath(l.toMecabPathT().lnext)
}

// Cost is a method to return a local cost
func (l *Path) Cost() int {
	return (int)(l.toMecabPathT().cost)
}

// Prob is a method to return a marginal probability
func (l *Path) Prob() float32 {
	return (float32)(l.toMecabPathT().prob)
}
