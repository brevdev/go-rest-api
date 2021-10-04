package album

import (
	"net/http"
	"testing"
	"time"

	"github.com/qiangxue/go-rest-api/internal/auth"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/test"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.Album{
		{"album_123", "album123", time.Now(), time.Now()},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{"get all", "GET", "/albums", "", nil, http.StatusOK, `*"total_count":1*`},
		{"get 123", "GET", "/albums/album_123", "", nil, http.StatusOK, `*album_123*`},
		{"get not album id", "GET", "/albums/random_123", "", nil, http.StatusBadRequest, "*invalid album id*"},
		{"get unknown", "GET", "/albums/album_1234", "", nil, http.StatusNotFound, ""},
		{"create ok", "POST", "/albums", `{"name":"test"}`, header, http.StatusCreated, "*test*"},
		{"create ok count", "GET", "/albums", "", nil, http.StatusOK, `*"total_count":2*`},
		{"create auth error", "POST", "/albums", `{"name":"test"}`, nil, http.StatusUnauthorized, ""},
		{"create input error", "POST", "/albums", `"name":"test"}`, header, http.StatusBadRequest, ""},
		{"update ok", "PUT", "/albums/album_123", `{"name":"albumxyz"}`, header, http.StatusOK, "*albumxyz*"},
		{"update verify", "GET", "/albums/album_123", "", nil, http.StatusOK, `*albumxyz*`},
		{"update auth error", "PUT", "/albums/album_123", `{"name":"albumxyz"}`, nil, http.StatusUnauthorized, ""},
		{"update input error", "PUT", "/albums/album_123", `"name":"albumxyz"}`, header, http.StatusBadRequest, ""},
		{"delete ok", "DELETE", "/albums/album_123", ``, header, http.StatusOK, "*albumxyz*"},
		{"delete verify", "DELETE", "/albums/album_123", ``, header, http.StatusNotFound, ""},
		{"delete auth error", "DELETE", "/albums/album_123", ``, nil, http.StatusUnauthorized, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
