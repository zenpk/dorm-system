package cache

import "context"

type Dorm struct{}

func (d *Dorm) GetRemainNum(ctx context.Context) (map[string]string, error) {
	res, err := Redis.HGetAll(ctx, "available").Result()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		if err := warmRemainCnt(); err != nil {
			return nil, err
		}
	}
	res, err = Redis.HGetAll(ctx, "remain").Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}
