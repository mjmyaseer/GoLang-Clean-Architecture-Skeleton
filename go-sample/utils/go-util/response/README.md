### Response

Response helper package of [core](http://gitlab.mytaxi.lk/pickme/go-util) library

#### Usage

Render an error

```go
import "go-sample/utils/go-util/response"

response.HandleError(ctx, err, w)

```
Above will check the error type and if 
- Error is a [DomainError]() api response render helpers will generate following output

```json
    
```


