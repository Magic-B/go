package main

import (
	"fmt"
	"io"
	"os"
)

const filePath = "reversedText.txt"

func main() {
	fWrite, err := os.Create(filePath)

	if err != nil {
		trowErr(err)
	}

	defer fWrite.Close()

	str := "Hello World"
	err = writeString(str, fWrite)

	if err != nil {
		trowErr(err)
	}

	fRead, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v", err)
		return
	}

	content, err := io.ReadAll(fRead)

	if err != nil {
		trowErr(err)
	}

	fmt.Printf("Content: %s", content)
}

func writeString(s string, w io.Writer) error {
	reversedStr := reverseStr(s)

	_, err := w.Write([]byte(reversedStr))

	if err != nil {
		return err
	}

	return nil
}

func reverseStr(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func trowErr(err error) {
	panic(fmt.Errorf("Error: %w", err))
}
