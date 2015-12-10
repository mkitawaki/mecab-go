package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "errors"
import "runtime"
import "sync"

// Lattice is the type defined in mecab_lattice_t
type Lattice struct {
	ptr      *_Ctype_struct_mecab_lattice_t
	isAlive  bool
	memorize *Memorize //map[string]*C.char
	mutex    *sync.Mutex
}

func latticeFinalizer(lattice *Lattice) {
	lattice.Destroy()
}

func newLattice(p *_Ctype_struct_mecab_lattice_t) (*Lattice, error) {
	if p == nil {
		return nil, ErrNotFoundLattice
	}
	lattice := &Lattice{ptr: p, isAlive: true, memorize: NewMemorize(), mutex: new(sync.Mutex)} //map[string](*C.char){}}
	runtime.SetFinalizer(lattice, latticeFinalizer)
	return lattice, nil
}

var (
	// ErrNotFoundLattice is the error that is returned when a lattice is not found
	ErrNotFoundLattice = errors.New("Lattice: lattice not found")
)

func (l *Lattice) toMecabLatticeT() *C.mecab_lattice_t {
	return (*C.mecab_lattice_t)(l.ptr)
}

// AddRequestType is a method to add a request type
func (l *Lattice) AddRequestType(requestType int) {
	C.mecab_lattice_add_request_type(l.toMecabLatticeT(), C.int(requestType))

}

// Clear is a method to clear all internal lattice data
func (l *Lattice) Clear() {
	C.mecab_lattice_clear(l.toMecabLatticeT())

}

// Destroy is a method to delete lattice object
func (l *Lattice) Destroy() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.isAlive {
		C.mecab_lattice_destroy(l.toMecabLatticeT())
		l.memorize.Clear()
		l.isAlive = false
	}
}

// GetAllBeginNodes is used internally
func (l *Lattice) GetAllBeginNodes() ([]*Node, error) {
	var (
		node   *Node
		length uint16
		err    error
	)
	if node, err = newNode(*(C.mecab_lattice_get_all_begin_nodes(l.toMecabLatticeT()))); err != nil {
		return nil, ErrNotFoundNode
	}
	if node, err = node.Bnext(); err != nil {
		return nil, ErrNotFoundNode
	}
	length = node.Length()
	beginNodes := make([]*Node, length)
	for i := range beginNodes {
		node, err = node.Next()
		if err != nil {
			return nil, ErrNotFoundNode
		}
		beginNodes[i] = node
	}
	return beginNodes, nil
}

// GetAllEndNodes is used internally
func (l *Lattice) GetAllEndNodes() ([]*Node, error) {
	var (
		node   *Node
		length uint16
		err    error
	)
	if node, err = newNode(*(C.mecab_lattice_get_all_end_nodes(l.toMecabLatticeT()))); err != nil {
		return nil, ErrNotFoundNode
	}
	if node, err = node.Bnext(); node == nil || err != nil {
		return nil, ErrNotFoundNode
	}
	length = node.Length()
	endNodes := make([]*Node, length)
	for i := range endNodes {
		node, err = node.Next()
		if err != nil {
			return nil, ErrNotFoundNode
		}
		endNodes[i] = node
	}
	return endNodes, nil
}

// GetBeginNodes is a method to return a node linked list starting at pos
func (l *Lattice) GetBeginNodes(pos uint64) (*Node, error) {
	return newNode(C.mecab_lattice_get_begin_nodes(l.toMecabLatticeT(), C.size_t(pos)))
}

// BosNode is a method to return a bos(begin of sentence) node
func (l *Lattice) BosNode() (*Node, error) {
	return newNode(C.mecab_lattice_get_bos_node(l.toMecabLatticeT()))
}

// BoundaryConstraint is a method to return the boundary constraint at the position
func (l *Lattice) BoundaryConstraint(pos uint64) int {
	return (int)(C.mecab_lattice_get_boundary_constraint(l.toMecabLatticeT(), C.size_t(pos)))
}

// GetEndNodes is used internally
func (l *Lattice) GetEndNodes(pos uint64) (*Node, error) {
	return newNode(C.mecab_lattice_get_end_nodes(l.toMecabLatticeT(), C.size_t(pos)))
}

// GetEosNode is a method to return a eos(end of sentence) node
func (l *Lattice) GetEosNode() (*Node, error) {
	return newNode(C.mecab_lattice_get_eos_node(l.toMecabLatticeT()))
}

// GetFeatureConstraint is a method to return the token constraint at the position
func (l *Lattice) GetFeatureConstraint(pos uint64) string {
	return C.GoString(C.mecab_lattice_get_feature_constraint(l.toMecabLatticeT(), C.size_t(pos)))
}

// GetRequestType is a method to return the current request type
func (l *Lattice) GetRequestType() int {
	return (int)(C.mecab_lattice_get_request_type(l.toMecabLatticeT()))
}

// GetSentence is a method to return sentence
func (l *Lattice) GetSentence() string {
	return C.GoString(C.mecab_lattice_get_sentence(l.toMecabLatticeT()))
}

// GetSize is a method to return sentence size
func (l *Lattice) GetSize() uint64 {
	return (uint64)(C.mecab_lattice_get_size(l.toMecabLatticeT()))
}

// GetTheta is a method to return temparature parameter theta
func (l *Lattice) GetTheta() float64 {
	return (float64)(C.mecab_lattice_get_theta(l.toMecabLatticeT()))
}

// GetZ is a method to return normalization factor of CRF
func (l *Lattice) GetZ() float64 {
	return (float64)(C.mecab_lattice_get_z(l.toMecabLatticeT()))
}

