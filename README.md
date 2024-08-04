# 北邮宿舍电费定时检测提醒

为了避免老旧校区电量用尽导致再次交电费后插座和空调没电的情况，特写此脚本用于定时检验电量剩余度数并邮件告警

## 运行方法

1. 进入 https://app.bupt.edu.cn/buptdf/wap/default/search 该网站获取自己宿舍的相关信息并完善 yml 配置文件
2. 配置 github action 的环境变量
3. 完成所有配置后 github action 即可自动运行
