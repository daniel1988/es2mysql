// 静态、不变、高可靠数据存放于ElasticSearch
package Store

import (
    Common "Common"
    "encoding/json"
    "flag"
    "fmt"
    es "github.com/mattbaird/elastigo/lib"
    "strings"
    "sync"
)

var (
    flgCalc = flag.String("calcindex", "", "calc index function method")
)

// 不是很确定ES的连接安全性，所以用了锁
type ElasticSearch struct {
    sync.Mutex
    *es.Conn
}

// ES的请求回应是HTTP-REST方式
func NewStaticStore(ip, port string) *ElasticSearch {
    defer Common.CheckPanic()
    c := es.NewConn()
    c.Domain = ip
    c.Port = port

    return &ElasticSearch{sync.Mutex{}, c}
}

// 关闭ES时刷入缓存
func (this *ElasticSearch) Close() {
    defer Common.CheckPanic()
    this.Flush()
}

// 索引生成，在需要优化索引性能时添加新方法
func (this *ElasticSearch) CalcIndex(index, _type, id string) string {
    switch *flgCalc {
    case "index-id1":
        return strings.ToLower(fmt.Sprintf("%s-%v", index, id[:1]))
    }

    return strings.ToLower(index)
}

// 添加新的内容
func (this *ElasticSearch) InsertDoc(index, _type, id string, ttl int, v interface{}) error {

    defer Common.CheckPanic()
    this.Lock()
    defer this.Unlock()

    var args map[string]interface{} = nil
    if ttl > 0 {
        args = map[string]interface{}{"ttl": ttl}
    }

    _, err := this.Index(this.CalcIndex(index, _type, id), _type, id, args, v)
    if err != nil {
        return err
    }

    return nil
}

// 修改内容
func (this *ElasticSearch) UpdateDoc(index, _type, id string, v map[string]interface{}) error {
    defer Common.CheckPanic()
    this.Lock()
    defer this.Unlock()

    doc := map[string]map[string]interface{}{"doc": v}
    _, err := this.Update(this.CalcIndex(index, _type, id), _type, id, nil, doc)
    if err != nil {
        return err
    }

    return nil
}

// 查找内容
func (this *ElasticSearch) GetDoc(index, _type, id string, doc interface{}) error {
    data, err := this.GetDocSource(index, _type, id)
    if err != nil {
        return err
    }

    if err = json.Unmarshal(data, doc); err != nil {
        return err
    }

    return nil
}

// 查找内容的源数据
func (this *ElasticSearch) GetDocSource(index, _type, id string) ([]byte, error) {
    defer Common.CheckPanic()
    this.Lock()
    defer this.Unlock()

    rsp, err := this.Get(this.CalcIndex(index, _type, id), _type, id, nil)
    if err != nil {
        return nil, err
    }

    return rsp.Source.MarshalJSON()
}

// 删除一条记录
func (this *ElasticSearch) DeleteDoc(index, _type, id string) error {
    defer Common.CheckPanic()
    this.Lock()
    defer this.Unlock()

    _, err := this.Delete(this.CalcIndex(index, _type, id), _type, id, nil)
    if err != nil {
        return err
    }

    return nil
}
