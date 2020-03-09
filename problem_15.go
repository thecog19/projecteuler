package main

import (
        "time"
        //"math"
        "fmt"
        "strconv"
)

type move struct {
	c_cell *cell
	multiplier int
}

type cell struct {
	x_coord int
	y_coord int
	down *cell
	right *cell
	name string
}

func main() {
	x,y := 21,21
        start := time.Now()
	maze := generate_maze(x,y)
        print(maze)
	res := a_star(maze,&maze[x-1][y-1])
	fmt.Printf("Result: %f", res)
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func print(maze [][]cell){
        var i, j int
        for  i = 0; i < len(maze); i++ {
                for j = 0; j < len(maze[0]); j++ {
                        fmt.Printf("[ ]")
                }
                fmt.Printf("\n" )
        }
}

func generate_maze(size_x int, size_y int) [][]cell {
	var maze [][]cell
	//starting at the top left corner, x = 0, y = 0
	// +y is down, +x is right
	y := 0
	x := 0
	for y < size_y {
		maze = append(maze, make([]cell,size_y))
		for x < size_x {
			c := cell {x,y, nil, nil, strconv.Itoa(x) + " " + strconv.Itoa(y)}
			maze[y][x] = c
		x += 1
		}
		y += 1
		x = 0
	}
	y = 0
	x = 0
	for y < size_y {
                for x < size_x {
			if(x + 1 < size_x){
				maze[y][x].right = &maze[y][x+1]
			}
			if(y + 1 <size_y){
				maze[y][x].down = &maze[y+1][x]
			}
			x += 1
                }
                y += 1
                x = 0
        }
	return maze
}

func a_star(maze [][]cell, target *cell) int{
	var list_of_moves []*move
	list_of_moves = make([]*move,0)
	branches := 1
	start := move{&maze[0][0],1}
	list_of_moves = append(list_of_moves, &start)
	node_visitors := make(map[string]bool)
	move_finder := make(map[string]*move)
	for len(list_of_moves)> 0 {
		//pop off the first element
		var path *cell
		var c_move *move
		c_move, list_of_moves = list_of_moves[0], list_of_moves[1:]
		path = c_move.c_cell
		multiplier := c_move.multiplier
		if(path.down == target || path.right == target){
			continue
		}else if(path.down != nil && path.right != nil){
			branches += multiplier
			if(!node_visitors[path.name]){
				node_visitors[path.name] = true
				move_finder[path.name] = c_move
				right_move := move{path.right,multiplier}
				down_move := move{path.down,multiplier}
				move_finder[path.right.name] = &right_move
				move_finder[path.down.name] = &down_move
				list_of_moves = append(list_of_moves, &right_move ,&down_move)
			}else{
				move_finder[path.right.name].multiplier += multiplier
				move_finder[path.down.name].multiplier += multiplier
			}
		}else if(path.down != nil){
			list_of_moves = append(list_of_moves,&move{path.down,multiplier})
		}else if(path.right != nil){
			list_of_moves = append(list_of_moves, &move{path.right,multiplier})
		}
	}
	return branches
}
