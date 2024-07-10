package elasticService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/mget"
	"net/http"
)

// ESClient service getting data from ElasticSearch.
type ESClient struct {
	client *elasticsearch.TypedClient
	prefix string
	config *fxconfig.Config
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

// Mget gets many documents from ElasticSearch.
func (es *ESClient) Mget(indexType string, ids []string) (any, error) {
	index := es.getIndexName(indexType)

	// Query ElasticSearch
	response, err := es.client.Mget().
		Index(index).
		Request(&mget.Request{Ids: ids}).
		Do(context.TODO())

	// Check for errors
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return es.decodeResponse(response)
}

// getIndexName gets the index name for a given index type.
func (es *ESClient) getIndexName(indexType string) string {
	return es.config.GetString("config.elastic.prefix") + indexType
}

// decodeResponse decode response from typed ES client.
func (es *ESClient) decodeResponse(response *http.Response) (any, error) {
	// Check response status
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("ElasticSearch responded with a bad status code: %s.", response.Status))
	}

	// Decode response
	var result map[string]any
	err := json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
