package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"src/github.com/faiface/pixel/imdraw"
	"src/golang.org/x/image/colornames"
	"time"
)

var cells = make([][]*imdraw.IMDraw, 100)

func generateCell(size, x, y float64) *imdraw.IMDraw {
	cell := imdraw.New(nil)
	cell.Color = colornames.Aliceblue
	cell.Push(pixel.V((x-1)*size+1, y*size))
	cell.Push(pixel.V(x*size, y*size))
	cell.Line(size)
	return cell
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of life",
		Bounds: pixel.R(0, 0, 1152, 648),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	initialPositions := [][]float64{
		[]float64{2, 10, 20},
		[]float64{2, 11, 20},
		[]float64{2, 12, 20},
		[]float64{2, 25, 60},
		[]float64{2, 25, 61},
		[]float64{2, 25, 62},
	}

	for index := range cells {
		cells[index] = make([]*imdraw.IMDraw, 100)
	}

	for _, position := range initialPositions {
		cells[int(position[1])][int(position[2])] = generateCell(position[0], position[1], position[2])
	}

	for !win.Closed() {
		win.Update()

		fmt.Printf("Update\n")

		time.Sleep(200 * time.Millisecond)

		for xIndex := range cells {
			for yIndex := range cells[xIndex] {

				neighborCount := 0

				for xOffset := -1; xOffset <= 1; xOffset++ {
					for yOffset := -1; yOffset <= 1; yOffset++ {

						if xOffset == 0 && yOffset == 0 {
							continue
						}

						neighborX := xIndex + xOffset
						neighborY := yIndex + yOffset

						if neighborX >= 0 && neighborX <= 99 && neighborY >= 0 && neighborY <= 99 {
							if cells[neighborX][neighborY] != nil {
								neighborCount++
							}
						}
					}
				}

				if neighborCount < 2 || neighborCount > 3 {
					cells[xIndex][yIndex] = nil
					continue
				} else if neighborCount == 3 {
					cells[xIndex][yIndex] = generateCell(2, float64(xIndex), float64(yIndex))
				}

				if cells[xIndex][yIndex] == nil {
					continue
				}

				cells[xIndex][yIndex].Draw(win)
			}
		}
	}
}

func main() {
	pixelgl.Run(run)
}
