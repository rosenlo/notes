# Kubernetes Pod 的水平自动伸缩机制（Horizontal Pod Autoscaler）

HPA 机制是基于 Pod 的历史 CPU 利用率（也支持自定义的 metric，如应用的 QPS 等）来计算副本数量。适用于 `replication controller`, `deployment` 和 `replica set`。有些对象不适合使用，例如：`DaemonSets`

HPA 被实现为 Kubernetes 的 API 资源和控制器（controller）, 该资源确定了控制器的行为，控制器周期性（默认15s）的计算当前平均 CPU 利用率与用户配置的阈值（target）是否相匹配

## HPA 如何工作？

在每个周期，HPA 控制器会查询预先定义的 metrics（包含自定义 metrics）的资源使用情况
- 针对每个 pod 内置的 metrics (比如 CPU)，控制器会从 metrics API 获取每个 Pod
  对应的 metrics 数据。如果配置了利用率值（target utilization value）控制器会转换为百分比值。如果配置为原始值（raw value），则直接计算。控制器会根据所有
  pods 计算一个平均值然后提供一个速率值（ratio value）用来控制伸缩副本的速度

注意：如果有些 pod 没有相关指标，那么控制器不会采取任何动作
- 针对每个 pod 自定义的 metrics，控制器行为相似，只有一点区别：**只支持始值，不支持利用率值**

HPA 通常从聚合（aggregated）APIs（`metrics.k8s.iod, `custom.metrics.k8s.io`, `external.metrics.k8s.io`） 获取 metrics。`metrics.k8s.io` API 通常由
`metrics-server` 提供，它需要独立部署。更多细节参考 [metrics-server](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/#metrics-server)


### HPA 算法

```
desiredReplicas = ceil[currentReplicas * ( currenMetricValue / desiredMetricValue )]
```
