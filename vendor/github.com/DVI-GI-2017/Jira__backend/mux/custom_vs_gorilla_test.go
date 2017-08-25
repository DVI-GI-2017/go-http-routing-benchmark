package mux

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"fmt"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/gorilla/mux"
)

var simpleBench = &benchmarkData{
	basePath: "/api/v1",

	patternGorilla: "/users/{id:[[:xdigit:]]{24}}",
	patternCustom:  "/users/:id",

	matchedPath: "/api/v1/users/599ce026ff64e74a60086508",

	pathParams: map[string]string{"id": "599ce026ff64e74a60086508"},
}

func BenchmarkSimpleGorilla(b *testing.B) {
	benchmark(b, simpleBench, initGorillaRouter(simpleBench))
}

func BenchmarkSimpleCustom(b *testing.B) {
	benchmark(b, simpleBench, initCustomRouter(simpleBench))
}

var nestedBench = &benchmarkData{
	basePath: "/api/v1",

	patternGorilla: "/projects/{project_id:[[:xdigit:]]{24}}/tasks/{task_id:[[:xdigit:]]{24}}",
	patternCustom:  "/projects/:project_id/tasks/:task_id",

	matchedPath: "/api/v1/projects/599ce026ff64e74a60086508/tasks/599ca654ff64e71ad83a1bc6",

	pathParams: map[string]string{
		"project_id": "599ce026ff64e74a60086508",
		"task_id":    "599ca654ff64e71ad83a1bc6",
	},
}

func BenchmarkNestedGorilla(b *testing.B) {
	benchmark(b, nestedBench, initGorillaRouter(nestedBench))
}

func BenchmarkNestedCustom(b *testing.B) {
	benchmark(b, nestedBench, initCustomRouter(nestedBench))
}

// Executes simple benchmark atop of given data
func benchmark(b *testing.B, data *benchmarkData, router http.Handler) {
	helper := newGetHelper(data.matchedPath)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(helper.w, helper.r)
		if !data.Ok(helper.w.Body.Bytes()) {
			b.Fatalf("%s", helper.w.Body.String())
		}
		helper.clear()
	}
}

// Helper structure to store bench data
type benchmarkData struct {
	basePath string

	patternGorilla string
	patternCustom  string

	matchedPath string

	pathParams map[string]string
}

func (b benchmarkData) Ok(response []byte) bool {
	for key, value := range b.pathParams {
		str, err := jsonparser.GetString(response, key)
		if err != nil {
			return false
		}
		if str != value {
			return false
		}
	}
	return true
}

// Helpers to work with gorilla router
func initGorillaRouter(data *benchmarkData) *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix(data.basePath).Subrouter()

	apiRouter.Path(data.patternGorilla).Methods(http.MethodGet).HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			vars := mux.Vars(req)

			data, err := json.Marshal(vars)
			if err != nil {
				log.Panicf("can not marshal data: %v", err)
			}
			w.Write(data)
		})

	return router
}

// Helper to work with custom router
func initCustomRouter(data *benchmarkData) *router {
	router, err := NewRouter(data.basePath)
	if err != nil {
		log.Panicf("can not create router: %v", err)
	}

	err = router.Get(data.patternCustom, func(w http.ResponseWriter, req *http.Request) {
		params := Params(req)
		data, err := json.Marshal(params.PathParams)
		if err != nil {
			log.Panicf("can not marshal data: %v", err)
		}
		w.Write(data)
	})
	if err != nil {
		log.Panicf("%v", err)
	}

	return router
}

// Converts flat map to json
func flatMapToJson(data map[string]string) []byte {
	template := fmt.Sprintf(`{%s}`, strings.Repeat(`"%s":"%s",`, len(data)))

	keys := make([]string, 0, len(data))
	values := make([]string, 0, len(data))

	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}

	toInsert := make([]string, 0, 2*len(data))
	for i := 0; i < len(data); i++ {
		toInsert = append(toInsert, keys[i], values[i])
	}

	str := fmt.Sprintf(template, toInsert)
	return bytes.NewBufferString(str).Bytes()
}

// Helper to mock get requests
type getHelper struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func newGetHelper(path string) *getHelper {
	r, _ := http.NewRequest(http.MethodGet, path, nil)

	return &getHelper{
		w: httptest.NewRecorder(),
		r: r,
	}
}

func (helper *getHelper) clear() {
	helper.w.Body = bytes.NewBuffer(make([]byte, 0, 100))
}
