package main

import (
        "time"
        //"math"
        "fmt"
        //"strconv"
)

type node struct{
  x_coord int
  y_coord int
  distance int
  next []node
  done bool
}

type cell struct {
	x_coord int
	y_corrd int
}

func main() {
	x,y := 13,13
        start := time.Now()
	maze := generate_maze(x,y)
        print(maze)
	res := a_star(maze,x,y)
	fmt.Printf("Result: %f", res)
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func print(maze [][]cell){
        var i, j int
        for  i = 0; i < len(maze); i++ {
                for j = 0; j < len(maze[0]); j++ {
                        fmt.Printf("%d ",maze[i][j] )
                }
                fmt.Printf("\n" )
        }
}

func generate_maze(size_x int, size_y int) [][]cell {
	var maze [][]cell
	//starting at the top left corner, x = 0, y = 9
	// +y is down, +x is right
	y := 0
	x := 0
	for y < size_y {
		maze = append(maze, make([]cell,size_y))
		for x < size_x {
			c := cell { x,y}
			maze[y][x] = c
		x += 1
		}
		y += 1
		x = 0
	}
	return maze
}

func a_star(maze [][]cell, target_x int, target_y int) int{
	//at a basic level, we want to keep a list of every path, sorted by which is the shortest 
	//each path should include what potential moves it can make 
	//the path cannot go back on any move its already made 
	//if it reaches the end, add it to the "solutions" list
	var list_of_paths []node
	//this is a queue, put the new longest path at the end, so the shortest paths wind up at the beginning
	list_of_paths = make([]node,0)
	var list_of_finished []node
	list_of_finished = make([]node,0)

	//x_neg_lim := 0
	//y_neg_lim := 0
	x_pos_lim := len(maze[0])
	y_pos_lim := len(maze)
	//a false move initalizes with the real first move as it "next". It's this or compute the first two moves by hand
	start := node{0,0,0, make([]node,1), false }
	list_of_paths = append(list_of_paths, start)
	fmt.Printf("%d", len(list_of_paths))
	for len(list_of_paths)> 0 {
		//pop off the first element
		var path node
		//fmt.Print("\n+++++++++++++++++")
		//fmt.Printf("\n%+v", list_of_paths)
		path, list_of_paths = list_of_paths[0], list_of_paths[1:]
		//the case where the next move finishes the path
		for _, move := range path.next {
			if(move.y_coord == target_y && move.x_coord == target_x){
				list_of_finished = append(list_of_finished, move)
				move.done = true
				path = move
				break
			}
		 }
		 if(path.done == true){
			continue
		}
		 //the case where the next move does not finish the path
		 for _, move := range path.next{
			 //each move becomes a new potential path
			 new_path := move
			 //add the moves to the new path, (only down and left right now)
			 if(new_path.x_coord < x_pos_lim){
				new_path.next = append(new_path.next, node{new_path.x_coord+1, new_path.y_coord,new_path.distance + 1, make([]node,0), false })
			 }
			 if(new_path.y_coord < y_pos_lim){
                                new_path.next = append(new_path.next, node{new_path.x_coord, new_path.y_coord+1,new_path.distance + 1, make([]node,0), false })
                         }
			 //put the path at the end of the queue 
			 list_of_paths = append(list_of_paths,new_path)
		 }
	}
	return len(list_of_finished)
}
