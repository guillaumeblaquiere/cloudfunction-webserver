package function_webserver

import (
	"fmt"
	"net/http"
	"path"
	"reflect"
	"strings"
)

var mux = newMux()

func Webserver(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/static/", serveStatic)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/subroute/login", login)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login from /subroute/login")
}

type Empty struct{}

var functionSourceCodeDir = "/workspace/src/" + reflect.TypeOf(Empty{}).PkgPath()

func serveStatic(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path
	if strings.HasSuffix(file, "/") {
		// Set the default page if the path finish by /
		file += "index.html"
	}
	http.ServeFile(w, r, path.Clean(functionSourceCodeDir+file))
}
