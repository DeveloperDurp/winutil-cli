package main

import (
	"fmt"
	"runtime"
	"winutil-cli/cmd"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This is a Windows Only Utility")
		return
	}
	cmd.Execute()
}
