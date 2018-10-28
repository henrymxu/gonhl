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

	// Output: Eastern
}

func ExampleClient_GetDivision() {
	client := gonhl.NewClient()
	division := client.GetDivision(18)
	fmt.Println(division.Name)

	// Output: Metropolitan
}

func ExampleClient_GetDraft() {
	client := gonhl.NewClient()
	draft := client.GetDraft(2018)
	fmt.Println(draft[0].DraftYear)

	// Output: 2018
}

