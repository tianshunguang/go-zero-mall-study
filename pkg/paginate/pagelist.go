package paginate

import go_paginate "github.com/raphaelvigee/go-paginate"

type PageInfo struct {
	go_paginate.PageInfo
	Hits  int64 `json:"hits"`
	Limit int   `json:"limit"`
}


type PageListRequest struct {
	PageNum int `json:"pageNum" binding:"required"`
	PerPage int `json:"perPage" binding:"required"`
}

type PageListResponse struct {
	HasNext      bool        `json:"hasNext"`
	HasPrev      bool        `json:"hasPrev"`
	NumPerPage   int         `json:"numPerPage"`
	TotalPageNum int         `json:"totalPageNum"`
	CurPageNum   int         `json:"curPageNum"`
	TotalNum     int64       `json:"totalNum"`
	Data         interface{} `json:"data"`
}

type SelectorListResponse struct {
	Data interface{} `json:"data"`
}

type CommonListResp struct {
	Data         interface{} `json:"data"`
	CurPageNum   int         `json:"curPageNum"`
	TotalPageNum int         `json:"totalPageNum"`
	NumPerPage   int         `json:"numPerPage"`
	HasPrev      bool        `json:"hasPrev"`
	HasNext      bool        `json:"hasNext"`
	TotalNum     int64       `json:"totalNum"`
}

