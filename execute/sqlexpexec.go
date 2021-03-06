package execute

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/weworksandbox/lingo"
)

type ExecSQLExpInTx = func(ctx context.Context, s ExpQuery) error

type SQLExp interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TxSQLExp, error)
	InTx(ctx context.Context, opts *sql.TxOptions, execThis ExecSQLExpInTx) error
	ExpQuery
}

type TxSQLExp interface {
	CommitOrRollback(ctx context.Context, err error) error
	Rollback(ctx context.Context, err error) error
	ExpQuery
}

type ExpQuery interface {
	Exec(ctx context.Context, exp lingo.Expression) (sql.Result, error)
	Query(ctx context.Context, exp lingo.Expression) (RowScanner, error)
	QueryRow(ctx context.Context, exp lingo.Expression, valuePtrs ...interface{}) error
}

func NewSQLExp(s SQL, d lingo.Dialect) SQLExp {
	return sqlExpExec{
		exec: s,
		d:    d,
	}
}

type sqlExpExec struct {
	exec SQL
	d    lingo.Dialect
}

func (sqlExec sqlExpExec) BeginTx(ctx context.Context, opts *sql.TxOptions) (TxSQLExp, error) {
	txSQL, err := sqlExec.exec.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewTxSQLExp(txSQL, sqlExec.d), nil
}

func (sqlExec sqlExpExec) InTx(ctx context.Context, opts *sql.TxOptions, execThis ExecSQLExpInTx) (err error) {
	var txSQL TxSQLExp

	panicked := true
	defer func() {
		r := recover()
		if r != nil || panicked { // Workaround for if someone throws `nil`
			panicked = true // Set it to true regardless because we checked r, or panicked was already true.
			err = fmt.Errorf("panicked with %v", r)
		}
		if txSQL != nil {
			err = txSQL.CommitOrRollback(ctx, err)
		}

		if panicked {
			panic(r) // Throw the same thing we caught. Do not change the type, value, etc.
		}
	}()

	txSQL, err = sqlExec.BeginTx(ctx, opts)
	if err != nil {
		panicked = false // Normal error condition short circuit, no panic happened
		return err       // Already Traced
	}
	err = execThis(ctx, txSQL)
	panicked = false
	return
}

func (sqlExec sqlExpExec) Exec(ctx context.Context, exp lingo.Expression) (result sql.Result, err error) {
	return execExp(ctx, sqlExec.exec, sqlExec.d, exp)
}

func (sqlExec sqlExpExec) QueryRow(ctx context.Context, exp lingo.Expression, queryIntoPtrs ...interface{}) error {
	return queryRowExp(ctx, sqlExec.exec, sqlExec.d, exp, queryIntoPtrs)
}

func (sqlExec sqlExpExec) Query(ctx context.Context, exp lingo.Expression) (RowScanner, error) {
	return queryExp(ctx, sqlExec.exec, sqlExec.d, exp)
}

func NewTxSQLExp(s TxSQL, d lingo.Dialect) TxSQLExp {
	return sqlExpTxExec{
		exec: s,
		d:    d,
	}
}

type sqlExpTxExec struct {
	exec TxSQL
	d    lingo.Dialect
}

func (txExec sqlExpTxExec) CommitOrRollback(ctx context.Context, err error) error {
	return txExec.exec.CommitOrRollback(ctx, err)
}

func (txExec sqlExpTxExec) Rollback(ctx context.Context, err error) error {
	return txExec.exec.Rollback(ctx, err)
}

func (txExec sqlExpTxExec) Exec(ctx context.Context, exp lingo.Expression) (result sql.Result, err error) {
	return execExp(ctx, txExec.exec, txExec.d, exp)
}

func (txExec sqlExpTxExec) QueryRow(ctx context.Context, exp lingo.Expression, queryIntoPtrs ...interface{}) error {
	return queryRowExp(ctx, txExec.exec, txExec.d, exp, queryIntoPtrs)
}

func (txExec sqlExpTxExec) Query(ctx context.Context, exp lingo.Expression) (RowScanner, error) {
	return queryExp(ctx, txExec.exec, txExec.d, exp)
}

func expandSQL(dialect lingo.Dialect, exp lingo.Expression) (string, []interface{}, error) {
	lSQL, err := exp.ToSQL(dialect)
	if err != nil {
		return "", nil, err
	}
	return lSQL.String(), lSQL.Values(), nil
}

func execExp(ctx context.Context, db SQLQuery, d lingo.Dialect, exp lingo.Expression) (result sql.Result, err error) {
	var tSQL string
	var sVals []interface{}
	tSQL, sVals, err = expandSQL(d, exp)
	if err != nil {
		return nil, traceErr(ctx, err)
	}
	return db.Exec(ctx, tSQL, sVals...)
}

func queryRowExp(
	ctx context.Context, e SQLQuery, d lingo.Dialect, exp lingo.Expression, queryIntoPtrs []interface{},
) error {
	tSQL, sVals, err := expandSQL(d, exp)
	if err != nil {
		return traceErr(ctx, err)
	}

	return e.QueryRow(ctx, tSQL, sVals, queryIntoPtrs...)
}

func queryExp(ctx context.Context, db SQLQuery, d lingo.Dialect, exp lingo.Expression) (RowScanner, error) {
	tSQL, sVals, err := expandSQL(d, exp)
	if err != nil {
		return nil, traceErr(ctx, err)
	}

	return db.Query(ctx, tSQL, sVals...)
}
