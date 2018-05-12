package persist

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"github.com/zhongliangshan/test/crawler2/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemService struct {
	Client *elastic.Client
	Index  string
}

func (i *ItemService) Saver(item engine.Item, result *string) error {
	err := persist.Save(i.Client, i.Index, item)
	if err != nil {
		return err
	}

	*result = "ok"
	return nil
}
