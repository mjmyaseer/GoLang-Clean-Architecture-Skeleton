package boot

import (
	stdMysql "github.com/go-sql-driver/mysql"
	"go-sample/utils/go-util/config"
	"go-sample/utils/go-util/mysql"
	"os"
	"os/signal"
	"go-sample/transport/http"
)

func Init() {

	config.LoadConfiguration(Configurations)
	//infrastructure.Backend = infrastructure.NewLocationBackend()

	config := stdMysql.NewConfig()
	config.Params = map[string]string{`parseTime`: `true`}

	database.Init(
		database.WithReadConfig(config),
		database.WithWriteConfig(config))
	defer database.Close(database.Connections.Read)
	defer database.Close(database.Connections.Write)

	//http.InitHttpRouter()



	go func() {
		// trap SIGINT to trigger a shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		defer close(signals)


		//http.StopHttpRouter()
	}()

	http.InitHttpRouter()

}
