package service

import "../../../db"

// CreateState ..
func (s *Server) CreateState(uid int) error {
	configstates := []db.Configstates{
		db.Configstates{
			UID:    uid,
			Types:  1,
			States: "运动",
		},
		db.Configstates{
			UID:    uid,
			Types:  2,
			States: "报警",
		},
		db.Configstates{
			UID:    uid,
			Types:  3,
			States: "警告",
		},
		db.Configstates{
			UID:    uid,
			Types:  4,
			States: "停止",
		},
		db.Configstates{
			UID:    uid,
			Types:  5,
			States: "启动",
		},
		db.Configstates{
			UID:    uid,
			Types:  6,
			States: "事件",
		},
		db.Configstates{
			UID:    uid,
			Types:  7,
			States: "提醒",
		},
		db.Configstates{
			UID:    uid,
			Types:  8,
			States: "关机",
		},
		db.Configstates{
			UID:    uid,
			Types:  9,
			States: "离线",
		},
	}
	var err error
	db := s.Begin()
	for _, v := range configstates {
		err = db.Create(&v).Error
		if err != nil {
			db.Rollback()
			return err
		}
	}
	err = db.Commit().Error
	return err
}
