package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/singapore-gophers/decode_json/stream" // go get github.com/singapore-gophers/decode_json/stream
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
)

type response struct {
	sync.RWMutex
	TeamName string   `json:"team"`
	Values   []string `json:"values"`
}

// for your convenience
func (self *response) addValue(value string) {
	self.Lock()
	defer self.Unlock()

	// we need 3 values only..
	if len(self.Values) == 3 {
		self.Values = self.Values[1:]
	}

	self.Values = append(self.Values, value)

}

// again, for your convenience. Use this to submit your solution.
func (self *response) submit(url, teamName string) error {

	self.TeamName = teamName
	data, _ := json.Marshal(self)
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Err %v - %s", res.StatusCode, string(responseData))
	}

	fmt.Println(string(responseData))
	return nil
}

func main() {
	conn, err := net.Dial("tcp", "10.0.2.235:4001")
	if err != nil {
		log.Fatal(err)
	}

	err = handleConn(conn)
	if err != nil {
		log.Println(err)
	}

	// done
}

func handleConn(conn io.ReadWriteCloser) error {
	// hints:
	// • Decode to stream.Packet objects. readme: http://.../../readme.md
	// • Use encoding/json's decoder
	// • To check if the received packet contains "magic"
	//   use stream.IsMagicValue(value string) bool
	// • Use response.Submit in the end, it posts your solution with your team's name.

	// create the json stream decoder
	dec := json.NewDecoder(conn)

	// you will collect 3 magic values in this struct
	solution := &response{
		Values: []string{},
	}

	for {

		// For checking a packet's value use stream.IsMagicValue()

		// You can use solution.addValue()

	}

	url := "http://10.0.2.235:4000/stage3/submit.json" // the server's url
	teamName := "YOUR TEAM NAME HERE !!!"                       // your team's name
	err := solution.submit(url, teamName)                       //
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// done
	return nil
}
