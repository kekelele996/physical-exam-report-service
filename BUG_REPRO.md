# BUG_REPRO

## Bug 是什么
套餐成交量用 sum 替代 count；异常指标统计把正常结果算进来；排序写反；登记默认状态写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestStatsAggregation 失败：套餐成交量与异常指标数量不正确。
