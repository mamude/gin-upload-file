package helpers

import "database/sql"

func CheckDBError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckDBErrorTx(tx *sql.Tx, err error) {
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
