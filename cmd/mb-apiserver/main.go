// Copyright 2025 武晓晨 <wuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iWuxc/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package main

import (
	"github.com/iWuxc/miniblog/cmd/mb-apiserver/app"
	_ "go.uber.org/automaxprocs"
	"os"
)

// Go 程序的默认入口函数。阅读项目代码的入口函数.
func main() {
	// 创建迷你博客命令
	command := app.NewMiniBlogCommand()

	//执行命令并处理错误
	if err := command.Execute(); err != nil {
		// 如果发生错误，则退出程序
		// 返回退出码，可以使其他程序（例如 bash 脚本）根据退出码来判断服务运行状态
		os.Exit(1)
	}
}
