package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		abort()
	}

	str := os.Args[1]
	if str == "" {
		abort()
	}

	fmt.Printf("    EUCJP : %s\n", encode(str, japanese.EUCJP.NewEncoder()))
	fmt.Printf("ISO2022JP : %s\n", encode(str, japanese.ISO2022JP.NewEncoder()))
	fmt.Printf(" ShiftJIS : %s\n", encode(str, japanese.ShiftJIS.NewEncoder()))
}

func abort() {
	fmt.Println("Usage:\n  jpn-bytes string\n\nExample:\n  jpn-bytes 日本語")
	os.Exit(0)
}

func encode(s string, t transform.Transformer) string {
	reader := strings.NewReader(s)
	r := transform.NewReader(reader, t)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return hex.EncodeToString(buf.Bytes())
}
