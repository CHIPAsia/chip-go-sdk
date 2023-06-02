package http

type RequestPaginate struct {
	Limit  int
	Page   int
	Offset int
}

func (r RequestPaginate) Init(initLimit int) RequestPaginate {
	limit := r.Limit
	page := r.Page
	if limit <= 0 {
		if initLimit <= 0 {
			limit = 30
		} else {
			limit = initLimit
		}
	}
	if page <= 0 {
		page = 1
	}
	offset := limit * (page - 1)

	r.Limit = limit
	r.Page = page
	r.Offset = offset

	return r
}
