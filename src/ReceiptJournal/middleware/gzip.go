package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleWare struct {
	Next http.Handler
}

func (gm *GzipMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encoding := r.Header.Get("Accept-Encoding")
	if !strings.Contains(encoding, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}

	w.Header().Add("Content-Encoding", "gzip")
	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()
	var rw http.ResponseWriter
	if pusher, ok := w.(http.Pusher); ok {
		rw = gzipPusherResponseWriter{
			gzipResponseWriter: gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzipwriter,
			},
			Pusher: pusher}
	} else {
		rw = gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzipwriter,
		}
	}

	gm.Next.ServeHTTP(rw, r)

}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}

type gzipPusherResponseWriter struct {
	gzipResponseWriter
	http.Pusher
}
