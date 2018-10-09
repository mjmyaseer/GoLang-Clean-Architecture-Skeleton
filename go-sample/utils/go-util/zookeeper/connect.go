package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"go-sample/utils/go-util/log"
	"time"
)

func Connect() *zk.Conn {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
