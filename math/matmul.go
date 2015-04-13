// Ke's library for matrix mulitiplication

package math

import ( 
	"fmt"
	"errors"
	"sync"
)

type Matrix struct {
	rows int
	cols int
	data []float32
}

func NewMatrix(rows, cols int) * Matrix {
	m := new(Matrix)
	m.rows = rows
	m.cols = cols
	m.data = make([]float32, rows*cols)
	return m
}

func (m *Matrix) Print() {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			fmt.Printf("%v ", m.Get(r, c))
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix) Set(r, c int, x float32) {
	m.data[r * m.cols + c] = x
}

func (m *Matrix) Get(r, c int) float32 {
	return m.data[r * m.cols + c]
}

// slow multiplication
func Mult(a, b Matrix) (c *Matrix, err error) {
	c = new(Matrix)
	// a is M x K
	// b is K x N
	// c should be M x N
	m, k, n := a.rows, a.cols, b.cols
	if k != b.rows {
		// dimension mismatch
		return c, errors.New("Dimension mismatch")
	}
	c = NewMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var x float32 = 0
			for v := 0; v < k; v++ {
				x += a.Get(i, v) * b.Get(v, j)
			}
			c.Set(i, j, x)
		}
	}
	return c, nil
}


// A faster version using go routines
func FastMul(a, b Matrix) (c *Matrix, err error) {
        c = new(Matrix)
        // a is M x K                                                                                                                              
        // b is K x N
        // c should be M x N                                                                                                                          
        m, k, n := a.rows, a.cols, b.cols
        if k != b.rows {
                // dimension mismatch 
                return c, errors.New("Dimension mismatch")
        }
        c = NewMatrix(m, n)
	var g sync.WaitGroup
        for i := 0; i < m; i++ {
                for j := 0; j < n; j++ {
			g.Add(1)
			// Use  go routine to run this in a thread
			go func(i, j int) {
				defer g.Done()
				var x float32 = 0
				for v := 0; v < k; v++ {
					x += a.Get(i, v) * b.Get(v, j)
				}
				c.Set(i, j, x)
			}(i, j)
		}
	}
	g.Wait()
	return c, nil
}
