package service

import "../../../db"

// CheckUserName 按账号名查找账号
func (s *Server) CheckUserName(name string) (db.Userinfo, error) {
	userinfo := db.Userinfo{}
	err := s.Where("name = ?", name).First(&userinfo).Error
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}

// CheckUserID 按ID名查找账号
func (s *Server) CheckUserID(ID int) (db.Userinfo, error) {
	userinfo := db.Userinfo{
		ID: ID,
	}
	err := s.Where("del != 1").First(&userinfo).Error
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}

// CreateUser 创建账号
func (s *Server) CreateUser(userinfo db.Userinfo) error {
	err := s.Create(&userinfo).Error
	return err
}

// CreateLoginfo 创建登陆信息
func (s *Server) CreateLoginfo(loginfo db.Logininfo) error {
	err := s.Create(&loginfo).Error
	return err
}

// UpdateUser 更新用户信息
func (s *Server) UpdateUser(id int, updates map[string]interface{}) (db.Userinfo, error) {
	userinfo := db.Userinfo{
		ID: id,
	}
	err := s.Model(&userinfo).Updates(updates).Error
	return userinfo, err
}
