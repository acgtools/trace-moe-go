# trace-moe-go

用于搜索番剧图片来源的 TUI 应用，使用 [trace.moe](https://trace.moe/) API。

如果这个程序对你有所帮助，可以帮忙给一个 star (o゜▽゜)o☆ ，谢谢 OwO。

> 随机 Wink OvO

<img src="https://waifu-getter.vercel.app/sfw?eps=wink" />

<br />
<!-- 
  If you want to use your own Moe-Counter
  please refer to the tutorial 
  in its original repo: https://github.com/journey-ad/Moe-Counter
  and deploy it to the Replit or Glitch
-->
![](https://political-capable-roll.glitch.me/get/@acgtooltracemoego?theme=rule34)

## 安装

### 使用 `go`

```sh
$ go install -ldflags "-s -w" github.com/acgtools/trace-moe-go
```

### 从 Release 页面下载

[release page](https://github.com/acgtools/trace-moe-go/releases)

## 快速开始

```sh
$ moe-go -h
A TUI app for finding anime scene by image, using trace.moe api

Usage:
  moe-go [command]

Available Commands:
  file        search image by file
  help        Help about any command

Flags:
  -h, --help      help for moe-go
  -v, --version   version for moe-go

Use "moe-go [command] --help" for more information about a command.
```

### 确保你的终端字符集为 UTF-8

#### Windows

```cmd
> chcp
Active code page: 65001

# 如果输出结果不是 65001(utf-8), 可以临时改变当前字符集
> chcp 65001
```

如果你想修改默认的字符集, 按照以下步骤:

1. 开始 -> 运行 -> regedit
2. 找到 `[HKEY_LOCAL_MACHINE\Software\Microsoft\Command Processor\Autorun]`
3. 将其值修改为 `@chcp 65001>nul`

如果 `Autorun` 不存在, 你可以创建一个新的字符串类型的键值对.

此方法将在`cmd` 启动时自动执行 `@chcp 65001>nul`。

#### Linux

```sh
$ echo $LANG
en_US.UTF-8
```

### 搜索图片文件

```sh
$ moe-go file <image file path>
```

按键:

- `up`, `down`: 移动
- `space` ,`enter`: 选择一个查找结果
- `q`: 退出程序

#### 示例

![gochiusa_rize](https://raw.githubusercontent.com/dreamjz/pics/main/pics/2023/202312042054552.jpg)

![1](https://raw.githubusercontent.com/dreamjz/pics/main/pics/2023/202312042051978.gif)

## Issue

欢迎创建 issue 来报告 bug 或者 请求添加新特性。
