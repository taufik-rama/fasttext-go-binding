package main

// #cgo LDFLAGS: -L${SRCDIR} -lfasttext -lstdc++ -lm
// #include <stdlib.h>
// void load_model(char *path);
// int predict(char *query, float *prob, char *buf, int buf_sz);
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {
	LoadModel("fastText/models/basic-model.bin")
	fmt.Println(Predict("pulsa"))
	fmt.Println(Predict("puls"))
	fmt.Println(Predict("Felix"))
	fmt.Println(Predict("felix"))
}

// LoadModel based on the `path`
func LoadModel(path string) {
	C.load_model(C.CString(path))
}

// Predict the `sentence`
func Predict(sentence string) (prob float32, label string, err error) {

	var cprob C.float
	var buf *C.char
	buf = (*C.char)(C.calloc(64, 1))

	ret := C.predict(C.CString(sentence), &cprob, buf, 64)

	if ret != 0 {
		err = errors.New("error in prediction")
	} else {
		label = C.GoString(buf)
		prob = float32(cprob)
	}
	C.free(unsafe.Pointer(buf))

	return prob, label, err
}