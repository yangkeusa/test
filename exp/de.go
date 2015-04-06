// defer

package exp 

import "log"

func A(log *log.Logger, x int) {
	log.Printf("A: x = %d\n", x)
}

func B(log *log.Logger, y int) {
	defer A(log, y+1)
	defer A(log, y+10)
	log.Printf("B: y = %d\n", y)
}

func C(log *log.Logger, z int) {
	defer A(log, z+1)
	defer B(log, z+10)
        log.Printf("C: z = %d\n", z)
}
