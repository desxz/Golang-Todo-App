package test

import (
	"fmt"
	"gunmurat7/todo-app-server/helpers"
	"gunmurat7/todo-app-server/server"
	"log"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

func startProvider(port string) {

	helpers.IsTestEnv = true

	err := server.StartServer(port)
	if err != nil {
		log.Fatal(err)
	}
}

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string // Basic authentication
	BrokerPassword  string // Basic authentication
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ProviderVersion string
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "TodoProvider"
	s.ConsumerName = "TodoConsumer"
	s.BrokerBaseURL = "http://localhost"
	s.ProviderVersion = "1.0.0"
	s.ConsumerVersion = "1.0.0"
}

func (s *Settings) getPactURL() string {
	var pactURL string

	if s.ConsumerVersion == "" {
		pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/latest/master.json", s.BrokerBaseURL, s.ProviderName, s.ConsumerName)
	} else {
		pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json", s.BrokerBaseURL, s.ProviderName, s.ConsumerName, s.ConsumerVersion)
	}

	return pactURL
}

func TestTodoPactProvider(t *testing.T) {

	port, portErr := utils.GetFreePort()
	if portErr != nil {
		log.Fatal(portErr)
	}

	go startProvider(fmt.Sprintf(":%d", port))

	settings := Settings{}
	settings.create()

	pact := dsl.Pact{
		Host:                     settings.Host,
		Provider:                 settings.ProviderName,
		Consumer:                 settings.ConsumerName,
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://%s:%d", settings.Host, port),
		ProviderVersion:            settings.ProviderVersion,
		BrokerUsername:             settings.BrokerUsername,
		BrokerURL:                  settings.BrokerBaseURL,
		BrokerPassword:             settings.BrokerPassword,
		PactURLs:                   []string{settings.getPactURL()},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(verifyResponses), "pact tests run")

}
