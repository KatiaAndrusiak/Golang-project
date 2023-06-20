package main

import (
	"bufio"
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func DrawGraph(graph2 Graph) {
	gg := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic())

	for _, v := range graph2.Vertices {
		_ = gg.AddVertex(v.Id, graph.VertexAttribute("color", v.Color))
	}

	for _, es := range graph2.Edges {
		_ = gg.AddEdge(es.Source, es.Destination, graph.EdgeAttribute("label", fmt.Sprintf("%d", es.Weight)), graph.EdgeAttribute("color", es.Color))
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

	e := os.Remove(fname)
	if e != nil {
		log.Fatal(e)
	}
}

func GetMatrixAndConvert(filename string) Graph {
	matrix := ReadMatrixFromFile(filename)
	return ConvertAdjacencyMatrixToGraph(matrix)
}

func ReadMatrixFromFile(filename string) [][]int {
	reader, err := os.Open(filename)
	if err != nil {
		Error.Println("Error during file opening")
		return nil
	}

	var sli [][]int
	s := bufio.NewScanner(reader)
	for s.Scan() {
		var row []int
		for _, w := range strings.Split(s.Text(), " ") {
			v, err := strconv.Atoi(w)
			if err != nil {
				Error.Println("Error during conversion")
				return nil
			}
			row = append(row, v)
		}
		sli = append(sli, row)
	}

	fmt.Println(sli)
	for _, row := range sli {
		if len(sli) != len(row) {
			Error.Println("Should be square matrix")
		}
	}
	return sli
}

func ConvertAdjacencyMatrixToGraph(matrix [][]int) Graph {
	var vertices []*Vertex
	var edges []*Edge
	for i, row := range matrix {
		vertex := Vertex{Id: i, Color: "black"}
		vertices = append(vertices, &vertex)
		for j, el := range row {
			if el != 0 {
				edge := Edge{Source: i, Destination: j, Weight: el}
				edges = append(edges, &edge)
			}
		}
	}
	gr := Graph{Vertices: vertices, Edges: edges}
	return gr
}
