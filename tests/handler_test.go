package tests

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/mamude/internal/server"
	"github.com/gin-gonic/gin/binding"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMECSV               = "text/csv"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
	MIMEYAML2             = "application/yaml"
	MIMETOML              = "application/toml"
)

type testFile struct {
	Fieldname string
	Filename  string
	Content   []byte
}

var s struct {
	FileValue multipart.FileHeader `form:"file"`
}

func TestHomeHandler(t *testing.T) {
	t.Setenv("TEMPLATE", "../cmd/web/templates/*")

	server := &server.Server{}
	router := server.RegisterRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSendFileHandler(t *testing.T) {
	t.Setenv("TEMPLATE", "../cmd/web/templates/*")

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	server := &server.Server{DB: mock}
	router := server.RegisterRoutes()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	file := testFile{"file", "file1.txt", []byte("hello")}

	mw := multipart.NewWriter(&body)
	fw, err := mw.CreateFormFile(file.Fieldname, file.Filename)
	assert.NoError(t, err)

	n, err := fw.Write(file.Content)
	assert.NoError(t, err)
	assert.Equal(t, len(file.Content), n)

	mw.Close()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/send_file", &body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+mw.Boundary())

	err = binding.FormMultipart.Bind(req, &s)
	assert.NoError(t, err)

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestSendEmptyFileHandler(t *testing.T) {
	t.Setenv("TEMPLATE", "../cmd/web/templates/*")

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	server := &server.Server{DB: mock}
	router := server.RegisterRoutes()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	file := testFile{}

	mw := multipart.NewWriter(&body)
	fw, err := mw.CreateFormFile(file.Fieldname, file.Filename)
	assert.NoError(t, err)

	n, err := fw.Write(file.Content)
	assert.NoError(t, err)
	assert.Equal(t, len(file.Content), n)

	mw.Close()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/send_file", &body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+mw.Boundary())

	err = binding.FormMultipart.Bind(req, &s)
	assert.NoError(t, err)

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "selecione um arquivo!", w.Body.String())
}

func TestSendInvalidFileHandler(t *testing.T) {
	t.Setenv("TEMPLATE", "../cmd/web/templates/*")

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	server := &server.Server{DB: mock}
	router := server.RegisterRoutes()
	w := httptest.NewRecorder()

	var body bytes.Buffer
	file := testFile{"file", "file1.csv", []byte("hello")}

	mw := multipart.NewWriter(&body)
	fw, err := mw.CreateFormFile(file.Fieldname, file.Filename)
	assert.NoError(t, err)

	n, err := fw.Write(file.Content)
	assert.NoError(t, err)
	assert.Equal(t, len(file.Content), n)

	mw.Close()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/send_file", &body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMECSV)

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
