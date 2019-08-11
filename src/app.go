package main

import (
	"bytes"
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

	checkValidNesRom(romBytes)
	getDataFromRom(romBytes)
	// fmt.Printf("File contents: %s", romBytes)

	fmt.Println("Complete!")
}

func checkArgs() {
	argCount := len(os.Args) - 1
	if argCount < 1 {
		log.Fatal("Requires path to NES ROM file!")
	} else if argCount > 1 {
		log.Fatal("Too many arguments!")
	}
}

func loadRom(filePath string) []byte {
	contents, ex := ioutil.ReadFile(filePath)

	if ex != nil {
		log.Fatal(ex)
	}

	return contents
}

func checkValidNesRom(data []byte) {
	header := make([]byte, 4)
	copy(header, data)

	if bytes.Compare(header, []byte("4e45531a")) == 0 {
		fmt.Printf("Appears to be a valid NES ROM...")
	} else {
		log.Fatal("This is an invalid NES ROM file!")
	}
	//fmt.Printf("%x", header)
}

func getDataFromRom(data []byte) {
	offset := 5
	pgr := data[offset]

	offset += 1
	chr := data[offset]

	offset += 1
	trainer := data[offset] & 0x4

	offset += 9

	if trainer != 0 {
		offset += 512
	}

	// if chrbanks is 0 then the sprites are embedded in pgr blocks instead
	var multiplier byte
	if chr == 0 {
		multiplier = pgr
	} else {
		// skip pgr section
		offset += int(pgr) * 16384
		multiplier = chr
	}

	chrbanks := make([]byte, 8192 * int(multiplier))
	copy(chrbanks[:], data)
}
