package main

import (
        "time"
        //"math"
        "fmt"
        "strconv"
        "strings"
)

struct node {
  value int,
  left_child *node, 
  right_child *node

}

func main() {
        start := time.Now()
        input := "75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23"
        tree := generate_tree(input)
        fmt.Printf("Result: %f", "ssss")
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func generate_tree(input string) *node{
  string_arr := strings.split(string, "\n")
  var previous_node_array []*node
  previous_node_array = make([]*node, 0)
  var upper_node *node 
  upper_node = nil
  for _, line := range string_arr {
    node_arr := strings.split(line, " ")
    for _,curr := range node_arr {
      num, _ := strconv.Atoi(curr)
      curr_node := node{num, nil, nil}
      previous_node_array = append(previous_node_array,&curr_node)
      if(upper_node != nil){
        if(upper_node.left == nil){
          upper_node.left = curr_node
        }else if(upper_node.right == nil){
          upper_node.right = curr_node
        }else{
          upper_node, previous_node_array = previous_node_array[0], previous_node_array[1:]
        }
      }else{
        upper_node, previous_node_array = previous_node_array[0], previous_node_array[1:]
      }
    }
  }
  return upper_node[0]
}

func print_tree(start *node){
  next_arr := make([]*node, 0)
  next_arr = append(next_arr, start)
  for len(next_arr) > 0 {

  }
  fmt.Printf("%d", start.value)


}