# 微服务 (GO-MICRO) 项目 Proto 文件管理

本项目依赖于B站开源的 `kratos` 框架, 目前服务有:

* 文章服务 (article)

* 用户服务 (user)

对应的服务项目 GitHub 仓库 -> git@github.com:raylin666/go-micro.git

代码拉下来后, 通过项目提供的 `Makefile` 文件可完成项目的初始化和代码的自动生成。

1. 初始化, 下载依赖操作命令:
> make init     

2. 编写完 `proto` 文件后, 生成代码文件命令:
> make api