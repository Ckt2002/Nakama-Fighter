package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func CreateMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	modulename := "standard_match"
	params := map[string]interface{}{}

	matchID, err := nk.MatchCreate(ctx, modulename, params)
	if err != nil {
		return "", err
	}

	logger.Info(matchID)

	return "", nil
}
