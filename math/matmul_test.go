// tests for matmul

package math

import ("math"; "testing")

func TestMatMul(t *testing.T) {
	m, k, n := 3, 4, 5

	a := NewMatrix(m, k)
        b := NewMatrix(k, n)
        for i := 0; i < m; i++ {
                for j := 0; j < k; j++ {
                        a.Set(i, j, float32(i + j))
                }
        }
        for i := 0; i < k; i++ {
                for j := 0; j < k; j++ {
                        b.Set(i, j, float32(i + j))
                }
        }
	c, err := Mult(*a, *b)
	if err != nil {
		t.Error("MatMul error: ", err)
	}
	for i := 0; i < m; i++ {
                for j := 0; j < n; j++ {
                        var x float64 = 0
                        for v := 0; v < k; v++ {
                                x += float64(a.Get(i, v)) * float64(b.Get(v, j))
                        }
                        if math.Abs(x - float64(c.Get(i, j))) > 1e-5 {
				t.Error("Numeric Error")
			}
                }
        }
}

func TestFastMaxMul(t *testing.T) {
	m, k, n := 3, 4, 5

        a := NewMatrix(m, k)
        b := NewMatrix(k, n)
        for i := 0; i < m; i++ {
                for j := 0; j < k; j++ {
                        a.Set(i, j, float32(i + j))
                }
	}
        for i := 0; i < k; i++ {
                for j := 0; j < k; j++ {
                        b.Set(i, j, float32(i + j))
		}
        }
	c, err := FastMul(*a, *b)
	if err != nil {
		t.Error("MatMul error: ", err)
	}
        for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var x float64 = 0
			for v := 0; v < k; v++ {
				x += float64(a.Get(i, v)) * float64(b.Get(v, j))
                        }
                        if math.Abs(x - float64(c.Get(i, j))) > 1e-5 {
                                t.Error("Numeric Error")
                        }
                }
        }
}


func BenchmarkMatMul(b *testing.B) {
	b.StopTimer()
	m, k, n := 256, 256, 256

        X := NewMatrix(m, k)
        Y := NewMatrix(k, n)
        for i := 0; i < m; i++ {
                for j := 0; j < k; j++ {
                        X.Set(i, j, float32(i + j))
                }
        }
        for i := 0; i < k; i++ {
                for j := 0; j < k; j++ {
                        Y.Set(i, j, float32(i + j))
                }
        }
//        begin := time.Now()
	b.StartTimer()
        for i := 0; i < b.N; i++ {
                c, _ := Mult(*X, *Y)
                _ = c
	}
	b.StopTimer()
//        end := time.Now()
//	dur := end.Sub(begin)

        num_ops := int64(b.N * m * k * n * 2)
	b.SetBytes(num_ops)
//        tflops := float64(num_ops) / float64(dur.Nanoseconds())
//        fmt.Printf("Took %v to finish %v operations = %v TFLOPS\n", dur, num_ops, tflops)
}

func BenchmarkFastMatMul(b *testing.B) {
	b.StopTimer()
        m, k, n := 256, 256, 256

        X := NewMatrix(m, k)
        Y := NewMatrix(k, n)
        for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			X.Set(i, j, float32(i + j))
		}
	}
        for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			Y.Set(i, j, float32(i + j))
		}
	}
	//        begin := time.Now()                                                                                                                          
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c, _ := FastMul(*X, *Y)
		_ = c
	}
        b.StopTimer()
	//        end := time.Now()                                                                                                                            
	//      dur := end.Sub(begin)                                                                                                                          

	num_ops := int64(b.N * m * k * n * 2)
        b.SetBytes(num_ops)
	//        tflops := float64(num_ops) / float64(dur.Nanoseconds())                                                                                      
	//        fmt.Printf("Took %v to finish %v operations = %v TFLOPS\n", dur, num_ops, tflops)                                                            
}
