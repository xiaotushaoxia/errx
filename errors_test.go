package errx

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWrap(t *testing.T) {
	_, err := myOpen3("xx")
	fmt.Println(Format(err))
}

func myOpen(path string) ([]byte, error) {
	_, err := http.Get(path)
	if err != nil {
		return nil, Wrap(err, "myOpen")
	}
	return []byte("Ok"), nil
}

func myOpen2(path string) ([]byte, error) {
	open, err := myOpen(path)
	if err != nil {
		return open, Wrap(err, "MyOpen2 failed")
	}
	return nil, nil
}

func myOpen3(path string) ([]byte, error) {
	open, err := myOpen2(path)
	if err != nil {
		return open, err
	}
	return nil, nil
}

//=== RUN   TestWrap
//Error: MyOpen2 failed: myOpen: Get "xx": unsupported protocol scheme ""
//Stack:
//github.com/xiaotusaoxia/errx.myOpen
//	/src/errx/errors_test.go:17
//github.com/xiaotusaoxia/errx.myOpen2
//	/src/errx/errors_test.go:23
//github.com/xiaotusaoxia/errx.myOpen3
//	/src/errx/errors_test.go:31
//github.com/xiaotusaoxia/errx.TestWrap
//	/src/errx/errors_test.go:10
//testing.tRunner
//	C:/Program Files/Go120/src/testing/testing.go:1576
//runtime.goexit
//	C:/Program Files/Go120/src/runtime/asm_amd64.s:1598
//--- PASS: TestWrap (0.00s)
//PASS
