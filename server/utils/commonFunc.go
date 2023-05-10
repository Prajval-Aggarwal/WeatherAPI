package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RapidApiCall(apiUrl string, data interface{}, host string) error {
	//fmt.Println("apiurl in common function is", apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return errors.New("failed to make API call")
	}
	req.Header.Add("X-RapidAPI-Key", "877d7321bdmsh5db7b7a54b66d8fp168429jsn4835e587d80e")
	req.Header.Add("X-RapidAPI-Host", host)
	fmt.Println("sdfjkd", req.Header.Get("X-RapidAPI-Host"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("failed to read response body")
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error is:", err)
		return err
	}
	//fmt.Println("data is from common", data)
	return nil
}
