package tree

import (
	"errors"
	"fmt"
	"sort"
)

const (
	// Inf is defined as the max value of an integer, used as +∞
	Inf = int(^uint(0) >> 1)
	// NegInf is defined as the min value of an integer, used as -∞
	NegInf = -Inf - 1
)

type interval struct {
	segment
	element interface{}
}

type segment struct {
	from int
	to   int
}

type node struct {
	segment     segment
	left, right *node
	intervals   []*interval
}

// SegmentTree 线段树
type SegmentTree struct {
	base     []interval
	elements map[interface{}]struct{}
	root     *node
}

func NewSegmentTree() *SegmentTree {
	return &SegmentTree{}
}

// Push pushes an interval to the interval stack
func (t *SegmentTree) Push(from, to int, element interface{}) {
	if to < from {
		from, to = to, from
	}
	if t.elements == nil {
		t.elements = make(map[interface{}]struct{})
	}
	t.elements[element] = struct{}{}
	t.base = append(t.base, interval{segment{from, to}, element})
}

// Clear clears the interval stack
func (t *SegmentTree) Clear() {
	t.root = nil
	t.base = nil
}

// BuildTree builds the segment tree from the interval stack
func (t *SegmentTree) BuildTree() error {
	if len(t.base) == 0 {
		return errors.New("No intervals in stack. Push intervals on the stack of the tree first.")
	}

	leaves := elementaryIntervals(t.endpoints())
	t.root = t.insertNodes(leaves)

	for i := range t.base {
		t.root.insertInterval(&t.base[i])
	}

	return nil
}

// Print prints a binary tree recursively to sdout
func (t *SegmentTree) Print() {
	endpoints := len(t.base)*2 + 2
	leaves := endpoints*2 - 3
	height := 1 + log2(leaves)

	fmt.Println("Height:", height, ", leaves:", leaves)
	levels := make([][]*node, height+1)

	traverse(t.root, 0, func(n *node, depth int) {
		levels[depth] = append(levels[depth], n)
	}, nil)

	for i, level := range levels {
		for j, n := range level {
			space(12 * (len(levels) - 1 - i))
			n.print()
			space(1 * (height - i))

			if j-1%2 == 0 {
				space(2)
			}
		}
		fmt.Println()
	}
}

// Removes duplicate entries from a sorted slice
func removedups(sorted []int) (unique []int) {
	unique = make([]int, 0, len(sorted))
	unique = append(unique, sorted[0])
	prev := sorted[0]
	for _, val := range sorted[1:] {
		if val != prev {
			unique = append(unique, val)
			prev = val
		}
	}
	return
}

// Creates a sorted slice of unique endpoints from a tree's base
func (t *SegmentTree) endpoints() []int {
	baseLen := len(t.base)
	endpoints := make([]int, baseLen*2)

	// When there are a lot of intervals, there is a big chance of big overlaps
	// Try to have the endpoints sorted as much as possible when putting them
	// in the slice to reduce the final sort time.
	// endpoints[0] = NegInf
	for i, interval := range t.base {
		endpoints[i] = interval.from
		endpoints[i+baseLen] = interval.to
	}
	// endpoints[baseLen*2+1] = Inf

	sort.Sort(sort.IntSlice(endpoints))

	return removedups(endpoints)
}

// Creates a slice of elementary intervals from a slice of (sorted) endpoints
// Input: [p1, p2, ..., pn]
// Output: [{p1 : p1}, {p1 : p2}, {p2 : p2},... , {pn : pn}
func elementaryIntervals(endpoints []int) []segment {
	if len(endpoints) == 1 {
		return []segment{segment{endpoints[0], endpoints[0]}}
	}

	intervals := make([]segment, len(endpoints)*2-1)

	for i := 0; i < len(endpoints); i++ {
		intervals[i*2] = segment{endpoints[i], endpoints[i]}
		if i < len(endpoints)-1 {
			intervals[i*2+1] = segment{endpoints[i], endpoints[i+1]}
		}
	}
	return intervals
}

