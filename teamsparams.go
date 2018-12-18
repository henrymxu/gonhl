package gonhl

func NewTeamsParams() *TeamsParams {
	return &TeamsParams{
		season: -1,
	}
}

func parseTeamsParams(params TeamsParams) map[string]string {
	query := map[string]string{}
	expand := expandQuery("team", map[string]bool{
		"roster":            params.roster,
		"schedule.next":     params.scheduleNext,
		"schedule.previous": params.schedulePrev,
		"stats":             params.stats,
	})
	query["expand"] = expand
	if params.season != -1 {
		query["season"] = createSeasonString(params.season)
	}
	query["teamId"] = combineIntArray(params.ids)
	query["stats"] = params.statsType
	return query
}

type TeamsParams struct {
	roster       bool
	scheduleNext bool
	schedulePrev bool
	stats        bool
	season       int
	ids          []int
	statsType    string
}

// ShowDetailedRoster makes response include the roster of active players for teams.
func (tParams *TeamsParams) ShowDetailedRoster() *TeamsParams {
	tParams.roster = true
	return tParams
}

// ShowScheduleNext makes response include details of the upcoming games for teams.
func (tParams *TeamsParams) ShowScheduleNext() *TeamsParams {
	tParams.scheduleNext = true
	return tParams
}

// ShowSchedulePrev makes response include details of the previous games for teams.
func (tParams *TeamsParams) ShowSchedulePrev() *TeamsParams {
	tParams.schedulePrev = true
	return tParams
}

// ShowTeamStats makes response include the teams stats for the season.
func (tParams *TeamsParams) ShowTeamStats() *TeamsParams {
	tParams.stats = true
	return tParams
}

// SetSeason specifies which season to use in response (use the year season started).
// The response will represent the roster for that season.
func (tParams *TeamsParams) SetSeason(season int) *TeamsParams {
	tParams.season = season
	return tParams
}

// SetTeamIds specifies which teams to include in response.
// Can string team id together to get multiple teams.
func (tParams *TeamsParams) SetTeamIds(ids ...int) *TeamsParams {
	tParams.ids = ids
	return tParams
}

// SetStatsType specifies which team stats to get.
// Retrieve Standings Types from <TBD> endpoint.
func (tParams *TeamsParams) SetStatsType(statsType string) *TeamsParams {
	tParams.statsType = statsType
	return tParams
}
