package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files, err := dirFiles(dir)
	if err != nil {
		panic(err)
	}
	line := countLine(files)

	fmt.Printf("Search in  %s\n", dir)
	fmt.Printf("line count: %d\n", line)
	fmt.Printf("file count: %d\n", len(files))

}

func dirFiles(pathName string) ([]string, error) {
	dir, err := ioutil.ReadDir(pathName)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, fileInfo := range dir {
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}
		if fileInfo.IsDir() {
			fs, err := dirFiles(path.Join(pathName, fileInfo.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, fs...)
		} else {
			//ignore hidden file
			if !checkIsText(path.Join(pathName, fileInfo.Name())) {
				continue
			}
			files = append(files, path.Join(pathName, fileInfo.Name()))
		}

	}
	return files, err

}
func checkIsText(fileName string) bool {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	contentType := http.DetectContentType(file)
	if !strings.Contains(contentType, "text/") {
		return false
	}
	return true
}
func countLine(files []string) int64 {
	var count int64
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			continue
		}
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			if len(line) != 0 {
				count++
			}
		}

	}
	return count

}