// insertNodes builds the tree structure from the elementary intervals
func (t *SegmentTree) insertNodes(leaves []segment) *node {
	var n *node
	if len(leaves) == 1 {
		n = &node{segment: leaves[0]}
		n.left = nil
		n.right = nil
	} else {
		n = &node{segment: segment{leaves[0].from, leaves[len(leaves)-1].to}}
		center := len(leaves) / 2
		n.left = t.insertNodes(leaves[:center])
		n.right = t.insertNodes(leaves[center:])
	}

	return n
}

func (s *segment) subsetOf(other *segment) bool {
	return other.from <= s.from && other.to >= s.to
}

func (s *segment) intersectsWith(other *segment) bool {
	return other.from <= s.to && s.from <= other.to || s.from <= other.to && other.from <= s.to
}

// Inserts interval into given tree structure
func (n *node) insertInterval(i *interval) {
	if n.segment.subsetOf(&i.segment) {
		if n.intervals == nil {
			n.intervals = make([]*interval, 0, 1)
		}
		n.intervals = append(n.intervals, i)
	} else {
		if n.left != nil && n.left.segment.intersectsWith(&i.segment) {
			n.left.insertInterval(i)
		}
		if n.right != nil && n.right.segment.intersectsWith(&i.segment) {
			n.right.insertInterval(i)
		}
	}
}

func (n *node) print() {
	from := fmt.Sprintf("%d", n.segment.from)
	switch n.segment.from {
	case Inf:
		from = "+∞"
	case NegInf:
		from = "-∞"
	}
	to := fmt.Sprintf("%d", n.segment.to)
	switch n.segment.to {
	case Inf:
		to = "Inf"
	case NegInf:
		to = "NegInf"
	}
	fmt.Printf("(%s,%s)", from, to)
	if n.intervals != nil {
		fmt.Print("->[")
		for _, intrvl := range n.intervals {
			fmt.Printf("(%v,%v)=[%v]", intrvl.from, intrvl.to, intrvl.element)
		}
		fmt.Print("]")
	}

}

// QueryIndex looks for all segments in the interval tree that contain
// a given index. The elements associated with the segments will be sent
// on the returned channel. No element will be sent twice.
// The elements will not be sent in any specific order.
func (t *SegmentTree) QueryIndex(index int) (<-chan interface{}, error) {
	if t.root == nil {
		return nil, errors.New("Tree is empty. Build the tree first")
	}

	intervals := make(chan *interval)

	go func(t *SegmentTree, index int, intervals chan *interval) {
		query(t.root, index, intervals)
		close(intervals)
	}(t, index, intervals)

	elements := make(chan interface{})

	go func(intervals chan *interval, elements chan interface{}) {
		defer close(elements)
		results := make(map[interface{}]struct{})
		for interval := range intervals {
			_, alreadyFound := results[interval.element]
			if !alreadyFound {
				// Store an empty struct in the map to minimize memory footprint
				results[interval.element] = struct{}{}
				elements <- interval.element
				if len(results) >= len(t.elements) {
					// found all elements that can be found
					return
				}
			}

		}
	}(intervals, elements)

	return elements, nil
}

func (s segment) contains(index int) bool {
	return s.from <= index && index <= s.to
}

func query(node *node, index int, results chan<- *interval) {
	if node.segment.contains(index) {
		for _, interval := range node.intervals {
			results <- interval
		}
		if node.left != nil {
			query(node.left, index, results)
		}
		if node.right != nil {
			query(node.right, index, results)
		}
	}
}

// Traverse tree recursively call enter when entering node, resp. leave
func traverse(node *node, depth int, enter, leave func(*node, int)) {
	if node == nil {
		return
	}
	if enter != nil {
		enter(node, depth)
	}
	traverse(node.left, depth+1, enter, leave)
	traverse(node.right, depth+1, enter, leave)
	if leave != nil {
		leave(node, depth)
	}
}

// Returs log with base 2 of an int.
func log2(num int) int {
	if num == 0 {
		return NegInf
	}
	i := -1
	for num > 0 {
		num = num >> 1
		i++
	}
	return i
}

func space(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func find(element interface{}, elements []interface{}) (bool, int) {
	for i, e := range elements {
		if element == e {
			return true, i
		}
	}
	return false, -1
}
