package main

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) { return false}
    f,cs,rt := make(map[rune]rune), make(map[rune]bool), []rune(t)
    for idx, v := range s {
        // if v == rt[idx] { continue }
        if c,ok := f[v];ok {
            if c != rt[idx] { return false}
        } else {
            if cs[rt[idx]] { return false}
            f[v] = rt[idx]
            cs[rt[idx]] = true
        }
        rt[idx] = v
    }
    fmt.Println(s, string(rt))
    return s == string(rt)
}
