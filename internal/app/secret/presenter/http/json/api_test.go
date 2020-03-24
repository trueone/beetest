package json_test

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/trueone/beetest/internal/app/secret"
	"github.com/trueone/beetest/internal/app/secret/data"
	"github.com/trueone/beetest/internal/app/secret/presenter/http/json"
)

var (
	registry = secret.New()
	handler  = json.New(registry)
)

func TestHandler_Create(t *testing.T) {
	e := echo.New()

	cases := []struct {
		name   string
		data   string
		status int
	}{
		{
			name:   "OK",
			data:   `{"secret":"secret","expireAfterViews":1,"expireAfter":0}`,
			status: http.StatusOK,
		},
		{
			name:   "Failed",
			data:   `{"secret":"secret","expireAfterViews":0,"expireAfter":0}`,
			status: http.StatusMethodNotAllowed,
		},
	}

	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)

	for _, c := range cases {
		req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)

		if assert.NoError(t, handler.Create(ctx), c.name) {
			assert.Equal(t, c.status, rec.Code)
		}
	}
}

func TestHandler_Get(t *testing.T) {
	e := echo.New()

	_, err := registry.Create().Execute(data.SecretRequest{
		Secret:           "secret_get",
		ExpireAfterViews: 1,
		ExpireAfter:      0,
	})

	if assert.NoError(t, err) {
		hash := md5.Sum([]byte("secret_get"))
		hashString := hex.EncodeToString(hash[:])

		cases := []struct {
			name   string
			data   string
			status int
		}{
			{
				name:   "OK",
				data:   hashString,
				status: http.StatusOK,
			},
			{
				name:   "View limit exceeded",
				data:   hashString,
				status: http.StatusNotFound,
			},
			{
				name:   "Not found",
				data:   "secret_not_found",
				status: http.StatusNotFound,
			},
		}

		var (
			req *http.Request
			rec *httptest.ResponseRecorder
			ctx echo.Context
		)

		for _, c := range cases {
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			rec = httptest.NewRecorder()
			ctx = e.NewContext(req, rec)

			ctx.SetPath("/secret/:hash")
			ctx.SetParamNames("hash")
			ctx.SetParamValues(c.data)

			if assert.NoError(t, handler.Get(ctx), c.name) {
				assert.Equal(t, c.status, rec.Code)
			}
		}
	}
}
