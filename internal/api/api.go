package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
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
		return nil, fmt.Errorf("request %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http body %s", err)
	}

	slog.Info(fmt.Sprintf("%s - %d", url, resp.StatusCode))

	defer resp.Body.Close()

	return body, nil
}

func UnmarshalJson(c []byte, s interface{}) error {
	if err := json.Unmarshal(c, s); err != nil {
		return fmt.Errorf("json %s", err)
	}

	return nil
}
