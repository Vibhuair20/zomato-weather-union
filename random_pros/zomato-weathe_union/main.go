package main

import (
	"fmt"
	"io"
	"net/http"
)

const apiKey = "812a7412e07a33ee95aa187bd3d53749"

func zomatoweatherunion(w http.ResponseWriter, r *http.Request) {

	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	if latitude == "" || longitude == "" {
		http.Error(w, "bhai kuch toh dal tune sab galat dala hai", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("https://www.weatherunion.com/gw/weather/external/v0/get_weather_data?latitude=%s&longitude=%s", latitude, longitude)
	// creating the http request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "fail ho gaya -> haww", http.StatusInternalServerError)
		return
	}
	// add api to the request header
	req.Header.Add("X-Zomato-Api-Key", apiKey)
	// reading the response made by the http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "failed response", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	// forwarding the api response
	w.Header().Set("content-type", "application-json")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "failed to read the response body", http.StatusInternalServerError)
		return
	}
	w.Write(body)

}

func main() {

	http.HandleFunc("/zomato", zomatoweatherunion)
	fmt.Println("server is running on http 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting the server:", err)
	}
}
