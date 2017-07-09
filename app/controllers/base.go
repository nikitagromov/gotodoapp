package controllers

import (
	"strconv"
	"github.com/jinzhu/gorm"
	"net/url"
	"strings"
	"io/ioutil"
	"github.com/revel/revel"
	"todoapp/app/models"
	"fmt"
)



func getItemsCollectionQuery(c url.Values) *gorm.DB {
	fmt.Println(models.Database)
	query := models.Database
	query.LogMode(true)
	_limit, limit_error := strconv.Atoi(c.Get("_limit"))
	_offset, offset_error := strconv.Atoi(c.Get("_offset"))

	if limit_error != nil {
		_limit = -1
	}

	if offset_error != nil {
		_offset = -1
	}

	if _limit >= 0 {
		query = query.Limit(_limit)
	}

	if _offset >= 0 {
		query = query.Order("id").Offset(_offset)
	}

	return query
}

func processQParam(c url.Values, query *gorm.DB) *gorm.DB{
	q := c.Get("q")
	if q != "" {

		param := strings.Split(q, ":")

		if strings.Contains(param[1], "*") {
			value := strings.Replace(param[1], "*", "", -1)
			query = query.Where(param[0] + " LIKE ?", "%" + value + "%")

		} else {
			query = query.Where(param[0] + " = ?", param[1])
		}
	}
	return query

}

func getBody(r *revel.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

