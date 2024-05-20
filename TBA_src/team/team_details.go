package team

import (
	"apiv1/SB_src"
	"apiv1/TBA_src"
	"context"
	"encoding/json"
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

func AddTeamRegisters(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID:   "get-teams-matches-by-event",
		Method:        http.MethodGet,
		Path:          "/team/{number}/event/{eventKey}",
		Summary:       "Shows the team's matches of the given event",
		Tags:          []string{"Team Event info"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		TeamNumber string `path:"number" maxLength:"5" example:"5887"`
		EventKey   string `path:"eventKey" maxLength:"8" example:"2024mxpu"`
	}) (*MatchDataListSender, error) {

		url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/team/frc%s/event/%v/matches", input.TeamNumber, input.EventKey)

		body, err := TBA_src.AccessTBA(url)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		var matches []MatchReciever
		if err := json.Unmarshal([]byte(body), &matches); err != nil {
			fmt.Println(string(err.Error()))
			return nil, fmt.Errorf(err.Error())
		}
		matchesData := MatchDataListReciever{}
		matchesData.Body.Matches = matches

		matchesSender := MatchDataListSender{}

		for _, match := range matchesData.Body.Matches {
			matchSender := MatchSender{}

			isRedAlliance := true
			matchSender.AllianceFinalScore = match.Alliances.Red.Score
			matchSender.OpponentFinalScore = match.Alliances.Blue.Score
			for _, teamKey := range match.Alliances.Blue.TeamKeys {
				matchSender.BlueTeamNumbers = append(matchSender.BlueTeamNumbers, teamKey[3:])
				if teamKey == "frc"+input.TeamNumber {
					isRedAlliance = false
					matchSender.AllianceFinalScore = match.Alliances.Blue.Score
					matchSender.OpponentFinalScore = match.Alliances.Red.Score

				}
			}
			matchSender.IsRedAlliance = isRedAlliance

			for _, teamKey := range match.Alliances.Red.TeamKeys {
				matchSender.RedTeamNumbers = append(matchSender.RedTeamNumbers, teamKey[3:])

			}

			if (match.WinningAlliance == "red" && isRedAlliance) || (match.WinningAlliance == "false" && !isRedAlliance) {
				matchSender.Won = true
			} else {
				matchSender.Won = false
			}

			switch match.CompLevel {
			case "qm":
				matchSender.CompetitionLevel = "Qualification"
			case "sf":
				matchSender.CompetitionLevel = "Semifinal"
			case "f":
				matchSender.CompetitionLevel = "Final"
			}

			matchSender.MatchKey = match.Key
			matchSender.MatchNumber = match.MatchNumber
			matchSender.MatchTitle = matchSender.CompetitionLevel + " " + fmt.Sprintf("%d", matchSender.MatchNumber)
			matchSender.ActualTime = match.ActualTime
			matchSender.PredictedTime = match.PredictedTime
			matchesSender.Body.Matches = append(matchesSender.Body.Matches, matchSender)
		}

		return &matchesSender, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "get-team-EPA",
		Method:        http.MethodGet,
		Path:          "/team/{number}/year/{year}",
		Summary:       "Get the EPA of a team",
		Tags:          []string{"Team Year info"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		Number string `path:"number" maxLength:"5" example:"5887"`
		Year   string `path:"year" maxLength:"4" example:"2024"`
	}) (*OPRsSender, error) {

		url := fmt.Sprintf("https://api.statbotics.io/v3/team_year/%s/%s", input.Number, input.Year)

		body, err := SB_src.AccessSB(url)
		if err != nil {
			fmt.Println(string(err.Error()))
			return nil, fmt.Errorf(err.Error())
		}
		var epas EPAGetter
		if err := json.Unmarshal([]byte(body), &epas); err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		output := OPRsSender{}
		output.Body.EPA = epas.Epa.Unitless

		return &output, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "get-teams-events",
		Method:        http.MethodGet,
		Path:          "/team/{number}/year/{year}",
		Summary:       "Get the list of events of a team given the year",
		Tags:          []string{"Team Year info"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		Number string `path:"number" maxLength:"5" example:"5887"`
		Year   string `path:"year" maxLength:"4" example:"2024"`
	}) (*OPRsSender, error) {

		url := fmt.Sprintf("https://api.statbotics.io/v3/team_year/%s/%s", input.Number, input.Year)

		body, err := SB_src.AccessSB(url)
		if err != nil {
			fmt.Println(string(err.Error()))
			return nil, fmt.Errorf(err.Error())
		}
		var epas EPAGetter
		if err := json.Unmarshal([]byte(body), &epas); err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		output := OPRsSender{}
		output.Body.EPA = epas.Epa.Unitless

		return &output, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "get-team-nickname",
		Method:        http.MethodGet,
		Path:          "/team/{number}",
		Summary:       "Get the nickname of a eam",
		Tags:          []string{"Team Simple info"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		Number string `path:"number" maxLength:"5" example:"5887"`
	}) (*TeamNicknameResponse, error) {

		url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/team/frc%s", input.Number)

		body, err := TBA_src.AccessTBA(url)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		var team Team
		if err := json.Unmarshal([]byte(body), &team); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		fmt.Println(string(team.Nickname))
		output := TeamNicknameResponse{}
		output.Body.Nickname = team.Nickname
		return &output, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "get-details-by-event",
		Method:        http.MethodGet,
		Path:          "/team/{teamNumber}/event{eventKey}",
		Summary:       "Get the general stats of a team by event",
		Tags:          []string{"Team Event info"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		TeamNumber string `path:"teamNumber" maxLength:"5" example:"5887"`
		EventKey   string `path:"eventKey" maxLength:"8" example:"2024mxpu"`
	}) (*AvgTeamsEvent, error) {

		url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/team/frc%s/event/%v/status", input.TeamNumber, input.EventKey)

		body, err := TBA_src.AccessTBA(url)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		type Stats struct {
			Qual struct {
				Ranking struct {
					SortOrders []float64 `json:"sort_orders"`
				} `json:"ranking"`
			} `json:"qual"`
		}
		var teamStats Stats
		if err := json.Unmarshal([]byte(body), &teamStats); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		output := AvgTeamsEvent{}
		output.Body.AvgMatch = teamStats.Qual.Ranking.SortOrders[2]
		output.Body.AvgAuto = teamStats.Qual.Ranking.SortOrders[3]
		output.Body.AvgStage = teamStats.Qual.Ranking.SortOrders[4]

		return &output, nil

	})
}
