// go语言实现广度优先算法，走迷宫。以二维数组的形式保存
package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col , end int
	fmt.Fscanf(file, "%d %d", &row, &col)


	// window 每个文件windows下换行符是\r\n的关系，它每个行末读进了一个0。其实你看下scanf的返回，应该有错。这个可以每行后面加一个Scanln，或者使用bufio里面的readstring。
	fmt.Fscanf(file , "%d" , &end)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%v", &maze[i][j])
		}
		fmt.Fscanf(file , "%d" , &end) // 获取文件的换行符
	}

	return maze
}

// 定义点
type point struct {
	i , j int
}

// 定义 前 后 左 右 4个步骤
var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1},
}

// 行走
func (p point)add(dir point) point {
	return point{p.i+dir.i , p.j+dir.j}
}

func (p point) at(grid [][]int) (int , bool) {
	// 判断数组是否越界了  往上  往下都越界了
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0  || p.j >= len(grid[p.i]) {
		return 0 , false
	}

	return grid[p.i][p.j] , true
}

func walk(maze [][]int , start , end point)[][]int {
	steps := make([][]int  , len(maze))
	for i := range steps {
		steps[i] = make([]int , len(maze[i]))
	}

	// 保存一个需要遍历的根节点的队列，初始为开始节点
	queue := []point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			break
		}
		// 每次探索 前后左右 4步
		for _ , dir := range dirs {
			// 获取下一个点,当前节点探索的步
			next := current.add(dir)
			// 达到终点 退出
			//if next == end {
			//	break
			//}
			// 判断下一个节点是否需要加入队列当中
			// 1 maze 在 next 节点是0
			// 2 steps 在 next 节点也是0 (不为0表示已经走过了)
			// 3 next节点不是start节点
			val , ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val , ok = next.at(steps)
			if !ok || val !=0 {
				continue
			}

			if next == start {
				continue
			}
			// 获取当前的步数在steps 中的位置
			val , _ = current.at(steps)

			steps[next.i][next.j] = val+ 1

			queue = append(queue, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%3d", val)
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze , point{0,0} , point{len(maze) - 1 , len(maze[0]) -1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
