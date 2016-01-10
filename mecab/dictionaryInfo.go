package mecab

// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "errors"

// DictionaryInfo is the type defined in mecab_t
type DictionaryInfo struct {
	ptr *_Ctype_struct_mecab_dictionary_info_t
}

var (
	// ErrNotFoundNode is the error that is returned when a node is not found
	ErrNotFoundDictionaryInfo = errors.New("DictionaryInfo: dictionaryInfo not found")
)

func newDictionaryInfo(p *_Ctype_struct_mecab_dictionary_info_t) (*DictionaryInfo, error) {
	if p == nil {
		return nil, ErrNotFoundDictionaryInfo
	}
	return &DictionaryInfo{ptr: p}, nil
}

func (d *DictionaryInfo) toMecabDictionaryInfoT() *C.mecab_dictionary_info_t {
	return (*C.mecab_dictionary_info_t)(d.ptr)
}

// Filename is a method to return filename of dictionary On Windows, filename is stored in UTF-8 encoding
func (d *DictionaryInfo) Filename() string {
	return (C.GoString)(d.toMecabDictionaryInfoT().filename)
}

// Charset is a method to return character set of the dictionary
func (d *DictionaryInfo) Charset() string {
	return (C.GoString)(d.toMecabDictionaryInfoT().charset)
}

// Size is a method to return how many words are registered in this dictionary
func (d *DictionaryInfo) Size() uint {
	return (uint)(d.toMecabDictionaryInfoT().size)
}

// Type is a method to return dictionary type this value should be MECAB_USR_DIC, MECAB_SYS_DIC, or MECAB_UNK_DIC
func (d *DictionaryInfo) Type() int {
	return (int)(d.toMecabDictionaryInfoT()._type)
}

// Lsize is a method to return left attributes size
func (d *DictionaryInfo) Lsize() uint {
	return (uint)(d.toMecabDictionaryInfoT().lsize)
}

// Rsize is a method to return right attributes size
func (d *DictionaryInfo) Rsize() uint {
	return (uint)(d.toMecabDictionaryInfoT().rsize)
}

// Version is a method to return version of this dictionary
func (d *DictionaryInfo) Version() uint16 {
	return (uint16)(d.toMecabDictionaryInfoT().version)
}

// Next is a method to pointer to the next dictionary info
func (d *DictionaryInfo) Next() (*DictionaryInfo, error) {
	return newDictionaryInfo(d.toMecabDictionaryInfoT().next)
}
