package connections

import (
	"net/http"

	"github.com/HawkinsThomas/connections-be/src/router"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("CloudFunctionEntry", CloudFunctionEntry)
}

func CloudFunctionEntry(w http.ResponseWriter, r *http.Request) {
	rtr := router.InitializeRouter()
	rtr.ServeHTTP(w, r)
}
