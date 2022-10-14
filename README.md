# go-whisper

使用 Go 编写的个人博客系统。

## 功能特性

- [x] Markdown 支持；
- [x] 快速部署：一个二进制文件丢到服务器上搞定，不需要额外的运行环境；
- [x] 低成本运行：其实也不是很低，大概会占用十多M内存；使用 SQLite 数据库； 
- [x] **随意书写：** 很多静态博客生成工具都满足以上几个需求，但用这些工具的体验就是：写东西是一件很严肃的事，我必须摆开架势创建一个 markdown 文件，然后生成静态网页，然后发布。 ~~太麻烦了！~~  不太适合懒人，我希望写东西可以随心所欲一些：可以一本正经的长篇大论，也可以像用随手贴一样随便记录点什么。 使用 Go-Whisper 可随手打开网页写东西，你甚至可以不用写标题（就像在发布一条 twitter 一样），也不用考虑写下的东西放到哪个分类下（当然，你可以灵活的使用标签来实现归整内容的目的）；
- [x] 自定义模板： 毕竟程序员的审美……
- [x] 数据库每日自动备份，目前仅支持备份到腾讯云（是的，腾讯云的永久免费额度是真香😆）； 

## 安装

- 克隆代码 `git clone git@github.com:go-whisper/go-whisper.git`
- 编译 `go build`
- 复制配置文件 `cp config.example.toml config.toml`
- 运行安装脚本生成表结构及初始化用户 `./go-whisper install`
- 运行程序 `./go-whisper` 
- 打开浏览器 `http://127.0.0.1:8080` 初始账号密码均为： `whisper`

**如果你需要用它正二八经的干活，你可能还需要做以下操作：**
- 让 `go-whisper` 以服务的方式运行（你可能需要借助 `supervisor`、 `systemd` 等工具）  
- 你可以编辑 `config.toml` 文件，修改一些配置项  

## 接下来做什么

- [ ] 通过管理面板恢复指定数据库
- [ ] 图片上传、管理


## 截图

![屏幕截图](./screenshot.jpg)

