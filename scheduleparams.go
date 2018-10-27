package gonhl

import "time"

func NewScheduleParams() *scheduleParams {
	return &scheduleParams{
		teamId: -1,
	}
}

func parseScheduleParams(params *scheduleParams) map[string]string {
	query := map[string]string{}
	expand := expandQuery("schedule", map[string]bool{
		"broadcasts": params.broadcasts,
		"ticket":     params.ticket,
		"linescore":  params.linescore,
	})
	query["expand"] = expand
	if params.teamId != -1 {
		query["teamId"] = string(params.teamId)
	}
	if &params.date != nil {
		query["date"] = CreateDateFromTime(params.date)
	} else if &params.startDate != nil {
		query["startDate"] = CreateDateFromTime(params.startDate)
		query["endDate"] = CreateDateFromTime(params.endDate)
	}
	return query
}

type scheduleParams struct {
	broadcasts bool      // Shows the broadcasts of the game
	linescore  bool      // Linescore for completed games
	ticket     bool      // Provides the different places to buy tickets for the upcoming games
	teamId     int       // Limit results to a specific team. Team ids can be found through the teams endpoint
	date       time.Time // Single defined date for the search
	startDate  time.Time // Start date for the search
	endDate    time.Time // End date for the search
}

func (sp *scheduleParams) ShowBroadcasts() *scheduleParams {
	sp.broadcasts = true
	return sp
}

func (sp *scheduleParams) ShowLinescore() *scheduleParams {
	sp.linescore = true
	return sp
}

func (sp *scheduleParams) ShowTicketRetailers() *scheduleParams {
	sp.ticket = true
	return sp
}

func (sp *scheduleParams) SetTeamId(teamId int) *scheduleParams {
	sp.teamId = teamId
	return sp
}

func (sp *scheduleParams) SetDate(date ...time.Time) *scheduleParams {
	if len(date) == 1 {
		sp.date = date[0]
	} else {
		sp.startDate = date[0]
		sp.endDate = date[1]
	}
	return sp
}
