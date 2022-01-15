## GO package管理方式

  该项目下所有代码都是通过GOPATH来管理的，所以有的代码如果你使用go mod运行可能会报错。
  建议将`GO111MODULE`设置为`off`或者`auto`。

```bash
go env -w GO111MODULE=auto
go env -w GO111MODULE=on
go env -w GO111MODULE=off
```