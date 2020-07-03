# v2ex-example

这是一个演示如何使用 [rod](https://github.com/go-rod/rod) 的项目。整个项目仅一个 `main.go` 文件。

## 安装运行

除了一个可执行文件无需任何依赖，可以脚本下载运行：

```bash
curl -L https://git.io/fjaxx | repo=go-rod/v2ex-example sh

v2ex-example # 不传任何配置就是每天中午自动签到

v2ex-example -topic https://www.v2ex.com/t/686655 -interval 1h # 每隔 1 小时点击一次置顶主题

v2ex-example -help # 查看更多设置
```

或者去[发布页面](https://github.com/go-rod/v2ex-example/releases)下载对应 OS 的可执行文件，

或者 `go run .` 运行源代码。

**＊ 第一次启动或者 cookie 失效时会自动弹出登陆页面，此时手动登陆下即可**

Cookie 会被保存到 `./tmp` 文件夹。

可以用 tmux 之类的工具让它在后台运行。

如果有人可以提供训练好的 AI 识别验证码就可以放 docker 里完全无介入运行了，有闲工夫的话我自己会试试 pytorch。

这个项目使用 [kit](https://github.com/ysmood/kit) 自动编译发布。
