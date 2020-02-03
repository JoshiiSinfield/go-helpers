package sns

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"io/ioutil"
	"log"
	"os"
)

// used for testing to read an event from a file as an unmarshalled events.SNSEvent
func ReadSNSEventFromFile(fileToLoad string) events.SNSEvent {
	dir, _ := os.Getwd()
	println(dir)
	file, err := ioutil.ReadFile(fileToLoad)
	if err != nil {
		log.Panicf("Unable to read file: %v\n", err)
	}

	var inputEvent events.SNSEvent
	err = json.Unmarshal(file, &inputEvent)
	if err != nil {
		log.Panicf("Unable to unmarshall event from file to object, %v\n", err)
	}
	return inputEvent
}
