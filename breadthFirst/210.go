func findOrder(numCourses int, prerequisites [][]int) []int {
	// 入度表和出边表
	inDegree := make([]int, numCourses)
	outEdges := make([][]int, numCourses)

	// 根据prerequisites进行根据入度表和出边表
	for _, v := range prerequisites {
		inDegree[v[0]]++
		outEdges[v[1]] = append(outEdges[v[1]], v[0])
	}

	queue := make([]int, 0)
	rsl := make([]int, 0)
	// 获取入度为0的点，加入队列
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	var node int
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:] // 将队头出队
		// 消除以队头为源点的边
		for _, v := range outEdges[node] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
		// 添加结果
		rsl = append(rsl, node)
	}
	// 检查有没有成环，如果成环，直接返回空
	if len(rsl) != numCourses {
		rsl = []int{}
	}
	return rsl
}