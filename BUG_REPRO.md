# BUG_REPRO

## Bug 是什么
AppError.Unwrap 被改成返回 nil；检查结果、体检人、登记查询未命中返回新错误，导致 errors.Is 失效。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestErrorChainSentinels 失败：FindByID 未命中不再匹配 ErrNotFound。
