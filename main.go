package main

import (
	"github.com/Tzzg/go-tool/Data"
	"log"
)

func main() {
	// test Data
	a := Data.StrToMd5("asdsa", 10)
	log.Printf(a)
}
