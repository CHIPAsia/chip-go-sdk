package http

import "math"

type FileResponse struct {
	Byte     []byte `json:"byte"`
	Filesize int    `json:"size"`
	Filename string `json:"filename"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponsePaginate struct {
	Results     interface{} `json:"results,omitempty"`
	NextPage    *int        `json:"nextPage,omitempty"`
	PrevPage    *int        `json:"prevPage,omitempty"`
	TotalPage   int         `json:"totalPage,omitempty"`
	Limit       int         `json:"limit,omitempty"`
	Facet       interface{} `json:"facet,omitempty"`
	CurrentPage int         `json:"currentPage"`
	Count       int         `json:"totalItems"`
}

func (r ResponsePaginate) Init() ResponsePaginate {
	nextPage := 0
	prevPage := 0
	r.TotalPage = int(math.Ceil(float64(r.Count) / float64(r.Limit)))
	if r.TotalPage > r.CurrentPage {
		nextPage = r.CurrentPage + 1
	}
	if r.CurrentPage > 1 {
		prevPage = r.CurrentPage - 1
	}

	if nextPage > 0 {
		r.NextPage = &nextPage
	}
	if prevPage > 0 {
		r.PrevPage = &prevPage
	}

	return r
}
