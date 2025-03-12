package main

import (
	"fmt"

	"github.com/tisonpatrik/HyZen/internal"
)

func main() {
	internal.CheckRoot()
	output := internal.RunCommand("echo 'HyZen is running'")
	fmt.Println(output)
}
