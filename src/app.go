package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting program...")
	checkArgs()

	fmt.Println("Loading ROM...")
	romPath := os.Args[1]
	romBytes := loadRom(romPath)

	fmt.Printf("File contents: %s", romBytes)
}

func checkArgs() {
	if len(os.Args) < 2 {
		log.Fatal("Requires path to NES ROM file!")
	}
}

func loadRom(filePath string) []byte {
	contents, ex := ioutil.ReadFile(filePath)

	if ex != nil {
		log.Fatal(ex)
	}

	return contents
}
