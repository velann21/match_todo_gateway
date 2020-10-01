package repository

import "database/sql"




type SqlClient interface {
	Begin()(*sql.Tx, error)
	Prepare(tx *sql.Tx, query string)(*sql.Stmt, error)
	Exec(stmt *sql.Stmt,args ...interface{})(sql.Result, error)
	LastInsertedID(result sql.Result)(int64, error)
	RowEffected(result sql.Result)(int64, error)
	Commit(tx *sql.Tx)error
	RollBack(tx *sql.Tx)error
}

type SqlClientImpl struct {
	sqlClient *sql.DB
}

func New(sql *sql.DB)SqlClient{
	return &SqlClientImpl{sqlClient:sql}
}

func (sql *SqlClientImpl) Begin()(*sql.Tx, error){
	return sql.sqlClient.Begin()
}

func (sql *SqlClientImpl) Prepare(tx *sql.Tx, query string)(*sql.Stmt, error){

	return tx.Prepare(query)
}

func (sql *SqlClientImpl) Exec(stmt *sql.Stmt,args ...interface{})(sql.Result, error){
	return stmt.Exec(args...)
}

func (sql *SqlClientImpl) LastInsertedID(result sql.Result)(int64, error){
	return result.LastInsertId()
}

func (sql *SqlClientImpl) RowEffected(result sql.Result)(int64, error){
	return result.RowsAffected()
}

func (sql *SqlClientImpl) Commit(tx *sql.Tx)error{
	return tx.Commit()
}

func (sql *SqlClientImpl) RollBack(tx *sql.Tx)error{
	return tx.Rollback()
}




