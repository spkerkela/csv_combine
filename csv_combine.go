package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// CombineFiles Combine files. The first row is skipped from files after the first
func CombineFiles(filesToCombine []string) {
	firstRead := false
	var firstRow string
	for index := 0; index < len(filesToCombine); index++ {
		fileName := filesToCombine[index]
		fileHandle, err := os.Open(fileName)
		if err != nil {
			log.Panic(err)
		}
		fileScanner := bufio.NewScanner(fileHandle)

		if !firstRead && fileScanner.Scan() {
			firstRow = fileScanner.Text()
			fmt.Println(firstRow)
			firstRead = true
		} else {
			fileScanner.Scan()
		}

		for fileScanner.Scan() {
			fmt.Println(fileScanner.Text())
		}
	}

}

func main() {
	filesToCombine := os.Args[1:]
	CombineFiles(filesToCombine)
}
