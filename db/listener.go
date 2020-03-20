package db

import "github.com/BOPR/types"

func (db *DB) StoreListenerLog(log types.ListenerLog) error {
	if err := db.Instance.Table("listener_logs").Assign(types.ListenerLog{LastRecordedBlock: log.LastRecordedBlock}).FirstOrCreate(&log).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetLastListenerLog() (log types.ListenerLog, err error) {
	if err := db.Instance.First(&log).Error; err != nil {
		return log, err
	}
	return log, nil
}
