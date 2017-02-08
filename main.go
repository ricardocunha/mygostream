package main

import (
	"fmt"

	"github.com/ricardocunha/twistream"
)

func main() {
	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		CONSUMERKEY,
		CONSUMERSECRET,
		ACCESSTOKEN,
		ACCESSTOKENSECRET,
	)

	// Listen timeline
	for {
		status := <-timeline.Listen()
		fmt.Println(status)
	}

	// Tweet to timeline
	status := twistream.Status{
		Text:              "@otiai10 How does Go work?",
		InReplyToStatusId: 493324823926288386,
	}
	_ := timeline.Tweet(status)
	fmt.Println("Teste")
}
