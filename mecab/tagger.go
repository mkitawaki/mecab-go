package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"
import "errors"
import "runtime"
import "sync"

// Tagger is the type defined in mecab_t
type Tagger struct {
	ptr      *_Ctype_struct_mecab_t
	memorize *Memorize
	isAlive  bool
	mutex    *sync.Mutex
}

var (
	// ErrNotFoundTagger is the error that is returned when a tagger is not found
	ErrNotFoundTagger = errors.New("Tagger: tagger not found")
)

func taggerFinalizer(tagger *Tagger) {
	tagger.mutex.Lock()
	defer tagger.mutex.Unlock()
	if tagger.isAlive {
		tagger.Destroy()
	}
}

func newTagger(p *_Ctype_struct_mecab_t) (*Tagger, error) {
	if p == nil {
		return nil, ErrNotFoundTagger
	}
	tagger := &Tagger{ptr: p, memorize: NewMemorize(), isAlive: true, mutex: new(sync.Mutex)}
	runtime.SetFinalizer(tagger, taggerFinalizer)
	return tagger, nil
}

// NewTagger is a factory method to create a new Tagger with a specified main's argc/argv-style parameters
func NewTagger(argc int32, argv string) (*Tagger, error) {
	v := (**C.char)(unsafe.Pointer(&argv))
	// defer C.free(unsafe.Pointer(v))
	tagger, err := newTagger(C.mecab_new(C.int(argc), v))
	if err != nil {
		return nil, err
	}
	tagger.memorize.Cache("NewTagger", *v)
	return tagger, nil
}

// NewTagger2 is a factory method to create new Tagger with a string parameter representation, i.e., "-d /user/local/mecab/dic/ipadic -Ochasen"
func NewTagger2(arg string) (*Tagger, error) {
	a := C.CString(arg)
	// defer C.free(unsafe.Pointer(a))
	tagger, err := newTagger(C.mecab_new2(a))
	if err != nil {
		return nil, err
	}
	tagger.memorize.Cache("NewTagger2", a)
	return tagger, nil
}

func (t *Tagger) toMecabT() *C.mecab_t {
	return (*C.mecab_t)(t.ptr)
}

// Version is a method to return a version string
func (t *Tagger) Version() string {
	return C.GoString(C.mecab_version())
}

// StrError is a method to return last error string
func (t *Tagger) StrError() string {
	return C.GoString(C.mecab_strerror(t.toMecabT()))
}

// Destroy is a method to delete Tagger object
// This method calles "delete tagger". In some environment, e.g., MS-Windows, an object allocated inside a DLL must be deleted in the same DLL too.
func (t *Tagger) Destroy() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	C.mecab_destroy(t.toMecabT())
	t.memorize.Clear()
	t.isAlive = false
}

// GetPartial is a method to set partial parsing mode
// This method is DEPRECATED. Use Lattice::add_request_type(MECAB_PARTIAL) or Lattice::remove_request_type(MECAB_PARTIAL)
func (t *Tagger) GetPartial() int {
	return int(C.mecab_get_partial(t.toMecabT()))
}

// SetPartial is a method to set partial parsing mode
// This method is DEPRECATED. Use Lattice::has_request_type(MECAB_PARTIAL).
func (t *Tagger) SetPartial(partial int) {
	C.mecab_set_partial(t.toMecabT(), C.int(partial))

}

// GetTheta is a method to set temparature parameter theta.
func (t *Tagger) GetTheta() float32 {
	return (float32)(C.mecab_get_theta(t.toMecabT()))
}

// SetTheta is a method to return temparature parameter theta.
func (t *Tagger) SetTheta(theta float32) {
	C.mecab_set_theta(t.toMecabT(), C.float(theta))

}

// GetLatticeLevel is a method to return lattice level.
// This method is DEPRECATED. Use Lattice::*_request_type()
func (t *Tagger) GetLatticeLevel() int {
	return (int)(C.mecab_get_lattice_level(t.toMecabT()))
}

// SetLatticeLevel is a method to set lattice level.
// This method is DEPRECATED. Use Lattice::*_request_type()
func (t *Tagger) SetLatticeLevel(level int) {
	C.mecab_set_lattice_level(t.toMecabT(), C.int(level))

}

