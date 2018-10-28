package gonhl

func NewPlayerParams() *PlayerParams {
	return &PlayerParams{}
}

func parseParams(params *PlayerParams) map[string]string {
	query := map[string]string{}
	query["id"] = string(params.id)
	query["season"] = createSeasonString(params.season)
	query["stats"] = combineStringArray(params.stat)
	return query
}

type PlayerParams struct {
	id     int
	season int
	stat   []string
}

// SetId specifies which player to include in response.
func (pParams *PlayerParams) SetId(id int) *PlayerParams {
	pParams.id = id
	return pParams
}

// SetSeason specifies which season to use in response (use the year season started).
// The response will represent the player stats for that season.
func (pParams *PlayerParams) SetSeason(season int) *PlayerParams {
	pParams.season = season
	return pParams
}

// SetStat specifies which single season player stats to get.
func (pParams *PlayerParams) SetStat(stat ...string) *PlayerParams {
	pParams.stat = stat
	return pParams
}
