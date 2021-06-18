package main

import (
	"github.com/spf13/cobra",
	"os",
	"fmt"
)

func main() {
	matiTokensCmd := &cobra.Command{
		Use: "matiTokens",
		Short: "Mati tokens CLI",
		Run: func(cmd *cobra.Command, args[]string) {

		}
	}

	err := matiTokensCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}



