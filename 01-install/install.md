# 安装

由于`golang`的泛型还没有正式发布，所以需要使用最新的开发版本的`go`来做泛型的演练。

## go包的代理设置

如果还没有设置过代理，则需要执行下面的代码（linux或者mac）：

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```



## 安装gotip

```bash
go install golang.org/dl/gotip@latest
```

得到`gotip`后(一般在`GOPATH`的`bin`目录下，如果没有设置`GOPATH`，则在`$HOME/go/bin`目录下），然后获取最新的`golang`。

```bash
gotip download
```

运行的时候，只要将原先用的`go`命令中的`go`，用新的`gotip`来替换。

