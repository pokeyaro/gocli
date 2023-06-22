# gocli

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
$ git tag -a v1.0.1 -m "Version 1.0.1"
$ git push origin v1.0.1
```

这些命令会创建一个名为`v1.0.1`的标签，并将其关联到最新的提交。然后将标签推送到远程仓库，使其在远程仓库中可见。

总结：先完成代码提交和推送，确保代码上传到主分支，然后再创建和推送标签。这样可以确保标签是基于已提交的代码创建的，并与特定的代码版本关联起来。

## 编译程序并注入版本信息

```bash
$ cd gocli/
$ make build
```

## 运行二进制程序

```bash
$ ./app version
gocli version information: 
  Version: v1.0.1
  Git Commit: c4a1566
  Build Time: Thu, 22 Jun 2023 16:28:36 +0800
  Go version: go1.20.4
  OS/Arch: darwin/arm64
```
