package arkctlModularityCmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// 定义 Modularity 命令
var ModularityCmd = &cobra.Command{
	Use:   "modularity",
	Short: "Run the Modularity subcommand",
	Long:  `Run the Modularity subcommand to deploy your application by running a Java JAR file.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectPath, _ := cmd.Flags().GetString("projectPath")
		applicationName, _ := cmd.Flags().GetString("applicationName")

		jarPath := "D:\\demo\\target\\demo-0.0.1-SNAPSHOT.jar"
		runJavaProgram(jarPath, projectPath, applicationName)
	},
}

// runJavaProgram 运行 Java 程序
func runJavaProgram(jarPath, projectPath, applicationName string) {
	var outBuffer bytes.Buffer
	javaCmd := exec.Command("java", "-jar", jarPath, projectPath, applicationName)
	javaCmd.Stdout = &outBuffer
	javaCmd.Stderr = &outBuffer

	if err := javaCmd.Run(); err != nil {
		fmt.Printf("Java 程序运行出错: %s\n", err)
		fmt.Printf("错误输出: %s\n", outBuffer.String())
		return
	}

	fmt.Printf("Java 程序输出:\n%s\n", outBuffer.String())
}

func main() {
	var rootCmd = &cobra.Command{Use: "arkctl"}
	rootCmd.AddCommand(ModularityCmd)

	// 定义命令行参数
	rootCmd.PersistentFlags().StringP("projectPath", "p", "", "Project path")
	rootCmd.PersistentFlags().StringP("applicationName", "a", "", "Application name")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
