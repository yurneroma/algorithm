package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	h  []int
	ph []int
	hp []int
)
var size int

const N = 100010

func main() {
	n := 0
	size = 0
	m := 0
	h = make([]int, N)
	ph = make([]int, N)
	hp = make([]int, N)
	rd := bufio.NewReader(os.Stdin)
	fline, _ := rd.ReadString('\n')

	n, _ = strconv.Atoi(strings.TrimSpace(fline))
	for i := 0; i < n; i++ {
		obj := readline(rd)
		if obj.Op == "I" {
			size++
			m++
			ph[m] = size
			hp[size] = m
			h[size] = obj.x
			up(size)
		} else if obj.Op == "PM" {
			fmt.Println(h[1])
		} else if obj.Op == "DM" {
			swap(1, size)
			size--
			down(1)
		} else if obj.Op == "D" {
			k := obj.k
			ki := ph[k]
			swap(ki, size)
			size--
			down(ki)
			up(ki)
		} else {
			k := obj.k
			x := obj.x
			ki := ph[k]
			h[ki] = x
			down(ki)
			up(ki)
		}
	}
}

type Data struct {
	Op string
	k  int
	x  int
}

func readline(rd *bufio.Reader) Data {
	line, _ := rd.ReadString('\n')
	params := strings.Fields(strings.TrimSuffix(line, "\n"))
	obj := Data{}
	if params[0] == "I" {
		obj.Op = "I"
		obj.x, _ = strconv.Atoi(params[1])
		return obj
	}

	if params[0] == "PM" {
		obj.Op = "PM"
		return obj
	}

	if params[0] == "DM" {
		obj.Op = "DM"
		return obj
	}

	if params[0] == "D" {
		obj.Op = "D"
		obj.k, _ = strconv.Atoi(params[1])
		return obj
	}

	if params[0] == "C" {
		obj.Op = "C"
		obj.k, _ = strconv.Atoi(params[1])
		obj.x, _ = strconv.Atoi(params[2])
	}
	return obj
}

func swap(a, b int) {
	// swap ph
	ph[hp[a]], ph[hp[b]] = ph[hp[b]], ph[hp[a]]
	// swap hp
	hp[a], hp[b] = hp[b], hp[a]
	// swap value
	h[a], h[b] = h[b], h[a]
}
func up(x int) {
	for x/2 != 0 && h[x] < h[x/2] {
		swap(x, x/2)
		x /= 2
	}
}
func down(x int) {
	t := x
	if 2*x <= size && h[2*x] < h[t] {
		t = 2 * x
	}
	if 2*x+1 <= size && h[2*x+1] < h[t] {
		t = 2*x + 1
	}

	if t != x {
		swap(t, x)
		down(t)
	}
}
