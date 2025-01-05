package main

import (
	"fmt"

	"github.com/cgcoder/gocode/jsonutility/util"
)

func main() {
	util.FormatFile("/Users/cgopi24/cgcoder/code/gocode/jsonutility/in.txt", "/Users/cgopi24/cgcoder/code/gocode/jsonutility/out.txt")
	fmt.Println("Done!")
}
