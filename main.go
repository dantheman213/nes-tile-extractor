package main

import "fmt"

func main() {
	fmt.Println("Welcome to NES Tile Extractor!")

	checkArgs()
	importRomDataFromFile()
}

func checkArgs() {

}

func importRomDataFromFile() {
	loadDataToArray()
	checkValidNesRom()
	extractTileDataFromRom()
}

func loadDataToArray() {

}

func checkValidNesRom() {

}

func extractTileDataFromRom() {
	// convertTileDataToImageData()

}

func convertTileDataToImageData() {

}

func saveTileAsImageAtPath() {

}
