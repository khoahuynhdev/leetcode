package main

// Fancy Sequence using a global affine transformation f(x) = mult*x + add (mod 10^9+7).
// Each appended value is stored in inverse-transformed form so that addAll and multAll
// are O(1) and getIndex simply applies the current transform.

const mod = 1_000_000_007

type Fancy struct {
	vals []int64
	mult int64
	add  int64
}

func Constructor() Fancy {
	return Fancy{mult: 1, add: 0}
}

func (f *Fancy) Append(val int) {
	// Store b such that mult*b + add ≡ val (mod mod)
	b := ((int64(val) - f.add) % mod + mod) % mod
	b = b * modInverse(f.mult) % mod
	f.vals = append(f.vals, b)
}

func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + int64(inc)) % mod
}

func (f *Fancy) MultAll(m int) {
	f.mult = f.mult * int64(m) % mod
	f.add = f.add * int64(m) % mod
}

func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.vals) {
		return -1
	}
	return int((f.mult*f.vals[idx]%mod + f.add) % mod)
}

func modPow(base, exp, m int64) int64 {
	result := int64(1)
	base %= m
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % m
		}
		exp /= 2
		base = base * base % m
	}
	return result
}

func modInverse(a int64) int64 {
	return modPow(a, mod-2, mod)
}
