# BUG_REPRO

## Bug 是什么
异常判定边界用 <=/>= 把临界值误判；>、< 参考值判定方向反了；异常等级阈值写反；异常指标默认状态写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestAbnormalJudgementAndLevel 失败：边界值、>、< 判定与等级猜测不正确。
