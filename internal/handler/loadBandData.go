package handler

import (
	"encoding/json"
	"os"

	"example.com/phoenix-the-band/internal/model"
)

func LoadBandData(filepath string) (*model.BandData, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var bandData model.BandData
	if err := json.Unmarshal(file, &bandData); err != nil {
		return nil, err
	}

	return &bandData, nil
}
