package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	errInternalError  = runtime.NewError("internal server error", 13) // INTERNAL
	errMarshal        = runtime.NewError("cannot marshal type", 13)   // INTERNAL
	errNoInputAllowed = runtime.NewError("no input allowed", 3)       // INVALID_ARGUMENT
	errNoUserIdFound  = runtime.NewError("no user ID in context", 3)  // INVALID_ARGUMENT
	errUnmarshal      = runtime.NewError("cannot unmarshal type", 13) // INTERNAL
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	// Register Match Handler
	err := initializer.RegisterMatch("standard_match", newMatch)
	if err != nil {
		logger.Error("[RegisterMatch] error: ", err.Error())
		return err
	}

	err = initializer.RegisterRpc("createMatch", CreateMatch)
	if err != nil {
		logger.Error("Error CreateMatch: ", err.Error())
		return err
	}

	err = initializer.RegisterRpc("getMatchPlayerNumber", GetMatchPlayerNumber)
	if err != nil {
		logger.Error("Error GetMatchPlayerNumber: ", err.Error())
		return err
	}

	err = initializer.RegisterRpc("dailyRewards", DailyReward)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("dailyRewards", DailyReward)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("updateWallet", UpdateWallet)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = initializer.RegisterRpc("deleteItem", DeleteItem)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Module loaded in %dms", time.Since(initStart).Milliseconds())

	return nil
}
