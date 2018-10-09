package datasource

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go-sample/utils/go-util/log"
	"go-sample/utils/go-util/mysql"
)

type transaction struct {
	tx *sql.Tx
}

type db struct {
	db *sql.DB
}

var (
	transactionKey            = "datasource: Transaction"
	TransactionStartFailed    = `datasource: Transaction Start Failed`
	TransactionRollbackFailed = `datasource: Transaction Rollback Failed`
	TransactionCommitFailed   = `datasource: Transaction Commit Failed`
)

//contextWithTransaction Assign a transaction to a context value
func contextWithTransaction(c context.Context, t *transaction) context.Context {
	return context.WithValue(c, &transactionKey, t)
}

//RunInTransaction Run a particular function inside transaction
func RunInTransaction(ctx context.Context, fn func(c context.Context) error, options *sql.TxOptions) error {

	//initiate an empty transaction
	t := new(transaction)
	//mutex.Lock()
	err := error(nil)

	//if context already contains a transaction then use it
	//Or create a new transaction
	if existingT := fromAnotherTransaction(ctx); existingT != nil {
		t = existingT
	} else {

		if database.Connections.Write == nil {
			log.Fatal(`transaction: DB Write Connection is empty`)
		}

		t.tx, err = database.Connections.Write.Begin()
		if err != nil {
			log.ErrorContext(ctx, TransactionStartFailed, `error :- `, err)
			return err
		}
	}

	//mutex.Unlock()

	err = fn(contextWithTransaction(ctx, t))

	//handle errors
	//TODO implement deadlock and lock wait timeout retry
	if sqlError, ok := err.(*mysql.MySQLError); ok {
		//ER_LOCK_DEADLOCK OR ER_LOCK_WAIT_TIMEOUT
		//Restart transaction
		if sqlError.Number == 1213 || sqlError.Number == 1205 {
			log.ErrorContext(ctx, `Deadlock happened`)
			//RunInTransaction(ctx, fn, options)
		}
	}

	if err != nil {
		//if connection is there
		if t.tx != nil {
			//try to rollback error
			rollBackErr := t.tx.Rollback()
			if rollBackErr != nil {
				return rollBackErr
			}
		}

		return err

	}

	if t.tx != nil {
		err = t.tx.Commit()
	}

	return err
}

//if context contains another transaction then use it
func fromAnotherTransaction(ctx context.Context) *transaction {
	t, _ := ctx.Value(&transactionKey).(*transaction)
	return t
}

//if context contains another transaction then use it
func FromAnotherTransaction(ctx context.Context) *sql.Tx {
	tx := fromAnotherTransaction(ctx)
	if tx != nil {
		return tx.tx
	}
	return nil
}
