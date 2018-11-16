package pipeline

import (
	"sort"

	"github.com/luoyeaichifan/pholcus/app/pipeline/collector"
	"github.com/luoyeaichifan/pholcus/common/kafka"
	"github.com/luoyeaichifan/pholcus/common/mgo"
	"github.com/luoyeaichifan/pholcus/common/mysql"
	"github.com/luoyeaichifan/pholcus/runtime/cache"
)

// 初始化输出方式列表collector.DataOutputLib
func init() {
	for out, _ := range collector.DataOutput {
		collector.DataOutputLib = append(collector.DataOutputLib, out)
	}
	sort.Strings(collector.DataOutputLib)
}

// 刷新输出方式的状态
func RefreshOutput() {
	switch cache.Task.OutType {
	case "mgo":
		mgo.Refresh()
	case "mysql":
		mysql.Refresh()
	case "kafka":
		kafka.Refresh()
	}
}
