package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAPI(url string) []byte {

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body

}

func UnmarshalJson(c []byte, s interface{}) error {

	if err := json.Unmarshal(c, s); err != nil {
		log.Fatalln(err)
	}
	return nil
}
