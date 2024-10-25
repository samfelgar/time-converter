package cmd

import (
	"fmt"
	"strconv"

	"github.com/samfelgar/time-converter/pkg"
	"github.com/spf13/cobra"
)

var fromDecimal = &cobra.Command{
	Use: "from-decimal",
	Aliases: []string{"time"},
	Short: "Get the time representation from decimal",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {		
		timeArg := args[0]
		timeAsFloat, err := strconv.ParseFloat(timeArg, 64)
	
		if err != nil {
			return fmt.Errorf("you must inform a time as decimal")
		}
	
		result, err := pkg.FromDecimal(timeAsFloat)
	
		if err != nil {
			return fmt.Errorf("unable to parse the informed time")
		}
	
		fmt.Printf("The decimal time %.2f in time is %s\n", timeAsFloat, result)
		return nil
	},
}
