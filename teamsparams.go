package gonhl

func NewTeamsParams() *teamsParams {
	return &teamsParams{
		season: -1,
	}
}

func parseTeamsParams(params teamsParams) map[string]string {
	query := map[string]string{}
	expand := expandQuery("team", map[string]bool{
		"roster":       params.roster,
		"scheduleNext": params.scheduleNext,
		"schedulePrev": params.schedulePrev,
		"stats":        params.stats,
	})
	query["expand"] = expand
	if params.season != -1 {
		query["season"] = createSeasonString(params.season)
	}
	query["teamId"] = combineIntArray(params.ids)
	query["stats"] = params.statsType
	return query
}

type teamsParams struct {
	roster       bool   // Shows roster of active players for the specified team
	scheduleNext bool   // Returns details of the upcoming game for a team
	schedulePrev bool   // Same as above but for the last game played
	stats        bool   // Returns the teams stats for the season
	season       int    // Adding the season identifier shows the roster for that season (use the year season started)
	ids          []int  // Can string team id together to get multiple teams
	statsType    string // Specify which stats to get. Retrieve Standings Types from <TBD> endpoint
}

func (tParams *teamsParams) SetDetailedRoster() *teamsParams {
	tParams.roster = true
	return tParams
}

func (tParams *teamsParams) SetScheduleNext() *teamsParams {
	tParams.scheduleNext = true
	return tParams
}

func (tParams *teamsParams) SetSchedulePrev() *teamsParams {
	tParams.schedulePrev = true
	return tParams
}

func (tParams *teamsParams) SetTeamStats() *teamsParams {
	tParams.stats = true
	return tParams
}

func (tParams *teamsParams) SetSeason(season int) *teamsParams {
	tParams.season = season
	return tParams
}

func (tParams *teamsParams) SetTeamIds(ids ...int) *teamsParams {
	tParams.ids = ids
	return tParams
}

func (tParams *teamsParams) SetStatsType(statsType string) *teamsParams {
	tParams.statsType = statsType
	return tParams
}
