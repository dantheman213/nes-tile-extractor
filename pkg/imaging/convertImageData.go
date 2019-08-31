package imaging

/* How a single CHR is represented:
   - Each CHR is 128 bits (16 bytes)
   - Each CHR is 8x8 pixels
   - Each pixel within a CHR is one of 4 colors
   - Each pixel within a CHR is 2 bits (0b00 = color0, 0b10 = color1, 0b01 = color2, 0b11 = color3)
   - The first 64 bits of the CHR represent the first color bit for each of the 8x8 pixels, starting from top-left, traversing horizontally through each of the 8 rows from left to right
   - The second 64 bits of the CHR represent the second color bit for each of the 8x8 pixels, starting from top-left, traversing horizontally through each of the 8 rows from left to right
*/
func ConvertChrDataToImageData(chr []byte) {

}

func saveChrAsImageAtPath(path string) {

}
