package placeholder

import (
	"FlexiDoc/internal/logger"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

// Map for replacing tags in templates
type Placeholders map[string]string

// Reads JSON file containing placeholders and returns a Placeholders map
func LoadPlaceholders(filename string) (Placeholders, error) {
	file, err := os.Open(filename)
	if err != nil {
		logger.Log.Error("Error opening JSON file", zap.String("filename", filename), zap.Error(err))
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			logger.Log.Error("Error closing JSON file", zap.String("filename", filename), zap.Error(cerr))
		}
	}()

	var placeholders Placeholders
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&placeholders); err != nil {
		logger.Log.Error("Error decoding JSON file", zap.String("filename", filename), zap.Error(err))
		return nil, err
	}

	return placeholders, nil
}
