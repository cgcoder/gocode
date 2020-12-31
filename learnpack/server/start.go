package server

import (
	"fmt"
	"os"

	"github.com/gocode/learnpack/config"
	"github.com/gocode/learnpack/dal"
)

// InitServer start up
func InitServer() {
	initErr := config.InitConfig()
	if initErr != nil {
		fmt.Fprintf(os.Stderr, "Config init failed %v\n", initErr)
		os.Exit(1)
	}
	initErr = dal.InitSQL()
	if initErr != nil {
		fmt.Fprintf(os.Stderr, "DB init failed %v\n", initErr)
		os.Exit(1)
	}
	fmt.Println("Server started... OK")
}

// UnInitServer tear down
func UnInitServer() {
	dal.UninitSQL()
	fmt.Println("Server exited...")
}
