package gonhl

import "fmt"

const endpointDrafts = "/draft"
const endpointDraft = endpointDrafts + "/%d"
const endpointDraftProspects = endpointDrafts + "/prospects"
const endpointDraftProspect = endpointDraftProspects + "/%d"

// GetCurrentDraft retrieves information about the current NHL draft.
func (c *Client) GetCurrentDraft() []Draft {
	var draft drafts
	status := c.makeRequest(endpointDrafts, nil, &draft)
	fmt.Println(status)
	return draft.Drafts
}

// GetDraft retrieves information about a NHL draft from a specific year.
func (c *Client) GetDraft(year int) []Draft {
	var draft drafts
	status := c.makeRequest(fmt.Sprintf(endpointDraft, year), nil, &draft)
	fmt.Println(status)
	return draft.Drafts
}

// GetProspects retrieves information about all NHL prospects (beware large response).
func (c *Client) GetProspects() []Prospect {
	var prospects prospects
	status := c.makeRequest(endpointDraftProspects, nil, &prospects)
	fmt.Println(status)
	return prospects.Prospects
}

// GetProspect retrieves information about a single NHL prospect using a prospect id.
func (c *Client) GetProspect(id int) Prospect {
	var prospects prospects
	status := c.makeRequest(fmt.Sprintf(endpointDraftProspect, id), nil, &prospects)
	fmt.Println(status)
	return prospects.Prospects[0]
}

// API Endpoint
type drafts struct {
	Drafts []Draft `json:"drafts"`
}

// API Endpoint
type prospects struct {
	Prospects []Prospect `json:"prospects"`
}

type Draft struct {
	DraftYear int     `json:"draftYear"`
	Rounds    []Round `json:"rounds"`
}

type Round struct {
	RoundNumber int    `json:"roundNumber"`
	Round       string `json:"round"`
	Picks       []Pick `json:"picks"`
}

type Pick struct {
	Year        int    `json:"year"`
	Round       string `json:"round"`
	PickOverall int    `json:"pickOverall"`
	PickInRound int    `json:"pickInRound"`
	Team        Team   `json:"team"`
	Prospect    Player `json:"prospect"`
}

type Prospect struct {
	ID                 int              `json:"id"`
	FullName           string           `json:"fullName"`
	Link               string           `json:"link"`
	FirstName          string           `json:"firstName"`
	LastName           string           `json:"lastName"`
	BirthDate          string           `json:"birthDate"`
	BirthCity          string           `json:"birthCity"`
	BirthStateProvince string           `json:"birthStateProvince"`
	BirthCountry       string           `json:"birthCountry"`
	Height             string           `json:"height"`
	Weight             int              `json:"weight"`
	ShootsCatches      string           `json:"shootsCatches"`
	PrimaryPosition    Position         `json:"primaryPosition"`
	DraftStatus        string           `json:"draftStatus"`
	ProspectCategory   ProspectCategory `json:"prospectCategory"`
	AmateurTeam        Team             `json:"amateurTeam"`
	AmateurLeague      League           `json:"amateurLeague"`
	Ranks              Ranks            `json:"ranks"`
}

type Ranks struct {
	Midterm   int `json:"midterm"`
	FinalRank int `json:"finalRank"`
	DraftYear int `json:"draftYear"`
}

type ProspectCategory struct {
	ID        int    `json:"id"`
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
}
