// 静态、不变、高可靠数据存放于ElasticSearch
package Store

import (
	Common "Common"
	"context"
	"flag"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"strings"
	"sync"
)

var (
	flgCalc = flag.String("calcindex", "", "calc index function method")
)

// 不是很确定ES的连接安全性，所以用了锁
type ElasticSearch struct {
	sync.Mutex
	*elasticsearch.Client
}

// ES的请求回应是HTTP-REST方式
func NewElasticSearch(ip, port string) *ElasticSearch {
	defer Common.CheckPanic()

	addr := fmt.Sprintf("http://%v:%v", ip, port)
	config := elasticsearch.Config{}
	config.Addresses = []string{addr}
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	return &ElasticSearch{sync.Mutex{}, client}
}

/**
 * 创建索引
 * @param  {[type]} this *             ElasticSearch) CreateIndex(idx string, body string [description]
 * @return {[type]}      [description]
 */
func (this *ElasticSearch) CreateIndex(idx string, body string) {
	req := esapi.IndicesCreateRequest{
		Index: idx,
		Body:  strings.NewReader(body),
	}

	resp, _ := req.Do(context.Background(), this.Client)
	return this.ReadResp(resp)
}

/**
 * 插入文档
 * @param  {[type]} this *ElasticSearch) InsertDoc(idx string, docid string, body string [description]
 * @return {[type]}      [description]
 */
func (this *ElasticSearch) InsertDoc(idx string, docid string, body string) {
	req := esapi.CreateRequest{
		Index:        idx,
		DocumentType: "doc",
		DocumentID:   docid,
		Body:         strings.NewReader(body),
	}
	resp, _ := req.Do(context.Background(), this.Client)
	return this.ReadResp(resp)
}

/**
 * 删除文档
 */
func (this *ElasticSearch) DeleteDoc(idx string, docid string) {
	req := esapi.DeleteRequest{
		Index:        idx,
		DocumentType: "doc",
		DocumentID:   docid,
	}

	resp, _ := req.Do(context.Background(), this.Client)
	return this.ReadResp(resp)
}

/**
 * ElasticSearch 搜索
 * @param  {[type]} this *ElasticSearch) Search(index string, body string [description]
 * @return {[type]}      [description]
 */
func (this *ElasticSearch) Search(index string, body string) ([]interface{}, error) {
	req := esapi.SearchRequest{
		Index:        []string{index},
		DocumentType: []string{"doc"},
		Body:         strings.NewReader(body),
	}
	resp, err := req.Do(context.Background(), this.Client)
	if err != nil {
		fmt.Println(err)
	}

	simpleJson, _ := this.ReadResp(resp)

	return simpleJson.Get("hits").Get("hits").Array()
}

func (this *ElasticSearch) Count(index string, body string) {
	req := esapi.SearchRequest{
		Index:        []string{index},
		DocumentType: []string{"doc"},
		Body:         strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), this.Client)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	simpleJson, _ := this.ReadResp(res)

	fmt.Println(simpleJson.Get("total"))
	// resp, _ := ioutil.ReadAll(res.Body)

	// jsonObj, err := simplejson.NewJson(resp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(jsonObj)

}

/**
 * 读取ElasticSearch 返回
 * @param  {[type]} this *ElasticSearch) ReadResp(resp *esapi.Response) (*simplejson.Json, error [description]
 * @return {[type]}      [description]
 */
func (this *ElasticSearch) ReadResp(resp *esapi.Response) (*simplejson.Json, error) {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return simplejson.NewJson(body)
}
