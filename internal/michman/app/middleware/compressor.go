package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func NewCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:   w,
		gzW: gzip.NewWriter(w),
	}
}

type compressWriter struct {
	w   http.ResponseWriter
	gzW *gzip.Writer
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	return c.gzW.Write(p)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Add("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера.
func (c *compressWriter) Close() error {
	return c.gzW.Close()
}

func Compressor(compressMethod string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if compressMethod != "" {

				//decompress r
				contentEncoding := r.Header.Get("Content-Encoding")
				contentIsGzip := strings.Contains(contentEncoding, "gzip")

				if contentIsGzip {
					decomBody, err := gzip.NewReader(r.Body)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					defer decomBody.Close()
					r.Body = decomBody
				} else {
					http.Error(w, "", http.StatusInternalServerError)
					return
				}

				//compress w
				acceptEncoding := r.Header.Get("Accept-Encoding")
				clientCanGzip := strings.Contains(acceptEncoding, "gzip")
				if clientCanGzip {
					cResp := NewCompressWriter(w)
					defer cResp.Close()
					w = cResp
				}

			}
			next.ServeHTTP(w, r)
		})
	}
}
