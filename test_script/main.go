package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/google/uuid"
	"github.com/hashicorp/go-retryablehttp"
	co "github.com/metal-toolbox/conditionorc/pkg/api/v1/client"
	coapiv1 "github.com/metal-toolbox/conditionorc/pkg/api/v1/types"
	rctypes "github.com/metal-toolbox/rivets/condition"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/oauth2/clientcredentials"
)

// returns a conditionorc retryable http client with Otel and Oauth wrapped in
func newConditionorcClientWithOAuthOtel(ctx context.Context) (*co.Client, error) {
	// init retryable http client
	retryableClient := retryablehttp.NewClient()

	// set retryable HTTP client to be the otel http client to collect telemetry
	retryableClient.HTTPClient = otelhttp.DefaultClient

	// setup oidc provider
	provider, err := oidc.NewProvider(ctx, "https://hydra.iam.equinixmetal.net/")
	if err != nil {
		return nil, err
	}

	// setup oauth configuration
	token := "faxxxxxxxxxxxx"
	oauthConfig := clientcredentials.Config{
		ClientID:       "fleetdb-migration",
		ClientSecret:   token,
		TokenURL:       provider.Endpoint().TokenURL,
		Scopes:         []string{"create:condition", "create:server"},
		EndpointParams: url.Values{"audience": []string{"https://hollow.equinixmetal.net"}},
	}

	// wrap OAuth transport, cookie jar in the retryable client
	oAuthclient := oauthConfig.Client(ctx)

	retryableClient.HTTPClient.Transport = oAuthclient.Transport
	retryableClient.HTTPClient.Jar = oAuthclient.Jar

	// requests taking longer than timeout value should be canceled.
	client := retryableClient.StandardClient()
	client.Timeout = 600 * time.Second

	return co.NewClient(
		"https://serverconditions.equinixmetal.net",
		co.WithAuthToken(token),
		co.WithHTTPClient(client),
	)
}

func main() {
	client, err := newConditionorcClientWithOAuthOtel(context.Background())
	if err != nil {
		fmt.Printf("failed to connect to conditionorc %v", err)
		return
	}

	serverID, err := uuid.Parse("c8c5263b-1980-44bf-bc84-6fc25f33c75b")
	if err != nil {
		log.Fatal(err)
	}

	params, err := json.Marshal(rctypes.NewInventoryTaskParameters(
		serverID,
		rctypes.OutofbandInventory,
		true,
		true,
	))
	if err != nil {
		log.Fatal(err)
	}

	conditionCreate := coapiv1.ConditionCreate{
		Parameters: params,
	}

	_, err = client.ServerConditionCreate(context.Background(), serverID, rctypes.Inventory, conditionCreate)
	if err != nil {
		fmt.Printf("failed to enroll server %v\n", err)
		return
	}
}
