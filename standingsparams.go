package gonhl

import "time"

func NewStandingsParams() *StandingsParams {
	return &StandingsParams{
		season: -1,
	}
}

func parseStandingsParams(params *StandingsParams) map[string]string {
	query := map[string]string{}
	expand := expandQuery("standings", map[string]bool{
		"record": params.record,
	})
	query["expand"] = expand
	if params.season != -1 {
		query["season"] = createSeasonString(params.season)
	}
	if &params.date != nil {
		query["date"] = CreateStringFromDate(params.date)
	}
	return query
}

type StandingsParams struct {
	season        int
	date          time.Time
	record        bool
	standingsType string
}

// ShowDetailedRecords makes response include detailed information for each team.
// This includes home and away records, record in shootouts, last ten games,
// and split head-to-head records against divisions and conferences.
func (sp *StandingsParams) ShowDetailedRecords() *StandingsParams {
	sp.record = true
	return sp
}

// SetSeason specifies which season to use in response (use the year season started).
// The response will represent the standings for that season.
func (sp *StandingsParams) SetSeason(season int) *StandingsParams {
	sp.season = season
	return sp
}

// SetDate specifies which date to use in response.
// The response will represent the standings on that date.
func (sp *StandingsParams) SetDate(date time.Time) *StandingsParams {
	sp.date = date
	return sp
}

// SetStandingsType specifies which standing type to use.
// Retrieve Standings Types from GetStandingsTypes endpoint
func (sp *StandingsParams) SetStandingsType(standingsType string) *StandingsParams {
	sp.standingsType = standingsType
	return sp
}
