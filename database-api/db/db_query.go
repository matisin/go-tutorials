package db

import (
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryParams struct {
	Select    []string          `form:"select"`
	Where     map[string]string `form:"where"`
	OrWhere   map[string]string `form:"orwhere"`
	Order     map[string]string `form:"order"`
	Relations map[string]string `form:"relations"`
	Limit     int               `form:"limit"`
	Offset    int               `form:"offset"`
	Cache     bool              `form:"cache"`
}

func BindGormQuery(ctx *gin.Context) (QueryParams, error) {
	order := ctx.QueryMap("order")
	relations := ctx.QueryMap("relations")
	where := ctx.QueryMap("where")
	orWhere := ctx.QueryMap("orwhere")

	var queryParams QueryParams

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		return QueryParams{}, errors.New("unsupported operation")
	}
	if order != nil {
		queryParams.Order = order
	}
	if relations != nil {
		queryParams.Relations = relations
	}
	if where != nil {
		queryParams.Where = where
	}
	if orWhere != nil {
		queryParams.OrWhere = orWhere
	}
	return queryParams, nil
}

func (q *QueryParams) GetQuery(db *gorm.DB) *gorm.DB {
	log.Println(q.Where)
	if q.Where != nil {
		for key, value := range q.Where {
			db = db.Where(key, value)
		}
	}
	if q.OrWhere != nil {
		for key, value := range q.OrWhere {
			db = db.Or(key, value)
		}
	}
	if q.Order != nil {
		for key, value := range q.Order {
			db = db.Order(key + " " + value)
		}
	}

	if q.Offset != 0 {
		db = db.Offset(q.Offset)
	}

	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}

	if q.Select != nil {
		if q.Select != nil {
			db = db.Select(q.Select)
		}
	}
	if q.Relations != nil {
		for field, value := range q.Relations {
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				log.Fatal(err)
			}
			if boolValue {
				db = db.Joins(field)
			}
		}
	}
	return db
}
