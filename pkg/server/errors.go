package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/identbase/getting/pkg/resource/representor"
)

/*
Standard Errors in Matrix errcode format. */
const CodeBadRequest = "M_BAD_REQUEST"                                       // 400
const CodeUnauthorized = "M_UNAUTHORIZED"                                    // 401
const CodePaymentRequired = "M_PAYMENT_REQUIRED"                             // 402
const CodeForbidden = "M_FORBIDDEN"                                          // 403
const CodeNotFound = "M_NOT_FOUND"                                           // 404
const CodeMethodNotAllowed = "M_METHOD_NOT_ALLOWED"                          // 405
const CodeNotAcceptable = "M_NOT_ACCEPTABLE"                                 // 406
const CodeProxyAuthRequired = "M_PROXY_AUTH_REQUIRED"                        // 407
const CodeRequestTimeout = "M_REQUEST_TIMEOUT"                               // 408
const CodeConflict = "M_CONFLICT"                                            // 409
const CodeGone = "M_GONE"                                                    // 410
const CodeLengthRequired = "M_LENGTH_REQUIRED"                               // 411
const CodePreconditionFailed = "M_PRECONDITION_FAILED"                       // 412
const CodeRequestEntityTooLarge = "M_REQUEST_ENTITY_TOO_LARGE"               // 413
const CodeRequestURITooLong = "M_REQUEST_URI_TOO_LONG"                       // 414
const CodeUnsupportedMediaType = "M_UNSUPPORTED_MEDIA_TYPE"                  // 415
const CodeRequestedRangeNotSatisfiable = "M_REQUESTED_RANGE_NOT_SATISFIABLE" // 416
const CodeExpectationFailed = "M_EXPECTATION_FAILED"                         // 417
const CodeTeapot = "M_TEAPOT"                                                // 418
const CodeMisdirectedRequest = "M_MISDIRECTED_REQUEST"                       // 421
const CodeUnprocessableEntity = "M_UNPROCESSABLE_ENTITY"                     // 422
const CodeLocked = "M_LOCKED"                                                // 423
const CodeFailedDependency = "M_FAILED_DEPENDENCY"                           // 424
const CodeTooEarly = "M_TOO_EARLY"                                           // 425
const CodeUpgradeRequired = "M_UPGRADE_REQUIRED"                             // 426
const CodePreconditionRequired = "M_PRECONDITION_REQUIRED"                   // 428
const CodeTooManyRequests = "M_TOO_MANY_REQUESTS"                            // 429
const CodeRequestHeaderFieldsTooLarge = "M_REQUEST_HEADER_FIELDS_TOO_LARGE"  // 431
const CodeUnavailableForLegalReasons = "M_UNAVAILABLE_FOR_LEGAL_REASONS"     // 451

const CodeInternalServerError = "M_INTERNAL_SERVER_ERROR"                     // 500
const CodeNotImplemented = "M_NOT_IMPLEMENTED"                                // 501
const CodeBadGateway = "M_BAD_GATEWAY"                                        // 502
const CodeServiceUnavailable = "M_SERVICE_UNAVAILABLE"                        // 503
const CodeGatewayTimeout = "M_GATEWAY_TIMEOUT"                                // 504
const CodeHTTPVersionNotSupported = "M_HTTP_VERSION_NOT_SUPPORTED"            // 505
const CodeVariantAlsoNegotiates = "M_VARIANT_ALSO_NEGOTIATES"                 // 506
const CodeInsufficientStorage = "M_INSUFFICIENT_STORAGE"                      // 507
const CodeLoopDetected = "M_LOOP_DETECTED"                                    // 508
const CodeNotExtended = "M_NOT_EXTENDED"                                      // 510
const CodeNetworkAuthenticationRequired = "M_NETWORK_AUTHENTICATION_REQUIRED" // 511

/*
Error provides the basic JSON fields that are expected by the Matrix
specification, and Extra for additional details.  */
type Error struct {
	Extra   interface{} `json:"extra,omitempty"`
	Code    string      `json:"errcode"`
	Message string      `json:"error"`
}

type HALError struct {
	representor.HALBody
}

func NewHALError(m string, c string) *HALError {
	e := HALError{
		representor.HALBody{
			Links: map[string][]representor.HALLink{
				"self": []representor.HALLink{
					representor.HALLink{
						HRef:  fmt.Sprintf("/_matrix/errors/%s", strings.ToLower(c)),
						Title: c,
					},
				},
			},
			Properties: map[string]interface{}{
				"errcode": c,
				"error":   m,
			},
			Embedded: map[string]representor.HALBody{},
		},
	}

	return &e
}

