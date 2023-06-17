package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

var Error *log.Logger
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

func Bellman_Ford(mygraph *Graph, start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)

	for _, vertex := range mygraph.Vertices {
		distances[vertex.Id] = math.MaxInt64
		visited[vertex.Id] = false
	}

	distances[start] = 0

	for _, vertex := range mygraph.Vertices {
		vertex.changeColor("green")
		drawGraph(*mygraph)

		for _, edge := range mygraph.Edges {
			edge.changeColor("red")
			drawGraph(*mygraph)
			u := edge.Source
			v := edge.Destination
			w := edge.Weight

			if distances[u] != math.MaxInt64 && distances[v] > (distances[u]+w) {
				distances[v] = distances[u] + w
				visited[v] = true
			}
			edge.changeColor("black")

			drawGraph(*mygraph)
		}
		vertex.changeColor("blue")
		drawGraph(*mygraph)
	}

	for _, edge := range mygraph.Edges {
		u := edge.Source
		v := edge.Destination
		w := edge.Weight
		if distances[v] > (distances[u] + w) {
			Error.Println("There is a negative cycle")
			os.Exit(-1)
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
			{Source: 2, Destination: 1, Weight: 2},
			{Source: 1, Destination: 3, Weight: 1},
			{Source: 2, Destination: 3, Weight: 5},
			{Source: 3, Destination: 4, Weight: 3},
		},
	}

	startVertex, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Error.Println("Error during conversion")
		return
	}

	distances_bf := Bellman_Ford(&gr, startVertex)

	fmt.Println(distances_bf)

	fmt.Println("\nBellman-Ford's Shortest Paths")
	fmt.Println("------------------------")
	for vertex, distance := range distances_bf {
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
		Error.Println(errBmp)
		os.Exit(-1)
	}
}

func init() {
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	if len(os.Args) != 2 {
		Error.Println("The initial vertex is not defined")
		os.Exit(-1)
	}
}
