package main

var graph = make(map[string]map[string]bool)
// graph, who has strings as key and map[string]bool (i.e., set of strings) as
// value

func addEdge(from, to string) {
  edges := graph[from]
  if edges == nil {
    edges = make(map[string]bool)
    graph[from] = edges
  }
  edges[to] = true
}

func hasEdge(from, to string) bool {
  return graph[from][to]
}
