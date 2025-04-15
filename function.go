package main

import (
	"image"
	"sync"
	
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var mutex sync.Mutex
var wg sync.WaitGroup

func FindCoordAround(nbDimentions int, nbAround int) *[][]int {
	find := make([][]int, 0)
	for i := 1; i <= nbDimentions; i++ {
		wg.Add(1)
		go GetNCoordDifferent(nbDimentions, i, nbAround, &find)
	}
	wg.Wait()
	return &find
}

func GetNCoordDifferent(nbDimentions int, nbDifferent int, nbAround int, find *[][]int) {
	defer wg.Done()
	coord := make([]int, nbDifferent)
	for i := 0; i < nbDifferent; i++ {
		coord[i] = -nbAround + i
	}
	mutex.Lock()
	wg.Add(1)
	go GetCompletVariationOfCoord(nbDimentions, find, &coord)
	for ;;{
		modified := false
		cursor := len(coord)-1
		for ;!modified; {
			if cursor < 0 || cursor == len(coord) {
				return
			}
			if coord[cursor] + len(coord) -1 - cursor < nbAround {
				modified = true
				coord[cursor]++
				val := coord[cursor]+1
				cursor++
				for ;cursor < len(coord); {
					coord[cursor] = val
					cursor++
					val++
				}
			} else {
				cursor--
			}
		}
		mutex.Lock()
		wg.Add(1)
		go GetCompletVariationOfCoord(nbDimentions, find, &coord)
	}
}

func GetCompletVariationOfCoord(nbDimentions int, find *[][]int, ogPartialCoord *[]int) {
	defer wg.Done()
	partialCoord := make([]int, len(*ogPartialCoord))
	copy(partialCoord, *ogPartialCoord)
	mutex.Unlock()
	coord := make([]int, nbDimentions)
	for i := 0; i < nbDimentions; i++ {
		if i < len(partialCoord) {
			coord[i] = partialCoord[i]
		} else {
			coord[i] = partialCoord[0]
		}
	}
	mutex.Lock()
	wg.Add(1)
	go getPermutationOfCoord(find, &coord)
	for ;;{
		modified := false
		cursor := len(coord)-1
		for ;!modified; {
			if cursor == len(partialCoord) -1{
				return
			}
			if coord[cursor] != partialCoord[len(partialCoord)-1] {
				modified = true
				afterCursor := 0
				for ;afterCursor < len(partialCoord) && coord[cursor] != partialCoord[afterCursor]; {
					afterCursor++
				}
				afterCursor++
				for ;cursor < len(coord); {
					coord[cursor] = partialCoord[afterCursor]
					cursor ++
				}
			} else {
				cursor--
			}
		}
		mutex.Lock()
		wg.Add(1)
		go getPermutationOfCoord(find, &coord)
	}
}

func getPermutationOfCoord(find *[][]int, coord *[]int) {
	defer wg.Done()
	//https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
	var helper func([]int, int)
	tempFind := make([][]int, 0)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			if !isDoublon(&tempFind, &tmp) {
				tempFind = append(tempFind, tmp)
			}
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	tcoord := make([]int, len(*coord))
	copy(tcoord, *coord)
	mutex.Unlock()
	helper(tcoord, len(tcoord))
	*find = append(*find, tempFind...)
}

func isDoublon(find *[][]int, coord *[]int) bool {
	for i := 0; i < len(*find); i++ {
		for j := 0; j < len(*coord); j++ {
			if (*find)[i][j] != (*coord)[j] {
				break
			}
			if j == len(*coord)-1 {
				return true
			}
		}
	}
	return false
}

func getCursorPosition(win *pixelgl.Window) (int,int) {
	return int(win.MousePosition().X), 280- int(win.MousePosition().Y)
}

func drawPixel(img *image.RGBA, x int, y int) {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c := 126 - (127 * (float64(i) + float64(j)) / 5)
			addColor(img, x+i, y+j, c)
			addColor(img, x-i, y+j, c)
			addColor(img, x+i, y-j, c)
			addColor(img, x-i, y-j, c)
		}
	}
}

func addColor(img *image.RGBA, x int, y int, c float64) {
	if x < 0 || y < 0 || x >= img.Bounds().Max.X || y >= img.Bounds().Max.Y {
		return
	}
	oldC := img.RGBAAt(x, y)
	c = c + float64(oldC.R)
	if c > 255 {
		c = 255
	}
	img.Set(x, y, pixel.RGB(c, c, c))
}

func imgToPoint(img *image.RGBA) *Point {
	point := Point{}
	for y := 0; y < 28; y++ {
		for x := 0; x < 28; x++ {
			point.coord = append(point.coord, getValueOfPoint(img, x*10, y*10))
		}
	}
	return &point
}

func getValueOfPoint(img *image.RGBA, x int,y int) float64 {
	c := float64(0)
	for i:=0; i<10; i++ {
		for j:=0; j<10; j++ {
			c += float64(img.RGBAAt(x+i, y+j).R)
		}
	}
	return c/100
}

func FindCoordAround2(nbDimentions int, nbAround int) *[][]int {
	find := make([][]int, 0)
	coord := make([]int,nbDimentions)
	for i:=0; i<nbDimentions; i++ {
		coord[i] = -nbAround;
	}
	wg.Add(1)
	go recFunc(coord,nbAround,nbDimentions-1,&find)
	wg.Wait()
	return &find;
}

func recFunc(coord []int, nbAround int, idx int, find *[][]int) {
	defer wg.Done()

	for i:= 0; i<=idx; i++ {
		coord[i] += 1
		if coord[i] <= nbAround {
			wg.Add(1)
			tcoord := make([]int, len(coord))
			copy(tcoord, coord)
			go recFunc(tcoord,nbAround,i,find)
		}
		coord[i] -= 1
	}
	mutex.Lock()
	*find = append(*find, coord)
	mutex.Unlock()
}