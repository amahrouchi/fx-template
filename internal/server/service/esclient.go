package service

import (
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/elastic/go-elasticsearch/v8"
)

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
