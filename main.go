package arkctlModularityCmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 创建一个 Scanner 对象来读取用户输入
	scanner := bufio.NewScanner(os.Stdin)

	// 提示用户输入项目路径
	fmt.Println("请输入工程的绝对路径：")
	scanner.Scan()                // 读取用户输入
	projectPath := scanner.Text() // 获取输入的文本

	// 提示用户输入应用名称
	fmt.Println("请输入要设置的应用名称：")
	scanner.Scan()                    // 再次读取用户输入
	applicationName := scanner.Text() // 获取输入的文本

	// 指定 JAR 文件的路径
	jarPath := "D:\\demo\\target\\demo-0.0.1-SNAPSHOT.jar"

	// 创建一个命令来运行 Java 程序
	javaCmd := exec.Command("java", "-jar", jarPath)

	// 创建一个管道，用于传递输入到 Java 程序
	stdinPipe, err := javaCmd.StdinPipe()
	if err != nil {
		fmt.Printf("创建标准输入管道失败: %s\n", err)
		return
	}

	// 将标准输出和标准错误重定向到缓冲区
	var outBuffer bytes.Buffer
	javaCmd.Stdout = &outBuffer
	javaCmd.Stderr = &outBuffer

	// 启动 Java 程序
	if err := javaCmd.Start(); err != nil {
		fmt.Printf("启动 Java 程序出错: %s\n", err)
		return
	}

	// 写入项目路径到 Java 程序的标准输入
	if _, err := stdinPipe.Write([]byte(projectPath + "\n")); err != nil {
		fmt.Printf("写入项目路径出错: %s\n", err)
		return
	}

	// 写入应用名称到 Java 程序的标准输入
	if _, err := stdinPipe.Write([]byte(applicationName + "\n")); err != nil {
		fmt.Printf("写入应用名称出错: %s\n", err)
		return
	}

	// 关闭标准输入管道
	if err := stdinPipe.Close(); err != nil {
		fmt.Printf("关闭标准输入管道出错: %s\n", err)
		return
	}

	// 等待 Java 程序执行完成
	if err := javaCmd.Wait(); err != nil {
		fmt.Printf("Java 程序运行出错: %s\n", err)
		fmt.Printf("错误输出: %s\n", outBuffer.String())
		return
	}

	// 打印 Java 程序的输出
	fmt.Printf("Java 程序输出:\n%s\n", outBuffer.String())
}
