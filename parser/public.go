package parser

import "sync"

const (
	// HTTPTypeUnknown is unknown http state. May occurs on errors
	HTTPTypeUnknown HTTPType = "unknown"
	// HTTPTypeRequest is a request type
	HTTPTypeRequest HTTPType = "request"
	// HTTPTypeResponse is a response type
	HTTPTypeResponse HTTPType = "response"
)

// HTTPType is a type that used to store http body type
// which can be request or response. can be empty if unknown state occurred
type HTTPType string

// Result of the parsing
type Result struct {
	body          []byte
	rawHeadersEnd int
	preambulaEnd  int

	// result caches

	httpType         HTTPType
	hasBody          bool
	allHeadersParsed bool
	isBroken         bool

	// remove padding if needed
	padding [5]bool

	lastParseHeadersState int
	parsedHeaders         map[string][]byte
	reqMethod             []byte
	reqPath               []byte
	proto                 []byte

	statusCode   int
	statusString []byte

	mu *sync.RWMutex
}

// Parse the given body
func Parse(body []byte) *Result {
	res := &Result{
		body: body,
		mu:   &sync.RWMutex{},
	}

	res.loadHeaders()
	res.loadBody()

	return res
}
