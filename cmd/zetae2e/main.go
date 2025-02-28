package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	_ "github.com/zeta-chain/node/pkg/sdkconfig/default"
)

func main() {
	// enable color output
	color.NoColor = false

	// initialize root command
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
