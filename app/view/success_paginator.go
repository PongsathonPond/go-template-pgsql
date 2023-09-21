package view

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"idev-cms-service/domain"

	"github.com/gin-gonic/gin"
)

const xContentLength = "X-Content-Length"

type SuccessPaginatorResp struct {
	Code   int                `json:"code" example:"200"`
	Status string             `json:"status" example:"Get data success"`
	Data   dataListWithOption `json:"data" swaggerignore:"true"`
} // @Name SuccessPaginatorResponse

type SuccessGetAllResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Get data success"`
	Data   interface{} `json:"data" swaggerignore:"true"`
} // @Name SuccessGetAllResponse

type dataListWithOption struct {
	Option   optionResp  `json:"option"`
	DataList interface{} `json:"lists" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessPaginatorDataListWithOption

type optionResp struct {
	Total       int           `json:"total" example:"1"`
	PerPage     int           `json:"per_page" example:"15"`
	CurrentPage int           `json:"current_page" example:"1"`
	TotalPage   int           `json:"total_page" example:"1"`
	Filters     []*filterData `json:"filters"`
	Search      []*searchData `json:"search"`
	Sorts       []*sortData   `json:"sorts"`
} // @Name SuccessPaginatorOptionResponse

type filterData struct {
	Key       string `json:"key" example:"name"`
	Operation string `json:"operation" example:"eq"`
	Value     string `json:"value" example:"John"`
	Raw       string `json:"raw" example:"name:eq:John"`
} // @Name SuccessPaginatorFilterData

type searchData struct {
	Key   string `json:"key" example:"name"`
	Value string `json:"value" example:"John"`
	Raw   string `json:"raw" example:"name:eq:John"`
} // @Name SuccessPaginatorSearchData

type sortData struct {
	Key   string `json:"key" example:"created_at"`
	Value string `json:"value" example:"desc"`
	Raw   string `json:"raw" example:"created_at:desc"`
} // @Name SuccessPaginatorSortData

func MakePaginatorResp(c *gin.Context, total int, opt *domain.PageOption, items interface{}) {
	status := http.StatusOK
	msg := MsgGetDataSuccess(c)
	totalPage := int(math.Ceil(float64(total) / float64(opt.PerPage)))

	if opt.Page > totalPage {
		status = http.StatusNoContent
		msg = MsgNotFoundData(c)
	}

	filters := make([]*filterData, 0)
	if len(opt.Filters) > 0 {
		for _, raw := range opt.Filters {
			if raw == "deleted_at:isNull" || raw == "is_super_role:isNull" {
				continue
			}
			keyVal := strings.Split(raw, ":")
			var data filterData
			if len(keyVal) > 2 {
				data = filterData{
					Key:       keyVal[0],
					Operation: keyVal[1],
					Value:     keyVal[2],
					Raw:       raw,
				}
			} else {
				data = filterData{
					Key:       keyVal[0],
					Operation: keyVal[1],
					Raw:       raw,
				}
			}
			filters = append(filters, &data)
		}
	}
	sorts := make([]*sortData, 0)
	if len(opt.Sorts) > 0 {
		for _, raw := range opt.Sorts {
			keyVal := strings.Split(raw, ":")
			data := &sortData{
				Key:   keyVal[0],
				Value: keyVal[1],
				Raw:   raw,
			}
			sorts = append(sorts, data)
		}
	}

	c.Header(xContentLength, strconv.Itoa(total))
	c.JSON(status, SuccessPaginatorResp{
		Code:   status,
		Status: msg,
		Data: dataListWithOption{
			Option: optionResp{
				Total:       total,
				PerPage:     opt.PerPage,
				CurrentPage: opt.Page,
				TotalPage:   totalPage,
				Filters:     filters,
				Sorts:       sorts,
			},
			DataList: items,
		},
	})
}

func MakeGetAllResp(c *gin.Context, total int, items interface{}) {
	status := http.StatusOK
	msg := MsgGetDataSuccess(c)

	if total == 0 {
		status = http.StatusNoContent
		msg = MsgNotFoundData(c)
	}

	c.Header(xContentLength, strconv.Itoa(total))
	c.JSON(status, SuccessGetAllResp{
		Code:   status,
		Status: msg,
		Data:   items,
	})
}
