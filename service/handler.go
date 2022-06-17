package service

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/GrokkingSystemDesign/snowflake"
	"github.com/gin-gonic/gin"
)

func init() {
	m, err := snowflake.NewMachine(0)
	if err != nil {
		panic(err)
	}
	machine = m
}

var (
	storage sync.Map
	machine *snowflake.Machine
)

type request struct {
	Long string `json:"long"`
}

func isURLValid(rawURL string) bool {
	_, e := url.ParseRequestURI(rawURL)
	return e == nil
}

func encodeURL(rawURL string) string {
	var (
		shortURL strings.Builder
		id, _    = machine.Generate()
		i64      = id.Int64()
		hash     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)
	for i64 != 0 {
		shortURL.WriteByte(hash[i64%int64(len(hash))])
		i64 /= int64(len(hash))
	}
	return shortURL.String()
}

// HandleURLShorten converts long URL to a shorter one
func HandleURLShorten(c *gin.Context) {
	var body request
	bindErr := c.BindJSON(&body)
	if bindErr != nil || len(body.Long) <= 0 || !isURLValid(body.Long) {
		debug.PrintStack()
		c.String(http.StatusBadRequest, "invalid body")
		return
	}
	shortURL := encodeURL(body.Long)
	storage.Store(shortURL, body.Long)
	c.String(http.StatusOK, fmt.Sprintf("http://localhost:8080/%s", shortURL))
}

// HandleRedirect retrieves related long URL and redirect
func HandleRedirect(c *gin.Context) {
	shortURL := c.Param("url")
	longURL, ok := storage.Load(shortURL)
	if !ok {
		c.String(http.StatusNotFound, "short url missing in storage")
		return
	}
	c.Redirect(http.StatusPermanentRedirect, longURL.(string))
}
