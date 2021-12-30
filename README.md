# 基金走势分析通知工具
使用yaml文法表达基金交易策略的基金走势分析通知工具。它允许你在yaml中指定基金，指定需要分析的时间跨度（过去某天为起点，当天为终点），并定义你自己的交易策略。交易策略被触发时，你会得到你自己定义的输出。

## 内容列表
- [背景](#背景)
- [已实现的可自定义功能](#已实现的可自定义功能)
- [交易策略配置说明](#交易策略配置说明)
- [使用者注意事项](#使用者注意事项)
- [使用](#使用)
- [维护者](#维护者)

## 背景
上班时，过渡关注基金容易导致消耗过多的精力，或头脑一热做出不理智判断，很可能导致整整一天心不在焉。所以想开发一款工具，能将自己的交易策略作为输入，代替自己判断，判断结果在一个交易日的某个指定时刻输出，并通知到自己。

## 已实现的可自定义功能
- `涨跌监视`： 定义上界或下界，时间跨度内的指定基金总涨幅超过该阈值后触发自定义消息通知
- `横盘检测`： 定义上届和下界，时间跨度内的指定基金从现在到过去的涨幅之和一直在范围内时，触发自定义消息通知

## 交易策略配置说明
```
myStrategies:       # 策略集合
- code: "519697"    # 基金编号
  subStrategies:    # 对应基金的多条子交易策略
  - border:         # 指定涨跌监视的上下界，每次只能定义上届或下界
      span: 7       # 时间跨度（当天为终点，7天前为起点）
      min: -2       # 跌幅下界负2个点，超出负2个点时输出消息
    msg: "交银优势行业混合一周内跌幅超过2个点，记得加仓哟"  # 需要输出的消息
  - border:
      span: 30
      max: 4
    msg: "交银优势行业混合一个月内涨幅超过4个点，记得减仓哟"

- code: "400015"
  subStrategies:
  - border:         # 一条策略，当所有阈值配置都被满足时，才会输出消息
      span: 365
      max: 50
    detectHPM:      # 横盘检测, 超出范围时输出消息
      span: 182     
      min: -20      
      max: 15
    msg: "东方新能源汽车主题混合去年涨幅超过50%，半年内在负20个点和15个点之间振荡，先别买入哟"
```


## 使用者注意事项
本工具只提供命令行内的标准输出，不打算嵌入定时计算、微信公众号、邮箱通知等功能。

如果你想定时计算，可以考虑使用[Github Action](https://docs.github.com/en/actions) + [Crontab](https://crontab.guru/)

如果你想让公众号或者邮箱通知你，可以考虑自己搭一个webhook服务，定时计算结束后触发该webhook，并将输出的消息作为输入。

## 使用
```golang
go run main.go
```

## 维护者
[@TonyShanc](https://github.com/TonyShanc)。

## 如何贡献
非常欢迎你的加入！[提一个 Issue](https://github.com/TonyShanc/fund-strategy/issues/new) 或者提交一个 Pull Request。

遵循 [Contributor Covenant](http://contributor-covenant.org/version/1/3/0/) 行为规范。

## 使用许可

[MIT](LICENSE) © Richard Littauer