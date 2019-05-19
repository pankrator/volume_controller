package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/itchyny/volume-go"
)

func main() {
	http.HandleFunc("/get", func(rw http.ResponseWriter, req *http.Request) {
		writeVolume(rw)
	})

	http.HandleFunc("/up", func(rw http.ResponseWriter, req *http.Request) {
		volume.IncreaseVolume(10)
		fmt.Println("Volume increased")
		writeVolume(rw)
	})

	http.HandleFunc("/down", func(rw http.ResponseWriter, req *http.Request) {
		volume.IncreaseVolume(-10)
		fmt.Println("Volume decreased")
		writeVolume(rw)
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.String()
		if path == "/" {
			path = "index.html"
		}
		fmt.Printf("path=%s\n", path)
		f, err := os.Open(fmt.Sprintf("public/%s", path))
		if err != nil {
			fmt.Printf("could not open file: %v\n", err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Something happenned"))
			return
		}
		result, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Printf("could not read file: %v\n", err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Something happenned2"))
			return
		}
		rw.Write(result)
	})

	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		panic(err)
	}
}

func writeVolume(rw http.ResponseWriter) {
	vol, err := volume.GetVolume()
	if err != nil {
		fmt.Printf("Could not get volume: %v\n", err)
	}
	rw.Write([]byte(fmt.Sprintf("Volume is: %d", vol)))
}
