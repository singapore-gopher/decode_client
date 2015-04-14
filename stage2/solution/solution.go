package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	server = "10.0.2.235:4000"
)

type sensorData struct {
	Id          *int  `json:"id"`
	Smoke       *bool `json:"have_smoke"`
	Humidity    *int  `json:"humidity"`
	Temperature *int  `json:"temperature"`
}

type sensorResponse struct {
	SensorData []*sensorData `json:"sensors"`
}

// {
//    "have_smoke" : false,
//    "id" : 234,
//    "humidity" : 40,
//    "temperature" : 25
// },

func main() {

	requestUrl := url.URL{}
	requestUrl.Host = server
	requestUrl.Scheme = "http"
	requestUrl.Path = "/stage2/data.json"

	stage1response, err := http.Get(requestUrl.String())
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	// YOUR CODE HERE
	data, err := ioutil.ReadAll(stage1response.Body)
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	var sensors sensorResponse
	err = json.Unmarshal(data, &sensors)
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	faulty := []int{}

	for _, sensor := range sensors.SensorData {

		if sensor.Smoke == nil ||
			sensor.Humidity == nil ||
			sensor.Temperature == nil {
			faulty = append(faulty, *sensor.Id) // assume, that is does not fail
		}

	}

	response := struct {
		Team   string `json:"team"`
		Faulty []int  `json:"faulty"`
	}{
		"zoltan",
		faulty,
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	log.Println(string(responseData))

	body := bytes.NewBuffer(responseData)

	responseUrl := requestUrl
	responseUrl.Path = "/stage2/submit.json"
	stage1SubmitResponse, err := http.Post(responseUrl.String(), "application/json", body)
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	responseData, _ = ioutil.ReadAll(stage1SubmitResponse.Body)
	log.Println(string(responseData))
}
