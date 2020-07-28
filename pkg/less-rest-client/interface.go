package lessrestclient

type LRC interface {
	Request(
		method,
		route string,
		inData interface{},
		outData interface{},
		expectedStatusCode int,
	) (statusCode int, respBody []byte, err error)

	GET(
		route string,
		outData interface{},
		expectedStatusCode int,
	) (statusCode int, respBody []byte, err error)

	POST(
		route string,
		inData interface{},
		outData interface{},
		expectedStatusCode int,
	) (statusCode int, respBody []byte, err error)

	PUT(
		route string,
		inData interface{},
		outData interface{},
		expectedStatusCode int,
	) (statusCode int, respBody []byte, err error)

	PATCH(
		route string,
		inData interface{},
		outData interface{},
		expectedStatusCode int,
	) (statusCode int, respBody []byte, err error)

	DELETE(
		route string,
		inData interface{},
		outData interface{},
		expectedStatusCode int,
	) (statusCode int, respBody []byte, err error)
}
