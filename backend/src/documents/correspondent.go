package documents

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sapp/paperless-accounting/paperless"

	"github.com/redis/go-redis/v9"
)

func (m *DocumentMgr) GetCorrespondent(id int) (*paperless.PaperlessCorrespondent, error) {
	ctx := context.Background()

	val, err := m.client.ZRangeByScore(ctx, "correspondents", &redis.ZRangeBy{
		Min:   fmt.Sprint(id),
		Max:   fmt.Sprint(id),
		Count: 1,
	}).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	} else if len(val) == 0 {
		return nil, nil
	}

	var res paperless.PaperlessCorrespondent
	err = json.Unmarshal([]byte(val[0]), &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *DocumentMgr) GetCorrespondents() ([]*paperless.PaperlessCorrespondent, error) {
	ctx := context.Background()

	val, err := m.client.ZRange(ctx, "correspondents", 0, -1).Result()

	log.Printf("N: %v\n", len(val))

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	} else if len(val) == 0 {
		return nil, nil
	}

	out := make([]*paperless.PaperlessCorrespondent, len(val))
	for i, r := range val {
		var res paperless.PaperlessCorrespondent
		err = json.Unmarshal([]byte(r), &res)

		if err != nil {
			return nil, err
		}
		out[i] = &res
	}

	return out, nil
}
