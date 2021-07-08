package fizzhttpMocks

import (
	"net/http"
)

func NewTestingRequestContext(shouldBindBodyWithJSONOutput error) *TestingRequestContext {
	return &TestingRequestContext{
		shouldBindBodyWithJSONOutput: shouldBindBodyWithJSONOutput,
	}
}

func (requestContext *TestingRequestContext) ShouldBindBodyWithJSON(targetObjectPointer interface{}) (err error) {
	requestContext.RecordFuncCall("ShouldBindBodyWithJSON", targetObjectPointer)
	return requestContext.shouldBindBodyWithJSONOutput
}

func (requestContext *TestingRequestContext) AbortWithStatusJSON(statusCode int, responseObject interface{}) {
	requestContext.RecordFuncCall("AbortWithStatusJSON", statusCode, responseObject)
}

func (requestContext *TestingRequestContext) SendJSONResponse(code int, obj interface{}) {
	requestContext.RecordFuncCall("SendJSONResponse", code, obj)
}

func (requestContext *TestingRequestContext) GetResponseWriter() http.ResponseWriter {
	requestContext.RecordFuncCall("GetResponseWriter")
	return nil
}

func (requestContext *TestingRequestContext) GetRequest() *http.Request {
	requestContext.RecordFuncCall("GetRequest")
	return nil
}
