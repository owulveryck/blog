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
	ID       int               `json:id`
	Name     string            `json:name,omitempty`
	Engine   string            `json:engine:omitempty` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:artifact`
	Args     []string          `json:args,omitempty`   // the arguments of the artifact, if needed
	Inputs   map[string]string `json:input,omitempty`  // the key is the node that gives the information, and  the value, its output (always a string)
	Outputs  map[string]string `json:output,omitempty` // the key is the name of the parameter, the value its value (always a string)
}

func main() {

}
