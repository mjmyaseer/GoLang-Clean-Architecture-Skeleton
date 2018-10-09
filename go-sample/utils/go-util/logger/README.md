### Logger

Logger package of [core]() library

#### Usage

simple loggin

```go

logger.log().ErrorContext(ctx, `Something went wrong`, `param1`, `param2`)

```
this will log following output to the terminal

```bash
!!! ERROR !!!
Request Id => eb429299-4528-11e7-9d6a-42010af00018
Message => Something went wrong
Params { "param1", "param2"}
```

#### config

```json
{
  "level": "INFO",
  "remote_log": true,
  "local_log": true,
  "file_logging": true,
  "log_to_file": false,
  "log_path": "logs"
}
```