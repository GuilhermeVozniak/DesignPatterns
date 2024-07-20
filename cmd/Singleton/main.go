package main

import (
	"fmt"

	"github.com/GuilhermeVozniak/DesignPatterns/pkg/singleton"
)

func main() {

	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
