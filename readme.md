# v2ex-clockin

这是一个演示如何使用 [rod](https://github.com/ysmood/rod) 的项目，全部代码只有约 50 行。

## 安装运行

除了一个可执行文件无需任何依赖，可以脚本下载运行：

```bash
curl -L https://git.io/fjaxx | repo=ysmood/v2ex-clockin sh

v2ex-clockin
```

或者去[发布页面](https://github.com/ysmood/v2ex-clockin/releases)下载对应 OS 的可执行文件，

或者 `go run .` 运行源代码。

**＊ 第一次启动或者 cookie 失效时会自动弹出登陆页面，此时手动登陆下即可**

Cookie 会被保存到 `./tmp` 文件夹。

可以用 tmux 之类的工具让它在后台运行。

如果有人可以提供训练好的 AI 识别验证码就可以放 docker 里完全无介入运行了，有闲工夫的话我自己会试试 pytorch。

这个项目使用 [kit](https://github.com/ysmood/kit) 自动编译发布。
