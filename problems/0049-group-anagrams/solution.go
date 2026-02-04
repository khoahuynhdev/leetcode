package solution

// Intuition: computing key without sorting

// func groupAnagrams(strs []string) [][]string {
//     tbl, ans := make(map[string][]string), make([][]string, 0)
//     for _,v := range strs {
//         tmp := []rune(v)
//         sort.Slice(tmp, func (i,j int) bool {
//             return tmp[i] < tmp[j]
//         })
//         tbl[string(tmp)] = append(tbl[string(tmp)], v)
//     }
//     for _, v := range tbl {
//         ans = append(ans, v)
//     }
//     return ans
// }

type Key [26]byte
func MakeKey(str string) (key Key) {
    for _,v := range str {
        key[v - 'a']++
    }
    return 
}
func groupAnagrams(strs []string) [][]string {
    tbl, ans := make(map[[26]byte][]string), make([][]string, 0)
    for _,v := range strs {
        key := MakeKey(v)
        tbl[key] = append(tbl[key], v)
    }
    for _, v := range tbl {
        ans = append(ans, v)
    }
    return ans
}
