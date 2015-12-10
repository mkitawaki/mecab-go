package mecab

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type Memorize map[string]*C.char

func NewMemorize() *Memorize {
	return &Memorize{}
}

func (m *Memorize) Cache(key string, value *C.char) {
	_, ok := (*m)[key]
	if ok {
		C.free(unsafe.Pointer((*m)[key]))
	}
	(*m)[key] = value
}

func (m *Memorize) Clear() {
	for _, value := range *m {
		C.free(unsafe.Pointer(value))
	}
}
