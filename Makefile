runDijkstra:
	go run graphUtil.go graphDraw_Dijkstra.go -v=${v} -matrix=${matrix}
	go run createGif.go Dijkstra

runBellman:
	go run graphUtil.go graphDraw_Bellman.go ${v}
	go run createGif.go Bellman

runBellmanNegative:
	go run graphUtil.go graphDraw_Bellman_negative.go ${v}
	go run createGif.go BellmanNegative
