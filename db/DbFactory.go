package db

import (
	"gorm.io/gorm"
)

// DBFactory provides resolver for reader and writer database instances.
type DBFactory interface {
	Name() string
	Reader() *gorm.DB
	Writer() *gorm.DB
}

// Gobal instance
var instance *dbFactory

type dbFactory struct {
	db DBFactory
}

func (db *dbFactory) Name() string {
	return db.db.Name()
}

func (db *dbFactory) Reader() *gorm.DB {
	return db.db.Reader()
}

func (db *dbFactory) Writer() *gorm.DB {
	return db.db.Writer()
}

// function method variables.
var (
	Name   = instance.Name
	Reader = instance.Reader
	Writer = instance.Writer
)

// Instance sets the global database instance
func Instance(db DBFactory) {
	instance = &dbFactory{db}
	Name = instance.Name
	Reader = instance.Reader
	Writer = instance.Writer
}
