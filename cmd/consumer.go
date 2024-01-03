package cmd

import "github.com/spf13/cobra"

type Consumer struct{}

func (c Consumer) Command() *cobra.Command {
	return nil
}

func (c Consumer) main() {

}
