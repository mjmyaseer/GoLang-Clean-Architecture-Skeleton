package request

import "context"
import (
	"fmt"
	"github.com/google/uuid"
)

var uuidKey = `Request: uuid`

func WithContext(ctx context.Context) context.Context {
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(`Cannot generate uuid`)
	}

	return context.WithValue(ctx, &uuidKey, id)

}

func IdFromContext(ctx context.Context) string {
	id, ok := ctx.Value(&uuidKey).(uuid.UUID)
	if !ok {
		//fmt.Println(ok)
	}
	uid, err := id.Value()
	if err != nil {
		//fmt.Println(err)
	}

	return fmt.Sprint(uid)
}
