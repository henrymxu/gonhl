package gonhl_test

import (
	"fmt"
	"github.com/henrymxu/gonhl"
)

func ExampleNewClient() {
	client := gonhl.NewClient()
	fmt.Println(client)
}

func ExampleClient_GetConference() {
	client := gonhl.NewClient()
	conference := client.GetConference(6)
	fmt.Println(conference.Name)
}
