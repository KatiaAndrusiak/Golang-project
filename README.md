# Golang-project
Graph algorithm animation

To run:

Install Graphviz https://graphviz.org/download/
```
go get github.com/dominikbraun/graph
```

To run Dijkstra algorithm and create gif:
```
make runDijkstra v=1 matrix=exampleGraph1.txt
```

To run Bellman Ford algorithm and create gif:
```
make runBellman v=1  
```

To run Bellman Ford algorithm with negative weights and create gif:
```
make runBellmanNegative v=1  
```