package main

import (
	"FlexiDoc/internal/logger"
	"FlexiDoc/internal/placeholder"
	"FlexiDoc/internal/processor"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"go.uber.org/zap"
)

func main() {
	// Create a log file
	logger.CreateLogFile()
	logger.Log.Info("Application Release: " + "1.0.0.20250114")

	// Create a new parser
	parser := argparse.NewParser("FlexiDoc", "Processes a DOCX file with placeholders and a JSON file")

	currentDate := time.Now().Format("2006-01-02_15-04")
	defaultOutput := fmt.Sprintf("output_%s.docx", currentDate)
	// Add arguments
	jsonFile := parser.String("j", "json", &argparse.Options{Required: true, Help: "Path to the JSON file containing placeholders"})
	templateFile := parser.String("t", "template", &argparse.Options{Required: true, Help: "Path to the DOCX template file"})
	outputFile := parser.String("o", "output", &argparse.Options{
		Required: false,
		Default:  defaultOutput, // Default location
		Help:     "Path to save the generated DOCX file (default: current directory)",
	})

	// Parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		logger.Log.Error("Error parsing arguments", zap.Error(err))
		fmt.Println(parser.Usage(err))
		return
	}

	// Load placeholders from the JSON file
	placeholders, err := placeholder.LoadPlaceholders(*jsonFile)
	if err != nil {
		logger.Log.Error("Error loading placeholders", zap.Error(err))
		log.Fatalf("Error loading placeholders: %v", err)
	}

	// Process the DOCX file
	err = processor.ProcessDocx(*templateFile, *outputFile, placeholders)
	if err != nil {
		logger.Log.Error("Error processing DOCX file", zap.Error(err))
		log.Fatalf("Error processing DOCX file: %v", err)
	}

	logger.Log.Info("DOCX file successfully generated", zap.String("output", *outputFile))
}
