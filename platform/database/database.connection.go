package database

import (
	"SOCIAL_MEDIA_APP/platform/database/postgress"
	"SOCIAL_MEDIA_APP/platform/database/redis"
)

func DatabaseConnections() {
	postgress.GetPostgressDBInstance()
	redis.GetRedisClient()
}
