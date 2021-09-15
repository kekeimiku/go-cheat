<img align="left" width="160" src="./assets/miku.png">

## go-cheat

### users friendly linux  hacking library   
[![Build & Test](https://github.com/kekeimiku/go-cheat/actions/workflows/test.yml/badge.svg)](https://github.com/kekeimiku/go-cheat/actions/workflows/test.yml)

## ！！！ Not reach usable state

## Install
```
go get -u github.com/kekeimiku/cheat
```

The current function is still very basic

### 目前的计划：

- [ ] 友好的使用libc api读写内存
- [x] 根据进程名获取pid
- [x] 解析/proc/self/maps
- [ ] 写/proc/self/mem
- [ ] 内存搜索
- [ ] 冻结内存
- [ ] 实用算法
- [ ] 其它实用功能
- [ ] 其它平台

这个项目目前专注于实现功能中，可能充斥很多无用代码，一旦功能成熟我会进行清理。

当前阶段api不会向后兼容。