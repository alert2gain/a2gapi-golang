package main

import (
	"fmt"
	"time"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

/* Here we define three struct.
   The first and second define the Sensor.
   The third contains all the data that we will send to the API. */
type Sensor struct {
	Station string
	Data  interface{}
}

type InternalData struct {
	Temperature float64
	Humidity float64
	LastRead string
}

type AllData struct {
	IKEY string
	Data string
}

/* In this method we initialize the sensor with some values. */
func createSensor(){
	/* Intialize the sensor with your own logic. */
	s := Sensor{
		Station: "Station 1",
		Data: InternalData{
			Temperature: 15,
			Humidity: 10,
			LastRead: time.Now().UTC().Format("2006-01-02T15:04:05")}}

	sendData(s)
}

/* In this method we send the data to A2G InputStream API */
func sendData(sensor Sensor){
	/* First we transform the sensor to a JSON */
	sen, _ := json.Marshal(sensor)

	/* Here we initialize the AllData struct */
	d := AllData{
		IKEY: "[YOUR_INPUTSTREAM_KEY]",
		Data: string(sen)}

	/* And here we transform all the data in a JSON */
	data, _ := json.Marshal(d)

	/* Finally, we send the data to the API */
	url := "https://listen.a2g.io/v1/testing/inputstream"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("x-api-key", "[YOUR_API_KEY]")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("")
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

/* Here we start with the process */
func main() {
	createSensor()
}