// HasConstraint is a method to return true if any parsing constraint is set
func (l *Lattice) HasConstraint() int {
	return (int)(C.mecab_lattice_has_constraint(l.toMecabLatticeT()))
}

// HasRequestType is a method to return true if the object has a specified request type
func (l *Lattice) HasRequestType(requestType int) int {
	return (int)(C.mecab_lattice_has_request_type(l.toMecabLatticeT(), (C.int)(requestType)))
}

// IsAvailable is a method to return true if result object is available
func (l *Lattice) IsAvailable() int {
	return (int)(C.mecab_lattice_is_available(l.toMecabLatticeT()))
}

// NbestTostr is a method to return string representation of the N-best results
// Returned object is managed by this instance. When clear/set_sentence() method is called, the returned buffer is initialized
func (l *Lattice) NbestTostr(n uint64) string {
	return (C.GoString)(C.mecab_lattice_nbest_tostr(l.toMecabLatticeT(), C.size_t(n)))
}

// NbestTostr2 is a method to return representation of the N-best result
// Result is saved in the specified
func (l *Lattice) NbestTostr2(n C.size_t, buf string, size C.size_t) string {
	b := (C.CString)(buf)
	l.memorize.Cache("NbestTostr2", b)
	// defer C.free(unsafe.Pointer(&b))
	return (C.GoString)(C.mecab_lattice_nbest_tostr2(l.toMecabLatticeT(), n, b, size))
}

// New is a function to return create new lattice object
func New() (*Lattice, error) {
	return newLattice(C.mecab_lattice_new())
}

// NewNode is a method to return new node
func (l *Lattice) NewNode() (*Node, error) {
	return newNode(C.mecab_lattice_new_node(l.toMecabLatticeT()))
}

// Next is a method to return obtain next-best result
// The internal linked list structure is updated. You should set MECAB_NBEST reques_type in advance. Return false if no more results are available or request_type is invalid
func (l *Lattice) Next() int {
	return int(C.mecab_lattice_next(l.toMecabLatticeT()))
}

// RemoveRequestType is a method to remove request type
func (l *Lattice) RemoveRequestType(requestType int) {
	C.mecab_lattice_remove_request_type(l.toMecabLatticeT(), (C.int)(requestType))

}

// SetBoundaryConstraint is a method to set parsing constraint for partial parsing mode
func (l *Lattice) SetBoundaryConstraint(pos uint64, boundaryType int) {
	C.mecab_lattice_set_boundary_constraint(l.toMecabLatticeT(), (C.size_t)(pos), C.int(boundaryType))

}

// SetFeatureConstraint is a method to set parsing constraint for partial parsing mode
func (l *Lattice) SetFeatureConstraint(beginPos uint64, endPos uint64, feature string) {
	f := C.CString(feature)
	l.memorize.Cache("SetFeatureConstraint", f)
	// defer C.free(unsafe.Pointer(&f))
	C.mecab_lattice_set_feature_constraint(l.toMecabLatticeT(), (C.size_t)(beginPos), (C.size_t)(endPos), f)

}

// SetRequestType is a method to set request type
func (l *Lattice) SetRequestType(requestType int) {
	C.mecab_lattice_set_request_type(l.toMecabLatticeT(), (C.int)(requestType))

}

// SetResult is a method to set golden parsing results for unittesting
func (l *Lattice) SetResult(result string) {
	r := C.CString(result)
	l.memorize.Cache("SetResult", r)
	//defer C.free(unsafe.Pointer(r))
	C.mecab_lattice_set_result(l.toMecabLatticeT(), r)

}

// SetSentence is a method to set sentence
func (l *Lattice) SetSentence(input string) {
	i := C.CString(input)
	l.memorize.Cache("SetSentence", i)
	// defer C.free(unsafe.Pointer(i))
	C.mecab_lattice_set_sentence(l.toMecabLatticeT(), i)

}

// SetSentence2 is a method to set sentence
// This method does not take the ownership of the object
func (l *Lattice) SetSentence2(sentence string, length uint64) {
	i := C.CString(sentence)
	l.memorize.Cache("SetSentence2", i)
	//defer C.free(unsafe.Pointer(i))
	C.mecab_lattice_set_sentence2(l.toMecabLatticeT(), i, (C.size_t)(length))

}

// SetTheta is a method to set temparature parameter theta
func (l *Lattice) SetTheta(theta float64) {
	C.mecab_lattice_set_theta(l.toMecabLatticeT(), (C.double)(theta))

}

// SetZ is a method to set normalization factor of CRF
func (l *Lattice) SetZ(z float64) {
	C.mecab_lattice_set_z(l.toMecabLatticeT(), (C.double)(z))

}

// Strerror is a method to return error string
func (l *Lattice) Strerror() string {
	return C.GoString(C.mecab_lattice_strerror(l.toMecabLatticeT()))
}

// ToStr is a method to return string representation of the lattice
// Returned object is managed by this instance. When clear/set_sentence() method is called, the returned buffer is initialized
func (l *Lattice) ToStr() string {
	return C.GoString(C.mecab_lattice_tostr(l.toMecabLatticeT()))
}

// Tostr2 is a method to return representation of the node
// Returned object is managed by this instance. When clear/set_sentence() method is called, the returned buffer is initialized
func (l *Lattice) Tostr2(buf string, size uint64) string {
	//defer C.free(unsafe.Pointer(i))
	i := C.CString(buf)
	l.memorize.Cache("Tostr2", i)
	return C.GoString(C.mecab_lattice_tostr2(l.toMecabLatticeT(), i, (C.size_t)(size)))
}
