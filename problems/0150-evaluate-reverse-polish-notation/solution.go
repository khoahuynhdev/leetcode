package solution
import(
  "strconv"
)
func evalRPN(tokens []string) int {
 stack,si := make([]int,6000), -1
 for i:=0;i<len(tokens);i++{
     if s, err := strconv.Atoi(tokens[i]); err == nil {
         si++
         stack[si] = s
	} else {
        st := stack[si]
        nd := stack[si-1]
        var res int
        switch tokens[i] {
            case "/":
                res = nd / st
                break
            case "*":
                res = nd * st
                break
            case "-":
                res = nd - st
                break
            case "+":
                res = nd + st
                break
        }
        si-=1
        stack[si] = res
    }
 }
 return stack[si]
}
