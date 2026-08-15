# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

依赖范围写成 <1.x 时，1.5.0 竟然被判定为满足约束；>1.x 也把 1.5.0 当成满足。看起来只要版本里带通配符，前面的比较运算符就完全不起作用了，不带通配符的 <1.0.0 是正常的。请修复带运算符的通配符范围解析，让 <、<=、>、>= 按语义作用到通配符对应的边界上，同时保持裸通配符（如 1.x）仍表示原来的区间，并保证全量测试通过。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/gogo-51
- 仓库地址：https://github.com/zhanglei10281852-gif/gogo-51.git
- parent SHA：d440042233b5090f74f0aa5bcd28d2bd041b47e6

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/gogo-51.git bug-repro
cd bug-repro
git checkout --detach d440042233b5090f74f0aa5bcd28d2bd041b47e6
go test ./internal/rail -run "^TestParseRangeAppliesOperatorToWildcardBounds$" -count=1 -v
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/rail -run "^TestParseRangeAppliesOperatorToWildcardBounds$" -count=1 -v
=== RUN   TestParseRangeAppliesOperatorToWildcardBounds
    wildcard_regression_test.go:30: ParseRange("<1.x").Contains(1.5.0) = true, want false
--- FAIL: TestParseRangeAppliesOperatorToWildcardBounds (0.00s)
FAIL
FAIL	releaserail/internal/rail	0.002s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/rail -run "^TestParseRangeAppliesOperatorToWildcardBounds$" -count=1 -v
=== RUN   TestParseRangeAppliesOperatorToWildcardBounds
    wildcard_regression_test.go:30: ParseRange("<1.x").Contains(1.5.0) = true, want false
--- FAIL: TestParseRangeAppliesOperatorToWildcardBounds (0.01s)
FAIL
FAIL	releaserail/internal/rail	0.115s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

<1.x 排除 1.5.0、包含 0.9.0；>1.x 排除 1.5.0、包含 2.0.0；<=1.x 包含 1.5.0；>=1.x 包含 1.0.0；裸 1.x 仍为 [1.0.0,2.0.0) 区间；双架构定向、全量、build/vet 通过。
