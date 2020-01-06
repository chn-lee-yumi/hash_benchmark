package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"reflect"
	"time"
)

func countFuncRunTime(d []byte, name string, funcs map[string]interface{}) {
	f := reflect.ValueOf(funcs[name])
	param := make([]reflect.Value, 1)
	param[0] = reflect.ValueOf(d)

	t0 := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		f.Call(param)
	}
	t1 := time.Now().UnixNano()

	dt := t1 - t0
	fmt.Printf("%s 用时: %.3f s\n", name, float32(dt)/1000000000)
}

func main() {

	funcs := map[string]interface{}{
		"md5":    md5.Sum,
		"sha1":   sha1.Sum,
		"sha224": sha256.Sum224,
		"sha256": sha256.Sum256,
		"sha384": sha512.Sum384,
		"sha512": sha512.Sum512,
	}

	d, err := ioutil.ReadFile("test.img")
	if err != nil {
		fmt.Println("read fail", err)
	}

	countFuncRunTime(d, "md5", funcs)
	countFuncRunTime(d, "sha1", funcs)
	countFuncRunTime(d, "sha224", funcs)
	countFuncRunTime(d, "sha256", funcs)
	countFuncRunTime(d, "sha384", funcs)
	countFuncRunTime(d, "sha512", funcs)
	countFuncRunTime(d, "md5", funcs)
	countFuncRunTime(d, "sha1", funcs)
	countFuncRunTime(d, "sha224", funcs)
	countFuncRunTime(d, "sha256", funcs)
	countFuncRunTime(d, "sha384", funcs)
	countFuncRunTime(d, "sha512", funcs)

}
