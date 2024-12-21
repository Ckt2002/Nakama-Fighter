package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func UpdateWallet(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	userID := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	// ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	// Phân tích cú pháp payload JSON
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		logger.WithField("err", err).Error("Failed to parse payload.")
		return "", err
	}

	// Lấy giá trị từ payload
	coin, ok := data["coin"].(float64)
	if !ok {
		logger.Error("Invalid coin value.")
		return "", nil
	}

	action, ok := data["action"].(string)
	if !ok {
		logger.Error("Invalid action value.")
		return "", nil
	}

	changeset := map[string]int64{
		"coins": int64(coin),
	}
	metadata := map[string]interface{}{
		"action": action,
	}
	_, _, err := nk.WalletUpdate(ctx, userID, changeset, metadata, true)
	if err != nil {
		logger.WithField("err", err).Error("Wallet update error.")
	}

	return "Wallet updated successfully", nil
}
