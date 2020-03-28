package main

import (
	"bytes"
)

// Mat ...
type Mat [][]int

var symbolMap = map[int]string{0: "o", 1: "x"}

// NewMat ...
func NewMat(x, y int) Mat {
	m := make(Mat, y)
	for i := 0; i < x; i++ {
		m[i] = make([]int, x)
	}
	return m
}

// Grow ...
func (m *Mat) Grow() {
	m.growY()
	m.growX()
}

func (m *Mat) growX() {
	for i := range *m {
		(*m)[i] = append((*m)[i], 0)
	}
}

func (m *Mat) growY() {
	*m = append(*m, make([]int, len(*m)))
}

func (m *Mat) String() string {
	var out bytes.Buffer
	for i, v := range *m {
		if i == 0 {
			out.WriteByte(' ')
			for i := range v {
				out.WriteByte(' ')
				out.WriteByte(byte(i + 97))
			}
			out.WriteByte('\n')
		}
		for j, v := range v {
			if j == 0 {
				out.WriteByte(byte(i + 97))
			}
			out.WriteString(" " + symbolMap[v])
		}
		if i < len(*m)-1 {
			out.WriteByte('\n')
		}
	}
	return out.String()
}
