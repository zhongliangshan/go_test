package persist

import "github.com/zhongliangshan/src/gopkg.in/olivere/elastic.v5"

type ItemService struct {
	Client elastic.Client
	Index string
}


func (i *ItemService) Save() {
	i.Client.Index().Index(i.Index).Type()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "" , err
	}

	response, err := client.Index().Index("profile").Type("zhenai").BodyJson(item).Do(context.Background())
	return response.Id , nil
}