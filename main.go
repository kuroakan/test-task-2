package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	if len(os.Args) < 5 {
		log.Fatal("Usage: go run main.go X Y Z S where:\nX - First Coordinate\nY - Second Coordinate\nZ - Matrix size, default 100\nS - show created matrix, 1 to show, 0 to don't")
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("incorrect X value")
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("incorrect Y value")
	}
	z, err := strconv.Atoi(os.Args[3])
	if err != nil || z <= 0 {
		fmt.Println("Incorrect Z value, using default(100)")
		z = 101
	}
	s, err := strconv.Atoi(os.Args[4])
	if err != nil || s > 1 || s < 0 {
		log.Fatal("Incorrect S value, please choose 0 or 1")
	}

	if x < 1 || x >= z+1 || y < 1 || y >= z+1 {
		log.Fatal("Coordinates must be From 1 to Z\nExample: X = 2, Y = 2 then Z must be atleast 2")
	}

	x--
	y--

	matrix := make([][]any, z)
	for i := range matrix {
		matrix[i] = make([]any, z)
	}

	for i := 0; i < z; i++ {
		for j := 0; j < z; j++ {
			switch rand.Intn(4) {
			case 0:
				matrix[i][j] = rand.Intn(10)
			case 1:
				matrix[i][j] = rand.Float32() * 10
			case 2:
				matrix[i][j] = rand.Float64() * 10
			case 3:
				matrix[i][j] = "str"
			}
		}
	}

	toFloat := func(val any) float64 {
		switch v := val.(type) {
		case int:
			return float64(v)
		case float32:
			return float64(v)
		case float64:
			return v
		default:
			return 0
		}
	}

	sum := 0.0
	for i := 0; i < z; i++ {
		sum += toFloat(matrix[i][y])
		sum += toFloat(matrix[x][i])
	}
	sum -= toFloat(matrix[x][y])

	if s == 1 {
		formatValue := func(value any) string {
			switch v := value.(type) {
			case int:
				return fmt.Sprintf("%6d", v)
			case float32, float64:
				return fmt.Sprintf("%6.2f", v)
			case string:
				return fmt.Sprintf("%6s", v)
			default:
				return "      "
			}
		}

		fmt.Println("\nGenerated Matrix:")
		for i := 0; i < z; i++ {
			for j := 0; j < z; j++ {
				fmt.Print(formatValue(matrix[i][j]), " ")
			}
			fmt.Println()
		}
	}

	fmt.Printf("Sum of row %d and column %d: %.2f\n", x+1, y+1, sum)
}
