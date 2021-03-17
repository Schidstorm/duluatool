package cli

import (
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

	helpCommand := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Help()
		},
		Use: "help",
	}

	decodeCommand := &cobra.Command{
		Use: "decode",
		RunE: func(cmd *cobra.Command, args []string) error {
			inputFilePath, err := cmd.PersistentFlags().GetString("file")
			if err != nil {
				return err
			}

			outputDirectory, err := cmd.PersistentFlags().GetString("dir")
			if err != nil {
				return err
			}

			return decoder.Run(&decoder.Options{
				InputFilePath:   inputFilePath,
				OutputDirectory: outputDirectory,
			})
		},
	}

	decodeCommand.PersistentFlags().String("file", path.Join(workingDir, "code.json"), "Input file name in json format.")
	decodeCommand.PersistentFlags().String("dir", workingDir, "Output directory where to put the file structure.")
	decodeCommand.AddCommand(helpCommand)

	encodeCommand := &cobra.Command{
		Use: "encode",
		RunE: func(cmd *cobra.Command, args []string) error {
			outputFilePath, err := cmd.PersistentFlags().GetString("file")
			if err != nil {
				return err
			}

			inputDirectory, err := cmd.PersistentFlags().GetString("dir")
			if err != nil {
				return err
			}

			return encoder.Run(&encoder.Options{
				OutputFilePath: outputFilePath,
				InputDirectory: inputDirectory,
			})
		},
	}
	encodeCommand.PersistentFlags().String("file", path.Join(workingDir, "code.json"), "Output file name.")
	encodeCommand.PersistentFlags().String("dir", workingDir, "Input directory.")
	encodeCommand.AddCommand(helpCommand)

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(helpCommand)
	rootCmd.AddCommand(decodeCommand)
	rootCmd.AddCommand(encodeCommand)

	err = rootCmd.Execute()

	if err != nil {
		logrus.Fatal(err)
	}
}
