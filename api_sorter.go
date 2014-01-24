package main

import (
	"sort"
)

// By is the type of a "less" function that defines the ordering of its Objects arguments.
type By func(o1, o2 *Object) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(objects []*Object) {
	os := &objectSorter{
		objects: objects,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(os)
}

// objectSorter joins a By function and a slice of Objects to be sorted.
type objectSorter struct {
	objects []*Object
	by      func(o1, o2 *Object) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (o *objectSorter) Len() int {
	return len(o.objects)
}

// Swap is part of sort.Interface.
func (o *objectSorter) Swap(i, j int) {
	o.objects[i], o.objects[j] = o.objects[j], o.objects[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (o *objectSorter) Less(i, j int) bool {
	return o.by(o.objects[i], o.objects[j])
}
