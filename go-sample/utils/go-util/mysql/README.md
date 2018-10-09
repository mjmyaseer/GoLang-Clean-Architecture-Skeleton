### Mysql

Mysql helper package of [core](http://gitlab.mytaxi.lk/pickme/go-util) library

#### Usage

Init a connection

```go
import "go-sample/utils/go-util/mysql"


//initiate db connections
database.Init()
defer database.Close(database.Connections.Read)
defer database.Close(database.Connections.Write)

```
#### config

```json
{
  "database": {
    "read": {
      "host": "127.0.0.1",
      "port": "7001",
      "database": "",
      "user": "root",
      "password": "mysql",
      "max_idle_connections": 50,
      "max_open_connections": 100
    },

    "write": {
      "host": "127.0.0.1",
      "port": "7001",
      "database": "eztaxi",
      "user": "root",
      "password": "",
      "max_idle_connections": 50,
      "max_open_connections": 100
    }
  }
}
```
Parse time
 
```go
type DateTime struct {
	TimeStamp string
}

var myTime database.DateTime
time.TimeStamp = `2006-01-02 15:04:05` //mysql timestamp
timeObj, err = myTime.Parse() //Create a go time object

```
 
Parse time object to mysql timestamp
```go
database.ParseToString(time.Now())
``` 

Parse time object to nullable mysql timestamp
```go
database.ParseToNullableString(time.Now())
``` 

Parse time object to mysql date
```go
database.ParseToDateString(time.Now())
``` 
