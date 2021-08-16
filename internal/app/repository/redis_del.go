package repository

import (
	"github.com/pkg/errors"
)

func (r *redisRepository) Del(key string) (int64, error) {
	res, err := r.client.Del(r.ctx, key).Result()

	if err != nil {
		return 0, errors.Wrap(err, "error while deleting key")
	}

	return res, err
}
