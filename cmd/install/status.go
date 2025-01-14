package install

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	mctl "github.com/metal-toolbox/mctl/cmd"
	"github.com/metal-toolbox/mctl/internal/app"
	rctypes "github.com/metal-toolbox/rivets/condition"
	"github.com/spf13/cobra"
)

var serverIDStr string

var installStatus = &cobra.Command{
	Use:   "status --server | -s <server uuid>",
	Short: "check the progress of a firmware install on a server",
	Run: func(cmd *cobra.Command, _ []string) {
		statusCheck(cmd.Context())
	},
}

func statusCheck(ctx context.Context) {
	theApp := mctl.MustCreateApp(ctx)

	client, err := app.NewConditionsClient(ctx, theApp.Config.Conditions, theApp.Reauth)
	if err != nil {
		log.Fatal(err)
	}

	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		log.Fatalf("parsing server id: %s", err.Error())
	}

	resp, err := client.ServerConditionStatus(ctx, serverID)
	if err != nil {
		log.Fatalf("querying server conditions: %s", err.Error())
	}

	s, err := mctl.FormatConditionResponse(resp, rctypes.FirmwareInstall)
	if err != nil {
		log.Fatalf("condition response error: %s", err.Error())
	}

	fmt.Println(s)
}

func init() {
	flags := installStatus.Flags()
	flags.StringVarP(&serverIDStr, "server", "s", "", "server id (typically a UUID)")

	if err := installStatus.MarkFlagRequired("server"); err != nil {
		log.Fatalf("marking server flag as required: %s", err.Error())
	}

	install.AddCommand(installStatus)
}
