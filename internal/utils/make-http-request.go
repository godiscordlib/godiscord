package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func MakeRequest(Method, URL string, Data any, Reason string) (any, error) {
	body, err := json.Marshal(Data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(Method, URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("GODISCORD_TOKEN"))
	req.Header.Set("X-Audit-Log-Reason", Reason)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == 204 {
		return nil, nil
	} else {
		res_body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return res_body, nil
	}
}
