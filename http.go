package RazerCache

const defaultBasePath = "/api/Razercache/"

type HTTPPool struct {

	self string

	BasePath string

}

func NewHTTPPool(self string) *HTTPPool  {
	return &HTTPPool{self: self,BasePath: defaultBasePath}
}