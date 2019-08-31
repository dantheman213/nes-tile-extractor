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

	//chrBankCount, prgBankCount := getRomHeaderMetadata(filePayloadBytes)
	getRomHeaderMetadata(filePayloadBytes)
	//extractTileDataFromRom(filePayloadBytes)
}

func loadRomFileDataToArray(filePath string) (contents []byte) {
	contents, ex := ioutil.ReadFile(filePath)

	if ex != nil {
		log.Fatal(ex)
	}

	return
}

func checkValidNesRom(romData []byte) (result bool) {
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

func getRomHeaderMetadata(romData []byte) (offsetBytesToChrData, chrDataSizeInBytes int) {
	prgBankCount := uint8(romData[4])
	chrBankCount := uint8(romData[5])

	fmt.Printf("Found %d PRG blocks and %d CHR blocks.\n", prgBankCount, chrBankCount)

	offsetBytesToChrData = 16 + 16384 * int(prgBankCount)
	chrDataSizeInBytes = 8192 * int(chrBankCount)

	fmt.Printf("Calculated - CHR byte offset: %d, CHR size in bytes: %d, ROM size in bytes: %d\n", offsetBytesToChrData, chrDataSizeInBytes, len(romData))

	if len(romData) < offsetBytesToChrData + chrDataSizeInBytes {
		log.Fatal("Invalid ROM payload or unable to calculate CHR bank location and size.\n")
	}

	return
}

func extractTileDataFromRom(romData []byte) {
	offsetBytes := 0

	// move past nes signature and top level information
	offsetBytes = 16

	pgr := romData[offsetBytes]
	offsetBytes += 1

	chr := romData[offsetBytes]
	offsetBytes += 1

	trainer := romData[offsetBytes] & 0x4
	offsetBytes += 1

	offsetBytes += 9

	if trainer != 0 {
		offsetBytes += 512
	}

	// if chrbanks is 0 then the sprites are embedded in pgr blocks instead
	var multiplier byte
	if chr == 0 {
		multiplier = pgr
	} else {
		// skip pgr section
		offsetBytes += int(pgr) * 16384
		multiplier = chr
	}

	dataSize := 8192 * int(multiplier)
	chrBanks := make([]byte, dataSize)
	copy(chrBanks[:], romData[offsetBytes:dataSize])

	fmt.Println("Reading ROM Data...")
}

func convertTileDataToImageData() {

}

func saveTileAsImageAtPath() {

}
