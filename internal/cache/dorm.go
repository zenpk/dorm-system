package cache

import "context"

type Dorm struct{}

func (d *Dorm) GetAvailableNum(ctx context.Context) (map[string]string, error) {
	res, err := Redis.HGetAll(ctx, "available").Result()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		if err := warmAvailable(); err != nil {
			return nil, err
		}
	}
	res, err = Redis.HGetAll(ctx, "available").Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}
