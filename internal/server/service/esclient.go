package service

import (
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/elastic/go-elasticsearch/v8"
)

// ESClientContract is the interface for the ESClient services.
type ESClientContract interface {
	Ping() (bool, error)
}

// NewESClient Creates a new ESClientContract service.
func NewESClient(
	config *fxconfig.Config,
	logger *fxlogger.Logger,
) ESClientContract {
	// Get config
	host := config.GetString("config.elastic.host")
	user := config.GetString("config.elastic.user")
	password := config.GetString("config.elastic.password")
	prefix := config.GetString("config.elastic.prefix")

	// Check config
	if host == "" || user == "" || password == "" || prefix == "" {
		logger.Fatal().
			Str("host", host).
			Str("user", user).
			Str("password", password).
			Str("prefix", prefix).
			Msg("The ElasticSearch config is invalid.")
	}

	// Build the ES client
	cfg := elasticsearch.Config{
		Addresses: []string{host},
		Username:  user,
		Password:  password,
	}
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}

	return &ESClient{
		client: client,
		prefix: prefix,
		logger: logger,
	}
}

// ESClient service getting data from ElasticSearch.
type ESClient struct {
	client *elasticsearch.TypedClient
	prefix string
	logger *fxlogger.Logger
}

// Ping pings the ElasticSearch server.
func (es *ESClient) Ping() (bool, error) {
	res, err := es.client.Ping().IsSuccess(nil)
	if err != nil {
		es.logger.Err(err).Msg("Error pinging ElasticSearch")
		return false, err
	}

	return res, nil
}
