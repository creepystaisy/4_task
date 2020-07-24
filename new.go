package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var dir string
var name *string = flag.String("name", "name", "строка")
var put *string = flag.String("put", "put", "f - вперёд, b - назад")

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//length = len(lines)
	return lines, scanner.Err()
}
func workWithString(a string) string {
	s := strings.Split(a, "/")
	str := strings.Join(s, "")
	s = strings.Split(str, "https:")
	a = strings.Join(s, "")
	fmt.Print(a)
	return a
}

func chanal(lines []string, path1 string, i int) {
	fmt.Print("Potok #")
	fmt.Println(i)
	resp, err := http.Get(lines[i])
	if err != nil {
		fmt.Println(err)

		//return err
	}
	a := lines[i]
	a = workWithString(a)
	dir = path.Join(path1, "/")
	dir = path.Join(dir, a)
	os.MkdirAll(dir, 0644)
	openedDir := path1 + "/" + a + "/" + a + ".html"
	file, err := os.OpenFile(openedDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //os.Create(path)

	if err != nil {
		fmt.Println(err)

		//return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	io.Copy(w, resp.Body)
	defer resp.Body.Close()
}

//var makedDir string

func main() {
	flag.Parse()
	//использавать библ флаг
	inputFile := *name
	outputFile := *put
	lines, err := readLines(inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
	for i := 0; i < len(lines); i++ {
		fmt.Print(i)
		fmt.Print("inner:")
		go chanal(lines, outputFile, i)
	}
	fmt.Println("")
	var input string
	fmt.Scanln(&input)
	//writeLines(lines, outputFile)
	//var text string = outputFile + "res.txt"
	//if err := writeLines(lines, outputFile); err != nil {
	//	log.Fatalf("writeLines: %s", err)
	//}
}
