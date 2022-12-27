package entity

import "time"

type MatchResult struct {
	teamAScore int
	teamBScore int
}

func NewMatchResult(teamAScore int, teamBScore int) *MatchResult {
	return &MatchResult{
		teamAScore: teamAScore,
		teamBScore: teamBScore
	}
}

func (m *MatchResult) GetResult() string {
	return strconv.Itoa(m.teamAScore) + "-" + strconv.Itoa(m.teamBScore)
}

type Match struct {
	ID string
	TeamA *Team
	TeamB *Team
	TeamAID string
	TeamBID string
	Date time.Time
	Status string
	Result MatchResult
	Actions []GameAction
}

func NewMatch(id string, teamA *Team, teamB *Team, teamAID string, teamBID string, date time.Time) *Match {
	return &Match{
		ID: id,
		TeamA: teamA,
		TeamB: teamB,
		TeamAID: teamAID,
		TeamBID: teamBID,
		Date: date,
	}
}