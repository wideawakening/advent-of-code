package day12

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Adjacent Vertex
type Vertex struct {
	key      Node
	adjacent []*Vertex
}

// AddVertext will add a vertex to a graph
func (g *Graph) AddVertex(vertex Node) error {
	if contains(g.vertices, vertex) {
		err := fmt.Errorf("Vertex %v already exists", vertex.position)
		return err
	} else {
		v := &Vertex{
			key: vertex,
		}
		g.vertices = append(g.vertices, v)
	}
	return nil
}

// AddEdge will add ad endge from a vertex to a vertex
func (g *Graph) AddEdge(to, from Node) (bool, error) {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil {
		return false, fmt.Errorf("Not a valid edge from %v ---> %v", from.position, to.position)
	} else if contains(fromVertex.adjacent, toVertex.key) {
		return false, fmt.Errorf("Edge from vertex %d ---> %d already exists", fromVertex.key.position, toVertex.key.position)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		return true, nil
	}
}

// getVertex will return a vertex point if exists or return nil
func (g *Graph) getVertex(vertex Node) *Vertex {
	for i, v := range g.vertices {
		if v.key.position == vertex.position {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(v []*Vertex, key Node) bool {
	for _, v := range v {
		if v.key.position == key.position {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%v : ", v.key.position)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.key.position)
		}
		fmt.Println()
	}
}

func (g *Graph) GetCountToFrom(from *Node, to *Node) int {
	candidates := g.getVertex(*to)
	for _, candidate := range candidates.adjacent {
		processed = make(map[*Vertex]bool)
		CountCandidatesToFrom(candidate.adjacent, from)
		fmt.Println()
		fmt.Println(len(processed))
	}

	return 0
}

var processed = make(map[*Vertex]bool)

func CountCandidatesToFrom(candidates []*Vertex, from *Node) {
	for _, candidate := range candidates {
		fmt.Printf("candidate node %v\n", candidate.key.position)
		if !processed[candidate] {
			processed[candidate] = true
		} else {
			continue
		}
		if candidate.key.position == from.position {
			continue
		}

		CountCandidatesToFrom(candidate.adjacent, from)
	}
	fmt.Println("no more nodes")
}
