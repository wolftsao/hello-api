package translation

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/wolftsao/hello-api/config"
	"github.com/wolftsao/hello-api/handlers/rest"
)

var _ rest.Translator = &Database{}

type Database struct {
	conn *redis.Client
}

func NewDatabaseService(cfg config.Configuration) *Database {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DatabaseURL, cfg.DatabasePort),
		Password: "",
		DB:       0,
	})
	return &Database{
		conn: rdb,
	}
}

func (s *Database) Close() error {
	return s.conn.Close()
}

func (s *Database) Translate(word string, language string) string {
	out := s.conn.Get(context.Background(), fmt.Sprintf("%s:%s", word, language))
	return out.Val()
}
