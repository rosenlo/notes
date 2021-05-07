# Memo

## K8S

- 在服务上 K8S 过程中发现 `cpu.switches` 大的应用会影响同 node 上的其他应用（具体表现为业务接口响应时间变大），应单独部署或隔离。
- `container_cpu_cfs_throttled_periods_total` - CPU Throttled 会影响业务接口的响应时间，具体表现为受 resources: limit  影响，虽然实际 usage 远远未到 limit ，做法是取消 CPU limit 只保留 CPU request

## Prometheus & VictoriaMetrics

- vmagent 在某个点之后内存快速增长，经过 `pprof` 分析， `github.com/VictoriaMetrics/VictoriaMetrics/lib/promscrape.appendLabels`
占用了大量内存，最终发现研发把 exception 的堆栈内容丢到 label，导致一次 scrape 数据量达到100M+ ， 由于 `maxScrapeSize` 设置为200M，没有 reject 直接保存到内存导致内存快速增长，最终调整 `maxScrapeSize` 为50MB恢复，建议生产环境 `maxScrapeSize` 不要调过大

## Go

- Goroutine 与 OS Threads 是 M:1 的关系，即多个 Goroutinue 对应一个 OS Threads
