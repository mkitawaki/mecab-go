package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "errors"
import "unsafe"

// Lattice is the type defined in mecab_lattice_t
type Lattice _Ctype_struct_mecab_lattice_t

var (
	// ErrNotFoundNode is the error that is returned when a node is not found
	ErrNotFoundNode = errors.New("Lattice: node not found")
)

func (l *Lattice) unalias() *C.mecab_lattice_t {
	return (*C.mecab_lattice_t)(l)
}

// AddRequestType is a method to add a request type
func (l *Lattice) AddRequestType(requestType int) {
	C.mecab_lattice_add_request_type(l.unalias(), C.int(requestType))
}

// Clear is a method to clear all internal lattice data
func (l *Lattice) Clear() {
	C.mecab_lattice_clear(l.unalias())
}

// Destroy is a method to delete lattice object
func (l *Lattice) Destroy() {
	C.mecab_lattice_destroy(l.unalias())
}

// GetAllBeginNodes is used internally
func (l *Lattice) GetAllBeginNodes() ([]*Node, error) {
	var node *Node
	if node = (*Node)(*(C.mecab_lattice_get_all_begin_nodes(l.unalias()))); node == nil {
		return nil, ErrNotFoundNode
	}
	if node = node.Bnext(); node == nil {
		return nil, ErrNotFoundNode
	}
	beginNodes := make([]*Node, node.Length())
	for i := range beginNodes {
		node = node.Next()
		beginNodes[i] = node
	}
	return beginNodes, nil
}

// GetAllEndNodes is used internally
func (l *Lattice) GetAllEndNodes() ([]*Node, error) {
	var node *Node
	if node = (*Node)(*(C.mecab_lattice_get_all_end_nodes(l.unalias()))); node == nil {
		return nil, ErrNotFoundNode
	}
	if node = node.Bnext(); node == nil {
		return nil, ErrNotFoundNode
	}
	endNodes := make([]*Node, node.Length())
	for i := range endNodes {
		node = node.Next()
		endNodes[i] = node
	}
	return endNodes, nil
}

// GetBeginNodes is a method to return a node linked list starting at pos
func (l *Lattice) GetBeginNodes(pos uint64) *Node {
	return (*Node)(C.mecab_lattice_get_begin_nodes(l.unalias(), C.size_t(pos)))
}

// BosNode is a method to return a bos(begin of sentence) node
func (l *Lattice) BosNode() *Node {
	return (*Node)(C.mecab_lattice_get_bos_node(l.unalias()))
}

// BoundaryConstraint is a method to return the boundary constraint at the position
func (l *Lattice) BoundaryConstraint(pos uint64) int {
	return (int)(C.mecab_lattice_get_boundary_constraint(l.unalias(), C.size_t(pos)))
}

// GetEndNodes is used internally
func (l *Lattice) GetEndNodes(pos uint64) *Node {
	return (*Node)(C.mecab_lattice_get_end_nodes(l.unalias(), C.size_t(pos)))
}

// GetEosNode is a method to return a eos(end of sentence) node
func (l *Lattice) GetEosNode() *Node {
	return (*Node)(C.mecab_lattice_get_eos_node(l.unalias()))
}

// GetFeatureConstraint is a method to return the token constraint at the position
func (l *Lattice) GetFeatureConstraint(pos uint64) string {
	return C.GoString(C.mecab_lattice_get_feature_constraint(l.unalias(), C.size_t(pos)))
}

// GetRequestType is a method to return the current request type
func (l *Lattice) GetRequestType() int {
	return (int)(C.mecab_lattice_get_request_type(l.unalias()))
}

// GetSentence is a method to return sentence
func (l *Lattice) GetSentence() string {
	return C.GoString(C.mecab_lattice_get_sentence(l.unalias()))
}

// GetSize is a method to return sentence size
func (l *Lattice) GetSize() uint64 {
	return (uint64)(C.mecab_lattice_get_size(l.unalias()))
}

// GetTheta is a method to return temparature parameter theta
func (l *Lattice) GetTheta() float64 {
	return (float64)(C.mecab_lattice_get_theta(l.unalias()))
}

// GetZ is a method to return normalization factor of CRF
func (l *Lattice) GetZ() float64 {
	return (float64)(C.mecab_lattice_get_z(l.unalias()))
}

// HasConstraint is a method to return true if any parsing constraint is set
func (l *Lattice) HasConstraint() int {
	return (int)(C.mecab_lattice_has_constraint(l.unalias()))
}

// HasRequestType is a method to return true if the object has a specified request type
func (l *Lattice) HasRequestType(requestType int) int {
	return (int)(C.mecab_lattice_has_request_type(l.unalias(), (C.int)(requestType)))
}

