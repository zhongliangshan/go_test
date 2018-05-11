package persist

import "github.com/zhongliangshan/src/gopkg.in/olivere/elastic.v5"

type ItemService struct {

}


func Save() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "" , err
	}

	response, err := client.Index().Index("profile").Type("zhenai").BodyJson(item).Do(context.Background())
	return response.Id , nil
}