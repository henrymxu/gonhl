package gonhl

func NewPlayerParams() *playerParams {
	return &playerParams{}
}

func parseParams(params *playerParams) map[string]string {
	query := map[string]string{}
	query["id"] = string(params.id)
	query["season"] = createSeasonString(params.season)
	query["stats"] = combineStringArray(params.stat)
	return query
}

type playerParams struct {
	id     int      // Player id
	season int      // Player stats for that season (use the year season started)
	stat   []string // Obtains single season statistics for a player
}

func (pParams *playerParams) SetId(id int) *playerParams {
	pParams.id = id
	return pParams
}

func (pParams *playerParams) SetSeason(season int) *playerParams {
	pParams.season = season
	return pParams
}

func (pParams *playerParams) SetStat(stat ...string) *playerParams {
	pParams.stat = stat
	return pParams
}
