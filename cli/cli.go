package cli

import (
	"errors"
	"github.com/atotto/clipboard"
	"github.com/schidstorm/duluatool/decoder"
	"github.com/schidstorm/duluatool/encoder"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func Run() {
	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = ""
	}

	decodeCommand := &cobra.Command{
		Use: "decode",
		RunE: func(cmd *cobra.Command, args []string) error {
			inputClipboard, errClipboard := cmd.PersistentFlags().GetBool("clipboard")
			inputFilePath, errFile := cmd.PersistentFlags().GetString("file")
			if errFile != nil && errClipboard != nil {
				return err
			}

			if inputClipboard && clipboard.Unsupported {
				return errors.New("clipboard is not supported")
			}

			outputDirectory, err := cmd.PersistentFlags().GetString("dir")
			if err != nil {
				return err
			}

			return decoder.Run(&decoder.Options{
				InputFilePath:   inputFilePath,
				InputClipboard:  inputClipboard,
				OutputDirectory: outputDirectory,
			})
		},
	}

	decodeCommand.PersistentFlags().String("file", path.Join(workingDir, "code.json"), "Input file name in json format.")
	decodeCommand.PersistentFlags().String("dir", workingDir, "Output directory where to put the file structure.")
	decodeCommand.PersistentFlags().Bool("clipboard", false, "Use clipboard as input file.")

	encodeCommand := &cobra.Command{
		Use: "encode",
		RunE: func(cmd *cobra.Command, args []string) error {
			outputFilePath, errClipboard := cmd.PersistentFlags().GetString("file")
			outputClipboard, errFile := cmd.PersistentFlags().GetBool("clipboard")
			if errFile != nil && errClipboard != nil {
				return err
			}

			if outputClipboard && clipboard.Unsupported {
				return errors.New("clipboard is not supported")
			}

			inputDirectory, err := cmd.PersistentFlags().GetString("dir")
			if err != nil {
				return err
			}

			return encoder.Run(&encoder.Options{
				OutputFilePath:  outputFilePath,
				OutputClipboard: outputClipboard,
				InputDirectory:  inputDirectory,
			})
		},
	}
	encodeCommand.PersistentFlags().String("file", path.Join(workingDir, "code.json"), "Output file name.")
	encodeCommand.PersistentFlags().String("dir", workingDir, "Input directory.")
	encodeCommand.PersistentFlags().Bool("clipboard", false, "Use clipboard as output file.")

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(decodeCommand)
	rootCmd.AddCommand(encodeCommand)

	err = rootCmd.Execute()

	if err != nil {
		logrus.Fatal(err)
	}
}
