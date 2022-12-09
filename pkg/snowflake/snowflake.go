package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake // 实例
	sonyMachineID uint16               // 机器ID
)

func getMachineID() (uint16, error) { // 返回全局定义的机器ID
	return sonyMachineID, nil
}

func Init(startTime string, machineID uint16) (err error) {
	sonyMachineID = machineID
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime) // 初始化一个开始的时间
	fmt.Println(startTime)
	if err != nil {
		return
	}
	// 生成全局配置
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID, // 指定机器ID
	}

	sonyFlake = sonyflake.NewSonyflake(settings) // 用配置生成sonyflake节点
	return
}

// GetID 返回生成的id值
func GenID() (id uint64, err error) {
	// 拿到sonyflake节点生成id值
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
