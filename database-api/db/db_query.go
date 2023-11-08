package db

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryParams struct {
	Select    []string          `form:"select"`
	Where     map[string]string `form:"where"`
	OrWhere   map[string]string `form:"orwhere"`
	Order     map[string]string `form:"order"`
	Relations []string          `form:"relations"`
	Limit     int               `form:"limit"`
	Offset    int               `form:"offset"`
	Cache     bool              `form:"cache"`
}

func BindGormQuery(ctx *gin.Context) (QueryParams, error) {

	var queryParams QueryParams

	if err := ctx.BindQuery(&queryParams); err != nil {
		return QueryParams{}, errors.New("unsupported operation")
	}
	order := ctx.QueryMap("order")
	where := ctx.QueryMap("where")
	orWhere := ctx.QueryMap("orwhere")
	if order != nil {
		queryParams.Order = order
	}
	if where != nil {
		queryParams.Where = where
	}
	if orWhere != nil {
		queryParams.OrWhere = orWhere
	}
	return queryParams, nil
}

func (query *QueryParams) GetQuery(db *gorm.DB) *gorm.DB {
	if query.Where != nil {
		for key, value := range query.Where {
			db = db.Where(key, value)
		}
	}
	if query.OrWhere != nil {
		for key, value := range query.OrWhere {
			db = db.Or(key, value)
		}
	}
	if query.Order != nil {
		for key, value := range query.Order {
			db = db.Order(key + " " + value)
		}
	}
	if query.Offset != 0 {
		db = db.Offset(query.Offset)
	}

	if query.Limit != 0 {
		db = db.Limit(query.Limit)
	}

	if query.Select != nil {
		if query.Select != nil {
			db = db.Select(query.Select)
		}
	}
	if query.Relations != nil {
		for _, s := range query.Relations {
			db = db.Joins(s)
		}
	}
	return db
}
