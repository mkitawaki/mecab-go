package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
import "C"

// Node is the type defined in mecab_node_t
type Node _Ctype_struct_mecab_node_t

func (l *Node) unalias() *C.mecab_node_t {
	return (*C.mecab_node_t)(l)
}

// Prev is a method to return a pointer to the previous node
func (l *Node) Prev() *Node {
	return (*Node)(l.unalias().prev)
}

// Next is a method to return a pointer to the next node
func (l *Node) Next() *Node {
	return (*Node)(l.unalias().next)
}

// Enext is a method to return a pointer to the node which ends at the same position
func (l *Node) Enext() *Node {
	return (*Node)(l.unalias().enext)
}

// Bnext is a method to return a pointer to the node which starts at the same position
func (l *Node) Bnext() *Node {
	return (*Node)(l.unalias().bnext)
}

// Rpath is a method to return a pointer to the right path
func (l *Node) Rpath() *Path {
	return (*Path)(l.unalias().rpath)
}

// Lpath is a method to return a pointer to the right path
func (l *Node) Lpath() *Path {
	return (*Path)(l.unalias().lpath)
}

// Surface is a method to return a surface string
func (l *Node) Surface() string {
	return C.GoString(l.unalias().surface)
}

// Feature is a method to return a feature string
func (l *Node) Feature() string {
	return C.GoString(l.unalias().feature)
}

// Id is a method to return a unique node id
func (l *Node) Id() uint {
	return (uint)(l.unalias().id)
}

// Length is a method to return a length of the surface form
func (l *Node) Length() uint16 {
	return (uint16)(l.unalias().length)
}

// Rlength is a method to return a length of the surface form including white space before the morph
func (l *Node) Rlength() uint16 {
	return (uint16)(l.unalias().rlength)
}

// RcAttr is a method to return a right attribute id
func (l *Node) RcAttr() uint16 {
	return (uint16)(l.unalias().rcAttr)
}

// LcAttr is a method to return a left attribute id
func (l *Node) LcAttr() uint16 {
	return (uint16)(l.unalias().lcAttr)
}

// Posid is a method to return a unique part of speech id
func (l *Node) Posid() uint16 {
	return (uint16)(l.unalias().posid)
}

// CharType is a method to return a character type
func (l *Node) CharType() byte {
	return (byte)(l.unalias().char_type)
}

// Stat is a method to return a status of this model
func (l *Node) Stat() byte {
	return (byte)(l.unalias().stat)
}

// Isbest is a method to return a set 1 if this node is best node
func (l *Node) Isbest() byte {
	return (byte)(l.unalias().isbest)
}

// Alpha is a method to return a forward accumulative log summation
func (l *Node) Alpha() float32 {
	return (float32)(l.unalias().alpha)
}

// Beta is a method to return a backward accumulative log summation
func (l *Node) Beta() float32 {
	return (float32)(l.unalias().beta)
}

// Prob is a method to return a marginal probability
func (l *Node) Prob() float32 {
	return (float32)(l.unalias().prob)
}

// Wcost is a method to return a word cost
func (l *Node) Wcost() int16 {
	return (int16)(l.unalias().wcost)
}

// Cost is a method to return a best accumulative cost from bos node to this node
func (l *Node) Cost() int64 {
	return (int64)(l.unalias().cost)
}
