package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Dijkstra(mygraph *Graph, start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)

	for _, vertex := range mygraph.Vertices {
		distances[vertex.Id] = math.MaxInt64
		visited[vertex.Id] = false
	}

	distances[start] = 0

	for _, vertex := range mygraph.Vertices {
		u := MinDistance(distances, visited)
		visited[u] = true
		vertex.changeColor("green")
		DrawGraph(*mygraph)

		for _, edge := range mygraph.Edges {
			edge.changeColor("red")
			DrawGraph(*mygraph)

			if edge.Source == u && !visited[edge.Destination] && distances[u] != math.MaxInt64 && distances[u]+edge.Weight < distances[edge.Destination] {
				distances[edge.Destination] = distances[u] + edge.Weight
			}
			edge.changeColor("black")

			DrawGraph(*mygraph)
		}
		vertex.changeColor("blue")
		DrawGraph(*mygraph)
	}

	return distances
}

func MinDistance(distances map[int]int, visited map[int]bool) int {
	min := math.MaxInt64
	minIndex := -1

	for vertex, distance := range distances {
		if !visited[vertex] && distance <= min {
			min = distance
			minIndex = vertex
		}
	}

	return minIndex
}

var matrixFilename = flag.String("matrix", "exampleGraph1.txt", "Adjacency matrix filename")
var initialVertex = flag.String("v", "0", "Initial vertex")

func main() {
	flag.Parse()
	var gr Graph
	if len(*matrixFilename) < 1 {
		gr = Graph{
			Vertices: []*Vertex{
				{Id: 0, Color: "black"},
				{Id: 1, Color: "black"},
				{Id: 2, Color: "black"},
				{Id: 3, Color: "black"},
				{Id: 4, Color: "black"},
			},
			Edges: []*Edge{
				{Source: 0, Destination: 1, Weight: 4},
				{Source: 0, Destination: 2, Weight: 1},
				{Source: 2, Destination: 1, Weight: 2},
				{Source: 1, Destination: 3, Weight: 1},
				{Source: 2, Destination: 3, Weight: 5},
				{Source: 3, Destination: 4, Weight: 3},
			},
		}
	} else {
		gr = GetMatrixAndConvert(*matrixFilename)
	}

	startVertex, err := strconv.Atoi(*initialVertex)
	if err != nil {
		Error.Println("Error during conversion")
		return
	}

	distances := Dijkstra(&gr, startVertex)

	fmt.Println(distances)

	fmt.Println("\nDijkstra's Shortest Paths")
	fmt.Println("------------------------")
	for vertex, distance := range distances {
		fmt.Printf("Vertex %d - Distance: ", vertex)
		if distance == math.MaxInt64 {
			fmt.Println("INF")
		} else {
			fmt.Println(distance)
		}
	}

}

func init() {
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
}
