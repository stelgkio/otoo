package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// CreateSchema creates database schema for User and Story models.
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*domain.User)(nil),
		(*domain.Project)(nil),
		// (*domain.WooCommerce_Topic)(nil),
		// (*domain.Shopify_Topic)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			//	Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
