package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/farahsrw/manageProject/utils"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateMiddleware(schema string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body map[string]interface{}
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}

			err = json.Unmarshal(bodyBytes, &body)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}

			schemaLoader := gojsonschema.NewStringLoader(schema)
			documentLoader := gojsonschema.NewGoLoader(body)

			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil || !result.Valid() {
				var errs []string
				for _, err := range result.Errors() {
					errs = append(errs, err.String())
				}
				utils.RespondError(w, http.StatusBadRequest, strings.Join(errs, ", "))
				return
			}

			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			next.ServeHTTP(w, r)
		})
	}
}
