package mongo

import (
	"fmt"
	"go-sample/utils/go-util/log"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Db      *mgo.Database
)

func Init() {
	log.Info(" Mongo host : ",Conf.Host)
	log.Info(" Mongo Port : ",Conf.Port)
	Session, err := mgo.Dial(fmt.Sprintf(`%s:%d`, Conf.Host, Conf.Port))

	if err != nil {
		log.Fatal(`Mongodb dial failed : `, err)
	}

	Session.SetMode(mgo.Monotonic, true)
	if Conf.Auth {
		err = Session.Login(&mgo.Credential{
			Username: Conf.User,
			Password: Conf.Password,
			Source:   Conf.AuthDb,
		})
		if err != nil {
			log.Fatal(`Mongodb login failed `, err)
		}
	}

	log.Info(`MongoDb connection initiated`)

	Db = Session.DB(Conf.Database)
}

func Close(s mgo.Session) {
	s.Close()
}
