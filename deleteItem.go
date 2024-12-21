package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func DeleteItem(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	if !ok {
		return "", errNoUserIdFound
	}

	// Phân tích cú pháp payload JSON
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		logger.WithField("err", err).Error("Failed to parse payload.")
		return "", err
	}

	// Lấy giá trị từ payload
	collectionsName, ok := data["collectionsName"].(string)
	if !ok {
		logger.Error("Invalid collectionsName value.")
		return "", nil
	}
	keyName, ok := data["keyName"].(string)
	if !ok {
		logger.Error("Invalid keyName value.")
		return "", nil
	}

	objectIds := []*runtime.StorageDelete{&runtime.StorageDelete{
		Collection: collectionsName,
		Key:        keyName,
		UserID:     userID,
	},
	}

	err := nk.StorageDelete(ctx, objectIds)
	if err != nil {
		return "", err
	}

	return "Item deleted succesfully", nil
}
