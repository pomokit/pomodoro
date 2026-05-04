//go:build !windows

package main

import "fmt"

func setConsoleTitle(title string) {
	fmt.Printf("\033]0;%s\007", title)
}
