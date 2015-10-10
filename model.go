package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// Model is the type defined in mecab_model_t
type Model _Ctype_struct_mecab_model_t

// NewModel is a factory method to create a new Model with a specified main's argc/argv-style parameters
// Return NULL if new model cannot be initialized. Use MeCab::getLastError() to obtain the cause of the errors
func NewModel(argc int32, argv string) *Model {
	v := (**C.char)(unsafe.Pointer(&argv))
	defer C.free(unsafe.Pointer(v))
	return (*Model)(C.mecab_model_new(C.int(argc), v))
}

// NewModel2 is a factory method to create new Model with a string parameter representation, i.e., "-d /user/local/mecab/dic/ipadic -Ochasen"
// Return NULL if new model cannot be initialized. Use MeCab::getLastError() to obtain the cause of the errors
func NewModel2(arg string) *Model {
	a := C.CString(arg)
	defer C.free(unsafe.Pointer(a))
	return (*Model)(C.mecab_model_new2(a))
}

// CreateLattice is a method to create a new Lattice object
func (m *Model) CreateLattice() *Lattice {
	return (*Lattice)(C.mecab_model_new_lattice(m.unalias()))
}

// CreateTagger is a method to create a new Tagger object
// All returned tagger object shares this model object as a parsing model. Never delete this model object before deleting tagger object
func (m *Model) CreateTagger() *Tagger {
	return (*Tagger)(C.mecab_model_new_tagger(m.unalias()))
}

func (m *Model) unalias() *C.mecab_model_t {
	return (*C.mecab_model_t)(m)
}
