package http

type HttpResponse Response

func (r HttpResponse) OK(data interface{}) HttpResponse {
	r.Status = 200
	r.Data = data
	return r
}

func (r HttpResponse) NotFound() HttpResponse {
	r.Status = 404
	return r
}

func (r HttpResponse) Unauthorized(message interface{}) HttpResponse {
	r.Status = 401
	r.Data = message
	return r
}

func (r HttpResponse) Forbidden(message interface{}) HttpResponse {
	r.Status = 403
	r.Data = message
	return r
}

func (r HttpResponse) BadRequest(message interface{}) HttpResponse {
	r.Status = 400
	r.Data = message
	return r
}
