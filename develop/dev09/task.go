package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//Реализовать утилиту wget с возможностью скачивать сайты целиком.

func main() {
	fullURLFile, fileName := getURLAndFileName()   //получаем URL, как аргумент и возвращаем URL и имя файла.
	file, client := createFile(fileName)           //создаем файл.
	size := writeToFile(client, file, fullURLFile) //получаем содержимое(HTML) и кладем в файл.

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)
}

func getURLAndFileName() (string, string) {
	urlFile := flag.String("url", "", "Введите адрес сайта")
	flag.Parse()
	fileURL, err := url.Parse(*urlFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	if fileName == "" && fileURL.Host != "" {
		fileName = fileURL.Host
	}
	return *urlFile, fileName
}

func createFile(fileName string) (*os.File, *http.Client) {
	file, err := os.Create("save_web/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return file, &client
}

func writeToFile(client *http.Client, file *os.File, fullURLFile string) int64 {
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	defer file.Close()

	return size
}
