package fizzhttpMocks

import "github.com/gin-gonic/gin/binding"

func NewTestingRequestContext(shouldBindBodyWithOutput error) *TestingRequestContext {
	return &TestingRequestContext{
		shouldBindBodyWithOutput: shouldBindBodyWithOutput,
	}
}

func (requestContext *TestingRequestContext) ShouldBindBodyWith(targetObjectPointer interface{}, bindingBody binding.BindingBody) (err error) {
	requestContext.RecordFuncCall("ShouldBindBodyWith", targetObjectPointer, bindingBody)
	return requestContext.shouldBindBodyWithOutput
}

func (requestContext *TestingRequestContext) AbortWithStatusJSON(statusCode int, responseObject interface{}) {
	requestContext.RecordFuncCall("AbortWithStatusJSON", statusCode, responseObject)
}

func (requestContext *TestingRequestContext) JSON(code int, obj interface{}) {
	requestContext.RecordFuncCall("JSON", code, obj)
}
