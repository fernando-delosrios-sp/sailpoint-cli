// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.
package root

import (
	"fmt"

	"github.com/sailpoint-oss/sailpoint-cli/cmd/connector"
	"github.com/sailpoint-oss/sailpoint-cli/cmd/transform"
	"github.com/sailpoint-oss/sailpoint-cli/cmd/va"
	"github.com/sailpoint-oss/sailpoint-cli/internal/client"
	"github.com/spf13/cobra"
)

var version = "0.4.0"

func NewRootCmd(client client.Client) *cobra.Command {
	root := &cobra.Command{
		Use:          "sail",
		Short:        "The SailPoint CLI allows you to administer your IdentityNow tenant from the command line.\n\nNavigate to developer.sailpoint.com to learn more.",
		Version:      version,
		SilenceUsage: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprint(cmd.OutOrStdout(), cmd.UsageString())
		},
	}
	root.AddCommand(
		newConfigureCmd(client),
		newDebugCommand(),
		newAuthCommand(),
		connector.NewConnCmd(client),
		transform.NewTransformCmd(client),
		va.NewVACmd(),
	)
	return root
}
