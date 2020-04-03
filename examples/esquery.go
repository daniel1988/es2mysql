package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"os"
)

var c *elasticsearch.Client

func init() {
	var err error
	config := elasticsearch.Config{}
	config.Addresses = []string{"http://10.105.249.250:9200"}
	c, err = elasticsearch.NewClient(config)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	selectBySearch()
}

func selectBySearch() {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"logtime": map[string]interface{}{
							"gt": 0,
						},
					},
				},
			},
		},
		"size": 1,
	}

	jsonStr := `{
        "query": {
            "bool": {
                "filter": {
                    "range": {
                        "logtime": {
                            "gt": 0
                        }
                    }
                }
            }
        },
        "size": 10
    }`
	jsonBody, _ := json.Marshal(query)

	jsonBody = []byte(jsonStr)

	req := esapi.SearchRequest{
		Index:        []string{"dc_acing_u8_2020"},
		DocumentType: []string{"doc"},
		Body:         bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), c)
	checkError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	ReadBody(body)
	return
	var jsonObj map[string]interface{}
	json.Unmarshal(body, &jsonObj)

	hits := jsonObj["hits"].(map[string]interface{})["hits"].([]interface{})

	for _, row := range hits {
		fmt.Println(row)
	}
	// fmt.Println(jsonObj["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"])
}

func ReadBody(body []byte) {
	res, err := simplejson.NewJson(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := res.Get("hits").Get("hits").Array()
	for _, row := range rows {
		fmt.Println(row.(map[string]interface{})["_source"].(map[string]interface{})["docid"])
	}
}
