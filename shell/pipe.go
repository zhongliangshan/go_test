package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"io"
)

func ShellCmd(cmds []*exec.Cmd) ([]string , error) {
	if len(cmds) <= 0 {
		return nil, fmt.Errorf("shell 命令为空")
	}
	var err error
	first := true
	var output []byte


	for len(cmds) > 0 {
		cmd := cmds[0]

		cmds = cmds[1:]

		if !first {
			var stdinput bytes.Buffer

			stdinput.Write(output)
			cmd.Stdin = &stdinput
		}


		var buffers bytes.Buffer

		cmd.Stdout = &buffers

		if err = cmd.Start(); err != nil {
			return nil, err
		}
		// 需要等到命令全部处理完 再去处理命令获取到的数据
		if err = cmd.Wait(); err != nil {
			return nil, err
		}

		output = buffers.Bytes()

		first = false

	}

	lines := make([]string, 0)
	var outputBuf bytes.Buffer
	outputBuf.Write(output)
	for {
		line, err := outputBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil

}

func main() {

	var cmds = []*exec.Cmd{
		exec.Command("ps" , "aux"),
		exec.Command("grep" , "bin"),
	}

	buffer, e := ShellCmd(cmds)
	if e != nil {
		panic(e)
	}

	fmt.Println(buffer)
}
