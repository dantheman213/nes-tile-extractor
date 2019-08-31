package common

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func GetCurrentApplicationDir() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(dir)
    return dir
}

func HexadecimalToBinary(in byte) string {
    var out []byte

    for i := 7; i >= 0; i-- {
        b := (in >> uint(i))
        out = append(out, (b%2)+48)
    }

    // fmt.Printf("%x -> %s\n", in, out)
    return string(out)
}