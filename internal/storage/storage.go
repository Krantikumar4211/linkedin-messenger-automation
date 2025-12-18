package storage

import (
	"encoding/json"
	"os"
)

type State struct {
	SentMessages []string `json:"sent_messages"`
}

func LoadState(path string) State {
	var state State

	file, err := os.ReadFile(path)
	if err != nil {
		return State{SentMessages: []string{}}
	}

	_ = json.Unmarshal(file, &state)

	// Ensure slice is initialized
	if state.SentMessages == nil {
		state.SentMessages = []string{}
	}

	return state
}

func SaveState(path string, state State) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