// GetAllMorphs is a method to return true if all morphs output mode is on.
// This method is DEPRECATED. Use Lattice::has_request_type(MECAB_ALL_MORPHS).
func (t *Tagger) GetAllMorphs() int {
	return (int)(C.mecab_get_all_morphs(t.toMecabT()))
}

// SetAllMorphs is a method to set all-morphs output mode.
// This method is DEPRECATED. Use Lattice::add_request_type(MECAB_ALL_MORPHS) or Lattice::remove_request_type(MECAB_ALL_MORPHS)
func (t *Tagger) SetAllMorphs(allMorphs int) {
	C.mecab_set_all_morphs(t.toMecabT(), C.int(allMorphs))

}

// ParseLattice is a method to parse lattice object.
// Return true if lattice is parsed successfully. A sentence must be set to the lattice with Lattice:set_sentence object before calling this method. Parsed node object can be obtained with Lattice:bos_node. This method is thread safe.
func (t *Tagger) ParseLattice(lattice *Lattice) int {
	l := lattice.toMecabLatticeT()
	return (int)(C.mecab_parse_lattice(t.toMecabT(), l))
}

// SparseTostr is a method to parse given sentence and return parsed result as string.
// You should not delete the returned string. The returned buffer is overwritten when parse method is called again. This method is NOT thread safe.
func (t *Tagger) SparseTostr(str string) string {
	s := C.CString(str)
	t.memorize.Cache("SparseTostr", s)
	// defer C.free(unsafe.Pointer(s))
	return (C.GoString)(C.mecab_sparse_tostr(t.toMecabT(), s))
}

// SparseTostr2 is a method to the same as parse() method, but input length can be passed.
func (t *Tagger) SparseTostr2(str string, l C.size_t) string {
	s := C.CString(str)
	t.memorize.Cache("SparseTostr2", s)
	// defer C.free(unsafe.Pointer(s))
	return (C.GoString)(C.mecab_sparse_tostr2(t.toMecabT(), s, l))
}

// SparseTostr3 is a method to the same as parse() method, but input length and output buffer are passed.
// Return parsed result as string. The result pointer is the same as |ostr|. Return NULL, if parsed result string cannot be stored within |olen| bytes.
func (t *Tagger) SparseTostr3(str string, l C.size_t, ostr string, olen C.size_t) string {
	s := C.CString(str)
	t.memorize.Cache("SparseTostr3a", s)
	// defer C.free(unsafe.Pointer(s))
	o := C.CString(ostr)
	t.memorize.Cache("SparseTostr3b", o)
	//defer C.free(unsafe.Pointer(o))
	return (C.GoString)(C.mecab_sparse_tostr3(t.toMecabT(), s, l, o, olen))
}

// SparseTonode is a method to parse given sentence and return Node object.
// You should not delete the returned node object. The returned buffer is overwritten when parse method is called again. You can traverse all nodes via Node::next member. This method is NOT thread safe.
func (t *Tagger) SparseTonode(str string) (*Node, error) {
	s := C.CString(str)
	t.memorize.Cache("SparseTonode", s)
	// defer C.free(unsafe.Pointer(s))
	return newNode(C.mecab_sparse_tonode(t.toMecabT(), s))
}

// SparseTonode2 is a method to the same as parseToNode(), but input lenth can be passed.
func (t *Tagger) SparseTonode2(str string, l C.size_t) (*Node, error) {
	s := C.CString(str)
	t.memorize.Cache("SparseTonode2", s)
	// defer C.free(unsafe.Pointer(s))
	return newNode(C.mecab_sparse_tonode2(t.toMecabT(), s, l))
}

// NbestSparseTostr is a method to parse given sentence and obtain N-best results as a string format.
// Currently, N must be 1 <= N <= 512 due to the limitation of the buffer size. You should not delete the returned string. The returned buffer is overwritten when parse method is called again. This method is DEPRECATED. Use Lattice class.
func (t *Tagger) NbestSparseTostr(n C.size_t, str string) string {
	s := C.CString(str)
	t.memorize.Cache("NbestSparseTostr", s)
	// defer C.free(unsafe.Pointer(s))
	return (C.GoString)(C.mecab_nbest_sparse_tostr(t.toMecabT(), n, s))
}

