package main

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func TestClientPact_Broker(t *testing.T) {
	pact := dsl.Pact{
		Consumer:                 "ledger",
		Provider:                 "guestLedger",
		DisableToolValidityCheck: true,
	}
	port := 0000
	pactUrl := "***"
	brokerToken := "***"
	broker := "***"
	version := "***"

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		BrokerURL:                  broker,
		BrokerToken:                brokerToken,
		PactURLs:                   []string{pactUrl},
		PublishVerificationResults: true,
		ProviderVersion:            version,
	})

	if err != nil {
		t.Fatal(err)
	}
}
