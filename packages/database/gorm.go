package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type DatabaseGorm struct {
	Db *gorm.DB
}

func NewDb(config *Config) (*DatabaseGorm, error) {
	gormConfig := &gorm.Config{
		AllowGlobalUpdate:      false,
		SkipDefaultTransaction: true,
	}
	var dialector gorm.Dialector
	switch config.Dialect {
	case "postgres":
		dialector = postgres.Open(config.ConnectToDb())
	case "mysql":
		dialector = mysql.Open(config.ConnectToDb())
	case "sqlite":
		dialector = sqlite.Open(config.ConnectToDb())
	default:
		return nil, fmt.Errorf("database dialect %s not supported", config.Dialect)

	}
	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, err
	}
	conn, err := db.DB()
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	return &DatabaseGorm{Db: db}, nil
}

func (db *DatabaseGorm) Create(model interface{}) (int64, error) {
	r := db.Db.Create(model)
	return r.RowsAffected, r.Error
}

func (db *DatabaseGorm) Update(model interface{}) (int64, error) {
	r := db.Db.Model(model).Updates(model)
	return r.RowsAffected, r.Error
}

func (db *DatabaseGorm) Delete(model interface{}) (int64, error) {
	r := db.Db.Delete(model)
	return r.RowsAffected, r.Error
}

// UpdateWithQuery updates records based on a custom WHERE condition
func (db *DatabaseGorm) UpdateWithQuery(model interface{}, conditions map[string]interface{}, updates map[string]interface{}) (int64, error) {
	r := db.Db.Model(model).Where(conditions).Updates(updates)
	return r.RowsAffected, r.Error
}

func (db *DatabaseGorm) FindByID(model interface{}, id string) (int64, error) {
	r := db.Db.Where("id=?", id).Find(model)
	return r.RowsAffected, r.Error
}

// FindWithQuery retrieves records based on a custom query builder with optional offset, limit, and order
func (db *DatabaseGorm) FindWithQuery(model interface{}, query map[string]interface{}) (int64, error) {
	r := db.Db.Model(model).Where(query).Find(model)
	return r.RowsAffected, r.Error
}

func (db *DatabaseGorm) FindWithQueryFilter(model interface{}, query *gorm.DB) (int64, error) {
	r := query.Find(model)
	return r.RowsAffected, r.Error
}

func (db *DatabaseGorm) DeleteWithWhere(model interface{}, conditions map[string]interface{}) error {
	return db.Db.Where(conditions).Delete(model).Error
}

func (db *DatabaseGorm) GetInstance() (*sql.DB, error) {
	return db.Db.DB()
}
