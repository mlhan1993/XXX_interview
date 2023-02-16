package api

import (
	"github.com/mlhan1993/league_interview/pkg/matrix"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

func TestMatrixHandlers_Echo(t *testing.T) {

	processor := matrix.Processor{}
	handlers := NewMatrixHandlers(&processor)
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	defer pr.Close()
	defer pw.Close()

	content := "1,2,3\n4,5,6\n7,8,9"
	go writeContent(writer, content)

	request := httptest.NewRequest("POST", "/echo", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	responseRecorder := httptest.NewRecorder()

	handlers.Echo(responseRecorder, request)
	response := responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, content, string(body))
}

func TestMatrixHandlers_Sum(t *testing.T) {

	processor := matrix.Processor{}
	handlers := NewMatrixHandlers(&processor)
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	defer pr.Close()
	defer pw.Close()

	content := "1,2,3\n4,5,6\n7,8,9"
	go writeContent(writer, content)

	request := httptest.NewRequest("POST", "/sum", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	responseRecorder := httptest.NewRecorder()

	handlers.Sum(responseRecorder, request)
	response := responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "45", string(body))
}

func TestMatrixHandlers_Invert(t *testing.T) {

	processor := matrix.Processor{}
	handlers := NewMatrixHandlers(&processor)
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	defer pr.Close()
	defer pw.Close()

	content := "1,2,3\n4,5,6\n7,8,9"
	go writeContent(writer, content)

	request := httptest.NewRequest("POST", "/invert", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	responseRecorder := httptest.NewRecorder()

	handlers.Invert(responseRecorder, request)
	response := responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	expected := "1,4,7\n2,5,8\n3,6,9"
	assert.Equal(t, expected, string(body))
}

func TestMatrixHandlers_Multiply(t *testing.T) {

	processor := matrix.Processor{}
	handlers := NewMatrixHandlers(&processor)
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	defer pr.Close()
	defer pw.Close()

	content := "1,1,1\n1,1,1\n1,1,9"
	go writeContent(writer, content)

	request := httptest.NewRequest("POST", "/multiply", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	responseRecorder := httptest.NewRecorder()

	handlers.Multiply(responseRecorder, request)
	response := responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	expected := "9"
	assert.Equal(t, expected, string(body))
}

func TestMatrixHandlers_Flatten(t *testing.T) {

	processor := matrix.Processor{}
	handlers := NewMatrixHandlers(&processor)
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	defer pr.Close()
	defer pw.Close()

	content := "1,1,1\n1,1,1\n1,1,9"
	go writeContent(writer, content)

	request := httptest.NewRequest("POST", "/flatten", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	responseRecorder := httptest.NewRecorder()

	handlers.Flatten(responseRecorder, request)
	response := responseRecorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fail()
	}

	expected := "1,1,1,1,1,1,1,1,9"
	assert.Equal(t, expected, string(body))
}

// writeContent helps with creates file key and write the content to be consumed by the corresponding reader.
// The function is blocking until the reader consumes the data and therefore should be used carefully.
func writeContent(writer *multipart.Writer, content string) {
	file, _ := writer.CreateFormFile("file", "test")
	io.WriteString(file, content)
	writer.Close()
}
