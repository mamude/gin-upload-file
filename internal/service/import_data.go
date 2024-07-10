package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ImportData(fileName string) {
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
