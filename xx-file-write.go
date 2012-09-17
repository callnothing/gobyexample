package main

import "os"

func main() {
	file, err := os.Create("xx-file-write.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    file.WriteString("contents\n")
}
