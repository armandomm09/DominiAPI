package team

type Options struct {
	port int `help:"Port to listen on" short:"p" default:"8888"`
}

type TeamNicknameResponse struct {
	Body struct {
		Nickname string `json:"nickname"`
	}
}

type Team struct {
	Nickname string `json:"nickname"`
}

type AvgTeamsEvent struct {
	Body struct {
		AvgMatch float64 `json:"avgMatch"`
		AvgAuto  float64 `json:"avgAuto"`
		AvgStage float64 `json:"avgStage"`
	}
}

//MATCH SENDER *********************************************************************

type MatchDataListSender struct {
	Body struct {
		Matches []MatchSender `json:matches`
	}
}

type MatchSender struct {
	Won                bool     `json:"won"`
	ActualTime         int64    `json:"actualTime"`
	PredictedTime      int64    `json:"predictedTime"`
	IsRedAlliance      bool     `json:"isRedAlliance"`
	BlueTeamNumbers    []string `json:"blueTeamNumbers"`
	RedTeamNumbers     []string `json:"RedTeamNumbers"`
	AllianceFinalScore int      `json:"AllianceFinalScore"`
	OpponentFinalScore int      `json:"OpponentFinalScore"`
	CompetitionLevel   string   `json:"CompetitionLevel"`
	MatchKey           string   `json:"matchKey"`
	MatchNumber        int      `json:"matchNumber"`
	MatchTitle         string   `json:"matchTitle"`
}

// MATCH RECIEVER *********************************************************************
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

// STATBOTICS EPA GETTER
type EPAGetter struct {
	Team string `json:"team"`
	Epa  struct {
		Unitless float64 `json:"unitless"`
	} `json:"epa"`
}

type OPRsSender struct {
	Body struct {
		EPA  float64 `json:"epa"`
		OPR  float64 `json:"opr"`
		DPR  float64 `json:"dpr"`
		CCWM float64 `json:"ccwm"`
	}
}
