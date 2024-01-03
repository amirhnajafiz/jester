package cmd

import "github.com/spf13/cobra"

type Publisher struct{}

func (p Publisher) Command() *cobra.Command {
	return nil
}

func (p Publisher) main() {

}
