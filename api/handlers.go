package api

import (
	"encoding/csv"
	"fmt"
	"github.com/mlhan1993/league_interview/pkg/matrix"
	"github.com/mlhan1993/league_interview/pkg/utils"
	"net/http"
	"strconv"
)

type MatrixProcessor interface {
	Invert(matrix.Matrix) (matrix.Matrix, error)
	Flatten(matrix.Matrix) []int
	Sum(matrix.Matrix) (int, error)
	Multiply(matrix.Matrix) (int, error)
}

type MatrixHandlers struct {
	processor MatrixProcessor
}

func NewMatrixHandlers(helper *matrix.Processor) *MatrixHandlers {
	return &MatrixHandlers{processor: helper}
}

func (m *MatrixHandlers) Echo(w http.ResponseWriter, r *http.Request) {
	myMatrix, err := m.readMatrix(r)
	if err != nil {
		m.writeError(w, err)
	}
	w.Write([]byte(myMatrix.ToString()))
}

func (m *MatrixHandlers) Invert(w http.ResponseWriter, r *http.Request) {
	myMatrix, err := m.readMatrix(r)
	if err != nil {
		m.writeError(w, err)
	}
	inverted, err := m.processor.Invert(myMatrix)
	if err != nil {
		m.writeError(w, err)
	}
	w.Write([]byte(inverted.ToString()))
}

// Flatten handler for POST /flatten
func (m *MatrixHandlers) Flatten(w http.ResponseWriter, r *http.Request) {
	myMatrix, err := m.readMatrix(r)
	if err != nil {
		m.writeError(w, err)
	}
	flatten := m.processor.Flatten(myMatrix)
	if err != nil {
		m.writeError(w, err)
	}

	w.Write([]byte(utils.IntArrToStr(flatten)))

}

// Sum handler for POST /sum
func (m *MatrixHandlers) Sum(w http.ResponseWriter, r *http.Request) {
	myMatrix, err := m.readMatrix(r)
	if err != nil {
		m.writeError(w, err)
	}
	sum, err := m.processor.Sum(myMatrix)
	if err != nil {
		m.writeError(w, err)
	}

	w.Write([]byte(strconv.Itoa(sum)))

}

// Multiply handler for POST /multiply
func (m *MatrixHandlers) Multiply(w http.ResponseWriter, r *http.Request) {

	myMatrix, err := m.readMatrix(r)
	if err != nil {
		m.writeError(w, err)
	}
	total, err := m.processor.Multiply(myMatrix)
	if err != nil {
		m.writeError(w, err)
	}

	w.Write([]byte(strconv.Itoa(total)))
}

// readMatrix reads form file and parse the file into Matrix.
func (m *MatrixHandlers) readMatrix(r *http.Request) (matrix.Matrix, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("error finding form key 'file': %s", err.Error())
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var result matrix.Matrix
	for _, row := range records {
		var tmp []int
		for _, col := range row {
			intVal, err := strconv.Atoi(col)
			if err != nil {
				return nil, fmt.Errorf("unable to convert %s to integer", col)
			}
			tmp = append(tmp, intVal)
		}
		result = append(result, tmp)
	}

	err = result.Validate()
	if err != nil {
		return nil, err
	}

	return result, err
}

// writeError writes response message in case of an error
func (m *MatrixHandlers) writeError(w http.ResponseWriter, err error) {
	w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
}