// IsAvailable is a method to return true if result object is available
func (l *Lattice) IsAvailable() int {
	return (int)(C.mecab_lattice_is_available(l.unalias()))
}

// NbestTostr is a method to return string representation of the N-best results
// Returned object is managed by this instance. When clear/set_sentence() method is called, the returned buffer is initialized
func (l *Lattice) NbestTostr(n C.size_t) string {
	return (C.GoString)(C.mecab_lattice_nbest_tostr(l.unalias(), n))
}

// NbestTostr2 is a method to return representation of the N-best result
// Result is saved in the specified
func (l *Lattice) NbestTostr2(n C.size_t, buf string, size C.size_t) string {
	b := (C.CString)(buf)
	defer C.free(unsafe.Pointer(&b))
	return (C.GoString)(C.mecab_lattice_nbest_tostr2(l.unalias(), n, b, size))
}

// New is a function to return create new lattice object
func New() *Lattice {
	return (*Lattice)(C.mecab_lattice_new())
}

// NewNode is a method to return new node
func (l *Lattice) NewNode() *Node {
	return (*Node)(C.mecab_lattice_new_node(l.unalias()))
}

// Next is a method to return obtain next-best result
// The internal linked list structure is updated. You should set MECAB_NBEST reques_type in advance. Return false if no more results are available or request_type is invalid
func (l *Lattice) Next() int {
	return (int)(C.mecab_lattice_next(l.unalias()))
}

// RemoveRequestType is a method to remove request type
func (l *Lattice) RemoveRequestType(requestType int) {
	C.mecab_lattice_remove_request_type(l.unalias(), (C.int)(requestType))
}

// SetBoundaryConstraint is a method to set parsing constraint for partial parsing mode
func (l *Lattice) SetBoundaryConstraint(pos uint64, boundaryType int) {
	C.mecab_lattice_set_boundary_constraint(l.unalias(), (C.size_t)(pos), C.int(boundaryType))
}

// SetFeatureConstraint is a method to set parsing constraint for partial parsing mode
func (l *Lattice) SetFeatureConstraint(beginPos uint64, endPos uint64, feature string) {
	f := C.CString(feature)
	defer C.free(unsafe.Pointer(&f))
	C.mecab_lattice_set_feature_constraint(l.unalias(), (C.size_t)(beginPos), (C.size_t)(endPos), f)
}

// SetRequestType is a method to set request type
func (l *Lattice) SetRequestType(requestType int) {
	C.mecab_lattice_set_request_type(l.unalias(), (C.int)(requestType))
}

// SetResult is a method to set golden parsing results for unittesting
func (l *Lattice) SetResult(result string) {
	r := C.CString(result)
	defer C.free(unsafe.Pointer(r))
	C.mecab_lattice_set_result(l.unalias(), r)
}

// SetSentence is a method to set sentence
func (l *Lattice) SetSentence(input string) {
	i := C.CString(input)
	defer C.free(unsafe.Pointer(i))
	C.mecab_lattice_set_sentence(l.unalias(), i)
}

// SetSentence2 is a method to set sentence
// This method does not take the ownership of the object
func (l *Lattice) SetSentence2(sentence string, length uint64) {
	i := C.CString(sentence)
	defer C.free(unsafe.Pointer(i))
	C.mecab_lattice_set_sentence2(l.unalias(), i, (C.size_t)(length))
}

// SetTheta is a method to set temparature parameter theta
func (l *Lattice) SetTheta(theta float64) {
	C.mecab_lattice_set_theta(l.unalias(), (C.double)(theta))
}

// SetZ is a method to set normalization factor of CRF
func (l *Lattice) SetZ(z float64) {
	C.mecab_lattice_set_z(l.unalias(), (C.double)(z))
}

// Strerror is a method to return error string
func (l *Lattice) Strerror() string {
	return C.GoString(C.mecab_lattice_strerror(l.unalias()))
}

// ToStr is a method to return string representation of the lattice
// Returned object is managed by this instance. When clear/set_sentence() method is called, the returned buffer is initialized
func (l *Lattice) ToStr() string {
	return C.GoString(C.mecab_lattice_tostr(l.unalias()))
}

// Tostr2 is a method to return representation of the node
// Returned object is managed by this instance. When clear/set_sentence() method is called, the returned buffer is initialized
func (l *Lattice) Tostr2(buf string, size uint64) string {
	i := C.CString(buf)
	defer C.free(unsafe.Pointer(i))
	return C.GoString(C.mecab_lattice_tostr2(l.unalias(), i, (C.size_t)(size)))
}
