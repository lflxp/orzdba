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

## Args

```
Usage of /var/folders/xs/s2b_zrvs7bz1jkl960nmvq240000gn/T/go-build631741448/b001/exe/main:
  -B	Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).
  -C int
    	运行时间 默认无限
  -H string
    	Mysql连接主机，默认127.0.0.1 (default "127.0.0.1")
  -L string
    	Print to Logfile. (default "none")
  -P string
    	Mysql连接端口,默认3306 (default "3306")
  -S string
    	mysql socket连接文件地址 (default "/tmp/mysql.sock")
  -T	Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).
  -c	打印Cpu info
  -com
    	Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).
  -d string
    	打印Disk info (default "none")
  -hit
    	Print Innodb Hit%.
  -i string
    	时间间隔 默认1秒 (default "1")
  -innodb
    	Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)
  -innodb_data
    	Print Innodb Data Status(Innodb_data_reads/writes/read/written)
  -innodb_log
    	Print Innodb Log  Status(Innodb_os_log_fsyncs/written)
  -innodb_pages
    	Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)
  -innodb_rows
    	Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).
  -innodb_status
    	Print Innodb Status from Command: 'Show Engine Innodb Status'
  -l	打印Load info
  -lazy
    	Print Info  (include -t,-l,-c,-s,-com,-hit).
  -logfile_by_day
    	One day a logfile,the suffix of logfile is 'yyyy-mm-dd';
  -mysql
    	Print MySQLInfo (include -t,-com,-hit,-T,-B).
  -n string
    	打印net info (default "none")
  -nocolor
    	不显示颜色
  -p string
    	mysql密码 (default "system")
  -rt
    	Print MySQL DB RT(us).
  -s	打印swap info
  -semi
    	半同步监控
  -slave
    	打印Slave info
  -sys
    	Print SysInfo   (include -t,-l,-c,-s).
  -t	打印当前时间
  -u string
    	mysql用户名 (default "root")
  ```
  

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
