# 一个基于 gin 开发的静态博客

## 为什么开发这个博客?

- 想要定制化，不想跟别人的一样
- 而且别人写的前端代码太复杂，看不懂
- 不喜欢在线编辑，想通过 markdown 写作

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

static 目录下是放静态的html 页面 , 在博客首页会自动列出 static 下所有文件的名称

conf/config.yaml 存放了可配置内容