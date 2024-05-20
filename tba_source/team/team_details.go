package team

import (
	"apiv1/tba_source"
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
		Tags:          []string{"FRC"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		TeamNumber string `path:"number" maxLength:"5" example:"5887"`
		EventKey   string `path:"eventKey" maxLength:"8" example:"2024mxpu"`
	}) (*MatchDataList, error) {

		url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/team/frc%s/event/%v/matches", input.TeamNumber, input.EventKey)

		body, err := tba_source.AccessTBA(url)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		var matches []MatchReciever
		if err := json.Unmarshal([]byte(body), &matches); err != nil {
			fmt.Println(string(err.Error()))
			return nil, fmt.Errorf(err.Error())
		}

		matchesList := MatchDataList{}

		matchesList.Body.Matches = matches

		return &matchesList, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "get-team-nickname",
		Method:        http.MethodGet,
		Path:          "/team/{number}",
		Summary:       "Get the nickname of a eam",
		Tags:          []string{"FRC"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		Number string `path:"number" maxLength:"5" example:"5887"`
	}) (*TeamNicknameResponse, error) {

		url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/team/frc%s", input.Number)

		body, err := tba_source.AccessTBA(url)
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
		Tags:          []string{"FRC"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		TeamNumber string `path:"teamNumber" maxLength:"5" example:"5887"`
		EventKey   string `path:"eventKey" maxLength:"8" example:"2024mxpu"`
	}) (*AvgTeamsEvent, error) {

		url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/team/frc%s/event/%v/status", input.TeamNumber, input.EventKey)

		body, err := tba_source.AccessTBA(url)
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
