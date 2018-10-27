package gonhl

import "time"

func NewStandingsParams() *standingsParams {
	return &standingsParams{
		season: -1,
	}
}

func parseStandingsParams(params *standingsParams) map[string]string {
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

type standingsParams struct {
	season        int       // Standings for a specified season (use the year season started)
	date          time.Time // Standings on a specified date
	record        bool      // Detailed information for each team including home and away records, record in shootouts, last ten games, and split head-to-head records against divisions and conferences
	standingsType string    // Get NHL standings for a specific standing type.  Retrieve Standings Types from GetStandingsTypes endpoint
}

func (sp *standingsParams) ShowDetailedRecords() *standingsParams {
	sp.record = true
	return sp
}

func (sp *standingsParams) SetSeason(season int) *standingsParams {
	sp.season = season
	return sp
}

func (sp *standingsParams) SetDate(date time.Time) *standingsParams {
	sp.date = date
	return sp
}

func (sp *standingsParams) SetStandingsType(standingsType string) *standingsParams {
	sp.standingsType = standingsType
	return sp
}
