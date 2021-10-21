package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// プリフライトリクエストへ対応
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body := make([]byte, length)
		_, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// json のフォーマット
		var buf bytes.Buffer
		err = json.Indent(&buf, []byte(body), "", "  ")
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		indentJson := buf.String()
		err = os.WriteFile("ladder.json", []byte(indentJson), 0666)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 非構造化データのアンマーシャル
		var jsonBody map[string]interface{}
		err = json.Unmarshal(body, &jsonBody)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Printf("%v\n", jsonBody)

  	w.WriteHeader(http.StatusOK)

	})

	server.ListenAndServe()
}
