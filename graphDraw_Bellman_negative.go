package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func BellmanFordNegative(mygraph *Graph, start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)

	for _, vertex := range mygraph.Vertices {
		distances[vertex.Id] = math.MaxInt64
		visited[vertex.Id] = false
	}

	distances[start] = 0

	for _, vertex := range mygraph.Vertices {
		vertex.changeColor("green")
		DrawGraph(*mygraph)

		for _, edge := range mygraph.Edges {
			edge.changeColor("red")
			DrawGraph(*mygraph)
			u := edge.Source
			v := edge.Destination
			w := edge.Weight

			if distances[u] != math.MaxInt64 && distances[v] > (distances[u]+w) {
				distances[v] = distances[u] + w
				visited[v] = true
			}
			edge.changeColor("black")

			DrawGraph(*mygraph)
		}
		vertex.changeColor("blue")
		DrawGraph(*mygraph)
	}

	for _, edge := range mygraph.Edges {
		u := edge.Source
		v := edge.Destination
		w := edge.Weight
		if distances[v] > (distances[u] + w) {
			Error.Println("There is a negative cycle")
			os.Exit(0)
		}
	}
	return distances
}

func main() {

	gr := Graph{
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
			{Source: 1, Destination: 2, Weight: 2},
			{Source: 3, Destination: 1, Weight: 1},
			{Source: 2, Destination: 3, Weight: -5},
			{Source: 3, Destination: 4, Weight: 3},
		},
	}

	startVertex, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Error.Println("Error during conversion")
		return
	}

	distancesBf := BellmanFordNegative(&gr, startVertex)

	fmt.Println(distancesBf)

	fmt.Println("\nBellman-Ford's Shortest Paths")
	fmt.Println("------------------------")
	for vertex, distance := range distancesBf {
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
	if len(os.Args) != 2 {
		Error.Println("The initial vertex is not defined")
		os.Exit(-1)
	}
}
