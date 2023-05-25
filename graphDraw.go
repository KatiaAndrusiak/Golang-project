package main

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"math"
	"os"
	"os/exec"
)

var (
	filenameCounter = 0
)

type Edge struct {
	Source      int
	Destination int
	Weight      int
	Color       string
}

type Vertex struct {
	Id    int
	Color string
}

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

func (vertex *Vertex) changeColor(color string) {
	vertex.Color = color
}

func (edge *Edge) changeColor(color string) {
	edge.Color = color
}

func Dijkstra(mygraph *Graph, start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)

	for _, vertex := range mygraph.Vertices {
		distances[vertex.Id] = math.MaxInt64
		visited[vertex.Id] = false
	}

	distances[start] = 0

	for _, vertex := range mygraph.Vertices {
		u := minDistance(distances, visited)
		visited[u] = true
		vertex.changeColor("green")
		drawGraph(*mygraph)

		for _, edge := range mygraph.Edges {
			edge.changeColor("red")
			drawGraph(*mygraph)

			if edge.Source == u && !visited[edge.Destination] && distances[u] != math.MaxInt64 && distances[u]+edge.Weight < distances[edge.Destination] {
				distances[edge.Destination] = distances[u] + edge.Weight
			}
			edge.changeColor("black")

			drawGraph(*mygraph)
		}
		vertex.changeColor("blue")
		drawGraph(*mygraph)
	}

	return distances
}

func minDistance(distances map[int]int, visited map[int]bool) int {
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
			{Source: 2, Destination: 1, Weight: 2},
			{Source: 1, Destination: 3, Weight: 1},
			{Source: 2, Destination: 3, Weight: 5},
			{Source: 3, Destination: 4, Weight: 3},
		},
	}

	startVertex := 0

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

func drawGraph(graph2 Graph) {
	gg := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic(), graph.Weighted())

	for _, v := range graph2.Vertices {
		_ = gg.AddVertex(v.Id, graph.VertexAttribute("color", v.Color))
	}

	for _, es := range graph2.Edges {
		_ = gg.AddEdge(es.Source, es.Destination, graph.EdgeAttribute("label", fmt.Sprintf("%d", es.Weight)), graph.EdgeWeight(es.Weight), graph.EdgeAttribute("color", es.Color))
	}

	fname := fmt.Sprintf("%d.gv", filenameCounter)
	filenameCounter++
	file, _ := os.Create(fname)
	_ = draw.DOT(gg, file, draw.GraphAttribute("size", "100,100"))

	cmd := exec.Command("dot", "-Tpng", "-O", fname)
	errBmp := cmd.Run()
	if errBmp != nil {
		panic(errBmp)
	}
}
