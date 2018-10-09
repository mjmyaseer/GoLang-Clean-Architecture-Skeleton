package postgre

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-sample/utils/go-util/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DbConnections struct {
	Write *sql.DB
}

//TODO Implement dynamic connections

var Connections DbConnections

type DbConfig struct {
	Host     string   `yaml:"host" json:"host"`         //Db host name
	Port     string   `yaml:"port" json:"port"`         //Db Port
	Db       string   `yaml:"database" json:"database"` //Db Name
	User     string   `yaml:"user" json:"user"`         //Db User
	Password string   `yaml:"password" json:"password"` //Db Password
	Services []string `yaml:"services" json:"services"`
}

type confFile struct {
	Write    DbConfig `yaml:"write" json:"write"`
	Timezone string   `yaml:"timezone" json:"timezone"`
}

var dbConfFile confFile

func Init() {

	parseConfig()
	Connections.Write, _ = open(dbConfFile.Write)

}

func (conf DbConfig) InitWrite() {
	Connections.Write, _ = open(conf)
}

func open(conf DbConfig) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.User, conf.Password, conf.Db)

	con, err := sql.Open(`postgres`, dbinfo)
	log.Info("Connections  : ", conf, con)

	if err != nil {
		log.Fatal(err)
	}
	if err := con.Ping(); err != nil {
		log.Fatal(err)

	}
	return con, err
}

func Close(connection *sql.DB) {
	err := connection.Close()
	if err != nil {
		log.Error(`Cannot close postgres connection :`, err)
	}
}

func parseConfig() {
	file, err := ioutil.ReadFile(`config/postgres.yaml`)
	if err != nil {
		log.Fatal(`Cannot open config file`, `, config/postgres.yaml, `, err)
	}

	err = yaml.Unmarshal(file, &dbConfFile)
	if err != nil {
		log.Fatal(`Cannot parse config file`, `config/postgres.yaml`, err)
	}

	log.Info(`Postgres connection establish`)

}
