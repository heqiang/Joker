package initlize

import (
	"context"
	"github.com/olivere/elastic/v7"
	"spider/global"
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
          "analyzer":"ik_max_word"
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

func InitEs() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	//检测 index是否存在
	exist, err := client.IndexExists("joker").Do(ctx)
	if !exist {
		_, err := client.CreateIndex("joker").BodyString(JokerMapping).Do(ctx)
		if err != nil {
			panic(err)
			return
		}

	}
	global.Es = client

}