/*
Errors is a mapping of net/http status codes to Error*'s. */
var Errors = map[int]*HALError{
	http.StatusBadRequest:                   ErrorBadRequest,
	http.StatusUnauthorized:                 ErrorUnauthorized,
	http.StatusPaymentRequired:              ErrorPaymentRequired,
	http.StatusForbidden:                    ErrorForbidden,
	http.StatusNotFound:                     ErrorNotFound,
	http.StatusMethodNotAllowed:             ErrorMethodNotAllowed,
	http.StatusNotAcceptable:                ErrorNotAcceptable,
	http.StatusProxyAuthRequired:            ErrorProxyAuthRequired,
	http.StatusRequestTimeout:               ErrorRequestTimeout,
	http.StatusConflict:                     ErrorConflict,
	http.StatusGone:                         ErrorGone,
	http.StatusLengthRequired:               ErrorLengthRequired,
	http.StatusPreconditionFailed:           ErrorPreconditionFailed,
	http.StatusRequestEntityTooLarge:        ErrorRequestEntityTooLarge,
	http.StatusRequestURITooLong:            ErrorRequestURITooLong,
	http.StatusUnsupportedMediaType:         ErrorUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable: ErrorRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed:            ErrorExpectationFailed,
	http.StatusTeapot:                       ErrorTeapot,
	http.StatusMisdirectedRequest:           ErrorMisdirectedRequest,
	http.StatusUnprocessableEntity:          ErrorUnprocessableEntity,
	http.StatusLocked:                       ErrorLocked,
	http.StatusFailedDependency:             ErrorFailedDependency,
	http.StatusTooEarly:                     ErrorTooEarly,
	http.StatusRequestHeaderFieldsTooLarge:  ErrorRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons:   ErrorUnavailableForLegalReasons,

	// 500s
	http.StatusInternalServerError:           ErrorInternalServerError,
	http.StatusNotImplemented:                ErrorNotImplemented,
	http.StatusBadGateway:                    ErrorBadGateway,
	http.StatusServiceUnavailable:            ErrorServiceUnavailable,
	http.StatusGatewayTimeout:                ErrorGatewayTimeout,
	http.StatusHTTPVersionNotSupported:       ErrorHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates:         ErrorVariantAlsoNegotiates,
	http.StatusInsufficientStorage:           ErrorInsufficientStorage,
	http.StatusLoopDetected:                  ErrorLoopDetected,
	http.StatusNotExtended:                   ErrorNotExtended,
	http.StatusNetworkAuthenticationRequired: ErrorNetworkAuthenticationRequired,
}

/*
Error* are net/http status codes to Error objects. */
var ErrorBadRequest = NewHALError("Bad Request", CodeBadRequest)
var ErrorUnauthorized = NewHALError("Unauthorized", CodeUnauthorized)
var ErrorPaymentRequired = NewHALError("Payment Required", CodePaymentRequired)
var ErrorForbidden = NewHALError("Forbidden", CodeForbidden)
var ErrorNotFound = NewHALError("Not Found", CodeNotFound)
var ErrorMethodNotAllowed = NewHALError("Method Not Allowed", CodeMethodNotAllowed)
var ErrorNotAcceptable = NewHALError("Not Acceptable", CodeNotAcceptable)
var ErrorProxyAuthRequired = NewHALError("Proxy Auth Required", CodeProxyAuthRequired)
var ErrorRequestTimeout = NewHALError("Request Timeout", CodeRequestTimeout)
var ErrorConflict = NewHALError("Conflict", CodeConflict)
var ErrorGone = NewHALError("Gone", CodeGone)
var ErrorLengthRequired = NewHALError("Length Required", CodeLengthRequired)
var ErrorPreconditionFailed = NewHALError("Precondition Failed", CodePreconditionFailed)
var ErrorRequestEntityTooLarge = NewHALError("Request Entity Too Large", CodeRequestEntityTooLarge)
var ErrorRequestURITooLong = NewHALError("Request URI Too Long", CodeRequestURITooLong)
var ErrorUnsupportedMediaType = NewHALError("Unsupported Media Type", CodeUnsupportedMediaType)
var ErrorRequestedRangeNotSatisfiable = NewHALError("Requested Range Not Satisfiable", CodeRequestedRangeNotSatisfiable)
var ErrorExpectationFailed = NewHALError("Expectation Failed", CodeExpectationFailed)
var ErrorTeapot = NewHALError("Teapot", CodeTeapot)
var ErrorMisdirectedRequest = NewHALError("Misdirected Request", CodeMisdirectedRequest)
var ErrorUnprocessableEntity = NewHALError("Unprocessable Entity", CodeUnprocessableEntity)
var ErrorLocked = NewHALError("Locked", CodeLocked)
var ErrorFailedDependency = NewHALError("Failed Dependency", CodeFailedDependency)
var ErrorTooEarly = NewHALError("Too Early", CodeTooEarly)
var ErrorUpgradeRequired = NewHALError("Upgrade Required", CodeUpgradeRequired)
var ErrorPreconditionRequired = NewHALError("Precondition Required", CodePreconditionRequired)
var ErrorTooManyRequests = NewHALError("Too Many Requests", CodeTooManyRequests)
var ErrorRequestHeaderFieldsTooLarge = NewHALError("Request Header Fields Too Large", CodeRequestHeaderFieldsTooLarge)
var ErrorUnavailableForLegalReasons = NewHALError("Unavailable For Legal Reasons", CodeUnavailableForLegalReasons)

var ErrorInternalServerError = NewHALError("Internal Server Error", CodeInternalServerError)
var ErrorNotImplemented = NewHALError("Not Implemented", CodeNotImplemented)
var ErrorBadGateway = NewHALError("Bad Gateway", CodeBadGateway)
var ErrorServiceUnavailable = NewHALError("Service Unavailable", CodeServiceUnavailable)
var ErrorGatewayTimeout = NewHALError("Gateway Timeout", CodeGatewayTimeout)
var ErrorHTTPVersionNotSupported = NewHALError("HTTP Version Not Supported", CodeHTTPVersionNotSupported)
var ErrorVariantAlsoNegotiates = NewHALError("Variant Also Negotiates", CodeVariantAlsoNegotiates)
var ErrorInsufficientStorage = NewHALError("Insufficient Storage", CodeInsufficientStorage)
var ErrorLoopDetected = NewHALError("Loop Detected", CodeLoopDetected)
var ErrorNotExtended = NewHALError("Not Extended", CodeNotExtended)
var ErrorNetworkAuthenticationRequired = NewHALError("Network Authentication Required", CodeNetworkAuthenticationRequired)
