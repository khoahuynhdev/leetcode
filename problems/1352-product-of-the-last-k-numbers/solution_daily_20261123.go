package main

// Approach: prefix products that reset on zero.
// prefix[i] holds the product of the first i numbers since the last zero
// (prefix[0] = 1 as a sentinel). The product of the last k numbers is
// prefix[n]/prefix[n-k]; if k reaches past the last reset, a zero is in the
// window and the answer is 0. Both operations run in O(1).

type ProductOfNumbers struct {
	prefix []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{prefix: []int{1}}
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.prefix = p.prefix[:1]
		return
	}
	p.prefix = append(p.prefix, p.prefix[len(p.prefix)-1]*num)
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	n := len(p.prefix) - 1
	if k > n {
		return 0
	}
	return p.prefix[n] / p.prefix[n-k]
}
