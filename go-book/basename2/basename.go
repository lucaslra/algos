package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}

func basename(path string) string {
	slash := strings.LastIndex(path,"/")
	path = path[slash+1:]
	if dot:=strings.LastIndex(path,".");dot >=0{
		path = path[:dot]
	}

	return path
}
