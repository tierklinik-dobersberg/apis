package spa

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tierklinik-dobersberg/apis/pkg/log"
)

type hookedResponseWriter struct {
	http.ResponseWriter
	got404 bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == http.StatusNotFound {
		// Don't actually write the 404 header, just set a flag.
		hrw.got404 = true
	} else {
		hrw.ResponseWriter.WriteHeader(status)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.got404 {
		// No-op, but pretend that we wrote len(p) bytes to the writer.
		return len(p), nil
	}

	return hrw.ResponseWriter.Write(p)
}

func intercept404(handler, on404 http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hookedWriter := &hookedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(hookedWriter, r)

		if hookedWriter.got404 {
			log.L(r.Context()).Infof("got 404 for %s", r.URL.String())
			on404.ServeHTTP(w, r)
		}
	})
}

func serveFileContents(file string, files http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Restrict only to instances where the browser is looking for an HTML file
		/*
			if !strings.Contains(r.Header.Get("Accept"), "text/html") {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 not found")

				return
			}
		*/

		// Open the file and return its contents using http.ServeContent
		index, err := files.Open(file)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)

			return
		}

		fi, err := index.Stat()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)

			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(w, r, fi.Name(), fi.ModTime(), index)
	}
}

func ServeSPA(fs http.FileSystem, indexFile string) http.Handler {
	return intercept404(http.FileServer(fs), serveFileContents(indexFile, fs))
}

func SetCORSHeaders(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()

	// FIXME
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Set("Access-Control-Allow-Origin", "*")

	reqMethod := r.Header.Get("Access-Control-Request-Method")
	reqHeaderList := strings.Join(r.Header.Values("Access-Control-Request-Headers"), ",")

	if reqMethod != "" {
		headers.Set("Access-Control-Allow-Methods", strings.ToUpper(reqMethod))
	}
	if reqHeaderList != "" {
		headers.Set("Access-Control-Allow-Headers", reqHeaderList)
	}

	headers.Set("Access-Control-Allow-Credentials", "true")
}
