package adapter

import "github.com/fit2cloudrd/operators-mysql-cluster/pkg/tools/log"

type LogIN interface {
	Info(string)
	Warn(string)
	Error(string)
}

type RealLog struct {
}

func (r *RealLog) info(s string) {
	log.Info(s)
}

func (r *RealLog) warn(s string) {
	log.Warning(s)
}

func (r *RealLog) error(s string) {
	log.Error(s)
}

type AdapterLog struct {
	log *RealLog
}

func (al *AdapterLog) Info(s string) {
	al.log.info(s)
}

func (al *AdapterLog) Warn(s string) {
	al.log.warn(s)
}

func (al *AdapterLog) Error(s string) {
	al.log.error(s)
}

func Test() {
	adapterLog := AdapterLog{log: &RealLog{}}
	adapterLog.Info("db test info")
	adapterLog.Warn(" db test warn")
	adapterLog.Error("db test error")
}
