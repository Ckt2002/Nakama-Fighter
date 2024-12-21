package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type MatchState struct {
	presences map[string]runtime.Presence
}

type Match struct{}

const MaxPlayers = 2

func newMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (m runtime.Match, err error) {
	// logger.Info("Welcome to match!")
	return &Match{}, nil
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	logger.Info("Match created!")
	state := &MatchState{
		presences: make(map[string]runtime.Presence),
	}

	tickRate := 1
	label := ""

	return state, tickRate, label
}

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	// matchState := state.(*MatchState)
	// logger.Info("Player number: %v", len(matchState.presences))
	// if len(matchState.presences) >= MaxPlayers {
	// 	logger.Info("Match is full: %v", presence.GetUserId())
	// 	return state, false, "Match is full"
	// }

	// matchState.presences[presence.GetUserId()] = presence
	logger.Info("Player joined: %v", presence.GetUserId())

	return state, true, ""
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// matchState := state.(*MatchState)
	// for _, presence := range presences {
	// 	matchState.presences[presence.GetUserId()] = presence
	// 	logger.Info("Player joined: %v", presence.GetUserId())
	// }
	return state
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// matchState := state.(*MatchState)
	// for _, presence := range presences {
	// 	delete(matchState.presences, presence.GetUserId())
	// 	logger.Info("Player left: %v", presence.GetUserId())
	// }
	return state
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	return state
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return state
}

func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, "signal received: " + data
}
