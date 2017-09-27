# orzdba
rebuild alibaba orzdba by golang

# 实现原理 

## Shell+Golang

> //基础:load+cpu+swap
//cat /proc/loadavg /proc/stat /proc/vmstat /proc/net/dev /proc/diskstats |grep -E '/|cpu|pswpin|pswpout'|sed 's/\// china /g'|grep -E -w 'china|cpu|pswpin|pswpout'|xargs echo|awk '{print $1,$2,$3,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$20/240,$22/240}'|sed 's/[[:space:]]/,/g'
//0.22,0.27,0.31,1673683,41,2065716,343819349,6613289,63,105102,0,0,0,0,0
//基础+网络
//cat /proc/loadavg /proc/stat /proc/vmstat /proc/net/dev /proc/diskstats |grep -E '/|cpu|pswpin|pswpout|eth0'|sed 's/\// china /g'|grep -E -w 'china|cpu|pswpin|pswpout|eth0'|xargs echo|awk '{print $1,$2,$3,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$20/240,$22/240,$24,$32}'|sed 's/[[:space:]]/,/g'
//基础+磁盘
//cat /proc/loadavg /proc/stat /proc/vmstat /proc/net/dev /proc/diskstats |grep -E '/|cpu|pswpin|pswpout|sda'|sed 's/\// china /g'|grep -E -w 'china|cpu|pswpin|pswpout|sda'|xargs echo|awk '{print $1,$2,$3,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$20/240,$22/240,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36}'|sed 's/[[:space:]]/,/g
//全数据：基础+网络+磁盘
// cat /proc/loadavg /proc/stat /proc/vmstat /proc/net/dev /proc/diskstats |grep -E '/|cpu|pswpin|pswpout|eth0|sda'|sed 's/\// china /g'|grep -E -w 'china|cpu|pswpin|pswpout|sda|eth0'|xargs echo|awk '{print $1,$2,$3,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$20/240,$22/240,$24,$32,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53}'|sed 's/[[:space:]]/,/g'

# 交叉编译

1. Linux 下编译 Mac 和 Windows 64位可执行程序

> CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go  
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go  

2. Windows 下编译 Mac 和 Linux 64位可执行程序

> SET CGO_ENABLED=0  
SET GOOS=darwin  
SET GOARCH=amd64  
go build main.go  
  
SET CGO_ENABLED=0  
SET GOOS=linux  
SET GOARCH=amd64  
go build main.go  
