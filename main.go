package main

func main() {
}


func levelOrder(root *TreeNode) [][]int {
    quene:=[]*TreeNode{root}
    result:=[][]int{}
    for len(quene) > 0 {
        ret:=[]int{}
        end:=len(quene)
        for i:=0;i<end;i++{
            ret = append(ret,quene[i].Val)
            if quene[i].Left !=nil {
                quene = append(quene,quene[i].Left)
            }
            if quene[i].Right !=nil {
                quene = append(quene,quene[i].Right)
            }
        }
        result = append(result,ret)
    }
    return result
}