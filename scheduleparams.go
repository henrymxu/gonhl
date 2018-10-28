package gonhl

import "time"

func NewScheduleParams() *ScheduleParams {
	return &ScheduleParams{
		teamId: -1,
	}
}

func parseScheduleParams(params *ScheduleParams) map[string]string {
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
		query["date"] = CreateStringFromDate(params.date)
	} else if &params.startDate != nil {
		query["startDate"] = CreateStringFromDate(params.startDate)
		query["endDate"] = CreateStringFromDate(params.endDate)
	}
	return query
}

type ScheduleParams struct {
	broadcasts bool
	linescore  bool
	ticket     bool
	teamId     int
	date       time.Time
	startDate  time.Time
	endDate    time.Time
}

// ShowBroadcasts makes response include the broadcasts for games.
func (sp *ScheduleParams) ShowBroadcasts() *ScheduleParams {
	sp.broadcasts = true
	return sp
}

// ShowLinescore makes response include linescore for completed games.
func (sp *ScheduleParams) ShowLinescore() *ScheduleParams {
	sp.linescore = true
	return sp
}

// ShowTicketRetailers makes response include the different places to buy tickets for the upcoming games.
func (sp *ScheduleParams) ShowTicketRetailers() *ScheduleParams {
	sp.ticket = true
	return sp
}

// SetTeamId limits the response to a specific team.
// Team ids can be found through the teams endpoint.
func (sp *ScheduleParams) SetTeamId(teamId int) *ScheduleParams {
	sp.teamId = teamId
	return sp
}

// SetDate specifies a single date for the response or a range of dates for the response.
// Providing a single date will return the schedule for that date.
// Providing multiple dates will return the schedule for dates starting from the 1st date to the 2nd date.
func (sp *ScheduleParams) SetDate(date ...time.Time) *ScheduleParams {
	if len(date) == 1 {
		sp.date = date[0]
	} else {
		sp.startDate = date[0]
		sp.endDate = date[1]
	}
	return sp
}
