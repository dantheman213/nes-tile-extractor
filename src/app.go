package main

import (
	"bytes"
	"encoding/hex"
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
	// fmt.Printf("File contents: %s\n", romBytes)

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

func checkValidNesRom(romData []byte) {
	romHeader := make([]byte, 4)
	copy(romHeader, romData)

	nesHeaderSignature, ex := hex.DecodeString("4e45531a")
	if ex != nil {
		log.Fatal(ex)
	}

	if bytes.Compare(romHeader, nesHeaderSignature) == 0 {
		fmt.Println("Appears to be a valid NES ROM...")
	} else {
		log.Fatal("This is an invalid NES ROM file!")
	}
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
