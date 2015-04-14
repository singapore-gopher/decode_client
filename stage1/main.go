package main

import (
	`bytes`
	`io/ioutil`
	`log`
	`net/http`
)

func main() {
	register(`mark`)
	stage1(`mark`)
}

func stage1(teamName string) {
	resp, err := http.Get(`http://localhost:4000/stage1/data.json`)
	if err != nil {
		log.Fatalf(`could not get stage1 data; err=%v`, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf(`could not read resp; err=%v`, err)
	}

	// do something with the body and figure out the result

	buffer := bytes.NewBufferString(`{"team":"` + teamName + `"}`) // send back your data
	resp, err = http.Post(`http://localhost:4000/stage1/submit.json`, `application/json`, buffer)
	if err != nil {
		log.Fatalf(`could submit to stage1; err=%v`, err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf(`could not stage 1 submit response; err=%v`, err)
	}
	log.Printf(`stage1: server responsed: %q`, body)
}

func register(teamName string) {
	buffer := bytes.NewBufferString(`{"name":"` + teamName + `"}`)
	resp, err := http.Post(`http://localhost:4000/register.json`, `application/json`, buffer)
	if err != nil {
		log.Fatalf(`could not register; err=%v`, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf(`could not read resp; err=%v`, err)
	}
	log.Printf(`server responsed: %q`, body)
}
