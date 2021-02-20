package response

type HttpResponse struct {
	code    int
	headers map[string]string
	body    string
}
