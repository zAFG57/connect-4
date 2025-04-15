package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func loadCsvToNnnarv(nnnarv *Nnnarv, fichier string) {
	fmt.Println("lecture du csv")
	defer fmt.Println("fin de la lecture du csv")
	file, err := os.Open(fichier)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			continue
		}
		fmt.Println(fields)
		value := fields[7*7]
		coords := make([]float64, len(fields)-1)
		for i := 0; i < len(fields)-1; i++ {
			coord, err := strconv.ParseFloat(fields[i], 64)
			if err != nil {
				coord = float64(0)
			}
			coords[i] = coord
		}
		point := Point{valueStr: value, coord: coords}
		nnnarv.AddPoint(point)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}