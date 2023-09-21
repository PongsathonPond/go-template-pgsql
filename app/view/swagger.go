package view

type SuccessReadResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Get data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessReadResponse

type SuccessCreateResp struct {
	Code   int         `json:"code" example:"201"`
	Status string      `json:"status" example:"Create data success"`
	Data   interface{} `json:"data" swaggertype:"string" example:"1400000000000000001"`
} // @Name SuccessCreateResponse

type SuccessUpdateResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Update data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessUpdateResponse

type SuccessDeleteResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Delete data success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessDeleteResponse

type SuccessLoginResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Login success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessLoginResponse

type SuccessLogoutResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Logout success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessLogoutResponse

type SuccessDefaultResp struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Success"`
	Data   interface{} `json:"data" swaggertype:"object" swaggerignore:"true"`
} // @Name SuccessDefaultResp

type Error400Resp struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"Bad request"`
} // @Name Error400Response

type Error404Resp struct {
	Code   int    `json:"code" example:"404"`
	Status string `json:"status" example:"Not found"`
} // @Name Error404Response

type Error422Resp struct {
	Code   int    `json:"code" example:"422"`
	Status string `json:"status" example:"Attempt error"`
} // @Name Error422Response

type Error500Resp struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"Internal server error"`
} // @Name Error500Response
