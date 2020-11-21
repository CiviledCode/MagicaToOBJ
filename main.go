package main

import (
	"fmt"
	"github.com/civiledcode/MagicaToOBJ/vox"
	"os"
)

func main() {
	path,_ := os.Getwd()
	tr, _ := vox.LoadVOXAndTriangulate(path + "/testing/testing.vox")
	fmt.Println(len(tr))
	for _, triangle := range tr {
		fmt.Printf("%+v", triangle)
	}
}
