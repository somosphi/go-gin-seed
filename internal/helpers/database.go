package helpers

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Driver   string `yaml:"driver"`
	Schema   string `yaml:"schema"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Pool     struct {
		Min uint8 `yaml:"min,omitempty"`
		Max uint8 `yaml:"max,omitempty"`
	} `yaml:"pool,omitempty"`
}

var database *gorm.DB
var once sync.Once

func _mountConnectionString(cfg *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Schema,
	)
}

// GetDatabase returns the connections to the database
//
// This function returns a singleton to access the database
func GetDatabase(cfg *DbConfig) (*gorm.DB, error) {
	var err error

	once.Do(func() {
		database, err = gorm.Open("mysql", _mountConnectionString(cfg))
		err = database.DB().Ping()
	})

	return database, err
}

func _setDbDefaults() {
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.pool.min", 1)
	viper.SetDefault("database.pool.max", 1)
}

func NewDbConfig() (cfg *DbConfig, err error) {
	_setDbDefaults()

	cfg = &DbConfig{}
	err = viper.UnmarshalKey("database", cfg)
	return
}
