package main

import (
	"fmt"

	"github.com/tisonpatrik/HyZen/internal"
)

func main() {
	internal.CheckRoot()
	internal.RunPreInstall()
	output, err := internal.RunCommand("echo 'HyZen is running'")
	if err != nil {
		fmt.Println("❌ Command execution failed:", err)
	}
	fmt.Println(output)
}
