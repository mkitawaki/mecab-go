package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
import "C"
import "errors"

// Node is the type defined in mecab_node_t
type Node struct {
	ptr *_Ctype_struct_mecab_node_t
}

var (
	// ErrNotFoundNode is the error that is returned when a node is not found
	ErrNotFoundNode = errors.New("Node: node not found")
)

func newNode(p *_Ctype_struct_mecab_node_t) (*Node, error) {
	if p == nil {
		return nil, ErrNotFoundNode
	}
	return &Node{ptr: p}, nil
}

func (l *Node) toMecabNodeT() *C.mecab_node_t {
	return (*C.mecab_node_t)(l.ptr)
}

// Prev is a method to return a pointer to the previous node
func (l *Node) Prev() (*Node, error) {
	return newNode(l.toMecabNodeT().prev)
}

// Next is a method to return a pointer to the next node
func (l *Node) Next() (*Node, error) {
	return newNode(l.toMecabNodeT().next)
}

// Enext is a method to return a pointer to the node which ends at the same position
func (l *Node) Enext() (*Node, error) {
	return newNode(l.toMecabNodeT().enext)
}

// Bnext is a method to return a pointer to the node which starts at the same position
func (l *Node) Bnext() (*Node, error) {
	return newNode(l.toMecabNodeT().bnext)
}

// Rpath is a method to return a pointer to the right path
func (l *Node) Rpath() (*Path, error) {
	return newPath(l.toMecabNodeT().rpath)
}

// Lpath is a method to return a pointer to the right path
func (l *Node) Lpath() (*Path, error) {
	return newPath(l.toMecabNodeT().lpath)
}

// Surface is a method to return a surface string
func (l *Node) Surface() string {
	return C.GoString(l.toMecabNodeT().surface)
}

// Feature is a method to return a feature string
func (l *Node) Feature() string {
	return C.GoString(l.toMecabNodeT().feature)
}

// Id is a method to return a unique node id
func (l *Node) Id() uint {
	return (uint)(l.toMecabNodeT().id)
}

// Length is a method to return a length of the surface form
func (l *Node) Length() uint16 {
	return (uint16)(l.toMecabNodeT().length)
}

// Rlength is a method to return a length of the surface form including white space before the morph
func (l *Node) Rlength() uint16 {
	return (uint16)(l.toMecabNodeT().rlength)
}

// RcAttr is a method to return a right attribute id
func (l *Node) RcAttr() uint16 {
	return (uint16)(l.toMecabNodeT().rcAttr)
}

// LcAttr is a method to return a left attribute id
func (l *Node) LcAttr() uint16 {
	return (uint16)(l.toMecabNodeT().lcAttr)
}

// Posid is a method to return a unique part of speech id
func (l *Node) Posid() uint16 {
	return (uint16)(l.toMecabNodeT().posid)
}

// CharType is a method to return a character type
func (l *Node) CharType() byte {
	return (byte)(l.toMecabNodeT().char_type)
}

// Stat is a method to return a status of this model
func (l *Node) Stat() byte {
	return (byte)(l.toMecabNodeT().stat)
}

// Isbest is a method to return a set 1 if this node is best node
func (l *Node) Isbest() byte {
	return (byte)(l.toMecabNodeT().isbest)
}

// Alpha is a method to return a forward accumulative log summation
func (l *Node) Alpha() float32 {
	return (float32)(l.toMecabNodeT().alpha)
}

// Beta is a method to return a backward accumulative log summation
func (l *Node) Beta() float32 {
	return (float32)(l.toMecabNodeT().beta)
}

// Prob is a method to return a marginal probability
func (l *Node) Prob() float32 {
	return (float32)(l.toMecabNodeT().prob)
}

// Wcost is a method to return a word cost
func (l *Node) Wcost() int16 {
	return (int16)(l.toMecabNodeT().wcost)
}

// Cost is a method to return a best accumulative cost from bos node to this node
func (l *Node) Cost() int64 {
	return (int64)(l.toMecabNodeT().cost)
}
