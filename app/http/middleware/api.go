package middleware

import (
	"github.com/lanvard/contract/inter"
	foundation "github.com/lanvard/foundation/http/middleware"
)

// This standard api middleware must be
// run before you can make adjustments to the request
var Api = []inter.HttpMiddleware{
	foundation.RequestJsonBody{},
}
