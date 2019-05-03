# 一个基于 gin 开发的静态博客

## 为什么开发这个博客?

- 想要定制化，想要不一样
- 而且别人写的前端代码太复杂，看不懂
- 不喜欢在线编辑，想通过 markdown 写作

## 本博客采用的方式

- 服务端语言用的是 golang
- web 框架用的 gin
- 静态页面使用 Typora 写好导出 .html 页面后放入 assets/static 目录下 (后期可能采用 makedown 转 html 工具批量渲染)
- 业界有好用的静态博客生成工具 hugo， hugo 写的很好且很多人用，不过我是嫌弃默认的皮肤不好看且不想去研究如何自定义皮肤 (懒...)

## 目录结构

```

├── Makefile
├── README.MD
├── admin.sh
├── assets
│   ├── css
│   │   ├── blog.css
│   │   └── bootstrap.min.css
│   ├── js
│   │   ├── bootstrap.min.js
│   │   └── jquery.min.js
│   └── static
│       ├──xxx.html
│       ├──xxx.html
│       ├── images
│       │   ├── design-database.jpg
│       │   ├── download.sh
│       │   ├── effective-window.jpg
│       │   ├── fin-ack.jpg
│       │   ├── osi.gif
│       │   ├── tcp-data-state.jpg
│       │   ├── tcp-open-close.jpeg
│       │   ├── tcp3.png
│       │   └── tcp_ip.jpg
│       ├── xxxx.html
│       ├── xxx.html
│       ├── xxx.html
│       └── xxx.html
├── conf
│   └── config.yaml
├── favicon.ico
├── handler
│   └── blog
│       ├── blog.go
│       └── blog_test.go
├── main.go
├── resources
├── router
│   ├── middleware
│   │   └── favicon.go
│   └── router.go
└── templates
    └── index.html

```

main.go 程序入口

router 加载路由，中间件 (middleware)

Makefile  封装常用操作， 如 build, clean ...

admin.sh  服务启动/停止/状态/重启 脚本

assets 一些静态资源

assets/static 目录下是放静态的html 页面 , 在博客首页会自动列出 static 下所有文件的名称

handler 对应路由的处理函数

## conf/config.yaml 存放了可配置内容

```yaml
runmode: debug  # Gin 的运行模式
port: :8080  # server port
max_ping_count: 3 # health check max try  count 
url: http://127.0.0.1:8080  # server url
blog_title: "xiaozefeng's blog" # blog title
```

