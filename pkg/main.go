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

	offsetBytesToChrData, chrDataSizeInBytes := getRomHeaderMetadata(filePayloadBytes)
	chrBankData := getChrBankData(filePayloadBytes, offsetBytesToChrData, chrDataSizeInBytes)
	extractTileDataFromRom(chrBankData)
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
	fmt.Println("Calculating ROM properties...")

	prgBankCount := uint8(romData[4])
	chrBankCount := uint8(romData[5])

	fmt.Printf("Found %d PRG blocks and %d CHR blocks.\n", prgBankCount, chrBankCount)

	offsetBytesToChrData = 16 + 16384 * int(prgBankCount)
	chrDataSizeInBytes = 8192 * int(chrBankCount)

	fmt.Printf("Calculated ROM properties:\nCHR byte offset: %d\nCHR size in bytes: %d\nROM size in bytes: %d\n", offsetBytesToChrData, chrDataSizeInBytes, len(romData))

	if len(romData) < offsetBytesToChrData + chrDataSizeInBytes {
		log.Fatal("Invalid ROM payload or unable to calculate CHR bank location and size.\n")
	}

	return
}

func getChrBankData(romData []byte, offsetBytesToChrData, chrDataSizeInBytes int) (chrBankData []byte) {
	fmt.Println("Getting CHR data bank...")

	chrBankData = make([]byte, chrDataSizeInBytes)
	copy(chrBankData[:], romData[offsetBytesToChrData:offsetBytesToChrData + chrDataSizeInBytes])

	return
}

func extractTileDataFromRom(chrBankData []byte) {
	chrCount := (len(chrBankData) + 1) / 16
	fmt.Printf("There are %d CHRs (sprites) in CHR data bank.\n", chrCount)

	fmt.Printf("Extracting")
	for chrIndex := 0; chrIndex < chrCount; chrIndex++ {
		fmt.Printf(".")
		chrData := make([]byte, 16)
		copy(chrData[:], chrBankData[chrIndex * 16:(chrIndex * 16) + 16])
		convertChrDataToImageData(chrData)
	}
}

func convertChrDataToImageData(chr []byte) {

}

func saveTileAsImageAtPath() {

}
