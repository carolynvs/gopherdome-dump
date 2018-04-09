package gopherdump

import (
	"fmt"

	"github.com/carolynvs/gopherdome-json"
)

func Dump(gopher interface{}) {
	fmt.Println("!!!DUMPING ALL THE PRETTY GOPHERS!!!")
	fmt.Println(gopherson.Pretty(gopher))
}
