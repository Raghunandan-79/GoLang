package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fileInfo() {
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInformation, err := f.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Println("Filename:", fileInformation.Name())
	fmt.Println("File is directory:", fileInformation.IsDir())
	fmt.Println("File size:", fileInformation.Size())
	fmt.Println("File modification time:", fileInformation.ModTime())
}

func readFile() {
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}

	fileSize := fileInfo.Size()

	buff := make([]byte, fileSize)
	data, err := f.Read(buff)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data inside file:", string(buff[:data]))
}

func readSmallFiles() {
	f, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Data inside file:", string(f))
}

func readFolders() {
	dir, err := os.Open("../")

	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}
}

func createFile() {
	f, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString("Hello")
	f.WriteString(", Golang")

	bytes := []byte(". Good Language")
	f.Write(bytes)
}

func replaceContent() {
	f, err := os.Open("example2.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	
	if err != nil {
		panic(err)
	}

	buff := make([]byte, fileInfo.Size())

	_, err = f.Read(buff)
	if err != nil {
		panic(err)
	}

	content := string(buff)

	content = strings.ReplaceAll(content, "Good Language", "It is a good language.")

	err = os.WriteFile("example2.txt", []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Replacement done!")
}

func copyContent() {
	sourceFile, err := os.Open("example2.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example3.txt")

	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		err1 := writer.WriteByte(b)

		if err1 != nil {
			panic(err1)
		}

		writer.Flush()
	}

	fmt.Println("Written")
}

func deleteFile() {
	err := os.Remove("example3.txt")

	if err != nil {
		panic(err)
	}


	fmt.Println("File Deleted Successfully")
}

func main() {
	fileInfo()
	readFile()
	readSmallFiles()
	readFolders()
	createFile()
	replaceContent()
	copyContent()
	deleteFile()
}
