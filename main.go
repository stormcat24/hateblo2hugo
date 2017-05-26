//package main
//
//import (
//	"github.com/catatsuy/movabletype"
//	"io/ioutil"
//	"bytes"
//	"fmt"
//)
//
//func main() {
//
//	data, err := ioutil.ReadFile("blog.stormcat.io.export.txt")
//	if err != nil {
//		panic(err)
//	}
//
//	reader := bytes.NewReader(data)
//
//	entries, err := movabletype.Parse(reader)
//	if err != nil {
//		panic(err)
//	}
//
//	for _, entry := range entries {
//		//stmt.(parser.Section)
//		fmt.Printf("%+v\n", entry)
//	}
//}


package main

import "github.com/stormcat24/hateblo2hugo2/cmd"

func main() {
	cmd.Execute()
}
