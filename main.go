package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	initialJson := `
	[
  {
    "actual_time": 1710525103,
    "alliances": {
      "blue": {
        "dq_team_keys": [],
        "score": 21,
        "surrogate_team_keys": [],
        "team_keys": [
          "frc6608",
          "frc9134",
          "frc6702"
        ]
      },
      "red": {
        "dq_team_keys": [],
        "score": 45,
        "surrogate_team_keys": [],
        "team_keys": [
          "frc9103",
          "frc5887",
          "frc9060"
        ]
      }
    },
    "comp_level": "qm",
    "event_key": "2024mxpu",
    "key": "2024mxpu_qm15",
    "match_number": 15,
    "predicted_time": 1710529672,
    "set_number": 1,
    "time": 1710524460,
    "winning_alliance": "red"
  },
  {
    "actual_time": 1710525103,
    "alliances": {
      "blue": {
        "dq_team_keys": [],
        "score": 21,
        "surrogate_team_keys": [],
        "team_keys": [
          "frc6608",
          "frc9134",
          "frc5887"
        ]
      },
      "red": {
        "dq_team_keys": [],
        "score": 45,
        "surrogate_team_keys": [],
        "team_keys": [
          "frc9103",
          "frc6702",
          "frc9060"
        ]
      }
    },
    "comp_level": "qm",
    "event_key": "2024mxpu",
    "key": "2024mxpu_qm15",
    "match_number": 15,
    "predicted_time": 1710529672,
    "set_number": 1,
    "time": 1710524460,
    "winning_alliance": "red"
  }
	]
	`

	var matches []MatchReciever
	if err := json.Unmarshal([]byte(initialJson), &matches); err != nil {
		fmt.Println(string(err.Error()))
	}

	matchesData := MatchDataListReciever{}
	matchesData.Body.Matches = matches

	matchesSender := MatchDataListSender{}

	for _, match := range matchesData.Body.Matches {
		matchSender := MatchSender{}

		isRedAlliance := true
		for _, teamKey := range match.Alliances.Blue.TeamKeys {
			if teamKey == "frc5887" {
				isRedAlliance = false
				break
			}
		}
		matchSender.IsRedAlliance = isRedAlliance

		if (match.WinningAlliance == "red" && isRedAlliance) || (match.WinningAlliance == "false" && !isRedAlliance) {
			matchSender.Won = true
		} else {
			matchSender.Won = false
		}
		matchesSender.Body.Matches = append(matchesSender.Body.Matches, matchSender)
	}

	marshaled, err := json.Marshal(matchesSender)
	if err != nil {
		fmt.Println("Error marshaling")
	}

	fmt.Println(string(marshaled))
}
