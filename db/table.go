package db

// Userinfo 用户列表
type Userinfo struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT;unique_index" json:"id"`
	Ownid       int    `gorm:"not null;index"`
	Name        string `gorm:"not null;size:100;unique_index"`
	Password    string `gorm:"size:32"`
	Photo       string `gorm:"size:255"`
	Permisson   int    `gorm:"not null"`
	VIP         int    `gorm:"not null"`
	Status      int    `gorm:"not null"`
	Createtime  int64  `gorm:"not null"`
	Updatetime  int64  `gorm:"not null;default:0"`
	Expiredtime int64  `gorm:"not null;default:0"`
	Mobile      string `gorm:"not null;size:11"`
	Email       string `gorm:"not null;size:100"`
	Address     string `gorm:"not null;size:255"`
	Details     string `gorm:"not null;size:255"`
	Del         int    `gorm:"not null;default:0"`
}

// Logininfo 登陆列表
type Logininfo struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	UID        int    `gorm:"not null;index"`
	Createtime int64  `gorm:"not null"`
	IP         string `gorm:"not null"`
}

// Deviceinfo 设备列表
type Deviceinfo struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	Name       string `gorm:"not null;size:80"` //设备名称
	UID        int    `gorm:"not null;index"`
	Uname      string `gorm:"not null"`
	Classid    int    `gorm:"not null;index"`
	Calssname  string `gorm:"not null"`
	DevEUI     string `gorm:"size:16"`  //设备UID
	Status     int    `gorm:"not null"` // 1启用 2禁用 3过期
	State      int    `gorm:"not null"`
	Createtime int64  `gorm:"not null"`
	Uptime     int64  `gorm:"not null;default:0"`
	Expiretime int64  `gorm:"not null;default:0"`
	Del        int    `gorm:"not null;default:0"`
}

// Devicedata 设备数据
type Devicedata struct {
	ID         int     `gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	UID        int     `gorm:"not null;index"`
	Classid    int     `gorm:"not null;index"`
	Longitude  float64 `gorm:"not null"`
	Latitude   float64 `gorm:"not null"`
	Createtime int64   `gorm:"not null"`
	Uptime     int64   `gorm:"not null"`
	Infos      string
}

// Logoinfo logo商标
type Logoinfo struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	Types      int    `gorm:"not null"` //种类 1背景 2登陆logo 3商标
	URL        string `gorm:"not null;size:200"`
	Createtime int64  `gorm:"not null;default:0"`
	Del        int    `gorm:"not null;default:0"`
}

// Configuration 用户配置
type Configuration struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT;unique_index"` //用户ID
	Message  int `gorm:"not null;default:1"`                      //消息推送开关
	Sound    int `gorm:"not null;default:1"`                      //声音开关
	Language int `gorm:"not null;default:1"`                      //语言选择 1简体中文 2繁体中文 3英文
}
