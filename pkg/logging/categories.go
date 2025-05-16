package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	RequestResponse Category = "RequestResponse"
)

const (
	// Startup General
	Startup SubCategory = "Startup"

	// Api Internal
	Api SubCategory = "Api"
)

const (
	ClientIp     ExtraKey = "ClientIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"
)
