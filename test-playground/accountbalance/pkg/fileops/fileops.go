package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Fileops struct {
	FileName string
	Value    float64
}

func (f *Fileops) WriteFloatToFile() {
	valueText := fmt.Sprint(f.Value)
	os.WriteFile(f.FileName, []byte(valueText), 0644)
}

func (f *Fileops) GetFloatFromFile() (float64, error) {
	data, err := os.ReadFile(f.FileName)
	if err != nil {
		return 0, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}
	return value, nil
}

func NewFileOps(f string, v float64) *Fileops {
	return &Fileops{
		FileName: f,
		Value:    v,
	}
}
