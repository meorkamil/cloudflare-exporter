package api

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func GetAPI(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	}

	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("ERROR: HTTP Request ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ERROR:", err)
	}

	log.Println("INFO:", resp.StatusCode, url)

	defer resp.Body.Close()

	return body, nil
}

func UnmarshalJson(c []byte, s interface{}) error {
	if err := json.Unmarshal(c, s); err != nil {
		log.Fatal("ERROR:", err)
	}

	return nil
}
