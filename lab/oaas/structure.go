package main

import (
//"encoding/json"
)

type Input struct {
	Name    string `json:name,omitempty`
	Digraph []int  `json:digraph`
	Nodes   []Node `json:nodes`
}

type Node struct {
	ID       int      `json:id`
	Name     string   `json:name,omitempty`
	Engine   string   `json:engine:omitempty` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string   `json:artifact`
	Args     []string `json:args`
}

func main() {

}
