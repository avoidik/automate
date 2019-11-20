package main

import (
	"github.com/spf13/cobra"

	"github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/api/config/gateway"
	"github.com/chef/automate/components/automate-deployment/pkg/client"
)

func init() {
	appsSubcmd := newApplicationsRootSubcmd()

	appsSubcmd.AddCommand(newApplicationsEnableCmd())
	appsSubcmd.AddCommand(newApplicationsDisableCmd())

	RootCmd.AddCommand(appsSubcmd)
}

func newApplicationsRootSubcmd() *cobra.Command {
	return &cobra.Command{
		Use:    "applications COMMAND",
		Short:  "Manage applications visibility features",
		Hidden: true,
	}
}

func newApplicationsEnableCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "enable",
		Short: "Enable applications visibility features",
		RunE:  runApplicationsEnableCmd,
		Args:  cobra.NoArgs,
	}
}

func newApplicationsDisableCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disable",
		Short: "Disable applications visibility features",
		RunE:  runApplicationsDisableCmd,
		Args:  cobra.NoArgs,
	}
}

func runApplicationsEnableCmd(*cobra.Command, []string) error {
	cfg := newApplicationsToggleConfig(true)
	if err := client.PatchAutomateConfig(configCmdFlags.timeout, cfg, writer); err != nil {
		return err
	}
	return nil
}

func runApplicationsDisableCmd(*cobra.Command, []string) error {
	cfg := newApplicationsToggleConfig(false)
	if err := client.PatchAutomateConfig(configCmdFlags.timeout, cfg, writer); err != nil {
		return err
	}
	return nil
}

// newApplicationsToggleConfig creates an AutomateConfig data structure
// equivalent to the following Automate config toml, with `$ENABLE` set
// according to the `enable` argument:
// # Gateway service configuration.
// [gateway.v1]
//   [gateway.v1.sys]
//     [gateway.v1.sys.service]
//       enable_apps_feature = $ENABLE
// # event-service configuration.
// [event_service.v1]
//   [event_service.v1.sys]
//     [event_service.v1.sys.service]
//       enable_nats_feature = $ENABLE
// # event-gateway
// [event_gateway]
//   [event_gateway.v1]
//     [event_gateway.v1.sys]
//       [event_gateway.v1.sys.service]
//         enable_nats_feature = $ENABLE
// # applications-service
// [applications]
//   [applications.v1]
//     [applications.v1.sys]
//       [applications.v1.sys.service]
//         enable_nats_feature = $ENABLE
// FIXME: remove
func newApplicationsToggleConfig(enable bool) *deployment.AutomateConfig {
	return &deployment.AutomateConfig{
		Gateway: &gateway.ConfigRequest{
			V1: &gateway.ConfigRequest_V1{
				Sys: &gateway.ConfigRequest_V1_System{
					Service: &gateway.ConfigRequest_V1_System_Service{},
				},
			},
		},
	}
}
