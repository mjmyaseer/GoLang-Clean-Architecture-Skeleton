package redis

import "context"
import "github.com/go-redis/redis"
import (
	"fmt"
	"go-sample/utils/go-util/config"
	"go-sample/utils/go-util/log"
	"os"
	"os/signal"
)

var (
	Client       *redis.Client
	KeySeparator = `:`
	Conf         Config
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database int    `json:"database"`
	Password string `json:"password"`
}

//init redis driver or start pool
func Init() {
	loadConfig()
	cl := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(`%s:%d`, Conf.Host, Conf.Port),
		Password: Conf.Password,
		DB:       Conf.Database,
	})

	_, err := cl.Ping().Result()
	if err != nil {
		log.Fatal(`Redis: Connection Ping Error : `, err)
	}

	log.Info(`Redis client establish`)

	Client = cl

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		select {
		case sig := <-signals:
			Close(Client)
			log.Info(`Redis connection aborted : `, sig)
			break
		}
	}()
}

func loadConfig() {
	config.DefaultConfigurator.Load(`config/redis`, &Conf, func(config interface{}) {})
}

func Close(client *redis.Client) {
	client.Close()
}

func Get(ctx context.Context, path string) (string, error) {
	rep, err := Client.Get(path).Result()
	if err != redis.Nil && err != nil {
		log.ErrorContext(ctx, err)
		return rep, err
	}

	return rep, err
}

func Keys(ctx context.Context, pattern string) ([]string, error) {
	rep, err := Client.Keys(pattern).Result()
	if err != nil {
		log.ErrorContext(ctx, err)
		return rep, err
	}

	return rep, err
}

func Exist(ctx context.Context, path string) (bool, error) {
	rep, err := Client.Exists(path).Result()
	if err != nil {
		log.ErrorContext(ctx, err)
		return false, err
	}

	isExist := rep == 1

	return isExist, err
}

func Set(ctx context.Context, path string, value interface{}) error {

	_, err := Client.Set(path, value, 0).Result()
	if err != nil {
		log.ErrorContext(ctx, err)
		return err
	}

	return err

}

func Delete(ctx context.Context, path string) error {
	_, err := Client.Del(path).Result()
	if err != nil {
		log.ErrorContext(ctx, err)
		return err
	}

	return err
}
