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
	fmt.Println("Welcome to NES Tile Extractor!")

	romPath := checkArgs()
	importRomDataFromFile(romPath)
}

func checkArgs() (romPath string) {
	romPath = os.Args[1]
	return
}

func importRomDataFromFile(filePath string) {
	filePayloadBytes := loadRomFileDataToArray(filePath)
	validRomResult := checkValidNesRom(filePayloadBytes)

	if validRomResult {
		fmt.Println("Appears to be a valid NES ROM...")
	} else {
		log.Fatal("This is an invalid NES ROM file!")
	}

	extractTileDataFromRom(filePayloadBytes)
}

func loadRomFileDataToArray(filePath string) (contents []byte) {
	contents, ex := ioutil.ReadFile(filePath)

	if ex != nil {
		log.Fatal(ex)
	}

	return
}

func checkValidNesRom(romData []byte) (result bool){
	romHeader := make([]byte, 4)
	copy(romHeader, romData)

	nesHeaderSignature, ex := hex.DecodeString("4e45531a")
	if ex != nil {
		log.Fatal(ex)
	}

	if bytes.Compare(romHeader, nesHeaderSignature) == 0 {
		return true;
	}

	return false;
}

func extractTileDataFromRom(romData []byte) {
	// start after nes signature
	offset := 4

	pgr := romData[offset]
	offset += 1

	chr := romData[offset]
	offset += 1

	trainer := romData[offset] & 0x4
	offset += 1

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

	dataSize := 8192 * int(multiplier)
	chrBanks := make([]byte, dataSize)
	copy(chrBanks[:], romData[offset:dataSize])

	fmt.Println("Reading ROM Data...")
}

func convertTileDataToImageData() {

}

func saveTileAsImageAtPath() {

}
