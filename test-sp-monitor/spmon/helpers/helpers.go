package helpers

import (
	"sp-monitoring/models"
	"time"
)

func CheckTime(kvkey models.KVKey) (time.Duration, error) {
	t2, err := time.Parse("2006-01-02 15:04:05", kvkey.Expireddate)
	if err != nil {
		return time.Duration(0), err
	}
	t1 := time.Now()
	duration := t2.Sub(t1)

	return duration, nil
}

func ConvertToTime(fromRedis models.KVKey, fromKv string) (fromRedisKey time.Time, fromKvKey time.Time, err error) {
	fromRedisKey, err = time.Parse("2006-01-02 15:04:05", fromRedis.Expireddate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	fromKvKey, err = time.Parse("2006-01-02 15:04:05", fromKv)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return
}
