package cmd

import (
	"io"

	"github.com/majestic-fox/swiss-army-knife/pkg/aws"
	"github.com/spf13/cobra"
)

type awsSetCredentialsCmd struct {
	profile bool
	out     io.Writer
}

func newAwsSetCredentialsCmd(out io.Writer) *cobra.Command {
	a := &awsSetCredentialsCmd{out: out}
	cmd := &cobra.Command{
		Use:     "credentials",
		Aliases: []string{"creds"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := a.run(); err != nil {
				return err
			}
			return nil
		},
	}
	p := cmd.Flags()
	p.BoolVarP(&a.profile, "profile", "p", false, "set just the AWS_PROFILE variable")
	return cmd
}

func (a *awsSetCredentialsCmd) run() error {
	err := aws.SetCredentials(a.profile)
	if err != nil && !settings.Quiet {
		return err
	}
	return nil
}
