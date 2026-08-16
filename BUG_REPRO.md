# BUG_REPRO

## Bug 是什么
复查状态校验被移除；复查更新条件写成 id<>id；异常指标列表排序写反；默认状态写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestFollowUpChain 失败：非法状态被接受、复查状态未更新。
