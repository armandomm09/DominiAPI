package main

type MatchReciever struct {
	ActualTime int64 `json:"actual_time"`
	Alliances  struct {
		Blue struct {
			DQTeamKeys        []string `json:"dq_team_keys"`
			Score             int      `json:"score"`
			SurrogateTeamKeys []string `json:"surrogate_team_keys"`
			TeamKeys          []string `json:"team_keys"`
		} `json:"blue"`
		Red struct {
			DQTeamKeys        []string `json:"dq_team_keys"`
			Score             int      `json:"score"`
			SurrogateTeamKeys []string `json:"surrogate_team_keys"`
			TeamKeys          []string `json:"team_keys"`
		} `json:"red"`
	} `json:"alliances"`
	CompLevel       string `json:"comp_level"`
	EventKey        string `json:"event_key"`
	Key             string `json:"key"`
	MatchNumber     int    `json:"match_number"`
	PredictedTime   int64  `json:"predicted_time"`
	SetNumber       int    `json:"set_number"`
	Time            int64  `json:"time"`
	WinningAlliance string `json:"winning_alliance"`
}

type MatchDataListReciever struct {
	Body struct {
		Matches []MatchReciever `json:"matches"`
	}
}

type MatchDataListSender struct {
	Body struct {
		Matches []MatchSender `json:matches`
	}
}

type MatchSender struct {
	Won                bool     `json:"won"`
	ActualTime         int64    `json:"actualTime"`
	IsRedAlliance      bool     `json:"isRedAlliance"`
	BlueTeamNumbers    []string `json:"blueTeamNumbers"`
	RedTeamNumbers     []string `json:"RedTeamNumbers"`
	AllianceFinalScore int64    `json:"AllianceFinalScore"`
	OpponentFinalScore int64    `json:"OpponentFinalScore"`
	CompetitionLevel   string   `json:"CompetitionLevel"`
	MatchKey           string   `json:"matchKey"`
	MatchNumber        int64    `json:"matchNumber"`
	MatchTitle         string   `json:"matchTitle"`
}
