package http

type Header map[string]string

type HttpRequest interface {
	PostJson(host string, headers Header, data *string) ([]byte, int, error)
	GetJson(host string, headers Header) ([]byte, int, error)
	PatchJson(host string, headers Header, data *string) ([]byte, int, error)
	PutJson(host string, headers Header, data *string) ([]byte, int, error)
	DeleteJson(host string, headers Header) ([]byte, int, error)
}

type DefaultHttpRequest struct{}
