package main

import (
	"fmt"
)

type Slope struct {
	row int
	col int
}

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println(calcEncounteredTrees(Slope{1,3}))
}

func part2() {
	slopes := []Slope{Slope{1, 1},
		Slope{1, 3},
		Slope{1, 5},
		Slope{1, 7},
		Slope{2, 1}}

	product := 1
	for _, slope := range slopes {
		product *= calcEncounteredTrees(slope)
	}
	fmt.Println(product)
}

const rowLength = 31

func calcEncounteredTrees(slope Slope) (encounteredTrees int) {
	encounteredTrees = 0
	col := 0
	row := 0
	for row < len(puzzleInputDay3) {
		if string(puzzleInputDay3[row][col]) == "#" {
			encounteredTrees++
		}
		row = row + slope.row
		col = (col + slope.col) % rowLength
	}
	return
}

var puzzleInputDay3 = []string{
".#..........#...#...#..#.......",
".###...#.#.##..###..#...#...#..",
"#.....#................#...#.#.",
"#.....#..###.............#....#",
"......#.....#....#...##.....###",
"....#........#.#......##....#.#",
"..#.......##..#.#.#............",
"#.............#..#...#.#...#...",
".#...........#.#....#..##......",
"......#..##..#....#....#...##..",
"....#.##.#####..#.##..........#",
"..#.#......#.#.#....#.....#....",
"...###.##......#..#.#...#...#..",
"...#..#.#..#..#.......#........",
"...#....#..#...........#.#.....",
"....#.........###.#....#...#...",
"....#..##.....#.##....##.#.....",
"........#.#.#.....#........#...",
"..#..#.....#.#...#.#...#.#.....",
"....#..........#....#....#...##",
".##...#..#...##....#..#.#....#.",
".#....##..#...#................",
"..#.###.........#.###.....#....",
"....#..#.......###.#...........",
"#...#...#.#...........#.#......",
".#..#.......##.....##...#......",
"....####.#..#.#.#...........#..",
".##...#..#..#.#....##.....#..##",
"...#......##....#...#.#.###....",
"##.#...........#.........#...#.",
"...........#...#...........##..",
".....#....#...........#........",
"...#..#.........#...#....#.##..",
".....##.........#...#........##",
"....#....#..#.#...#...##.#.....",
"...#.#..#...#...........#..#...",
".....#.#.....#....#...#....#...",
".#.............#..##..........#",
"..........#......#..##.....###.",
"..#....#........#.#.....##...#.",
"#..#......#.#.##......#.#.##...",
".....#..#.........#...#.#.#.#.#",
"#.#...#.......#.#..##.##.....##",
".....#......##......#.......#..",
"#.....#...##.#.#........#......",
"#..........#.#...#.......#.....",
"..#..#........#........#.......",
"...#....#....#..####.#....#...#",
"#.............#.....##....#..#.",
"##....#.....###..##....#......#",
"#.....#...#.#.............#....",
".#.#..##..##.#..#....#.#.#...#.",
".#...#..#.....#..#.#.#..#...##.",
"..#.#.#.#.#.#....##...#........",
".......##.....#..........#...#.",
"...#..#...#...........#....#...",
".....#..#....#..#.##...#.......",
"..##..#.......#.#..#....#......",
"...#...............#.#..#......",
"....#........#...#....#...#.#..",
"...#...#..........##....##.#...",
"..###.#.##.............#..#.#.#",
"##.......##.#..#.#.#.....#.#.#.",
"..#####...#......##...#........",
"...#.##...#................#..#",
"..#......#...#....#.#..##..#...",
"#.#.........#............#.....",
"##.............#.#.....#......#",
"....#.......#..#..##....#.#....",
"...#...##....#.........#..#....",
"...####.....#...........#....#.",
"#.#........##....#..#..#...#...",
"....#.#.###..........#........#",
"#.#......#.....#.##....#.#...#.",
"#....##.#..##..#.#.............",
".#.....##..#..................#",
"...#.#........#...#.#........#.",
"..#....#......#.....##........#",
"....#...#....#...#.....#.##....",
"...#........#.......##.........",
".#.##......#......#....##......",
".#...#...###.#............#..#.",
".#...........#.#.#....#...#..#.",
".#.....#....#.....#...#........",
".#..#.....#............#.#.##.#",
"...###.#.............#..##.....",
"...#.#.##.#..#..........#..#...",
".#.#.#....#..#...............##",
".......#.#..#...#.#.#........#.",
"....#.#...#..##....#........#.#",
"..........#...#.......#..#....#",
"...###.....#.#....#.....##.....",
"#......#..#..#........#.#...#..",
"#......#....#..#.#.............",
"...#....#........#...#..#......",
"...#..###........#.#.........##",
"#......#.#..###..#........###..",
".#.#......#.#..#.#.#.#.....#..#",
"#....#.....#..##.....#.........",
"....#......#...#..#..#.#.##.#..",
"........#.#...#...#..#...#.#..#",
".....##........#...#....#...#..",
"....#...##..#........#....##.#.",
"...............#.....#......##.",
"..##.....#.....#.#.............",
".....#.#...........##.#.....#..",
".#..##..#.##.#...##.#....#....#",
".##.....#.##......#....#..#..#.",
".......#.##......#....#...#.#..",
".#........#......#...##.#....#.",
".........#..........#.......###",
"#.#.........#..#..#....#...#...",
".......#.........#......#.#.#..",
".......#...........#....#....#.",
".###...##.#.#..........#...#..#",
"....#.....#...#..#.............",
".......##........#..#.......#..",
"....##..#.#....#....#..#...#..#",
"..#.####.....#.........#.#....#",
"..............#.#..#.....#...#.",
".....#.............#..........#",
"..##.#...#.....#....#.#....##..",
".#...#.......#..####..#..#...#.",
"#..........#................##.",
"......##.....#.................",
"..##...#.#..........##.#...#...",
"....#.#.#.#...##...#...#...####",
".............##..#.###...#.....",
"#.#....#.#..#..##........#..##.",
".....#.#...............#.......",
"...#..##......#..##...........#",
"#..#....#...........##..#......",
".##....#.#....###.......#..#...",
".....#..#.#....##...#......#...",
".#.........#####......#...#...#",
".......#.#.....#.....#.......#.",
"#....#.......###.......#..#....",
"#......##.###...#.......#......",
".......#...#......#....#..#....",
".#.####.......#...#.##.........",
"................##.#......#....",
"......##....#.#......#......#..",
"....##...##....#.........#.....",
"......#.#..............##.#...#",
"....#.#......#.#.............#.",
".#.#..####...#................#",
"....#.#.#.#......##...##......#",
".....#.#..#......#....#......#.",
"..........#.#.....#.......#...#",
"..##......##.#...##.#......#..#",
"...#............#..#...###.....",
".#.#..###..#.......##...#.....#",
".#....#.#.......#.....##....#..",
"#.............###...##.#.#...#.",
"#........#.#........#.#...#.#.#",
"##..#.................#....#...",
"...#.#...#..#.#..##....#...#...",
"#.....#.......#..............#.",
".......###...##..#.....#.......",
"#.#.........#..#.#.........#...",
".#.#............#.....##.....#.",
"........#....#....#.......#....",
"...#.#....#..#.##....#.#......#",
".#.....#.#..#...........#.#.#..",
"#......#..#......##.#.#.#.#..#.",
".......#.#..#......#.#.#..#.#.#",
"..........#...#..........#.##..",
".#.#..####.......#..........#..",
"......#.#.....#..#..#..#.....#.",
".....##..#.#.#..#..#...#.....##",
"............#.#....#.#....#....",
"..............#..#...#...#.....",
".....#......#.......#.....#....",
"..##....#..#...........#..##...",
"###...#.##..#.#...####....###..",
"..#.#.....#.........#....#..###",
"##...........##.............#..",
"....##..............#.........#",
"...#...##....#.#..#...##.....#.",
"..#..##...#.......#..#..#.....#",
"...#...#....####........##.#...",
"....#........#..#.#.........#..",
".#..........#...#..#.#.#......#",
"....#.#.....#.........#....#...",
"...#....#...##.......#...#.....",
"....#..#.......#.##.##.##...#..",
"##....##........#........##....",
".#.#..#...........#.....#...#..",
"...#.##...##..#...#...##.......",
".....#..###................#.#.",
"...#........##.#....##.....#.##",
"...#...#..##...#...#.#...#.....",
".#......#...#..#.##.......#...#",
".....#.......###.##...#........",
"#.....#..#........##.##.#.##..#",
"....#..............##.##...#...",
"#..........#..................#",
"..##.......#..........#..#..##.",
".#....###.#..#.........###....#",
".#....#.##..............#.##.##",
".#.##.#....#.......#.#......#..",
".#............#.#.....#........",
"..#......#.......#.............",
"#.#...#........##...#.#......#.",
"....#.........#........##..#...",
"..........##.....#.#......#....",
".##.#..#....#.......#...#...##.",
".#................#...#.##.....",
"....###.......#..#..#.........#",
".#.....#..##...###......#.....#",
".#.##..........#..#..#........#",
".......#.##..............#...##",
"#...#.#.#.......#..#......#.##.",
".#....#.#......#...#..........#",
".....#........##....#.##.....#.",
".#....................#..#.#.#.",
".....#.........#....#.......#.#",
".....#.#..##..#.....#..#.......",
"...#..#..#...#.....#....#....#.",
"#.....#.#.#..........#..#.#.#..",
".....##..##.....#.#..#.........",
"#.#..##....##......##...#.##..#",
"..##..#.....#..#..........##...",
"......#.#...#..#.......##.....#",
"..#.#.......#.#......#.........",
".....#........##..#.....####.#.",
".#.....#........#.......#..##..",
"......#...#....#.##...#.......#",
"..##..................#..#.....",
".....###.#..##...#.............",
"...##...##...#......#....#....#",
"#........#.#..........##..#....",
"#........#....#..........#...#.",
"...##.#.##..#...##......#......",
"#........##....#.#..##.....#..#",
"...####......#..#......#.#.....",
".#......#...#...#.#.....##....#",
".....###..##..#...#..........##",
"##.##....#...#.................",
"...##.#.......#.###......#..#..",
".....#.#.#.......#.......#..#.#",
"#...#...#.##..#....###.......#.",
".#.#..##.....#....#...##.......",
".....#..........#....#...#.##..",
"..........#....#...#...........",
".#....#..#...#...#.......#....#",
"#..#..............#.....####.##",
".......#....###....#....#.#.#..",
"###.#........##.#.......#......",
"#..#...#..#......#.............",
"#...###..#...#..#..##.#.###.#..",
"..#..#...##......##............",
".#..#.......#..###..##...#.....",
"....#..#..##.#.#.....##...#.#.#",
"....#....#.....#..#....#.......",
"..##..#....#.#...##..#.........",
".....#....#...........#.#......",
"...#........#.#..#..#......#..#",
".#...##....#....#.#.##......#.#",
"..#...........#..###.##.....#..",
".#.######.#..##.......#..#.....",
".....#..#......##.#.#...#......",
"....#....#..#.....#.......#.#.#",
".....#........##.....#.....#.##",
"........#....#...#...#.#.#...#.",
"...#.#.....#...........#.....#.",
"#.#.#...###......#.....#.....#.",
".#..........#.....#.......##...",
"#................#.#.....#.####",
".#......#......#.#..##.#.##....",
"..........#....#...........###.",
".##....#..####..#####..........",
"##.......##............#.....#.",
"...#.....#...#....#.......#....",
".#....##......#.#...#....#.....",
"....#............##..........#.",
".#....#....#.....#.#...........",
".............##.#.##...#.#.#...",
"..#............#.#..##.#....##.",
"#.....#...##..........#.#.#...#",
"......#............#..........#",
"..##..#.....#........#.##..#..#",
"#..#.#..##.#.....##.#..........",
"#..#...#.#..#......##.......##.",
".##......#...........##.....#..",
"...#.....#.....#..#....#.......",
".....#...............#........#",
".......#.....##..#..##..#.#.#..",
"#.#.....#..#..........##...#...",
"#..#......#.................#.#",
".##...#....#...#...#.......#...",
".#........##........#..........",
"........#..........#.........#.",
".....#.##..#.......#........#..",
"..##..#..#...##..#.#....#......",
"......#........#.##.....#.#....",
".#...#.#.........#..#.#.#.#..#.",
".#..#.#...#............#.#..#..",
"....#.................#...#..##",
".........##.....#.#.#......####",
"...............#....##.#.#.....",
"....##..#....#......#....#.....",
"....##.#...#....#.#..#...#..#..",
"..##......#.#..#........#.#.#..",
".........#.#................##.",
"##.....#.....##..##.#........#.",
"###....#..#..#..#..#.##..##.#..",
".....##..#...........##..#.#...",
"....#..#..#..#....#...#.#....#.",
"#....#............#..#....###..",
"....#..#.............#....##.#.",
"...#.................#...#.....",
".##...#....#..#..#........#....",
"...#.#..#...#.#......#....#....",
"...#.......##..........#...#.#.",
"...##..#.......#........#...#..",
".....#.#.#....#..##......##...#",
"....##......#........##....##..",
"..#..........#.#.##.....#......",
"..................#..#..#..###.",
".#..............#.#..#.#..#.###",
"..#....#....#......#..##..#...#",
"#.........#..#..#...........#.."}
