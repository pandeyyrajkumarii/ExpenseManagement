package database

import "fmt"

type Config struct {
	Dialect         string
	Protocol        string
	Host            string
	Port            int
	Username        string
	Password        string
	SSLMode         string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

func (c *Config) ConnectToDb() string {
	switch c.Dialect {
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			c.Host,
			c.Port,
			c.Username,
			c.Password,
			c.Name,
			c.SSLMode)
	case "mysql":
		return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			c.Username,
			c.Password,
			c.Protocol,
			c.Host,
			c.Port,
			c.Name)
	case "sqlite3":
		return fmt.Sprintf("file:%s?cache=shared&mode=memory",
			c.Name)
	}
	return ""
}
