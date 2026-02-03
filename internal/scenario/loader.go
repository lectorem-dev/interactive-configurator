package scenario

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(path string) (*Scenario, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read scenario file: %w", err)
	}

	var sc Scenario
	if err := json.Unmarshal(data, &sc); err != nil {
		return nil, fmt.Errorf("parse scenario json: %w", err)
	}

	if len(sc.Steps) == 0 {
		return nil, fmt.Errorf("scenario contains no steps")
	}

	return &sc, nil
}