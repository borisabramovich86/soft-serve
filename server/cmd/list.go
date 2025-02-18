package cmd

import (
	"github.com/charmbracelet/soft-serve/server/backend"
	"github.com/spf13/cobra"
)

// listCommand returns a command that list file or directory at path.
func listCommand() *cobra.Command {
	var all bool

	listCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List repositories",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, s := fromContext(cmd)
			repos, err := cfg.Backend.Repositories()
			if err != nil {
				return err
			}
			for _, r := range repos {
				if cfg.Backend.AccessLevelByPublicKey(r.Name(), s.PublicKey()) >= backend.ReadOnlyAccess {
					if !r.IsHidden() || all {
						cmd.Println(r.Name())
					}
				}
			}
			return nil
		},
	}

	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List all repositories")

	return listCmd
}
