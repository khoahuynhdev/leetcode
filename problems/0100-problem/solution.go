package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Walk(root *TreeNode, ch chan int) {
	defer close(ch)
	var walker func(t *TreeNode)
	walker = func(t *TreeNode) {
		if t == nil {
			ch <- 1e5
			return
		}
		ch <- t.Val
		walker(t.Left)
		walker(t.Right)
	}
	walker(root)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(p, ch1)
	go Walk(q, ch2)

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		// fmt.Println(val1,val2,ok1,ok2)
		if ok1 != ok2 || val1 != val2 {
			return false
		}
		if !ok1 || !ok2 {
			break
		}
	}
	return true
}
