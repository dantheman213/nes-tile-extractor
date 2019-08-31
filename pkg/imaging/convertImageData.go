package imaging

import (
    "../common"
    "image"
    "image/color"
    "image/png"
    "os"
)

/* How a single CHR is represented:
   - Each CHR is 128 bits (16 bytes)
   - Each CHR is 8x8 pixels
   - Each pixel within a CHR is one of 4 colors
   - Each pixel within a CHR is 2 bits (0b00 = color0, 0b10 = color1, 0b01 = color2, 0b11 = color3)
   - The first 64 bits of the CHR represent the first color bit for each of the 8x8 pixels, starting from top-left, traversing horizontally through each of the 8 rows from left to right
   - The second 64 bits of the CHR represent the second color bit for each of the 8x8 pixels, starting from top-left, traversing horizontally through each of the 8 rows from left to right
*/
func ConvertChrDataToImageData(chr []byte) [8][8]int {
    chrBytesChannelA := make([]byte, 8)
    copy(chrBytesChannelA[:], chr[0:7])

    chrBytesChannelB := make([]byte, 8)
    copy(chrBytesChannelB[:], chr[7:15])

    // This will convert 64 bits (8 bytes) x2 channels of data in CHR into easy to read 0 and 1 binary.
    chrBitStrChannelA := ConvertChrChannelBytesToBinaryStr(chrBytesChannelA)
    chrBitStrChannelB := ConvertChrChannelBytesToBinaryStr(chrBytesChannelB)

    // Sprites are 8x8 pixel bitmap graphics stored in 16 byte blocks in the CHR data banks.
    // Those 16 bytes are stored as 2 channels of 8 bytes each. Since a byte is 8 bits,
    // each byte of the channel represents one row of the sprite.
    xPos := 0
    yPos := 0
    var imgRawColorStruct[8][8] int
    for channelsIndexPos := 0; channelsIndexPos < 64; channelsIndexPos++ {
        if chrBitStrChannelA[channelsIndexPos:channelsIndexPos + 1] == "0" &&
            chrBitStrChannelB[channelsIndexPos:channelsIndexPos + 1] == "0" {
            imgRawColorStruct[xPos][yPos] = 0
        } else if chrBitStrChannelA[channelsIndexPos:channelsIndexPos + 1] == "1" &&
            chrBitStrChannelB[channelsIndexPos:channelsIndexPos + 1] == "0" {
            imgRawColorStruct[xPos][yPos] = 1
        } else if chrBitStrChannelA[channelsIndexPos:channelsIndexPos + 1] == "0" &&
            chrBitStrChannelB[channelsIndexPos:channelsIndexPos + 1] == "1" {
            imgRawColorStruct[xPos][yPos] = 2
        } else if chrBitStrChannelA[channelsIndexPos:channelsIndexPos + 1] == "1" &&
            chrBitStrChannelB[channelsIndexPos:channelsIndexPos + 1] == "1" {
            imgRawColorStruct[xPos][yPos] = 3
        }

        xPos += 1
        if xPos > 7 {
            xPos = 0
            yPos += 1
        }
    }

    return imgRawColorStruct
}

func GenerateModernImgFromChrData(data [8][8]int, targetImagePath string, scale int) {
    imgData := image.NewRGBA(image.Rect(0, 0, 7, 7))
    // TODO use scale

    for chrPosY := 0; chrPosY < 8; chrPosY++ {
        for chrPosX := 0; chrPosX < 8; chrPosX++ {
            switch data[chrPosX][chrPosY] {
            case 0:
                // black
                imgData.Set(chrPosX, chrPosY, color.RGBA{0, 0, 0, 255})
            case 1:
                // dark black
                imgData.Set(chrPosX, chrPosY, color.RGBA{33, 33, 33, 255})
            case 2:
                // dark gray
                imgData.Set(chrPosX, chrPosY, color.RGBA{94, 94, 94, 255})
            case 3:
                // gray
                imgData.Set(chrPosX, chrPosY, color.RGBA{221, 221, 221, 255})
            default:
                // TODO
            }
        }
    }

    // generate the image
    imageFile, fileErr := os.OpenFile(targetImagePath, os.O_WRONLY|os.O_CREATE, 0644)
    if fileErr != nil {
        // TODO
    }

    encodeErr := png.Encode(imageFile, imgData)
    if encodeErr != nil {
        // TODO
    }

    _ = imageFile.Close()
}

func ConvertChrChannelBytesToBinaryStr(channel []byte) string {
    result := "";

    for _, singleByte := range channel {
        result += common.HexadecimalToBinary(singleByte)
    }

    return result
}
