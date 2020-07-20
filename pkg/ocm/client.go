package ocm

import (
	"fmt"

	sdkClient "github.com/openshift-online/ocm-sdk-go"
)

type Client struct {
	config     *Config
	logger     sdkClient.Logger
	connection *sdkClient.Connection

	Authentication OCMAuthentication
}

type Config struct {
	BaseURL      string `envconfig:"OCM_BASE_URL" default:"https://api-integration.6943.hive-integration.openshiftapps.com"`
	ClientID     string `envconfig:"OCM_CLIENT_ID" default:""`
	ClientSecret string `envconfig:"OCM_CLIENT_SECRET" default:""`
	SelfToken    string `envconfig:"OCM_SELF_TOKEN" default:""`
	TokenURL     string `envconfig:"OCM_TOKEN_URL" default:"https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"`
	Debug        bool   `envconfig:"OCM_DEBUG" default:"false"`
}

func NewClient(config Config) (*Client, error) {
	// Create a logger that has the debug level enabled:
	logger, err := sdkClient.NewGoLoggerBuilder().
		Debug(config.Debug).
		Build()
	if err != nil {
		return nil, fmt.Errorf("Unable to build OCM logger: %s", err.Error())
	}

	client := &Client{
		config: &config,
		logger: logger,
	}
	err = client.newConnection()
	if err != nil {
		return nil, fmt.Errorf("Unable to build OCM connection: %s", err.Error())
	}
	client.Authentication = &authentication{client: client}
	return client, nil
}

func (c *Client) newConnection() error {
	builder := sdkClient.NewConnectionBuilder().
		Logger(c.logger).
		URL(c.config.BaseURL).
		Metrics("api_outbound")

	if c.config.ClientID != "" && c.config.ClientSecret != "" {
		builder = builder.Client(c.config.ClientID, c.config.ClientSecret)
	} else if c.config.SelfToken != "" {
		builder = builder.Tokens(c.config.SelfToken)
	} else {
		return fmt.Errorf("Can't build OCM client connection. No Client/Secret or Token has been provided")
	}

	connection, err := builder.Build()

	if err != nil {
		return fmt.Errorf("Can't build OCM client connection: %s", err.Error())
	}
	c.connection = connection
	return nil
}

func (c *Client) Close() {
	if c.connection != nil {
		c.connection.Close()
	}
}

type service struct {
	client *Client
}
