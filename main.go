package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := logFile.Write([]byte("Menambahkan data\n")); err != nil {
		log.Fatal(err)
	}
	if err := logFile.Close(); err != nil {
		log.Fatal(err)
	}
	data := logFile
	output := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(output)

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("sample.json", file, 0644)
}
