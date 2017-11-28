package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func getRawText(fileNames []string) string {
	buffer := bytes.NewBufferString("")
	for _, fileName := range fileNames {
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
			continue
		}
		_, err = buffer.Write(fileBytes)
		if err != nil {
			fmt.Printf("error writing to buffer: %v", err)
			continue
		}
	}
	return buffer.String()
}
