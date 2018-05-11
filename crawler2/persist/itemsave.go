package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
)

func ItemSaver() chan interface{} {
	// 创建 一个channel 并且返回出去 进行接收得到的 item
	itemChan := make(chan interface{})
	// 然后开启一个 go routine 进行循环处理item
	go func() {
		for  {
			item := <-itemChan
			// 暂时先打印 看看是不是正确的
			// 开始存储
			_, err := save(item)
			if err != nil {
				log.Printf("Item saver:Error" +
					" item : %v %v" , item , err)
			}
		}
	}()
	return itemChan
}

// 传递任意类型的值
func save(item interface{})(id string , err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "" , err
	}

	response, err := client.Index().Index("profile").Type("zhenai").BodyJson(item).Do(context.Background())
	return response.Id , nil
}