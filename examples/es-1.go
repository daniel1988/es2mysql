package main

import (
	"Store"
	"fmt"
)

func main() {

	client := Store.NewElasticSearch("10.105.249.250", "9200")
	body := `{
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
	hits, _ := client.Search("dc_acing_u8_2020", body)
	fmt.Println(hits)
	// fmt.Println("count") //
	// client.Count("dc_acing_u8_2020", body)

}
