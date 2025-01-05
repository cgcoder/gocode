package main

import (
	"fmt"
	"time"

	"github.com/cgcoder/gocode/jsonutility/util"
)

func main() {
	start := time.Now()
	util.FormatFile("/Users/cgopi24/cgcoder/code/gocode/jsonutility/message_group3.json", "/Users/cgopi24/cgcoder/code/gocode/jsonutility/out.txt")
	fmt.Printf("Done in %v\n", time.Since(start))
}
