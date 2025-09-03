package core

type IdParam struct {
	Id int64 `params:"id"`
}

type IdBody struct {
	Id int64 `json:"id"`
}
