package cmd

import (
	"fmt"
	"regexp"

	"github.com/samfelgar/time-converter/pkg"
	"github.com/spf13/cobra"
)

var formatExample = "(10h24m)"

var toDecimal = &cobra.Command{
	Use: "to-decimal",
	Short: fmt.Sprintf("Get the decimal representation from a time string %s", formatExample),
	Args: cobra.ExactArgs(1),
	Aliases: []string{"dec"},
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := regexp.Compile(`^\d+h\d+m(?:\d+s)?$`)

		if err != nil {
			return fmt.Errorf("invalid regex")
		}

		if !exp.MatchString(args[0]) {
			return fmt.Errorf("invalid time format: %s. You must follow the %s format", args[0], formatExample)
		}

		result, err := pkg.ToDecimal(args[0])

		if err != nil {
			return fmt.Errorf("unable to convert to decimal: %s", err.Error())
		}

		cmd.Printf("The time %s converted to decimal is %.2f\n", args[0], result)
		return nil
	},
}
