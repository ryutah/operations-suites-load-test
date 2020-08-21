# operations-suites-load-test

Cloud Trace, Cloud Profilerがシステムのレイテンシにどの程度影響を与えるのかを調査してみた

## Load test results

### Hello, World

Hello, Worldを返すだけの簡単なエンドポイントで検証。

この結果を見る感じ、全然影響ないように見える

#### Default

```txt
Requests      [total, rate, throughput]         300, 10.03, 10.01
Duration      [total, attack, wait]             29.973s, 29.9s, 72.965ms
Latencies     [min, mean, 50, 90, 95, 99, max]  60.496ms, 133.897ms, 113.358ms, 221.454ms, 293.956ms, 551.867ms, 723.208ms
Bytes In      [total, mean]                     3900, 13.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:300  
Error Set:
```

#### With Trace

```txt
Requests      [total, rate, throughput]         300, 10.03, 10.00
Duration      [total, attack, wait]             30.003s, 29.9s, 102.557ms
Latencies     [min, mean, 50, 90, 95, 99, max]  66.089ms, 133.874ms, 120.636ms, 196.339ms, 243.607ms, 392.285ms, 492.676ms
Bytes In      [total, mean]                     3900, 13.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:300  
Error Set:
```

#### With Profiler

```txt
Requests      [total, rate, throughput]         300, 10.03, 10.00
Duration      [total, attack, wait]             30.001s, 29.899s, 101.378ms
Latencies     [min, mean, 50, 90, 95, 99, max]  61.041ms, 115.306ms, 105.714ms, 170.865ms, 194.634ms, 285.692ms, 457.401ms
Bytes In      [total, mean]                     3900, 13.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:300  
Error Set:
```

#### With Both

```txt
Requests      [total, rate, throughput]         300, 10.03, 10.01
Duration      [total, attack, wait]             29.972s, 29.9s, 71.437ms
Latencies     [min, mean, 50, 90, 95, 99, max]  64.511ms, 91.881ms, 88.342ms, 116.732ms, 130.593ms, 218.557ms, 282.221ms
Bytes In      [total, mean]                     3900, 13.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:300  
Error Set:
```
