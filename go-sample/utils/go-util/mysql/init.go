package database

import (
	"database/sql"
	stdMysql "github.com/go-sql-driver/mysql"
	"go-sample/utils/go-util/config"
	"go-sample/utils/go-util/log"
	"os"
	"os/signal"
	"time"
)

var defaultOptions = new(sqlOptions)

type DbConnections struct {
	Read  *sql.DB
	Write *sql.DB
}

type sqlOptions struct {
	readOptions  *stdMysql.Config
	writeOptions *stdMysql.Config
}

type Options func(c *sqlOptions)

func (opts *sqlOptions) Apply(options ...Options) {
	opts.writeOptions = stdMysql.NewConfig()
	opts.readOptions = stdMysql.NewConfig()
	for _, option := range options {
		option(opts)
	}
}

func WithReadConfig(conf *stdMysql.Config) Options {
	return func(c *sqlOptions) {
		c.readOptions = conf
	}
}

func WithWriteConfig(conf *stdMysql.Config) Options {
	return func(c *sqlOptions) {
		c.writeOptions = conf
	}
}

//TODO Implement dynamic connections

var Connections DbConnections

type DbConfig struct {
	Host        string   `yaml:"host" json:"host"`                                 //Db host name
	Port        string   `yaml:"port" json:"port"`                                 //Db Port
	Db          string   `yaml:"database" json:"database"`                         //Db Name
	User        string   `yaml:"user" json:"user"`                                 //Db User
	Password    string   `yaml:"password" json:"password"`                         //Db Password
	MaxOpenCons int      `yaml:"max_open_connections" json:"max_open_connections"` //Max maximum opened connections in the pool
	MaxIdleCons int      `yaml:"max_idle_connections" json:"max_idle_connections"` //Max idle connections in the pool
	Services    []string `yaml:"services" json:"services"`
}

type confFile struct {
	Read     DbConfig `yaml:"read" json:"read"`
	Write    DbConfig `yaml:"write" json:"write"`
	Timezone string   `yaml:"timezone" json:"timezone"`
}

var dbConfFile confFile

func Init(options ...Options) {

	defaultOptions.Apply(options...)
	parseConfig()

	Connections.Read, _ = open(dbConfFile.Read, defaultOptions.readOptions)
	Connections.Write, _ = open(dbConfFile.Write, defaultOptions.writeOptions)

	go func() {

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		select {
		case sig := <-signals:
			Close(Connections.Read)
			Close(Connections.Write)
			log.Info(`Mysql connection aborted : `, sig)
			return
		}
	}()
}

func (conf DbConfig) InitRead(options *stdMysql.Config) {
	Connections.Read, _ = open(conf, options)
}

func (conf DbConfig) InitWrite(options *stdMysql.Config) {
	Connections.Write, _ = open(conf, options)
}

func open(conf DbConfig, options *stdMysql.Config) (*sql.DB, error) {
	options.DBName = conf.Db
	options.User = conf.User
	options.Passwd = conf.Password
	options.Addr = conf.Host + `:` + conf.Port
	options.Net = `tcp`

	tZone, err := time.LoadLocation(dbConfFile.Timezone)
	if err != nil {
		log.Fatal(err)
	}
	options.Loc = tZone

	//log.Fatal(options.FormatDSN())

	//con, err := sql.Open(`mysql`, conf.User+`:`+conf.Password+`@tcp(`+conf.Host+`:`+conf.Port+`)/`+conf.Db+`?loc=`+url.QueryEscape(dbConfFile.Timezone))
	con, err := sql.Open(`mysql`, options.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := con.Ping(); err != nil {
		log.Fatal(err)

	}
	con.SetMaxIdleConns(conf.MaxIdleCons)
	con.SetMaxOpenConns(conf.MaxOpenCons)

	log.Info(`Mysql connection establish`, conf.Host+`:`+conf.Port)

	return con, err
}

func Close(connection *sql.DB) {
	err := connection.Close()
	if err != nil {
		log.Error(`Cannot close mysql connection :`, err)
	}
}

func parseConfig() {
	config.DefaultConfigurator.Load(`config/database`, &dbConfFile, func(config interface{}) {
		conf, _ := config.(*confFile)

		if conf.Timezone == `` {
			log.Fatal(`config/database : timezone cannot be empty`)
		}

		if conf.Read.Port == `` {
			log.Fatal(`config/database/write : port cannot be empty`)
		}

		if conf.Read.Db == `` {
			log.Fatal(`config/database/db : db cannot be empty`)
		}

		if conf.Read.User == `` {
			log.Fatal(`config/database/user : user cannot be empty`)
		}

		if conf.Read.Host == `` {
			log.Fatal(`config/database/host : host cannot be empty`)
		}

		if conf.Read.MaxOpenCons < 1 {
			log.Fatal(`config/database/MaxOpenCons : MaxOpenCons should be greater than zero`)
		}

		if conf.Write.Port == `` {
			log.Fatal(`config/database/write : port cannot be empty`)
		}

		if conf.Write.Db == `` {
			log.Fatal(`config/database/db : db cannot be empty`)
		}

		if conf.Write.User == `` {
			log.Fatal(`config/database/user : user cannot be empty`)
		}

		if conf.Write.Host == `` {
			log.Fatal(`config/database/host : host cannot be empty`)
		}

		if conf.Write.MaxOpenCons < 1 {
			log.Fatal(`config/database/MaxOpenCons : MaxOpenCons should be greater than zero`)
		}
	})
}