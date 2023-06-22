# gocli

## 初始化项目

```bash
# 初始化 git 仓库
$ git init

# 初始化 cobra 的项目结构
$ cobra-cli init
```

## 发布版本与打tag

在进行版本发布时，通常的流程是先完成代码的提交和推送，确保代码已经上传到远程仓库的主分支（例如`master`分支）。这包括执行以下命令：

```bash
$ git add .  
$ git commit -m "update"
$ git push origin master
```

这些命令将暂存的文件添加到提交中，创建一个新的提交并将其推送到远程仓库的`master`分支。

然后，可以使用以下命令创建一个新的标签并将其推送到远程仓库：

```bash
$ git tag -a v1.0.2 -m "Version 1.0.2"
$ git push origin v1.0.2
```

这些命令会创建一个名为`v1.0.2`的标签，并将其关联到最新的提交。然后将标签推送到远程仓库，使其在远程仓库中可见。

总结：先完成代码提交和推送，确保代码上传到主分支，然后再创建和推送标签。这样可以确保标签是基于已提交的代码创建的，并与特定的代码版本关联起来。

## 编译程序并注入版本信息

```bash
$ cd gocli/
$ make build
```

## 运行二进制程序

```bash
$ ./gocli
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  gocli [flags]
  gocli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  start       Start the project
  version     Print the version number

Flags:
  -h, --help      help for gocli
  -v, --version   Print the version number

Use "gocli [command] --help" for more information about a command.

$ ./gocli version
gocli version information: 
  Version: v1.0.2
  Git Commit: c4a1566
  Build Time: Thu, 22 Jun 2023 16:39:42 +0800
  Go version: go1.20.4
  OS/Arch: darwin/arm64

$ ./gocli -v
gocli v1.0.2

$ ./gocli start
Starting the project...
```
