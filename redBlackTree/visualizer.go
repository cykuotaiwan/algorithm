package redBlackTree

import (
	"algorithms/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (tree *RedBlackTree) Visualize(imgName string) {
	formatStr := "digraph rbt{\n" + createRBTString(tree.Root) + "}\n"

	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		err := os.Mkdir("./tmp", os.ModePerm)
		if err != nil {
			fmt.Println("create dir error")
			return
		}
	}

	f, err := os.Create("tmp/tmp.dot")
	if err != nil {
		fmt.Println("open tmp file error")
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	len, err := writer.WriteString(formatStr)
	if err != nil || len == 0 {
		fmt.Println("write tmp file error")
		return
	}
	writer.Flush()

	util.ShellCommand("dot", []string{"-Tpng", "tmp/tmp.dot", "-o", "tmp/" + imgName + ".png"})

	err = os.Remove("tmp/tmp.dot")
	if err != nil {
		fmt.Println("remove tmp file error")
		return
	}
}

func createRBTString(node *TreeNode) string {
	if node == nil {
		return ""
	}
	var currentNodeStr string
	if node.Left != nil {
		currentNodeStr += strconv.Itoa(node.Key)
		currentNodeStr += "->"
		currentNodeStr += strconv.Itoa(node.Left.Key)
		currentNodeStr += ";\n"
	}
	if node.Right != nil {
		currentNodeStr += strconv.Itoa(node.Key)
		currentNodeStr += "->"
		currentNodeStr += strconv.Itoa(node.Right.Key)
		currentNodeStr += ";\n"
	}
	if GetColor(node) == BLACK {
		currentNodeStr += strconv.Itoa(node.Key)
		currentNodeStr += "[style = \"filled\" fillcolor = \"black\" fontcolor=\"white\"];\n"
	} else {
		currentNodeStr += strconv.Itoa(node.Key)
		currentNodeStr += "[style = \"filled\" fillcolor = \"red\" fontcolor=\"white\"];\n"
	}

	currentNodeStr += createRBTString(node.Left)
	currentNodeStr += createRBTString(node.Right)

	return currentNodeStr
}
