package main

import (
	"log"
	"tools/Data"
)

func main() {
	// test Data
	a := Data.StrToMd5("asdsa", 10)
	log.Printf(a)
}
