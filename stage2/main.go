package main

import (
	`bytes`
	`io/ioutil`
	`log`
	`net/http`
)

func main() {
	register(`mark`)
	stage2(`mark`)
}

func stage2(teamName string) {
	resp, err := http.Get(`http://10.0.2.235:4000/stage2/data.json`)
	if err != nil {
		log.Fatalf(`could not get stage2 data; err=%v`, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf(`could not read resp; err=%v`, err)
	}

	// do something with the body and figure out the result

	buffer := bytes.NewBufferString(`{"team":"` + teamName + `","faulty":[]}`) // send back your data
	resp, err = http.Post(`http://10.0.2.235:4000/stage2/submit.json`, `application/json`, buffer)
	if err != nil {
		log.Fatalf(`could submit to stage2; err=%v`, err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf(`could not stage 1 submit response; err=%v`, err)
	}
	log.Printf(`stage2: server responsed: %q`, body)
}

func register(teamName string) {
	buffer := bytes.NewBufferString(`{"name":"` + teamName + `"}`)
	resp, err := http.Post(`http://10.0.2.235:4000/register.json`, `application/json`, buffer)
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
