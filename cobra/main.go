package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type test struct {
	nums []int
}

func main() {
	//t := &test{}

	var nums []int

	// 创建一个新的 cobra 命令
	var rootCmd = &cobra.Command{
		Use:   "example",
		Short: "An example CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			// 打印解析后的标志值
			fmt.Println("Parsed numbers:", &nums)
		},
	}

	// 使用 cobra 的 Flags 方法定义一个整数切片标志
	rootCmd.PersistentFlags().IntSliceVar(&nums, "numbers", []int{1, 2, 3}, "A list of integers")
	fmt.Println(nums)
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
