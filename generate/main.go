package main

import "github.com/pupuk/addr/generate/autoCode"

//go:generate go run main.go
func main() {
	autoCode.AutoAreaMap()
}
