package main

import (
	"fmt"
)

type matriks []([]int)

func valid(n []int, kota int) bool {
	var check bool
	check = false
	for _, v := range n {
		if v == kota {
			check = true
			break
		}
	}
	return check
}

func main() {
	var (
		dMatrix                                                                                       matriks
		skippedCities, tempList, f                                                                    []int
		numofCities, occupiedCity, arrive, depart, totalDistance, tempDistance, distance, newDistance int
		check                                                                                         bool
	)
	fmt.Scan(&numofCities)
	for i := 0; i < numofCities; i += 1 {
		tempList = []int{}
		for j := 0; j < numofCities; j += 1 {
			fmt.Scan(&distance)
			tempList = append(tempList, distance)
		}
		dMatrix = append(dMatrix, tempList)
		f = append(f, -1)
	}
	fmt.Scan(&depart)
	fmt.Scan(&arrive)
	occupiedCity = depart
	skippedCities = append(skippedCities, occupiedCity)
	totalDistance = 0
	for skippedCities[len(skippedCities)-1] != arrive {
		newDistance = -1
		for n := 0; n < numofCities; n += 1 {
			check = valid(skippedCities, n)
			if !check {
				if dMatrix[occupiedCity][n] != -1 {
					if f[n] == -1 {
						f[n] = totalDistance + dMatrix[occupiedCity][n]
					} else {
						tempDistance = totalDistance + dMatrix[occupiedCity][n]
						if tempDistance < f[n] {
							f[n] = tempDistance
						}
					}
				}
				if newDistance == -1 || (f[n] < f[newDistance] && f[n] != -1) {
					newDistance = n
				}
			}
		}
		totalDistance = f[newDistance]
		occupiedCity = newDistance
		skippedCities = append(skippedCities, occupiedCity)
	}
	fmt.Println(totalDistance)
}
