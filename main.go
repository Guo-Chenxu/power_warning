package main

import (
	"fmt"
	"power_warning/conf"
	"power_warning/logic"
)

const star = "如果觉得本项目好用的话, 还请给个star吧 φ(゜▽゜*)♪ <br> 项目地址: https://github.com/Guo-Chenxu/power_warning"

func main() {
	config := conf.GetConfig()

	power, err := logic.GetPower(config.RoomConfig)
	if err != nil {
		fmt.Println("获取电量余额出现错误：", err)
		return
	}
	fmt.Println("当前电量信息：", power)

	if power.D.Data.Surplus < config.WarningThreshold {
		content := fmt.Sprintf("电量余额低于阈值，请及时充电！告警阈值：%.2f 度，剩余电量：%.2f 度", config.WarningThreshold, power.D.Data.Surplus)
		fmt.Println(content)
		config.MailConfig.Body = content + "<br><br>" + star
		
		if err := logic.SendEmail(config.MailConfig); err != nil {
			fmt.Println("发送邮件警告出现错误：", err)
		} else {
			fmt.Println("邮件警告发送成功。")
		}
	} else {
		fmt.Println("电量余额充足，无需充电。")
	}
}
