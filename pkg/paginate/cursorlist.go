package paginate

import go_paginate "github.com/raphaelvigee/go-paginate"

type CursorPageInfo struct {
	go_paginate.PageInfo
	Hits  int64 `json:"hits"`
	Limit int   `json:"limit"`
}

type CursorResponse struct {
	Data          interface{} `json:"data"`
	NextPage      bool        `json:"nextPage"`
	PreviousPage  bool        `json:"previousPage"`
	StartCursor   string      `json:"startCursor"`
	EndCursor     string      `json:"endCursor"`
	CurrentCursor string      `json:"currentCursor"`
	Hits          int64       `json:"hits"`
	Limit         int         `json:"limit"`
}

type ListResponse struct {
	Data interface{} `json:"data"`
}



