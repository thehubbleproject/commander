package db

import "github.com/BOPR/types"

func (db *DB) StoreListenerLog(log types.ListenerLog) {
	db.Instance.Update(&log)
	return
}

func (db *DB) GetLastListenerLog() (log types.ListenerLog) {
	db.Instance.First(&log)
	return
}