// NbestSparseTostr2 is a method to the same as parseNBest(), but input length can be passed.
func (t *Tagger) NbestSparseTostr2(n C.size_t, str string, l C.size_t) string {
	s := C.CString(str)
	t.memorize.Cache("NbestSparseTostr2", s)
	// defer C.free(unsafe.Pointer(s))
	return (C.GoString)(C.mecab_nbest_sparse_tostr2(t.toMecabT(), n, s, l))
}

// NbestSparseTostr3 is a method to the same as parseNBest(), but input length and output buffer can be passed.
// Return NULL if more than |olen| buffer is required to store output string.
func (t *Tagger) NbestSparseTostr3(n C.size_t, str string, l C.size_t, ostr string, olen C.size_t) string {
	s := C.CString(str)
	t.memorize.Cache("NbestSparseTostr3a", s)
	// defer C.free(unsafe.Pointer(s))
	o := C.CString(ostr)
	t.memorize.Cache("NbestSparseTostr3b", o)
	// defer C.free(unsafe.Pointer(o))
	return (C.GoString)(C.mecab_nbest_sparse_tostr3(t.toMecabT(), n, s, l, o, olen))
}

// NbestInit is a method to initialize N-best enumeration with a sentence.
// Return true if initialization finishes successfully. N-best result is obtained by calling next() or nextNode() in sequence. This method is NOT thread safe. This method is DEPRECATED. Use Lattice class.
func (t *Tagger) NbestInit(str string) int {
	s := C.CString(str)
	t.memorize.Cache("NbestInit", s)
	// defer C.free(unsafe.Pointer(s))
	return (int)(C.mecab_nbest_init(t.toMecabT(), s))
}

// NbestInit2 is a method to the same as parseNBestInit(), but input length can be passed.
func (t *Tagger) NbestInit2(str string, l C.size_t) int {
	s := C.CString(str)
	t.memorize.Cache("NbestInit2", s)
	// defer C.free(unsafe.Pointer(s))
	return (int)(C.mecab_nbest_init2(t.toMecabT(), s, l))
}

// NbestNextTostr is a method to return next-best parsed result.
// You must call parseNBestInit() in advance. Return NULL if no more reuslt is available. This method is NOT thread safe. This method is DEPRECATED. Use Lattice class.
func (t *Tagger) NbestNextTostr() string {
	return (C.GoString)(C.mecab_nbest_next_tostr(t.toMecabT()))
}

// NbestNextTostr2 is a method to return next-best parsed result.
// You must call parseNBestInit() in advance. Return NULL if no more reuslt is available. This method is NOT thread safe. This method is DEPRECATED. Use Lattice class
func (t *Tagger) NbestNextTostr2(ostr string, olen C.size_t) string {
	o := C.CString(ostr)
	t.memorize.Cache("NbestNextTostr2", o)
	// defer C.free(unsafe.Pointer(o))
	return (C.GoString)(C.mecab_nbest_next_tostr2(t.toMecabT(), o, olen))
}

// NbestNextTonode is a method to return formatted node object.
// The format is specified with --unk-format, --bos-format, --eos-format, and --eon-format respectively. You should not delete the returned string. The returned buffer is overwritten when parse method is called again. This method is NOT thread safe. This method is DEPRECATED. Use Lattice class.
func (t *Tagger) NbestNextTonode() (*Node, error) {
	return newNode(C.mecab_nbest_next_tonode(t.toMecabT()))
}

// FormatNode is a method to return formatted node object.
// The format is specified with --unk-format, --bos-format, --eos-format, and --eon-format respectively. You should not delete the returned string. The returned buffer is overwritten when parse method is called again. This method is NOT thread safe. This method is DEPRECATED. Use Lattice class.
func (t *Tagger) FormatNode(node *Node) string {
	return (C.GoString)(C.mecab_format_node(t.toMecabT(), node.toMecabNodeT()))
}

// DictionaryInfo is a method to return DictionaryInfo linked list.
func (t *Tagger) DictionaryInfo() (*DictionaryInfo, error) {
	return newDictionaryInfo(C.mecab_dictionary_info(t.toMecabT()))
}

func (t *Tagger) parse(l *Lattice) int {
	return int(C.mecab_parse_lattice(t.toMecabT(), l.toMecabLatticeT()))
}
