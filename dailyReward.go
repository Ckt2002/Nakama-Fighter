package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

var Coins struct {
	CoinsReceived int64 `json:"coins_received"`
}

type lastDailyReward struct {
	LastClaim int64 `json:"last_claim"` // The last time the user claimed the reward in UNIX time.
}

func DailyReward(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return "", errNoUserIdFound
	}

	if len(payload) > 0 {
		return "", errNoInputAllowed
	}

	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{{
		Collection: "reward",
		Key:        "daily",
		UserID:     userID,
	}})
	if err != nil {
		logger.Error("StorageRead error: %v", err)
		return "", errInternalError
	}

	dailyReward := &lastDailyReward{
		LastClaim: 0,
	}

	for _, object := range objects {
		if object.GetKey() == "daily" {
			if err := json.Unmarshal([]byte(object.GetValue()), dailyReward); err != nil {
				logger.Error("Unmarshal error: %v", err)
				return "", errUnmarshal
			}
			break
		}
	}
	Coins.CoinsReceived = int64(0)

	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	if time.Unix(dailyReward.LastClaim, 0).Before(midnight) {
		Coins.CoinsReceived = 10

		// Update player wallet.
		changeset := map[string]int64{
			"coins": Coins.CoinsReceived,
		}
		metadata := map[string]interface{}{
			"action": "Daily Reward",
		}

		if _, _, err := nk.WalletUpdate(ctx, userID, changeset, metadata, true); err != nil {
			logger.Error("WalletUpdate error: %v", err)
			return "", errInternalError
		}

		err = nk.NotificationsSend(ctx, []*runtime.NotificationSend{{
			Code: 1001,
			Content: map[string]interface{}{
				"coins": changeset["coins"],
			},
			Persistent: true,
			Sender:     "", // Server sent.
			Subject:    "You've received your daily reward!",
			UserID:     userID,
		}})
		if err != nil {
			logger.Error("NotificationsSend error: %v", err)
			return "", errInternalError
		}

		dailyReward.LastClaim = time.Now().Unix()

		object, err := json.Marshal(dailyReward)
		if err != nil {
			logger.Error("Marshal error: %v", err)
			return "", errInternalError
		}

		version := ""
		if len(objects) > 0 {
			// Use OCC to prevent concurrent writes.
			version = objects[0].GetVersion()
		}

		// Update daily reward storage object for user.
		_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{{
			Collection:      "reward",
			Key:             "daily",
			PermissionRead:  1,
			PermissionWrite: 0, // No client write.
			Value:           string(object),
			Version:         version,
			UserID:          userID,
		}})
		if err != nil {
			logger.Error("StorageWrite error: %v", err)
			return "", errInternalError
		}
	}

	out, err := json.Marshal(Coins)
	if err != nil {
		logger.Error("Marshal error: %v", err)
		return "", errMarshal
	}

	logger.Debug("rpcRewards resp: %v", string(out))
	return string(out), nil
}
