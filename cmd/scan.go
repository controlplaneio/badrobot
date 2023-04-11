package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/controlplaneio/badrobot/pkg/report"
	"github.com/controlplaneio/badrobot/pkg/ruler"
	"github.com/spf13/cobra"
)

type ScanFailedValidationError struct {
	ExitCode int
}

func (e *ScanFailedValidationError) Error() string {
	return "BadRobot scan failed"
}

var debug bool
var absolutePath bool
var format string
var template string
var schemaDir string
var outputLocation string
var exitCode int

func init() {
	scanCmd.Flags().BoolVar(&debug, "debug", false, "turn on debug logs")
	scanCmd.Flags().BoolVar(&absolutePath, "absolute-path", false, "use the absolute path for the file name")
	scanCmd.Flags().StringVarP(&format, "format", "f", "json", "Set output format (json, template)")
	scanCmd.Flags().StringVar(&schemaDir, "schema-dir", "", "Sets the directory for the json schemas")
	scanCmd.Flags().StringVarP(&template, "template", "t", "", "Set output template, it will check for a file or read input as the")
	scanCmd.Flags().StringVarP(&outputLocation, "output", "o", "", "Set output location")
	scanCmd.Flags().IntVar(&exitCode, "exit-code", 2, "Set the exit-code to use on failure")
	rootCmd.AddCommand(scanCmd)
}

// File holds the name and contents
type File struct {
	fileName  string
	fileBytes []byte
}

func getInput(args []string) (File, error) {
	var file File

	if len(args) == 1 && (args[0] == "-" || args[0] == "/dev/stdin") {
		fileBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return file, err
		}
		file = File{
			fileName:  "STDIN",
			fileBytes: fileBytes,
		}
		return file, nil
	}
	fileName := args[0]
	filePath, err := filepath.Abs(fileName)
	if err != nil {
		return file, err
	}
	if absolutePath {
		fileName = filePath
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return file, err
	}
	file = File{
		fileName:  fileName,
		fileBytes: fileBytes,
	}
	return file, nil
}

var scanCmd = &cobra.Command{
	Use:     `scan [file]`,
	Short:   "Scans Kubernetes Operator resource YAML or JSON",
	Example: `  badrobot scan ./operator.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("file path is required")
		}

		if debug {
			z, err := NewLogger("debug", "console")
			if err != nil {
				return fmt.Errorf("can't initialize zap logger: %v", err)
			}
			logger = z
		}

		rootCmd.SilenceErrors = true
		rootCmd.SilenceUsage = true

		file, err := getInput(args)
		if err != nil {
			return err
		}

		reports, err := ruler.NewRuleset(logger).Run(file.fileName, file.fileBytes, schemaDir)
		if err != nil {
			return err
		}

		if len(reports) == 0 {
			return fmt.Errorf("invalid input %s", file.fileName)
		}

		var lowScore bool
		for _, r := range reports {
			if r.Score <= 0 {
				lowScore = true
				break
			}
		}

		var buff bytes.Buffer
		err = report.WriteReports(format, &buff, reports, template)
		if err != nil {
			return err
		}

		if outputLocation != "" {
			err = ioutil.WriteFile(outputLocation, buff.Bytes(), 0644)
			if err != nil {
				return fmt.Errorf("writing output to %s, error: %s", outputLocation, err)
			}
		}

		out := buff.String()
		fmt.Println(out)

		if len(reports) > 0 && !lowScore {
			return nil
		}

		return &ScanFailedValidationError{ExitCode: exitCode}
	},
}
