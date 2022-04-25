package middlewares

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/thoas/go-funk"
	"tdd.com/v1/utils"
)

type UserClaimsInfo struct {
	AgreedToTermsAndConditions bool
	AcknowledgedToAdvisoryNote bool
	Permissions                Permissions
	AdditionalUserInfo         *AdditionalUserInfo `json:",omitempty"`
}

type AdditionalUserInfo struct {
	Name          string
	Email         string
	Agency        string
	Team          string
	OpsUnit       string
	Designation   string
	OpsUnitType   string
	LastLoginDate *time.Time
}

type Permissions struct {
	AgencyId             int `json:",omitempty"`
	OpsUnitId            int `json:",omitempty"`
	TeamId               int `json:",omitempty"`
	UserId               int `json:",omitempty"`
	AccessLevel          string
	AccessibleSubDomains []int `json:"SubDomains"`
	DataUploadAllowedFor []string
	CanScreenClient      bool
}

const TxContextKey utils.ContextKey = "TRANSACTION_IN_REQUEST_CONTEXT"
const UsernameContextKey utils.ContextKey = "USERNAME_CONTEXT_KEY"
const AccessLevelContextKey utils.ContextKey = "ACCESS_LEVEL_CONTEXT_KEY"
const AgreedToTnCContextKey utils.ContextKey = "AGREED_TO_TERMS_AND_CONDITIONS_CONTEXT_KEY"
const PermissionsKey utils.ContextKey = "PERMISSIONS_KEY"
const UserTypeKey utils.ContextKey = "USER_TYPE_KEY"
const AdditionalUserInfoKey utils.ContextKey = "ADDITIONAL_USER_INFO_KEY"
const ApplicationLoggerKey utils.ContextKey = "APPLICATION_LOGGER_KEY"
const TokenIssuedAtKey utils.ContextKey = "TOKEN_ISSUED_AT_KEY"

const CorrelationIDHeader string = "X-Correlation-ID"



func OptionsRequestHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_WHITELIST"))
	header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	header.Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization")
	header.Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusNoContent)
}

func MiddlewareChain(next func(http.ResponseWriter, *http.Request, httprouter.Params), middlewares ...func(func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers := funk.Reverse(middlewares).([]func(func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params))
		current := next

		for _, middleware := range handlers {
			current = middleware(current)
		}

		current = AttachAppLogger(current)

		//Required for all requests
		AllowCORSForWhitelist(current)(w, r, ps)
	}
}

// AllowCORSForWhitelist set the Access-Control-Allow-Origin and Access-Control-Allow-Credentials values in the header
func AllowCORSForWhitelist(next func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_WHITELIST"))
		header.Set("Access-Control-Allow-Credentials", "true")
		next(w, r, ps)
	}
}

// AddRouterParams decorate handler's function signature with the httprouter.Params argument
func AddRouterParams(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		next(w, r)
	}
}





// RefreshRDSConnectToken trigger a new request for a fresh RDS IAM access token
func RefreshRDSConnectToken(connector *utils.DBConnector) func(func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(next func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			if connector.Session != nil {
				connector.DB.Options().Password = utils.GetRDSConnectToken(connector.Options, connector.Creds)
			}
			next(w, r, ps)
		}
	}
}


// AttachAppLogger attach a AppLogger object to the request context if correlation id header is present in the request
func AttachAppLogger(next func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		correlationID := r.Header.Get(CorrelationIDHeader)

		if correlationID != "" {
			appLogger := utils.CreateAppLogger(correlationID)
			ctx := context.WithValue(r.Context(), ApplicationLoggerKey, appLogger)

			next(w, r.WithContext(ctx), ps)
		} else {
			next(w, r, ps)
		}
	}
}
