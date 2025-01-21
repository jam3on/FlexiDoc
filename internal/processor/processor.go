package processor

import (
	"fmt"
	"strings"

	"FlexiDoc/internal/placeholder"
	"FlexiDoc/internal/logger"

	"github.com/nguyenthenguyen/docx"
	"go.uber.org/zap"
)

// Replace placeholders in a DOCX file and save the result
func ProcessDocx(inputFile, outputFile string, placeholders placeholder.Placeholders) error {
	// Load the DOCX file
	r, err := docx.ReadDocxFile(inputFile)
	if err != nil {
		logger.Log.Error("Error opening DOCX file", zap.Error(err))
		return fmt.Errorf("error opening DOCX file: %v", err)
	}
	defer r.Close()

	// Read the document content
	doc := r.Editable()
	content := doc.GetContent()

	// Replace placeholders
	for key, value := range placeholders {
		placeholder := fmt.Sprintf("{{%s}}", key)
		content = strings.ReplaceAll(content, placeholder, value)
	}

	// Apply the changes
	doc.SetContent(content)

	// Save the modified DOCX file
	err = doc.WriteToFile(outputFile)
	if err != nil {
		logger.Log.Error("Error saving DOCX file", zap.Error(err))
		return fmt.Errorf("error saving DOCX file: %v", err)
	}

	return nil
}
