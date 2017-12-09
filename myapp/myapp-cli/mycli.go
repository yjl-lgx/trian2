package main

import (
	"fmt"
	"flag"
)

var My_file *string = flag.String("file", "musicfile", "Use -file <filesource>")

func main () {
	flag.Parse()
	fmt.Println(*My_file)
}