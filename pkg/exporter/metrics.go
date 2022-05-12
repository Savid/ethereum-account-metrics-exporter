package exporter

import (
	"github.com/samcm/ethereum-metrics-exporter/pkg/exporter/consensus"
	"github.com/samcm/ethereum-metrics-exporter/pkg/exporter/disk"
	"github.com/samcm/ethereum-metrics-exporter/pkg/exporter/execution"
	"github.com/sirupsen/logrus"
)

type Metrics interface {
	Consensus() consensus.Metrics
	Disk() disk.Metrics
}

type metrics struct {
	consensus consensus.Metrics
	execution execution.Metrics
	disk      disk.Metrics
}

func NewMetrics(log logrus.FieldLogger, executionName, consensusName, namespace string) Metrics {
	return &metrics{
		consensus: consensus.NewMetrics(log, consensusName, namespace),
		disk:      disk.NewMetrics(namespace),
	}
}

func (m *metrics) Consensus() consensus.Metrics {
	return m.consensus
}
func (m *metrics) Execution() execution.Metrics {
	return m.execution
}

func (m *metrics) Disk() disk.Metrics {
	return m.disk
}
