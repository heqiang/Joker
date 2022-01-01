package initlize

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"jokerweb/config"
	"jokerweb/global"
)

const JokerMapping = `
{
	"mappings": {
      "properties": {
        "articlid":{
          "type":"integer"
        },
        "domain":{
          "type":"keyword"
        },
        "url":{
          "type":"keyword"
        },
        "title":{
          "type":"text",
        },
        "content":{
          "type":"text"
        },
        "pubtime":{
          "type":"date"
        },
        "ClickNum":{
          "type":"integer"
        },
        "category":{
          "type":"keyword"
        }
      }
    }
}`

func InitEs(conf *config.EsConfig) error {
	//ctx := context.Background()
	url := fmt.Sprintf("http://%s:%d", conf.Host, conf.Port)
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		return err
	}
	//检测 index是否存在
	//exist, err := client.IndexExists("joker").Do(ctx)
	//if !exist {
	//	_, err := client.CreateIndex("joker").BodyString(JokerMapping).Do(ctx)
	//	if err != nil {
	//		panic(err)
	//		return err
	//	}
	//}
	global.Es = client
	return nil
}
