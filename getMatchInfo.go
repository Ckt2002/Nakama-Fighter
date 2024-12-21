package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func GetMatchPlayerNumber(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	var data map[string]interface{}
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		return "", err
	}

	matchIdJSON, ok := data["matchId"].(string)
	if !ok {
		logger.Error("Invalid matchId value.")
		return "", nil
	}

	logger.Info("matchIdJSON: %v", matchIdJSON)

	matchId := matchIdJSON

	match, err := nk.MatchGet(ctx, matchId)
	if err != nil {
		logger.Error(err.Error())
		return err.Error(), nil
	}

	logger.Info("Player number: %v", match.Size)

	response := map[string]interface{}{
		"playerNumber": match.Size,
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(responseJson), nil
}
