// The digraph command performs queries over unlabelled directed graphs
// represented in text form.  It is intended to integrate nicely with
// typical UNIX command pipelines.
//
// Since directed graphs (import graphs, reference graphs, call graphs,
// etc) often arise during software tool development and debugging, this
// command is included in the go.tools repository.
//
// TODO(adonovan):
// - support input files other than stdin
// - suport alternative formats (AT&T GraphViz, CSV, etc),
//   a comment syntax, etc.
// - allow queries to nest, like Blaze query language.
//
package main // import "golang.org/x/tools/cmd/digraph"

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/owulveryck/toscalib"
	"io"
	"os"
	"sort"
)

const Usage = `digraph: queries over directed graphs in text form.

Graph format:

  Each line contains zero or more words.  Words are separated by
  unquoted whitespace; words may contain Go-style double-quoted portions,
  allowing spaces and other characters to be expressed.

  Each field declares a node, and if there are more than one,
  an edge from the first to each subsequent one.
  The graph is provided on the standard input.

  For instance, the following (acyclic) graph specifies a partial order
  among the subtasks of getting dressed:

	% cat clothes.txt
	socks shoes
	"boxer shorts" pants
	pants belt shoes
	shirt tie sweater
	sweater jacket
	hat

  The line "shirt tie sweater" indicates the two edges shirt -> tie and
  shirt -> sweater, not shirt -> tie -> sweater.

Supported queries:

  somepath <label> <label>
`

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout

type nodelist []string
type nodeset map[string]bool

// A graph maps nodes to the non-nil set of their immediate successors.
type graph map[string]nodeset

// main function
func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(Usage)
		return
	}

	if err := digraph(args[0], args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "digraph: %s\n", err)
		os.Exit(1)
	}
}

// obviously a println function
func (l nodelist) println(sep string) {
	for i, label := range l {
		if i > 0 {
			fmt.Fprint(stdout, sep)
		}
		fmt.Fprint(stdout, label)
	}
	fmt.Fprintln(stdout)
}

// sort returns a node list in increasing order
func (s nodeset) sort() nodelist {
	labels := make(nodelist, len(s))
	var i int
	for label := range s {
		labels[i] = label
		i++
	}
	sort.Strings(labels)
	return labels
}

func (g graph) addNode(label string) nodeset {
	edges := g[label]
	if edges == nil {
		edges = make(nodeset)
		g[label] = edges
	}
	return edges
}

func (g graph) addEdges(from string, to ...string) {
	edges := g.addNode(from)
	for _, to := range to {
		g.addNode(to)
		edges[to] = true
	}
}

// Parses the graph
func parse(rd io.Reader) (graph, error) {
	g := make(graph)

	var linenum int
	in := bufio.NewScanner(rd)
	for in.Scan() {
		linenum++
		// Split into words, honoring double-quotes per Go spec.
		words, err := split(in.Text())
		if err != nil {
			return nil, fmt.Errorf("at line %d: %v", linenum, err)
		}
		if len(words) > 0 {
			g.addEdges(words[0], words[1:]...)
		}
	}
	if err := in.Err(); err != nil {
		return nil, err
	}
	return g, nil
}

func digraph(cmd string, args []string) error {
	// Parse the input graph.
	var toscaTemplate toscalib.ToscaDefinition
	err := toscaTemplate.Parse(stdin)
	if err != nil {
		return err
	}
	g := make(graph)
	// a map containing the ID and the corresponding action
	ids := make(map[int]string)
	// Fill in the graph with the toscaTemplate via the adjacency matrix
	for node, template := range toscaTemplate.TopologyTemplate.NodeTemplates {
		// Find the edges of the current node and add them to the graph

		ids[template.GetConfigureIndex()] = fmt.Sprintf("%v:Configure", node)
		ids[template.GetCreateIndex()] = fmt.Sprintf("%v:Create", node)
		ids[template.GetDeleteIndex()] = fmt.Sprintf("%v:Delete", node)
		ids[template.GetInitialIndex()] = fmt.Sprintf("%v:Initial", node)
		ids[template.GetPostConfigureSourceIndex()] = fmt.Sprintf("%v:PostConfigureSource", node)
		ids[template.GetPostConfigureTargetIndex()] = fmt.Sprintf("%v:PostconfigureTarget", node)
		ids[template.GetPreConfigureSourceIndex()] = fmt.Sprintf("%v:PreConfigureSource", node)
		ids[template.GetPreConfigureTargetIndex()] = fmt.Sprintf("%v:PreConfigureTarget", node)
		ids[template.GetStartIndex()] = fmt.Sprintf("%v:Start", node)
		ids[template.GetStopIndex()] = fmt.Sprintf("%v:Stop", node)
	}

	adjacencyMatrix := toscaTemplate.AdjacencyMatrix
	//g.addEdges(node)
	row, col := adjacencyMatrix.Dims()
	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if adjacencyMatrix.At(r, c) == 1 {
				g.addEdges(ids[r], ids[c])
			}
		}
	}
	// Parse the command line.
	switch cmd {
	case "somepath":
		if len(args) != 2 {
			return fmt.Errorf("usage: digraph somepath <from> <to>")
		}
		from, to := args[0], args[1]
		if g[from] == nil {
			return fmt.Errorf("no such 'from' node %q", from)
		}
		if g[to] == nil {
			return fmt.Errorf("no such 'to' node %q", to)
		}

		seen := make(nodeset)
		var visit func(path nodelist, label string) bool
		visit = func(path nodelist, label string) bool {
			if !seen[label] {
				seen[label] = true
				if label == to {
					append(path, label).println("\n")
					return true // unwind
				}
				for e := range g[label] {
					if visit(append(path, label), e) {
						return true
					}
				}
			}
			return false
		}
		if !visit(make(nodelist, 0, 100), from) {
			return fmt.Errorf("no path from %q to %q", args[0], args[1])
		}

	case "nodes":
		if len(args) != 0 {
			return fmt.Errorf("usage: nodes")

		}
		nodes := make(nodeset)
		for label := range g {
			nodes[label] = true

		}
		nodes.sort().println("\n")

	default:
		return fmt.Errorf("no such command %q", cmd)
	}

	return nil
}
