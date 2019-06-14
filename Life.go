package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"src/github.com/faiface/pixel/imdraw"
	"src/golang.org/x/image/colornames"
)

type Cell struct {
	x, y, size float64
	color      color.Color
}

func (c Cell) draw(drawer *imdraw.IMDraw) {
	drawer.Color = c.color
	size := c.size + 1 //+1 for next calculations, so the size is correct on screen
	drawer.Push(pixel.V((c.x-1)*size+1, c.y*size))
	drawer.Push(pixel.V(c.x*size, c.y*size))
	drawer.Line(size - 1)
}

func (c Cell) print() {
	fmt.Printf("Cell : %f;%f\n", c.x, c.y)
}

func generateCell(size, x, y float64) *Cell {
	var c = &Cell{x, y, size, colornames.White}
	return c
}

func run() {
	const cellSize = 3
	const resolutionX = 1280
	const resolutionY = 720
	const gridX = resolutionX / cellSize
	const gridY = resolutionY / cellSize

	cfg := pixelgl.WindowConfig{
		Title:  "Game of life",
		Bounds: pixel.R(0, 0, resolutionX, resolutionY),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	initialPositions := [][]float64{
		[]float64{cellSize, 3, 30},
		[]float64{cellSize, 4, 30},
		[]float64{cellSize, 5, 30},
		[]float64{cellSize, 6, 60},
		[]float64{cellSize, 6, 61},
		[]float64{cellSize, 6, 62},
	}

	var cells = make([][]*Cell, gridX)
	for i := 0; i < len(cells); i++ {
		cells[i] = make([]*Cell, gridY)
	}

	for _, position := range initialPositions {
		cells[int(position[1])][int(position[2])] = generateCell(position[0], position[1], position[2])
	}

	var draw = imdraw.New(nil)

	for !win.Closed() {

		win.Update()
		draw.Clear()
		win.Clear(colornames.Black)

		for x, cellTable := range cells {
			for y, cell := range cellTable {

				neighborsNumber := 0

				if cell != nil {
					fmt.Printf("Checking cell %d;%d\n", x, y)
				}

				for i := -1; i < 2; i++ {
					for j := -1; j < 2; j++ {

						if j == 0 && i == 0 {
							continue
						}

						if x+i < gridX && y+j < gridY && x+i >= 0 && y+j >= 0 {
							if cells[x+i][y+j] != nil {
								neighborsNumber++
								fmt.Printf("\tNeighbor n° %d on %d;%d\n", neighborsNumber, x+i, y+j)
							}
						}

					}
				}

				if cell != nil {
					fmt.Printf("----Neighbor n°%d----\n", neighborsNumber)
				}

				if neighborsNumber < 2 || neighborsNumber > 3 {
					cells[x][y] = nil
				} else if cell == nil && neighborsNumber == 3 {
					cells[x][y] = generateCell(cellSize, float64(x), float64(y))
				}

				if cell == nil {
					continue
				}

				cell.draw(draw)
			}
		}
		draw.Draw(win)
	}
}

func main() {
	pixelgl.Run(run)
}
