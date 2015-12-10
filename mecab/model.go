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

// Model is the type defined in mecab_model_t
type Model struct {
	ptr      *_Ctype_struct_mecab_model_t
	isAlive  bool
	memorize *Memorize
	mutex    *sync.Mutex
}

var (
	// ErrNotFoundNode is the error that is returned when a node is not found
	ErrNotFoundModel = errors.New("Model: model not found")
)

func modelFinalizer(model *Model) {
	model.mutex.Lock()
	defer model.mutex.Unlock()
	if model.isAlive {
		model.Destroy()
	}
}

func newModel(p *_Ctype_struct_mecab_model_t) (*Model, error) {
	if p == nil {
		return nil, ErrNotFoundModel
	}
	model := &Model{ptr: p, isAlive: true, memorize: NewMemorize(), mutex: new(sync.Mutex)}
	runtime.SetFinalizer(model, modelFinalizer)
	return model, nil
}

// NewModel1 is a factory method to create a new Model with a specified main's argc/argv-style parameters
// Return NULL if new model cannot be initialized. Use MeCab::getLastError() to obtain the cause of the errors
func NewModel1(argc int32, argv string) (*Model, error) {
	v := (**C.char)(unsafe.Pointer(&argv))
	//defer C.free(unsafe.Pointer(v))
	model, err := newModel(C.mecab_model_new(C.int(argc), v))
	if err != nil {
		return nil, err
	}
	model.memorize.Cache("NewModel1", *v)
	return model, nil
}

// NewModel2 is a factory method to create new Model with a string parameter representation, i.e., "-d /user/local/mecab/dic/ipadic -Ochasen"
// Return NULL if new model cannot be initialized. Use MeCab::getLastError() to obtain the cause of the errors
func NewModel2(arg string) (*Model, error) {
	a := C.CString(arg)
	//	defer C.free(unsafe.Pointer(a))
	model, err := newModel(C.mecab_model_new2(a))
	if err != nil {
		return nil, err
	}
	model.memorize.Cache("NewModel2", a)
	return model, nil
}

// CreateLattice is a method to create a new Lattice object
func (m *Model) CreateLattice() (*Lattice, error) {
	return newLattice(C.mecab_model_new_lattice(m.toMecabModelT()))
}

// CreateTagger is a method to create a new Tagger object
// All returned tagger object shares this model object as a parsing model. Never delete this model object before deleting tagger object
func (m *Model) CreateTagger() (*Tagger, error) {
	return newTagger(C.mecab_model_new_tagger(m.toMecabModelT()))
}

// Destroy is a method to delete model object
func (m *Model) Destroy() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	C.mecab_model_destroy(m.toMecabModelT())
	m.memorize.Clear()
	m.isAlive = false
}

func (m *Model) toMecabModelT() *C.mecab_model_t {
	return (*C.mecab_model_t)(m.ptr)
}
