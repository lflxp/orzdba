package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var black  string
var red    string
var green  string
var yellow string
var blue   string
var purple string
var dgreen string
var white  string
var gunit  bool

type basic struct {
	//basic info
	hostname                           string
	ip                                 string
	db                                 string
	var_binlog_format                  string
	var_max_binlog_cache_size          int
	var_max_binlog_size                int
	var_max_connect_errors             string
	var_max_connections                string
	var_max_user_connections           string
	var_open_files_limit               string
	var_sync_binlog                    string
	var_table_definition_cache         string
	var_table_open_cache               string
	var_thread_cache_size              string
	var_innodb_adaptive_flushing       string
	var_innodb_adaptive_hash_index     string
	var_innodb_buffer_pool_size        int
	var_innodb_file_per_table          string
	var_innodb_flush_log_at_trx_commit string
	var_innodb_flush_method            string
	var_innodb_io_capacity             string
	var_innodb_lock_wait_timeout       string
	var_innodb_log_buffer_size         int
	var_innodb_log_file_size           int
	var_innodb_log_files_in_group      string
	var_innodb_max_dirty_pages_pct     string
	var_innodb_open_files              string
	var_innodb_read_io_threads         string
	var_innodb_thread_concurrency      string
	var_innodb_write_io_threads        string
	//loadavg
	load_1         float64
	load_5         float64
	load_15        float64
	//cpu
	cpu_core       float64
	cpu_usr        int
	cpu_nice       int
	cpu_sys        int
	cpu_idl        int
	cpu_iow        int
	cpu_irq        int
	cpu_softirq    int
	cpu_steal      int
	cpu_guest      int
	cpu_guest_nice int
	//swap
	swap_in        int
	swap_out       int
	//net
	net_recv       int
	net_send       int
	//disk
	io_1           int
	io_2           int
	io_3           int
	io_4           int
	io_5           int
	io_6           int
	io_7           int
	io_8           int
	io_9           int
	io_10          int
	io_11          int
	//tcprstat rt
	rt_count       int
	rt_avg         int
	rt_a5          int
	rt_a9          int
	//mysql -e "show global status" 不用\G
	//mysql -com
	Com_select     int
	Com_insert     int
	Com_update     int
	Com_delete     int
	Com_commit     int
	Com_rollback   int
	//mysql -hit
	//while true;do s1=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $2}'`;s2=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $4}'`;sleep 1;ss1=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $2}'`;ss2=`mysql -e 'show global status'|grep -w -E 'Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads'|xargs echo|awk '{print $4}'`;rs1=$(($ss1-$s1+1));rs2=$(($ss2-$s2));rs3=$((1000000*($rs1-$rs2)/$rs1));echo $rs1,$rs2,$rs3;done
	// (Innodb_buffer_pool_read_requests - Innodb_buffer_pool_reads) / Innodb_buffer_pool_read_requests * 100%,每秒的计算
	Innodb_buffer_pool_read_requests int
	Innodb_buffer_pool_reads         int
	//mysql -innodb_rows
	Innodb_rows_inserted             int
	Innodb_rows_updated              int
	Innodb_rows_deleted              int
	Innodb_rows_read                 int
	//mysql -innodb_pages
	Innodb_buffer_pool_pages_data    int
	Innodb_buffer_pool_pages_free    int
	Innodb_buffer_pool_pages_dirty   int
	Innodb_buffer_pool_pages_flushed int
	//mysql --innodb_data
	Innodb_data_reads   int
	Innodb_data_writes  int
	Innodb_data_read    int
	Innodb_data_written int
	//mysql --innodb_log
	Innodb_os_log_fsyncs  int
	Innodb_os_log_written int
	//mysql --threads
	Threads_running   int
	Threads_connected int
	Threads_created   int
	Threads_cached    int
	//mysql --bytes
	Bytes_received int
	Bytes_sent     int
	//mysql --innodb_status show engine innodb status
	//log unflushed = Log sequence number - Log flushed up to
	//uncheckpointed bytes = Log sequence number - Last checkpoint at
	//mysql -e "show engine innodb status\G"|grep -n -E -A4 -B1 "^TRANSACTIONS|LOG|ROW OPERATIONS"
	//mysql -e "show engine innodb status\G"|grep -E "Last checkpoint|read view|queries inside|queue"
	Log_sequence    int
	Log_flushed     int
	History_list    int
	Last_checkpoint int
	Read_view       int
	Query_inside    int
	Query_queue     int
	//addition
	//show status
	Max_used_connections  int
	Aborted_connects      string
	Aborted_clients       string
	Select_full_join      string
	Binlog_cache_disk_use string
	Binlog_cache_use      string
	Opened_tables         string
	//Thread_cache_hits = (1 - Threads_created / connections ) * 100%
	Connections             int
	Qcache_hits             int
	Handler_read_first      int
	Handler_read_key        int
	Handler_read_next       int
	Handler_read_prev       int
	Handler_read_rnd        int
	Handler_read_rnd_next   int
	Handler_rollback        int
	Created_tmp_tables      int
	Created_tmp_disk_tables int
	Slow_queries            string
	Key_read_requests       int
	Key_reads               int
	Key_write_requests      int
	Key_writes              int
	Select_scan             string
	//半同步
	Rpl_semi_sync_master_net_avg_wait_time int
	Rpl_semi_sync_master_no_times          int
	Rpl_semi_sync_master_no_tx             int
	Rpl_semi_sync_master_status            string
	Rpl_semi_sync_master_tx_avg_wait_time  int
	Rpl_semi_sync_master_wait_sessions     int
	Rpl_semi_sync_master_yes_tx            int
	Rpl_semi_sync_slave_status             string
	rpl_semi_sync_master_timeout           string
	Rpl_semi_sync_master_clients		int
	Rpl_semi_sync_master_net_wait_time	int
	Rpl_semi_sync_master_tx_wait_time	int
	Rpl_semi_sync_master_net_waits		int
	Rpl_semi_sync_master_tx_waits		int
	//Slave状态监控
	Master_Host           string
	Master_User           string
	Master_Port           string
	Slave_IO_Running      string
	Slave_SQL_Running     string
	Master_Server_Id      string
	Seconds_Behind_Master int
	Read_Master_Log_Pos   int
	Exec_Master_Log_Pos   int
}

type flags struct {
	ununit		bool   //Print Unit for Byte
        ext             bool   //Print Extent info
	myid            string //MySQL Instance name, Default: m3361
	interval        string //Interval, Default: 1 second(s)
	count           int    //Running Time, Default: never stop
	time            bool   //Print current Time
        dtime           bool   //Print current Date Time
	nocolor         bool   //Has no color
	load            bool   //Print Load  info
	cpu             bool   //Print Cpu   info
	swap            bool   //Print Swap  info
	disk            string //Print Disk  info
	net             string //Print Net   info
	slave           bool   //Print Slave info
	username        string //MySQL Username
	password        string //MySQL Password
	host            string //MySQL Host, Default: 127.0.0.1
	port            string //MySQL Port, Default:3306
	socket          string //MySQL Socket
	com             bool   //Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).
	hit             bool   //Print Innodb Hit%.
	innodb_rows     bool   //Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).
	innodb_pages    bool   //Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)
	innodb_data     bool   //Print Innodb Data Status(Innodb_data_reads/writes/read/written)
	innodb_log      bool   //Print Innodb Log  Status(Innodb_os_log_fsyncs/written)
	innodb_status   bool   //Print Innodb Status from Command: 'Show Engine Innodb Status'
	//(history list/ log unflushed/uncheckpointed bytes/ read views/ queries inside/queued)
	threads         bool   //Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).
	rt              bool   //Print MySQL DB RT(us).
	bytes           bool   //Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).

	mysql           bool   //Print MySQLInfo (include -t,-com,-hit,-T,-B).
	innodb          bool   //Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)
	sys             bool   //Print SysInfo   (include -t,-l,-c,-s).
	lazy            bool   //Print Info      (include -t,-l,-c,-s,-com,-hit).

	logfile         string //Print to Logfile.
	logfile_by_day  bool   //One day a logfile,the suffix of logfile is 'yyyy-mm-dd';
	semi            bool   //MySQL Semi-Sync
	other           []string
	//and is valid with -L.
}

func (e *flags) init() {
	ununit	 := flag.Bool("un", false, "Print info by NO Unit")
        ext      := flag.Bool("ext", false, "Print Extend info")
	myid     := flag.String("myid", "MySQL", "MySQL instance, Default: m3361")
	interval := flag.String("i", "1", "Interval, Default: 1 second(s)")
	count    := flag.Int("C", 0, "Running Time, Default: never stop")
	time     := flag.Bool("t", false, "Print current Time")
	dtime    := flag.Bool("dt", false, "Print current Date Time")
	nocolor  := flag.Bool("nocolor", false, "Has no color")
	load     := flag.Bool("l", false, "Print Load info")
	cpu      := flag.Bool("c", false, "Print Cpu info")
	swap     := flag.Bool("s", false, "Print Swap info")
	disk     := flag.String("d", "none", "Print Disk info")
	net      := flag.String("n", "none", "Print net info")
	slave    := flag.Bool("slave", false, "Print Slave info")
	username := flag.String("u", "sys_root", "MySQL Username")
	password := flag.String("p", "dbaops.com", "MySQL Password")
	host     := flag.String("H", "127.0.0.1", "MySQL Host, Default: 127.0.0.1")
	port     := flag.String("P", "3306", "MySQL Port, Default:3306")
	socket   := flag.String("S", "/tmp/mysql.sock", "MySQL Socket")
	com      := flag.Bool("com", false, "Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete).")
	hit      := flag.Bool("hit", false, "Print Innodb Hit%.")
	innodb_rows   := flag.Bool("innodb_rows",   false, "Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read).")
	innodb_pages  := flag.Bool("innodb_pages",  false, "Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)")
	innodb_data   := flag.Bool("innodb_data",   false, "Print Innodb Data Status(Innodb_data_reads/writes/read/written)")
	innodb_log    := flag.Bool("innodb_log",    false, "Print Innodb Log  Status(Innodb_os_log_fsyncs/written)")
	innodb_status := flag.Bool("innodb_status", false, "Print Innodb Status from Command: 'Show Engine Innodb Status'")
	threads := flag.Bool("T",      false, "Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached).")
	rt      := flag.Bool("rt",     false, "Print MySQL DB RT(us).")
	bytes   := flag.Bool("B",      false, "Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent).")
	mysql   := flag.Bool("mysql",  false, "Print MySQLInfo (include -t,-com,-hit,-T,-B).")
	innodb  := flag.Bool("innodb", false, "Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)")
	sys     := flag.Bool("sys",    false, "Print SysInfo   (include -t,-l,-c,-s).")
	lazy    := flag.Bool("lazy",   false, "Print Info  (include -t,-l,-c,-s,-com,-hit).")
	semi    := flag.Bool("semi",   false, "半同步监控")
	logfile := flag.String("L", "none", "Print to Logfile.")
	logfile_by_day := flag.Bool("logfile_by_day", false, "One day a logfile,the suffix of logfile is 'yyyy-mm-dd';")

	flag.Parse()
	gunit	    = *ununit
	e.ununit    = *ununit
	e.ext       = *ext
	e.myid      = *myid
	e.interval  = *interval
	e.count     = *count
	e.time      = *time
	e.dtime     = *dtime
	e.nocolor   = *nocolor
	e.load      = *load
	e.cpu       = *cpu
	e.swap      = *swap
	e.disk      = *disk
	e.net       = *net
	e.slave     = *slave
	e.username  = *username
	e.password  = *password
	e.host      = *host
	e.port      = *port
	e.socket    = *socket
	e.com       = *com
	e.hit       = *hit
	e.innodb_rows   = *innodb_rows
	e.innodb_pages  = *innodb_pages
	e.innodb_data   = *innodb_data
	e.innodb_log    = *innodb_log
	e.innodb_status = *innodb_status
	e.threads    = *threads
	e.rt         = *rt
	e.bytes      = *bytes
	e.mysql      = *mysql
	e.innodb     = *innodb
	e.sys        = *sys
	e.lazy       = *lazy
	e.logfile     = *logfile
	e.logfile_by_day = *logfile_by_day
	e.semi       = *semi
	e.other      = flag.Args()

	if flag.NFlag() == 0 {
		fmt.Println("请输入【-h】查看帮助！\n Sample :shell> nohup ./orzdba -lazy -d sda -C 5 -i 2 -L /tmp/orzdba.log  > /dev/null 2>&1 &")
		os.Exit(1)
	}

	if flag.NArg() != 0 {
		fmt.Println("无FLAG区域预留给MySQL:远端数据接收服务器【USER:PASSWORD@IP:PORT/DBNAME】")
	}
}

func GetValue() map[string]interface{} {
	u := flags{}
	u.init()
	info := map[string]interface{}{
		"ununit":	  u.ununit,
		"ext":            u.ext,
		"myid":           u.myid,
		"interval":       u.interval,
		"count":          u.count,
		"time":           u.time,
		"dtime":          u.dtime,
		"nocolor":        u.nocolor,
		"load":           u.load,
		"cpu":            u.cpu,
		"swap":           u.swap,
		"disk":           u.disk,
		"net":            u.net,
		"slave":          u.slave,
		"username":       u.username,
		"password":       u.password,
		"host":           u.host,
		"port":           u.port,
		"socket":         u.socket,
		"com":            u.com,
		"hit":            u.hit,
		"innodb_rows":    u.innodb_rows,
		"innodb_pages":   u.innodb_pages,
		"innodb_data":    u.innodb_data,
		"innodb_log":     u.innodb_log,
		"innodb_status":  u.innodb_status,
		"threads":        u.threads,
		"rt":             u.rt,
		"bytes":          u.bytes,
		"mysql":          u.mysql,
		"innodb":         u.innodb,
		"sys":            u.sys,
		"semi":           u.semi,
		"lazy":           u.lazy,
		"logfile":        u.logfile,
		"logfile_by_day": u.logfile_by_day,
		"other":          u.other,
	}
	return info
}

func checkErr(errinfo error) {
	if errinfo != nil {
		fmt.Println(errinfo.Error())
		// panic(errinfo.Error())
	}
}

func execCommand(commands string) string {
	// basic := "for x in {1};do load=`top -bn 1|sed -n '1p'|awk '{print $12,$13,$14}'|sed 's/[[:space:]]//g'`;cpu=`top -bn 1|sed -n '3p'|awk '{print $2,$4,$6,$8,$10,$12,$14,$16}'|sed 's/[[:space:]]/,/g'`;echo $load,$cpu;done"
	// cmd := exec.Command("bash", "-c", commands)
	// stdout, err := cmd.StdoutPipe()
	// checkErr(err)
	//fmt.Println(commands)
	out, err := exec.Command("bash", "-c", commands).Output()
	checkErr(err)

	// bytesErr, err := ioutil.ReadAll(out)
	// checkErr(err)
	return string(out)
}

func createCommand(info map[string]interface{}, count int) basic {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\033[1;4;31m数据获取异常,请检查输入参数:%s\033[0m\n", err)
			os.Exit(1)
		}
	}()
	ss := basic{}
	myConn := "mysql -u" + info["username"].(string) + " -p" + info["password"].(string) + " --host=" + info["host"].(string) + " --socket=" + info["socket"].(string) + " --port=" + info["port"].(string)
	if count == 0 {
		host_cmd := "hostname"
		host_string := execCommand(host_cmd)
		// host_string = strings.Replace(host_string, "\n", "", -1)
		// host_result := strings.Split(rt_string, ",")
		ss.hostname = host_string

		ip_cmd := "ip a|grep global|grep -v lo:|head -1|awk '{print $2}'|cut -d '/' -f1"
		ip_string := execCommand(ip_cmd)
		// ip_string = strings.Replace(ip_string, "\n", "", -1)
		ss.ip = ip_string

		db_cmd := myConn + " -e 'show databases;'|xargs echo|sed 's/ /|/g'"
		db_string := execCommand(db_cmd)
		db_string = strings.Replace(db_string, "\n", "", -1)
		ss.db = db_string

		variables_cmd := myConn + " -e 'show global variables'|grep -E -w 'binlog_format|max_binlog_cache_size|max_binlog_size|max_connect_errors|max_connections|max_user_connections|open_files_limit|sync_binlog|table_definition_cache|table_open_cache|thread_cache_size|innodb_adaptive_flushing|innodb_adaptive_hash_index|innodb_buffer_pool_size|innodb_file_per_table|innodb_flush_log_at_trx_commit|innodb_io_capacity|innodb_lock_wait_timeout|innodb_log_buffer_size|innodb_log_file_size|innodb_log_files_in_group|innodb_max_dirty_pages_pct|innodb_open_files|innodb_read_io_threads|innodb_thread_concurrency|innodb_write_io_threads'|awk '{print $2}'|xargs echo|sed 's/ /,/g'"
		variables_string := execCommand(variables_cmd)
		variables_string  = strings.Replace(variables_string, "\n", "", -1)
		variables_result := strings.Split(variables_string, ",")

		ss.var_binlog_format                  = variables_result[0]
		ss.var_innodb_adaptive_flushing       = variables_result[1]
		ss.var_innodb_adaptive_hash_index     = variables_result[2]
		ss.var_innodb_buffer_pool_size, _     = strconv.Atoi(variables_result[3])
		ss.var_innodb_file_per_table          = variables_result[4]
		ss.var_innodb_flush_log_at_trx_commit = variables_result[5]
		ss.var_innodb_io_capacity             = variables_result[6]
		ss.var_innodb_lock_wait_timeout       = variables_result[7]
		ss.var_innodb_log_buffer_size, _      = strconv.Atoi(variables_result[8])
		ss.var_innodb_log_file_size, _        = strconv.Atoi(variables_result[9])
		ss.var_innodb_log_files_in_group      = variables_result[10]
		ss.var_innodb_max_dirty_pages_pct     = variables_result[11]
		ss.var_innodb_open_files              = variables_result[12]
		ss.var_innodb_read_io_threads         = variables_result[13]
		ss.var_innodb_thread_concurrency      = variables_result[14]
		ss.var_innodb_write_io_threads        = variables_result[15]
		ss.var_max_binlog_cache_size, _       = strconv.Atoi(variables_result[16])
		ss.var_max_binlog_size, _             = strconv.Atoi(variables_result[17])
		ss.var_max_connect_errors             = variables_result[18]
		ss.var_max_connections                = variables_result[19]
		ss.var_max_user_connections           = variables_result[20]
		ss.var_open_files_limit               = variables_result[21]
		ss.var_sync_binlog                    = variables_result[22]
		ss.var_table_definition_cache         = variables_result[23]
		ss.var_table_open_cache               = variables_result[24]
		ss.var_thread_cache_size              = variables_result[25]

		innodb_flush_method_cmd := myConn + " -e 'show variables'|grep -E -w 'innodb_flush_method'|awk '{print $2}'"
		innodb_flush_method_string := execCommand(innodb_flush_method_cmd)
		innodb_flush_method_string = strings.Replace(innodb_flush_method_string, "\n", "", -1)

		ss.var_innodb_flush_method = innodb_flush_method_string

		//semi
		semi_tmp := myConn + " -e 'show variables'|grep -E -w 'rpl_semi_sync_master_timeout'|awk '{print $2}'"
		semi_string := execCommand(semi_tmp)
		semi_string = strings.Replace(semi_string, "\n", "", -1)

		ss.rpl_semi_sync_master_timeout = semi_string

		// //mysql global status
		innodb_cmd := myConn + " -e 'show global status'|grep -w -E 'Max_used_connections|Aborted_connects|Aborted_clients|Select_full_join|Binlog_cache_disk_use|Binlog_cache_use|Opened_tables|Connections|Qcache_hits|Handler_read_first|Handler_read_key|Handler_read_next|Handler_read_prev|Handler_read_rnd|Handler_read_rnd_next|Handler_rollback|Created_tmp_tables|Created_tmp_disk_tables|Slow_queries|Key_read_requests|Key_reads|Key_write_requests|Key_writes|Select_scan|Rpl_semi_sync_master_status|Rpl_semi_sync_slave_status'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"

		innodb_string := execCommand(innodb_cmd)
		innodb_string  = strings.Replace(innodb_string, "\n", "", -1)
		innodb_result := strings.Split(innodb_string, ",")

		lens := len(innodb_result)

		ss.Aborted_clients           = innodb_result[0]
		ss.Aborted_connects          = innodb_result[1]
		ss.Binlog_cache_disk_use     = innodb_result[2]
		ss.Binlog_cache_use          = innodb_result[3]
		ss.Connections,            _ = strconv.Atoi(innodb_result[4])
		ss.Created_tmp_disk_tables,_ = strconv.Atoi(innodb_result[5])
		ss.Created_tmp_tables,     _ = strconv.Atoi(innodb_result[6])
		ss.Handler_read_first,     _ = strconv.Atoi(innodb_result[7])
		ss.Handler_read_key,       _ = strconv.Atoi(innodb_result[8])
		ss.Handler_read_next,      _ = strconv.Atoi(innodb_result[9])
		ss.Handler_read_prev,      _ = strconv.Atoi(innodb_result[10])
		ss.Handler_read_rnd,       _ = strconv.Atoi(innodb_result[11])
		ss.Handler_read_rnd_next,  _ = strconv.Atoi(innodb_result[12])
		ss.Handler_rollback,       _ = strconv.Atoi(innodb_result[13])
		ss.Key_read_requests,      _ = strconv.Atoi(innodb_result[14])
		ss.Key_reads,              _ = strconv.Atoi(innodb_result[15])
		ss.Key_write_requests,     _ = strconv.Atoi(innodb_result[16])
		ss.Key_writes,             _ = strconv.Atoi(innodb_result[17])
		ss.Max_used_connections,   _ = strconv.Atoi(innodb_result[18])
		ss.Opened_tables             = innodb_result[19]
		ss.Qcache_hits,            _ = strconv.Atoi(innodb_result[20])
		if lens == 26 {
			ss.Rpl_semi_sync_master_status = innodb_result[21]
			ss.Rpl_semi_sync_slave_status  = innodb_result[22]
			ss.Select_full_join            = innodb_result[23]
			ss.Select_scan                 = innodb_result[24]
			ss.Slow_queries                = innodb_result[25]
		} else {
			ss.Select_full_join            = innodb_result[21]
			ss.Select_scan                 = innodb_result[22]
			ss.Slow_queries                = innodb_result[23]
		}

		slave_cmd := myConn + " -e 'show slave status\\G'|grep -E -w 'Master_Host|Master_User|Master_Port|Slave_IO_Running|Slave_SQL_Running|Seconds_Behind_Master|Master_Server_Id|Read_Master_Log_Pos|Exec_Master_Log_Pos'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"
		slave_string := execCommand(slave_cmd)
		slave_string = strings.Replace(slave_string, "\n", "", -1)
		slave_result := strings.Split(slave_string, ",")
		if slave_result[0] == "" {
			ss.Master_Host = ""
		} else {
			ss.Master_Host       = slave_result[0]
			ss.Master_User       = slave_result[1]
			ss.Master_Port       = slave_result[2]
			ss.Slave_IO_Running  = slave_result[4]
			ss.Slave_SQL_Running = slave_result[5]
			ss.Master_Server_Id  = slave_result[8]
		}
		// fmt.Println(semi_cmd)
		// fmt.Println(semi_result)

	}

	//rt
	if info["rt"] == true {
		if count == 0 {
			ss.rt_count = 0
			ss.rt_avg = 0
			ss.rt_a5 = 0
			ss.rt_a9 = 0
		} else {
			var rt_cmd string
			rt_cmd = "tail -1 /tmp/orzdba_tcprstat.log |awk '{print $2,$5,$9,$12}'|sed 's/[[:space:]]/,/g'"

			rt_string := execCommand(rt_cmd)
			// fmt.Println(rt_string)
			rt_string = strings.Replace(rt_string, "\n", "", -1)
			rt_result := strings.Split(rt_string, ",")

			ss.rt_count, _ = strconv.Atoi(rt_result[0])
			ss.rt_avg, _ = strconv.Atoi(rt_result[1])
			ss.rt_a5, _ = strconv.Atoi(rt_result[2])
			ss.rt_a9, _ = strconv.Atoi(rt_result[3])
		}
	}

	// swap  bool   //打印swap info
	if info["swap"] == true || info["load"] == true || info["cpu"] == true {
		basic_cmd := "cat /proc/loadavg /proc/stat /proc/vmstat |sed 's/\\// china /g'|grep -w -E 'china|cpu|pswpin|pswpout'|xargs echo|awk '{print $1,$2,$3,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$20,$22}'|sed 's/[[:space:]]/,/g'"
		basic_string := execCommand(basic_cmd)
		basic_string = strings.Replace(basic_string, "\n", "", -1)
		basic_result := strings.Split(basic_string, ",")

		ss.load_1, _ = strconv.ParseFloat(basic_result[0], 64)
		ss.load_5, _ = strconv.ParseFloat(basic_result[1], 64)
		ss.load_15, _ = strconv.ParseFloat(basic_result[2], 64)
		ss.cpu_usr, _ = strconv.Atoi(basic_result[3])
		ss.cpu_nice, _ = strconv.Atoi(basic_result[4])
		ss.cpu_sys, _ = strconv.Atoi(basic_result[5])
		ss.cpu_idl, _ = strconv.Atoi(basic_result[6])
		ss.cpu_iow, _ = strconv.Atoi(basic_result[7])
		ss.cpu_irq, _ = strconv.Atoi(basic_result[8])
		ss.cpu_softirq, _ = strconv.Atoi(basic_result[9])
		ss.cpu_steal, _ = strconv.Atoi(basic_result[10])
		ss.cpu_guest, _ = strconv.Atoi(basic_result[11])
		ss.cpu_guest_nice, _ = strconv.Atoi(basic_result[12])
		ss.swap_in, _ = strconv.Atoi(basic_result[13])
		ss.swap_out, _ = strconv.Atoi(basic_result[14])
	}

	//disk
	if info["disk"] != "none" {
		disk_cmd := "cat /proc/diskstats |grep -w -E '" + info["disk"].(string) + "'|awk '{print $4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14}'|sed 's/[[:space:]]/,/g'"
		// fmt.Println(disk_cmd)
		disk_string := execCommand(disk_cmd)
		disk_string = strings.Replace(disk_string, "\n", "", -1)
		// fmt.Println(disk_string)
		disk_result := strings.Split(disk_string, ",")
		// fmt.Println(disk_result)
		ss.io_1, _ = strconv.Atoi(disk_result[0])
		ss.io_2, _ = strconv.Atoi(disk_result[1])
		ss.io_3, _ = strconv.Atoi(disk_result[2])
		ss.io_4, _ = strconv.Atoi(disk_result[3])
		ss.io_5, _ = strconv.Atoi(disk_result[4])
		ss.io_6, _ = strconv.Atoi(disk_result[5])
		ss.io_7, _ = strconv.Atoi(disk_result[6])
		ss.io_8, _ = strconv.Atoi(disk_result[7])
		ss.io_9, _ = strconv.Atoi(disk_result[8])
		ss.io_10, _ = strconv.Atoi(disk_result[9])
		ss.io_11, _ = strconv.Atoi(disk_result[10])
		// fmt.Println("is ok?")
	}

	//net
	if info["net"] != "none" {
		//坑 centos和debian的 /proc/net/dev的文件格式不一样 多了一个空格
		net_cmd := "ifconfig " + info["net"].(string) + "|grep bytes|sed 's/:/ /g'|awk '{print $3,$8}'|sed 's/ /,/g'"
		// net_cmd := "cat /proc/net/dev |grep -E -w '" + info["net"].(string) + "'|awk '{print $2*8,$10*8}'|sed 's/[[:space:]]/,/g'"
		// fmt.Println(net_cmd)
		net_string := execCommand(net_cmd)
		net_string = strings.Replace(net_string, "\n", "", -1)
		net_result := strings.Split(net_string, ",")

		ss.net_recv, _ = strconv.Atoi(net_result[0])
		ss.net_send, _ = strconv.Atoi(net_result[1])
	}

	// //mysql engine innodb status
	if info["innodb_status"] == true {
		engine_cmd := myConn + " -e 'show engine innodb status\\G'|grep -w -E 'History list|Log sequence|Log flushed|queries inside|queue|read views|Last checkpoint'|xargs echo|awk '{print $4,$8,$13,$17,$18,$22,$26}'|sed 's/[[:space:]]/,/g'"
		engine_string := execCommand(engine_cmd)
		engine_string = strings.Replace(engine_string, "\n", "", -1)
		engine_result := strings.Split(engine_string, ",")

		ss.History_list, _ = strconv.Atoi(engine_result[0])
		ss.Log_sequence, _ = strconv.Atoi(engine_result[1])
		ss.Log_flushed, _ = strconv.Atoi(engine_result[2])
		ss.Last_checkpoint, _ = strconv.Atoi(engine_result[3])
		ss.Query_inside, _ = strconv.Atoi(engine_result[4])
		ss.Query_queue, _ = strconv.Atoi(engine_result[5])
		ss.Read_view, _ = strconv.Atoi(engine_result[6])
	}
	// //mysql global status
	if info["com"] == true || info["hit"] == true || info["innodb_rows"] == true || info["innodb_pages"] == true || info["innodb_data"] == true || info["innodb_log"] == true || info["threads"] == true || info["bytes"] == true {
		global_cmd := myConn + " -e 'show global status'|grep -w -E 'Com_select|Com_insert|Com_update|Com_delete|Com_commit|Com_rollback|Innodb_buffer_pool_read_requests|Innodb_buffer_pool_reads|Innodb_rows_inserted|Innodb_rows_updated|Innodb_rows_deleted|Innodb_rows_read|Innodb_buffer_pool_pages_data|Innodb_buffer_pool_pages_free|Innodb_buffer_pool_pages_dirty|Innodb_buffer_pool_pages_flushed|Innodb_data_reads|Innodb_data_writes|Innodb_data_read|Innodb_data_written|Innodb_os_log_fsyncs|Innodb_os_log_written|Threads_running|Threads_connected|Threads_created|Threads_cached|Bytes_received|Bytes_sent|Max_used_connections|Aborted_connects|Aborted_clients|Select_full_join|Binlog_cache_disk_use|Binlog_cache_use|Opened_tables|Connections|Qcache_hits|Handler_read_first|Handler_read_key|Handler_read_next|Handler_read_prev|Handler_read_rnd|Handler_read_rnd_next|Handler_rollback|Created_tmp_tables|Created_tmp_disk_tables|Slow_queries|Key_read_requests|Key_reads|Key_write_requests|Key_writes'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"

		global_string := execCommand(global_cmd)
		global_string = strings.Replace(global_string, "\n", "", -1)
		global_result := strings.Split(global_string, ",")

		// ss.Aborted_clients = global_result[0]
		// ss.Aborted_connects = global_result[1]
		// ss.Binlog_cache_disk_use  = global_result[2]
		// ss.Binlog_cache_use, _ = strconv.Atoi(global_result[3])
		ss.Bytes_received, _ = strconv.Atoi(global_result[4])
		ss.Bytes_sent, _ = strconv.Atoi(global_result[5])
		ss.Com_commit, _ = strconv.Atoi(global_result[6])
		ss.Com_delete, _ = strconv.Atoi(global_result[7])
		ss.Com_insert, _ = strconv.Atoi(global_result[8])
		ss.Com_rollback, _ = strconv.Atoi(global_result[9])
		ss.Com_select, _ = strconv.Atoi(global_result[10])
		ss.Com_update, _ = strconv.Atoi(global_result[11])
		ss.Connections, _ = strconv.Atoi(global_result[12])
		ss.Created_tmp_disk_tables, _ = strconv.Atoi(global_result[13])
		ss.Created_tmp_tables, _ = strconv.Atoi(global_result[14])
		ss.Handler_read_first, _ = strconv.Atoi(global_result[15])
		ss.Handler_read_key, _ = strconv.Atoi(global_result[16])
		ss.Handler_read_next, _ = strconv.Atoi(global_result[17])
		ss.Handler_read_prev, _ = strconv.Atoi(global_result[18])
		ss.Handler_read_rnd, _ = strconv.Atoi(global_result[19])
		ss.Handler_read_rnd_next, _ = strconv.Atoi(global_result[20])
		ss.Handler_rollback, _ = strconv.Atoi(global_result[21])
		ss.Innodb_buffer_pool_pages_data, _ = strconv.Atoi(global_result[22])
		ss.Innodb_buffer_pool_pages_dirty, _ = strconv.Atoi(global_result[23])
		ss.Innodb_buffer_pool_pages_flushed, _ = strconv.Atoi(global_result[24])
		ss.Innodb_buffer_pool_pages_free, _ = strconv.Atoi(global_result[25])
		ss.Innodb_buffer_pool_read_requests, _ = strconv.Atoi(global_result[26])
		ss.Innodb_buffer_pool_reads, _ = strconv.Atoi(global_result[27])
		ss.Innodb_data_read, _ = strconv.Atoi(global_result[28])
		ss.Innodb_data_reads, _ = strconv.Atoi(global_result[29])
		ss.Innodb_data_writes, _ = strconv.Atoi(global_result[30])
		ss.Innodb_data_written, _ = strconv.Atoi(global_result[31])
		ss.Innodb_os_log_fsyncs, _ = strconv.Atoi(global_result[32])
		ss.Innodb_os_log_written, _ = strconv.Atoi(global_result[33])
		ss.Innodb_rows_deleted, _ = strconv.Atoi(global_result[34])
		ss.Innodb_rows_inserted, _ = strconv.Atoi(global_result[35])
		ss.Innodb_rows_read, _ = strconv.Atoi(global_result[36])
		ss.Innodb_rows_updated, _ = strconv.Atoi(global_result[37])
		ss.Key_read_requests, _ = strconv.Atoi(global_result[38])
		ss.Key_reads, _ = strconv.Atoi(global_result[39])
		ss.Key_write_requests, _ = strconv.Atoi(global_result[40])
		ss.Key_writes, _ = strconv.Atoi(global_result[41])
		ss.Max_used_connections, _ = strconv.Atoi(global_result[42])
		// ss.Opened_tables, _ = strconv.Atoi(global_result[43])
		ss.Qcache_hits, _ = strconv.Atoi(global_result[44])
		// ss.Select_full_join, _ = strconv.Atoi(global_result[45])
		ss.Slow_queries = global_result[46]
		ss.Threads_cached, _ = strconv.Atoi(global_result[47])
		ss.Threads_connected, _ = strconv.Atoi(global_result[48])
		ss.Threads_created, _ = strconv.Atoi(global_result[49])
		ss.Threads_running, _ = strconv.Atoi(global_result[50])

	}

	// //mysql engine innodb status
	if info["semi"] == true {
		semi_cmd := myConn + " -e 'show status'|grep -E Rpl_semi|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"
		semi_string := execCommand(semi_cmd)
		semi_string = strings.Replace(semi_string, "\n", "", -1)
		semi_result := strings.Split(semi_string, ",")
		if semi_result[0] == "" {
			fmt.Println(Colorize("semi半同步未开启", red, "", "", "y"))
			os.Exit(1)
		}
		// fmt.Println(semi_cmd)
		// fmt.Println(semi_result)
		ss.Rpl_semi_sync_master_clients, _ =strconv.Atoi(semi_result[0])
		ss.Rpl_semi_sync_master_net_avg_wait_time, _ = strconv.Atoi(semi_result[1])
		ss.Rpl_semi_sync_master_net_wait_time, _ = strconv.Atoi(semi_result[2])
		ss.Rpl_semi_sync_master_net_waits, _ = strconv.Atoi(semi_result[3])
		ss.Rpl_semi_sync_master_no_times, _ = strconv.Atoi(semi_result[4])
		ss.Rpl_semi_sync_master_no_tx, _ = strconv.Atoi(semi_result[5])
		ss.Rpl_semi_sync_master_status = semi_result[6]
		ss.Rpl_semi_sync_master_tx_avg_wait_time, _ = strconv.Atoi(semi_result[8])
		ss.Rpl_semi_sync_master_tx_wait_time, _ = strconv.Atoi(semi_result[9])
		ss.Rpl_semi_sync_master_tx_waits, _ = strconv.Atoi(semi_result[10])
		ss.Rpl_semi_sync_master_wait_sessions, _ = strconv.Atoi(semi_result[12])
		ss.Rpl_semi_sync_master_yes_tx, _ = strconv.Atoi(semi_result[13])
		ss.Rpl_semi_sync_slave_status = semi_result[14]
	}

	// slave status
	if info["slave"] == true {
		slave_cmd := myConn + " -e 'show slave status\\G'|grep -E -w 'Master_Host|Master_User|Master_Port|Slave_IO_Running|Slave_SQL_Running|Seconds_Behind_Master|Master_Server_Id|Read_Master_Log_Pos|Exec_Master_Log_Pos'|awk '{print $2}'|xargs echo|sed 's/[[:space:]]/,/g'"
		slave_string := execCommand(slave_cmd)
		slave_string = strings.Replace(slave_string, "\n", "", -1)
		slave_result := strings.Split(slave_string, ",")
		if slave_result[0] == "" {
			// fmt.Println(Colorize("IsNotSlave", red, "", "", "y"))
			// os.Exit(1)
                        slave_result = strings.Split("xxx.xxx.xxx.xxx,replica,3306,-1,  N,  N,-1,-1,-1", ",")
		}
		// fmt.Println(semi_cmd)
		// fmt.Println(semi_result)
		// ss.Master_Host = slave_result[0]
		// ss.Master_User = slave_result[1]
		// ss.Master_Port = slave_result[2]
		ss.Read_Master_Log_Pos, _ = strconv.Atoi(slave_result[3])
		ss.Slave_IO_Running = slave_result[4]
		ss.Slave_SQL_Running = slave_result[5]
		ss.Exec_Master_Log_Pos, _ = strconv.Atoi(slave_result[6])
		ss.Seconds_Behind_Master, _ = strconv.Atoi(slave_result[7])
		// ss.Master_Server_Id = slave_result[8]
	}

	return ss
}

// 文字字体 参数介绍：文本内容 文字颜色 背景颜色 是否下划线 是否高亮
// http://www.cnblogs.com/frydsh/p/4139922.html
func Colorize(text string, status string, background string, underline string, highshow string) string {
	out_one := "\033["
	out_two := ""
	out_three := ""
	out_four := ""
	// 可动态配置字体颜色 背景色 高亮
	// 显示：0(默认)、1(粗体/高亮)、22(非粗体)、4(单条下划线)、24(无下划线)、5(闪烁)、25(无闪烁)、7(反显、翻转前景色和背景色)、27(无反显)
	// 颜色：0(黑)、1(红)、2(绿)、 3(黄)、4(蓝)、5(洋红)、6(青)、7(白)
	// 前景色为30+颜色值，如31表示前景色为红色；背景色为40+颜色值，如41表示背景色为红色。
	if underline == "y" && highshow == "y" {
		out_four = ";1;4m" //高亮
	} else if underline != "y" && highshow == "y" {
		out_four = ";1m"
	} else if underline == "y" && highshow != "y" {
		out_four = ";4m"
	} else {
		out_four = ";22m"
	}

	switch status {
	case "black":
		out_two = "30"
	case "red":
		out_two = "31"
	case "green":
		out_two = "32"
	case "yellow":
		out_two = "33"
	case "blue":
		out_two = "34"
	case "purple":
		out_two = "35"
	case "dgreen":
		out_two = "36"
	case "white":
		out_two = "37"
	default:
		out_two = ""
	}

	switch background {
	case "black":
		out_three = "40;"
	case "red":
		out_three = "41;"
	case "green":
		out_three = "42;"
	case "yellow":
		out_three = "43;"
	case "blue":
		out_three = "44;"
	case "purple":
		out_three = "45;"
	case "dgreen":
		out_three = "46;"
	case "white":
		out_three = "47;"
	default:
		out_three = ""
	}
	return out_one + out_three + out_two + out_four + text + "\033[0m"
}

func perSecond_Int(before int, after int, time string) (string, bool) {
	var result interface{}
	var ok bool
	var rs string
	//转换时间为int
	seconds, err := strconv.Atoi(time)
	checkErr(err)
	result = (after - before) / seconds
	// fmt.Println(result.(int))
	// tmp = fmt.Sprintf("%s", reflect.TypeOf(result))
	switch result.(type) {
	case int:
		// fmt.Println("int")
		rs = strconv.Itoa(result.(int))
	case int64:
		fmt.Println("int64")
		rs = strconv.FormatInt(result.(int64), 64)
	case float64:
		fmt.Println("float64")
		rs = strconv.FormatFloat(result.(float64), 'E', 5, 64)
	default:
		panic("not fount number type in perSecond")
	}
	if result.(int) > 0 {
		ok = true
	} else {
		ok = false
	}

	return rs, ok
}

func floatToString(x float64, f int) string {
	rs := strconv.FormatFloat(x, 'f', f, 64)
	return rs
}

func perSecond_Float(before float64, after float64, time string) (string, bool) {
	var result interface{}
	var ok bool
	var rs string
	//转换时间为float64
	seconds, err := strconv.ParseFloat(time, 64)
	checkErr(err)
	result = (after - before) / seconds
	// tmp = fmt.Sprintf("%s", reflect.TypeOf(result))
	switch result.(type) {
	case int:
		// fmt.Println("int")
		rs = strconv.Itoa(result.(int))
	case int64:
		fmt.Println("int64")
		rs = strconv.FormatInt(result.(int64), 64)
	// case float32:
	//  fmt.Println("float32")
	//  rs = strconv.FormatFloat(result.(float32), 'f', 4, 32)
	case float64:
		fmt.Println("float64")
		rs = strconv.FormatFloat(result.(float64), 'f', 4, 64)
	default:
		panic("not fount number type in perSecond_Float")
	}

	if result.(int) > 0 {
		ok = true
	} else {
		ok = false
	}

	return rs, ok
}
func getColor(in int, l1 int, l2 int) string {
	var result string
	if in > 0 && in < l1 {
		result = green
	} else if in >= l1 && in < l2 {
		result = yellow
	} else if in >= l2 {
		result = red
	} else {
		result = ""
	}
	return result
}
func changeUnits(in int) string {
	var result string
	if gunit == true {
		result = strconv.Itoa(in)
	} else {
		if in/1024 < 1 {
			tmp := strconv.Itoa(in)
			result = tmp
		} else if in/1024 >= 1 && in/1024/1024 < 1 {
			tmp := strconv.Itoa(in / 1024)
			result = tmp + "k"
		} else if in/1024/1024 >= 1 && in/1024/1024/1024 < 1 {
			tmp := strconv.Itoa(in / 1024 / 1024)
			result = tmp + "m"
		} else if in/1024/1024/1024 >= 1 && in/1024/1024/1024/1024 < 1 {
			tmp := strconv.Itoa(in / 1024 / 1024 / 1024)
			result = tmp + "g"
		} else if in/1024/1024/1024/1024 >= 1 {
			tmp := strconv.Itoa(in / 1024 / 1024 / 1024 / 1024)
			result = tmp + "t"
		}
	}
	return result
}

func getNowTime() string {
	f := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
	timeformatdate, _ := time.Parse("2006-01-02 15:04:05", f)
	convtime := fmt.Sprintf("%s", timeformatdate.Format("15:04:05"))
	return convtime
}

func getNowDate() string {
	f := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
	timeformatdate, _ := time.Parse("2006-01-02 15:04:05", f)
	convtime := fmt.Sprintf("%s", timeformatdate.Format("2006-01-02"))
	return convtime
}

// func bigOrsmall(in int) string {
// var tmp_used string
// tmp_x, _ := strconv.Atoi(second.var_max_connections)
// if second.Max_used_connections > (tmp_max_connections * 7 / 10) {
// tmp_used = Colorize(strconv.Itoa(second.Max_used_connections), red, "", "", "y")
// } else {
// tmp_used = Colorize(strconv.Itoa(second.Max_used_connections), "", "", "", "")
// }
// }

func hitFloat(num int, in float64) string {
	var result string
        if  len(floatToString(in, 2)) > num {
		num=len(floatToString(in, 2))
	}
	if in > 99.0 {
		result = Colorize(strings.Repeat(" ", num-len(floatToString(in, 2)))+floatToString(in, 2), green, "", "", "")
	} else if in > 90.0 && in <= 99.0 {
		result = Colorize(strings.Repeat(" ", num-len(floatToString(in, 2)))+floatToString(in, 2), yellow, "", "", "")
	} else if in < 0.01 {
		result = Colorize(strings.Repeat(" ", num-len("100.00"))+"100.00", green, "", "", "")
	} else {
		result = Colorize(strings.Repeat(" ", num-len(floatToString(in, 2)))+floatToString(in, 2), red, "", "", "y")
	}
	return result
}
func hitNum(num int, in int) int {
        var result int
        if  in >= num {
                result=1
        } else {
		result=num-in
	}
        return result
}

func hitAbs(in int) int {
	if in > 0 {
		return in
	} else {
		return 0 - in
	}
}

func formatStr(num int, inStr string, status string, background string, underline string, highshow string) string {
	return Colorize(strings.Repeat(" ", hitNum(num,len(inStr)))+inStr, status, background, underline, highshow)
}
//参数解析 输入参数字典 结果集字典 运行次数字典
func gotNumber(flag_info map[string]interface{}, first basic, second basic, count int) {
	var title_summit string
	var title_detail string
	var data_detail string
	var pic string
	interval, _ := strconv.Atoi(flag_info["interval"].(string))

	if count == 0 {
		var tmp_used string
		tmp_max_connections, _ := strconv.Atoi(second.var_max_connections)
		if second.Max_used_connections > (tmp_max_connections * 7 / 10) {
			tmp_used = Colorize(strconv.Itoa(second.Max_used_connections), red, "", "", "y")
		} else {
			tmp_used = Colorize(strconv.Itoa(second.Max_used_connections), "", "", "", "")
		}

		var tmptable string
		tmp_table_x := float64(second.Created_tmp_disk_tables) / float64(second.Created_tmp_tables) * 100
		if tmp_table_x < 10.0 {
			tmptable = Colorize(floatToString(tmp_table_x, 2), green, "", "", "")
		} else {
			tmptable = Colorize(floatToString(tmp_table_x, 2), red, "", "", "y")
		}

		pic += Colorize(".==========================================================================================================.\n", green, "", "", "")
                pic += Colorize("|", green, "", "", "") + "Author:" + Colorize("Li", green, "", "", "") + Colorize("Xue", red, "", "", "") + Colorize("Ping", purple, "", "", "") + Colorize("|", green, "", "", "") + "\n"
		pic += Colorize("|", green, "", "", "") + "Author:" + Colorize("Chen", red, "", "", "") + Colorize("Xing", green, "", "", "")  + Colorize("Long", purple, "", "", "") + Colorize("|", green, "", "", "") + "\n"
		pic += Colorize("'=========================================================================================================='\n\n", green, "", "", "")
		pic += Colorize("HOST: ", red, "", "", "") + Colorize(strings.Replace(second.hostname, "\n", "", -1), yellow, "", "", "") + Colorize("IP: ", red, "", "", "") + Colorize(strings.Replace(second.ip, "\n", "", -1), yellow, "", "", "") + "\n"
		pic += Colorize("DB  : ", red, "", "", "") + Colorize(second.db, yellow, "", "", "") + "\n"
		pic += Colorize("Var : ", red, "", "", "") + Colorize("binlog_format", purple, "", "", "") + "[" + second.var_binlog_format + "]" + Colorize(" max_binlog_cache_size", purple, "", "", "") + "[" + changeUnits(second.var_max_binlog_cache_size) + "]" + Colorize(" max_binlog_size", purple, "", "", "") + "[" + changeUnits(second.var_max_binlog_size) + "]" + Colorize(" sync_binlog", purple, "", "", "") + "[" + second.var_sync_binlog + "]" + "\n"
		pic += Colorize("  max_connect_errors", purple, "", "", "") + "[" + second.var_max_connect_errors + "]" + Colorize(" max_connections", purple, "", "", "") + "[" + second.var_max_connections + "]" + Colorize(" max_user_connections", purple, "", "", "") + "[" + second.var_max_user_connections + "]" + Colorize(" max_used_connections", purple, "", "", "") + "[" + tmp_used + "]" + "\n"
		pic += Colorize("  open_files_limit", purple, "", "", "") + "[" + second.var_open_files_limit + "]" + Colorize(" table_definition_cache", purple, "", "", "") + "[" + second.var_table_definition_cache + "]" + Colorize(" Aborted_connects", purple, "", "", "") + "[" + second.Aborted_connects + "]" + Colorize(" Aborted_clients", purple, "", "", "") + "[" + second.Aborted_clients + "]" + "\n"
		pic += Colorize("  Binlog_cache_disk_use", purple, "", "", "") + "[" + second.Binlog_cache_disk_use + "]" + Colorize(" Select_scan", purple, "", "", "") + "[" + second.Select_scan + "]" + Colorize(" Select_full_join", purple, "", "", "") + "[" + second.Select_full_join + "]" + Colorize(" Slow_queries", purple, "", "", "") + "[" + second.Slow_queries + "]\n"
		if second.Rpl_semi_sync_master_status != "" {
			pic += Colorize("  Rpl_semi_sync_master_status", purple, "", "", "") + "[" + second.Rpl_semi_sync_master_status + "]" + Colorize(" Rpl_semi_sync_slave_status", purple, "", "", "") + "[" + second.Rpl_semi_sync_slave_status + "]" + Colorize(" rpl_semi_sync_master_timeout", purple, "", "", "") + "[" + second.rpl_semi_sync_master_timeout + "]\n"
		}
		if second.Master_Host != "" {
			pic += Colorize("  Master_Host", purple, "", "", "") + "[" + second.Master_Host + "]" + Colorize(" Master_User", purple, "", "", "") + "[" + second.Master_User + "]" + Colorize(" Master_Port", purple, "", "", "") + "[" + second.Master_Port + "]" + Colorize(" Master_Server_Id", purple, "", "", "") + "[" + second.Master_Server_Id + "]\n"
			io := ""
			sql := ""
			if second.Slave_IO_Running != "Yes" {
				io = Colorize("No", red, "", "", "y")
			} else {
				io = Colorize("Yes", green, "", "", "")
			}
			if second.Slave_SQL_Running != "Yes" {
				sql = Colorize("No", red, "", "", "y")
			} else {
				sql = Colorize("Yes", green, "", "", "")
			}
			pic += Colorize("  Slave_IO_Running", purple, "", "", "") + "[" + io + "]" + Colorize(" Slave_SQL_Running", purple, "", "", "") + "[" + sql + "]\n"
		}
		pic += Colorize("  table_open_cache", purple, "", "", "") + "[" + second.var_table_open_cache + "]" + Colorize(" thread_cache_size", purple, "", "", "") + "[" + second.var_thread_cache_size + "]" + Colorize(" Opened_tables", purple, "", "", "") + "[" + second.Opened_tables + "]" + Colorize(" Created_tmp_disk_tables_ratio", purple, "", "", "") + "[" + tmptable + "]\n\n"

		pic += Colorize("  innodb_adaptive_flushing", purple, "", "", "") + "[" + second.var_innodb_adaptive_flushing + "]" + Colorize(" innodb_adaptive_hash_index", purple, "", "", "") + "[" + second.var_innodb_adaptive_hash_index + "]" + Colorize(" innodb_buffer_pool_size", purple, "", "", "") + "[" + changeUnits(second.var_innodb_buffer_pool_size) + "]" + "\n"
		pic += Colorize("  innodb_file_per_table", purple, "", "", "") + "[" + second.var_innodb_file_per_table + "]" + Colorize(" innodb_flush_log_at_trx_commit", purple, "", "", "") + "[" + second.var_innodb_flush_log_at_trx_commit + "]" + Colorize(" innodb_flush_method", purple, "", "", "") + "[" + second.var_innodb_flush_method + "]" + "\n"
		pic += Colorize("  innodb_io_capacity", purple, "", "", "") + "[" + second.var_innodb_io_capacity + "]" + Colorize(" innodb_lock_wait_timeout", purple, "", "", "") + "[" + second.var_innodb_lock_wait_timeout + "]" + Colorize(" innodb_log_buffer_size", purple, "", "", "") + "[" + changeUnits(second.var_innodb_log_buffer_size) + "]" + "\n"
		pic += Colorize("  innodb_log_file_size", purple, "", "", "") + "[" + changeUnits(second.var_innodb_log_file_size) + "]" + Colorize(" innodb_log_files_in_group", purple, "", "", "") + "[" + second.var_innodb_log_files_in_group + "]" + Colorize(" innodb_max_dirty_pages_pct", purple, "", "", "") + "[" + second.var_innodb_max_dirty_pages_pct + "]\n"
		pic += Colorize("  innodb_open_files", purple, "", "", "") + "[" + second.var_innodb_open_files + "]" + Colorize(" innodb_read_io_threads", purple, "", "", "") + "[" + second.var_innodb_read_io_threads + "]" + Colorize(" innodb_thread_concurrency", purple, "", "", "") + "[" + second.var_innodb_thread_concurrency + "]" + "\n"
		pic += Colorize("  innodb_write_io_threads", purple, "", "", "") + "[" + second.var_innodb_write_io_threads + "]" + "\n"
	}

	// myid 信息
        if len(flag_info["myid"].(string)) > 0 {
                title_summit = Colorize("------ ", dgreen, "", "", "")
                title_detail = Colorize("  myid|", dgreen, "", "y", "")
                data_detail  = formatStr(6, flag_info["myid"].(string), yellow, "", "", "") + Colorize("|", dgreen, "", "", "")
        }
        // fulltime 信息
        if flag_info["dtime"] == true {
                title_summit += Colorize("-----------", dgreen, "", "", "")
                title_detail += Colorize("      date ", dgreen, "", "y", "")
                data_detail  += Colorize(getNowDate() + " ", yellow, "", "", "")
        }

	// time 信息
	if flag_info["time"] == true {
		title_summit += Colorize("-------- ", dgreen, "", "", "")
		title_detail += Colorize("    time|", dgreen, "", "y", "")
		data_detail  += Colorize(getNowTime(), yellow, "", "", "") + Colorize("|", dgreen, "", "", "")
	}

	//loadavg 信息
	if flag_info["load"] == true {
		title_summit += Colorize("-----load-avg---- ", dgreen, "", "", "")
		title_detail += Colorize("   1m   5m   15m |", dgreen, "", "y", "")
		// fmt.Println(strings.Repeat(" ", 5-len(floatToString(first.load_1, 2)))+floatToString(first.load_1, 2), floatToString(first.load_1, 2), len(floatToString(first.load_1, 2)))
		//load 1 min
		if first.load_1 > first.cpu_core {
			if first.load_1 >= 10.0 {
				data_detail += formatStr(5, floatToString(first.load_1, 2), red, "", "", "y")
			} else {
				data_detail += formatStr(5, floatToString(first.load_1, 2), yellow, "", "", "y")
			}
		} else {
			data_detail += formatStr(5, floatToString(first.load_1, 2), "", "", "", "")
		}

		if first.load_5 > first.cpu_core {
			if first.load_1 >= 10.0 {
				data_detail += formatStr(6, floatToString(first.load_5, 2), red, "", "", "y")
			} else {
				data_detail += formatStr(6, floatToString(first.load_5, 2), yellow, "", "", "y")
			}
		} else {
			data_detail += formatStr(6, floatToString(first.load_5, 2), "", "", "", "")
		}

		if first.load_15 > first.cpu_core {
			if first.load_1 >= 10.0 {
				data_detail += formatStr(6, floatToString(first.load_15, 2), red, "", "", "y") + Colorize("|", dgreen, "", "", "")
			} else {
				data_detail += formatStr(6, floatToString(first.load_15, 2), yellow, "", "", "y") + Colorize("|", dgreen, "", "", "")
			}
		} else {
			data_detail += formatStr(6, floatToString(first.load_15, 2), "", "", "", "") + Colorize("|", dgreen, "", "", "")
		}

	}

	//cpu-usage
	if flag_info["cpu"] == true {
		title_summit += Colorize("---cpu-usage--- ", dgreen, "", "", "")
		title_detail += Colorize("usr sys idl iow|", dgreen, "", "y", "")

		cpu_total1 := first.cpu_usr + first.cpu_nice + first.cpu_sys + first.cpu_idl + first.cpu_iow + first.cpu_irq + first.cpu_softirq
		cpu_total2 := second.cpu_usr + second.cpu_nice + second.cpu_sys + second.cpu_idl + second.cpu_iow + second.cpu_irq + second.cpu_softirq

		usr := (second.cpu_usr - first.cpu_usr) * 100 / (cpu_total2 - cpu_total1)
		sys := (second.cpu_sys - first.cpu_sys) * 100 / (cpu_total2 - cpu_total1)
		idl := (second.cpu_idl - first.cpu_idl) * 100 / (cpu_total2 - cpu_total1)
		iow := (second.cpu_iow - first.cpu_iow) * 100 / (cpu_total2 - cpu_total1)
		//usr

		if usr > 10 {
			data_detail += formatStr(3, strconv.Itoa(usr) , red, "", "", "y")
		} else {
			data_detail += formatStr(3, strconv.Itoa(usr) , "", "", "", "")
		}

		if sys > 10 {
			data_detail += formatStr(4, strconv.Itoa(sys) , red, "", "", "y")
		} else {
			data_detail += formatStr(4, strconv.Itoa(sys) , "", "", "", "")
		}

		if 1 != 1 {
			data_detail += formatStr(4, strconv.Itoa(idl) , red, "", "", "y")
		} else {
			data_detail += formatStr(4, strconv.Itoa(idl) , "", "", "", "")
		}

		if iow > 10 {
			data_detail += formatStr(4, strconv.Itoa(iow), red, "", "", "")
		} else {
			data_detail += formatStr(4, strconv.Itoa(iow), "", "", "", "")
		}
		data_detail += Colorize("|", dgreen, "", "", "")
	}

	//swap
	if flag_info["swap"] == true {
		title_summit += Colorize("---swap--- ", dgreen, "", "", "")
		title_detail += Colorize("   si   so|", dgreen, "", "y", "")
		if flag_info["interval"] == "1" && count == 0 {
			data_detail += "00" + Colorize("|", dgreen, "", "", "y")
		} else if flag_info["interval"] == "1" && count > 0 {
			si := second.swap_in - first.swap_in
			so := second.swap_out - first.swap_out
			// fmt.Println(second.swap_in, first.swap_in, si, second.swap_out, first.swap_out, so)
			si_string := strconv.Itoa(si)
			so_string := strconv.Itoa(so)

			in := strings.Repeat(" ", 5-len(si_string)) + si_string
			out := strings.Repeat(" ", 5-len(so_string)) + so_string
			if si > 0 {
				data_detail += Colorize(in, red, "", "", "y")
			} else {
				data_detail += Colorize(in, "", "", "", "")
			}

			if so > 0 {
				data_detail += Colorize(out, red, "", "", "y")
			} else {
				data_detail += Colorize(out, "", "", "", "")
			}

			data_detail += Colorize("|", dgreen, "", "", "")
		}
	}

	//net
	//swap
	if flag_info["net"] != "none" {
		title_summit += Colorize("---net(B)--- ", dgreen, "", "", "")
		title_detail += Colorize("  recv  send|", dgreen, "", "y", "")
		if flag_info["interval"] == "1" && count == 0 {
			data_detail += "  0  0" + Colorize("|", dgreen, "", "", "y")
		} else if flag_info["interval"] == "1" && count > 0 {
			net_in := float64(second.net_recv-first.net_recv) / 0.99
			net_out := float64(second.net_send-first.net_send) / 0.99
			if int(net_in) > 10*1024*1024 {
				data_detail += formatStr(6, changeUnits(int(net_in)), red, "", "", "")
			} else {
				data_detail += formatStr(6, changeUnits(int(net_in)), "", "", "", "")
			}
			if int(net_out) > 10*1024*1024 {
				data_detail += formatStr(6, changeUnits(int(net_out)), red, "", "", "")
			} else {
				data_detail += formatStr(6, changeUnits(int(net_out)), "", "", "", "")
			}
			data_detail += Colorize("|", dgreen, "", "", "")
		}
	}

	//disk
	if flag_info["disk"] != "none" {
		title_summit += Colorize("-------------------------io-usage----------------------- ", dgreen, "", "", "")
		title_detail += Colorize("   r/sw/srkB/swkB/s  queue await svctm "+"%"+"util|", dgreen, "", "y", "")
		if count == 0 {
			data_detail += Colorize("0.00.0 0.0  0.0   0.00.0   0.0   0.0|", "", "", "", "")
		} else {
			// fmt.Printf("rs_disk is float64(%d-%d)/0.999\n", second.io_1, first.io_1)
			rs_disk := float64(second.io_1-first.io_1) / 0.9999
			// fmt.Printf("ws_disk is float64(%d-%d)/0.999\n", second.io_5, first.io_5)
			ws_disk := float64(second.io_5-first.io_5) / 0.9999

			// fmt.Printf("rkbs_disk is float64(%d-%d)/1.999\n", second.io_3, first.io_3)
			rkbs_disk := float64(second.io_3-first.io_3) / 1.9999
			// fmt.Printf("wkbs_disk is float64(%d-%d)/1.999\n", second.io_7, first.io_7)
			wkbs_disk := float64(second.io_7-first.io_7) / 1.9999

			queue_disk := strconv.Itoa(second.io_9)

			var await_disk float64
			var svctm_disk float64
			if (rs_disk + ws_disk) == 0.0 {
				await_disk = float64(second.io_4+second.io_8-first.io_4-first.io_8) / (rs_disk + ws_disk + 1)
				svctm_disk = float64(second.io_10-first.io_10) / (rs_disk + ws_disk + 1)
			} else {
				await_disk = float64(second.io_4+second.io_8-first.io_4-first.io_8) / (rs_disk + ws_disk)
				svctm_disk = float64(second.io_10-first.io_10) / (rs_disk + ws_disk)
			}

			util_disk := float64(second.io_10-first.io_10) / 10
			//usr
			// fmt.Println(rs_disk, ws_disk, rkbs_disk, wkbs_disk, queue_disk, await_disk, svctm_disk, util_disk)
			// fmt.Println(strings.Repeat(" ", 6-len(floatToString(rs_disk, 1))) + floatToString(rs_disk, 1))
			if 1 != 1 {
				// data_detail += Colorize(strings.Repeat(" ", 6-len(floatToString(rs_disk, 1)))+floatToString(rs_disk, 1), red, "", "", "y")
				data_detail += formatStr(6, floatToString(rs_disk, 1), red, "", "", "y")
			} else {
				// data_detail += Colorize(strings.Repeat(" ", 6-len(floatToString(rs_disk, 1)))+floatToString(rs_disk, 1), "", "", "", "")
				data_detail += formatStr(6, floatToString(rs_disk, 1), "", "", "", "")
			}

			if 1 != 1 {
				// data_detail += Colorize(strings.Repeat(" ", 7-len(floatToString(ws_disk, 1)))+floatToString(ws_disk, 1), red, "", "", "y")
				data_detail += formatStr(7, floatToString(ws_disk, 1), red, "", "", "y")
			} else {
				// data_detail += Colorize(strings.Repeat(" ", 7-len(floatToString(ws_disk, 1)))+floatToString(ws_disk, 1), "", "", "", "")
				data_detail += formatStr(7, floatToString(ws_disk, 1), "", "", "", "")
			}

			if rkbs_disk > 1024.0 {
				// data_detail += Colorize(strings.Repeat(" ", 9-len(floatToString(rkbs_disk, 1)))+floatToString(rkbs_disk, 1), red, "", "", "y")
				data_detail += formatStr(9, floatToString(rkbs_disk, 1), red, "", "", "y")
			} else {
				// data_detail += Colorize(strings.Repeat(" ", 9-len(floatToString(rkbs_disk, 1)))+floatToString(rkbs_disk, 1), "", "", "", "")
				data_detail += formatStr(9, floatToString(rkbs_disk, 1), "", "", "", "")
			}

			if wkbs_disk > 1024.0 {
				data_detail += formatStr(9, floatToString(wkbs_disk, 1), red, "", "", "y")
			} else {
				data_detail += formatStr(9, floatToString(wkbs_disk, 1), "", "", "", "")
			}

			if second.io_9 > 10 {
				data_detail += formatStr(4, queue_disk+".0 ", red, "", "", "y")
			} else {
				data_detail += formatStr(4, queue_disk+".0 ", "", "", "", "")
			}

			if await_disk > 5.0 {
				data_detail += formatStr(6, floatToString(await_disk, 1), red, "", "", "y")
			} else {
				data_detail += formatStr(6, floatToString(await_disk, 1), green, "", "", "")
			}

			if svctm_disk > 5.0 {
				data_detail += formatStr(6, floatToString(svctm_disk, 1), red, "", "", "y")
			} else {
				data_detail += formatStr(6, floatToString(svctm_disk, 1), "", "", "", "")
			}

			if util_disk > 80.0 {
				data_detail += formatStr(6, floatToString(util_disk, 1), red, "", "", "y")
			} else if util_disk > 100.0 {
				data_detail += Colorize(" 100.0", green, "", "", "")
			} else {
				data_detail += formatStr(6, floatToString(util_disk, 1), green, "", "", "")
			}

			data_detail += Colorize("|", dgreen, "", "", "")
		}

	}

	//-com
	if flag_info["com"] == true {
		title_summit += Colorize("----MySQLcom----- -QPS- -TPS- ", green, blue, "", "")
		title_detail += Colorize("  ins   upd   del   sel   iud|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("0 0 0  0 0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			insert_diff := (second.Com_insert - first.Com_insert) / interval
			update_diff := (second.Com_update - first.Com_update) / interval
			delete_diff := (second.Com_delete - first.Com_delete) / interval
			select_diff := (second.Com_select - first.Com_select) / interval
			// commit_diff := (second.Com_commit - first.Com_commit) / interval
			// rollback_diff := (second.Com_rollback - first.Com_rollback) / interval
                        // tps := commit_diff + rollback_diff
			tps := insert_diff + update_diff + delete_diff

			data_detail += formatStr(5,strconv.Itoa(insert_diff), "", "", "", "")
			data_detail += formatStr(6,strconv.Itoa(update_diff), "", "", "", "")
			data_detail += formatStr(6,strconv.Itoa(delete_diff), "", "", "", "")
			data_detail += formatStr(6,strconv.Itoa(select_diff), yellow, "", "", "")
			data_detail += formatStr(6,strconv.Itoa(tps), yellow, "", "", "")
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//hit
	if flag_info["hit"] == true {
		if flag_info["ext"] == true {
			title_summit += Colorize("----KeyBuffer------Index----Qcache", green, blue, "", "")
			title_detail += Colorize("  read  write    cur  total    hit", green, "", "y", "")
			if count == 0 {
		                data_detail += Colorize("100.00 100.00 100.00 100.00 100.00", "", "", "", "") + Colorize("|", green, "", "", "")
	                } else {
				key_read := (second.Key_reads - first.Key_reads) / interval
	                        key_write := (second.Key_writes - first.Key_writes) / interval
	                        key_read_request := (second.Key_read_requests - first.Key_read_requests) / interval
	                        key_write_request := (second.Key_write_requests - first.Key_write_requests) / interval
	                        //innodb hit
	                        hrr := (second.Handler_read_rnd - first.Handler_read_rnd) / interval
	                        hrrn := (second.Handler_read_rnd_next - first.Handler_read_rnd_next) / interval
	                        hrf := (second.Handler_read_first - first.Handler_read_first) / interval
	                        hrk := (second.Handler_read_key - first.Handler_read_key) / interval
	                        hrn := (second.Handler_read_next - first.Handler_read_next) / interval
	                        hrp := (second.Handler_read_prev - first.Handler_read_prev) / interval
	                        //key buffer read hit
	                        key_read_hit := (float64(key_read_request-key_read) + 0.0001) / (float64(key_read_request) + 0.0001) * 100
	                        key_write_hit := (float64(key_write_request-key_write) + 0.0001) / (float64(key_write_request) + 0.0001) * 100
	                        index_total_hit := (100 - (100 * (float64(second.Handler_read_rnd+second.Handler_read_rnd_next) + 0.0001) / (0.0001 + float64(second.Handler_read_first+second.Handler_read_key+second.Handler_read_next+second.Handler_read_prev+second.Handler_read_rnd+second.Handler_read_rnd_next))))
	                        index_current_hit := 100.00
	                        if hrr+hrrn != 0 {
	                                index_current_hit = (100 - (100 * (float64(hrr+hrrn) + 0.0001) / (0.0001 + float64(hrf+hrk+hrn+hrp+hrr+hrrn))))
	                        }
	                        query_hits_s := (second.Qcache_hits - first.Qcache_hits) / interval
	                        com_select_s := (second.Com_select - first.Com_select) / interval
	                        query_hit := (float64(query_hits_s) + 0.0001) / (float64(query_hits_s+com_select_s) + 0.0001) * 100
				data_detail += hitFloat(6, key_read_hit)
	                        data_detail += hitFloat(7, key_write_hit)
	                        data_detail += hitFloat(7, index_current_hit)
	                        data_detail += hitFloat(7, index_total_hit)
	                        data_detail += hitFloat(7, query_hit)
			}
		}
		title_summit += Colorize("---Innodb------ ", green, blue, "", "")
		title_detail += Colorize("     lor   %hit|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("   0 100.00|", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			read_request := (second.Innodb_buffer_pool_read_requests - first.Innodb_buffer_pool_read_requests) / interval
			read := (second.Innodb_buffer_pool_reads - first.Innodb_buffer_pool_reads) / interval
			//innodb_hit
			innodb_hit := ((float64(read_request-read) + 0.0001) / (float64(read_request) + 0.0001)) * 100
			// lor = read_request
			data_detail += Colorize(strings.Repeat(" ", 8-len(strconv.Itoa(read_request)))+strconv.Itoa(read_request), "", "", "", "")
			data_detail += hitFloat(7, innodb_hit)

			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//innodb_rows
	if flag_info["innodb_rows"] == true {
		title_summit += Colorize("---innodb rows status--- ", green, blue, "", "")
		title_detail += Colorize("   ins   upd   del  read|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("0 0 0  0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			innodb_rows_inserted_diff := (second.Innodb_rows_inserted - first.Innodb_rows_inserted) / interval
			innodb_rows_updated_diff := (second.Innodb_rows_updated - first.Innodb_rows_updated) / interval
			innodb_rows_deleted_diff := (second.Innodb_rows_deleted - first.Innodb_rows_deleted) / interval
			innodb_rows_read_diff := (second.Innodb_rows_read - first.Innodb_rows_read) / interval

			data_detail += formatStr(6,strconv.Itoa(innodb_rows_inserted_diff), "", "", "", "")
                        data_detail += formatStr(6,strconv.Itoa(innodb_rows_updated_diff), "", "", "", "")
                        data_detail += formatStr(6,strconv.Itoa(innodb_rows_deleted_diff), "", "", "", "")
                        data_detail += formatStr(6,changeUnits(innodb_rows_read_diff), getColor(innodb_rows_read_diff,500000,2000000), "", "", "")
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//innodb_pages
	if flag_info["innodb_pages"] == true {
		title_summit += Colorize("---innodb bp pages status--- ", green, blue, "", "")
		title_detail += Colorize("    data   free  dirty flush|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("  0  0  0 0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			flush := (second.Innodb_buffer_pool_pages_flushed - first.Innodb_buffer_pool_pages_flushed) / interval

                        data_detail += formatStr(8,changeUnits(second.Innodb_buffer_pool_pages_data), "", "", "", "")
                        data_detail += formatStr(7,changeUnits(second.Innodb_buffer_pool_pages_free), "", "", "", "")
                        data_detail += formatStr(7,changeUnits(second.Innodb_buffer_pool_pages_dirty), getColor(second.Innodb_buffer_pool_pages_dirty, 1024, 1048576) , "", "", "")
                        data_detail += formatStr(6,changeUnits(flush), getColor(flush, 1000, 3000), "", "", "")
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//innodb_data
	if flag_info["innodb_data"] == true {
		title_summit += Colorize("------innodb data status------ ", green, blue, "", "")
		title_detail += Colorize(" reads writes    read  written|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize(" 0  0  0  0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			innodb_data_reads_diff := (second.Innodb_data_reads - first.Innodb_data_reads) / interval
			innodb_data_writes_diff := (second.Innodb_data_writes - first.Innodb_data_writes) / interval
			innodb_data_read_diff := (second.Innodb_data_read - first.Innodb_data_read) / interval
			innodb_data_written_diff := (second.Innodb_data_written - first.Innodb_data_written) / interval

			data_detail += formatStr(6, strconv.Itoa(innodb_data_reads_diff), "", "", "", "")
			data_detail += formatStr(7, strconv.Itoa(innodb_data_writes_diff), "", "", "", "")

			data_detail += formatStr(8, changeUnits(innodb_data_read_diff), "", "", "", "")
			data_detail += formatStr(9, changeUnits(innodb_data_written_diff), "", "", "", "")

			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//innodb_log
	if flag_info["innodb_log"] == true {
		title_summit += Colorize("--innodb log--- ", green, blue, "", "")
		title_detail += Colorize("fsyncs  written|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize(" 0   0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {

			innodb_os_log_fsyncs_diff := (second.Innodb_os_log_fsyncs - first.Innodb_os_log_fsyncs) / interval
			innodb_os_log_written_diff := (second.Innodb_os_log_written - first.Innodb_os_log_written) / interval

			data_detail += Colorize(strings.Repeat(" ", 6-len(strconv.Itoa(innodb_os_log_fsyncs_diff)))+strconv.Itoa(innodb_os_log_fsyncs_diff), "", "", "", "")
			data_detail += formatStr(9, changeUnits(innodb_os_log_written_diff), getColor(innodb_os_log_written_diff,1023,1048576), "", "", "")
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//innodb_status
	if flag_info["innodb_status"] == true {
		title_summit += Colorize("--his---log(byte)------  read ---query--- ", green, blue, "", "")
		title_detail += Colorize(" list  uflush     uckpt  view inside  que|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("0  0   0  0 0 0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {

			//mysql --innodb_status show engine innodb status
			//log unflushed = Log sequence number - Log flushed up to
			//uncheckpointed bytes = Log sequence number - Last checkpoint at
			//mysql -e "show engine innodb status\G"|grep -n -E -A4 -B1 "^TRANSACTIONS|LOG|ROW OPERATIONS"
			//mysql -e "show engine innodb status\G"|grep -E "Last checkpoint|read view|queries inside|queue"
			// Log_sequence int
			// Log_flushed  int
			// History_list int
			// Last_checkpoint int
			// Read_view    int
			// Query_inside int
			// Query_queue  int

			unflushed_log := second.Log_sequence - second.Log_flushed
			uncheckpointed_bytes := second.Log_sequence - second.Last_checkpoint
			//History_list
			data_detail += formatStr(5, changeUnits(second.History_list), "", "", "", "")
			//unflushed_log
			data_detail += formatStr(8, changeUnits(unflushed_log), yellow, "", "", "")

			//uncheckpointed_bytes
			data_detail += formatStr(10, changeUnits(uncheckpointed_bytes), yellow, "", "", "")

			//Read_views
			data_detail += formatStr(6, strconv.Itoa(second.Read_view), "", "", "", "")
			//inside
			data_detail += formatStr(6, strconv.Itoa(second.Query_inside), "", "", "", "")
			//queue
			data_detail += formatStr(6, strconv.Itoa(second.Query_queue), "", "", "", "")

			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//threads ------threads------
	if flag_info["threads"] == true {
		title_summit += Colorize("----------threads--------- ", green, blue, "", "")
		title_detail += Colorize(" run  con  cre  cac   "+"%"+"hit|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("   0000 0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			connections_diff := (second.Connections - first.Connections) / interval

			threads_created_diff := (second.Threads_created - first.Threads_created) / interval

			thread_cache_hit := (1 - float64(threads_created_diff)/float64(connections_diff)) * 100

			data_detail += formatStr(4, strconv.Itoa(second.Threads_running), "", "", "", "")

			data_detail += formatStr(5, strconv.Itoa(second.Threads_connected), "", "", "", "")

			data_detail += formatStr(5, strconv.Itoa(threads_created_diff), "", "", "", "")

			data_detail += formatStr(5, strconv.Itoa(second.Threads_cached), "", "", "", "")
			if thread_cache_hit > 99.0 {
				data_detail += formatStr(7, floatToString(thread_cache_hit, 2), green, "", "", "")
			} else if thread_cache_hit <= 99.0 && thread_cache_hit > 90.0 {
				data_detail += formatStr(7, floatToString(thread_cache_hit, 2), yellow, "", "", "")
			} else {
				data_detail += formatStr(7, floatToString(thread_cache_hit, 2), red, "", "", "")
			}

			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//bytes
	if flag_info["bytes"] == true {
		title_summit += Colorize("-----bytes---- ", green, blue, "", "")
		title_detail += Colorize("   recv   send|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("  0  0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {

			bytes_received_diff := (second.Bytes_received - first.Bytes_received) / interval
			bytes_sent_diff := (second.Bytes_sent - first.Bytes_sent) / interval

			data_detail += formatStr(7, changeUnits(bytes_received_diff), "", "", "", "")
			data_detail += formatStr(7, changeUnits(bytes_sent_diff), "", "", "", "")
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//semi
	if flag_info["semi"] == true {
		if flag_info["ext"] == true {
			title_summit += Colorize("------wait_time(s)----- ----------", green, blue, "", "")
			title_detail += Colorize("  net    tx  nets   txs   yes   no", green, "", "y", "")
			if count == 0 {
				data_detail += Colorize("100ms 100ms 1000 1000  1000 1000", "", "", "", "")
			} else {
				// 主库网络等待时间
				semi_net_diff  := (second.Rpl_semi_sync_master_net_wait_time - first.Rpl_semi_sync_master_net_wait_time) / interval
				// 主库事务的等待时间
				semi_tx_diff := (second.Rpl_semi_sync_master_tx_wait_time - first.Rpl_semi_sync_master_tx_wait_time) / interval
				// 主库网络等待次数
				semi_nets_diff := (second.Rpl_semi_sync_master_net_waits - first.Rpl_semi_sync_master_net_waits) / interval
				// 主库事务的等待次数
				semi_txs_diff  := (second.Rpl_semi_sync_master_tx_waits - first.Rpl_semi_sync_master_tx_waits) / interval
				// 半同步成功的事务数
				semi_yestx_diff :=(second.Rpl_semi_sync_master_yes_tx - first.Rpl_semi_sync_master_yes_tx) / interval
				// 半同步失败的是无数
				semi_notx_diff  :=(second.Rpl_semi_sync_master_no_tx - first.Rpl_semi_sync_master_no_tx) / interval
				data_detail += formatStr(5, strconv.Itoa(semi_net_diff), "", "", "", "")
				data_detail += formatStr(6, changeUnits(semi_tx_diff), "", "", "", "")
				data_detail += formatStr(6, changeUnits(semi_nets_diff), "", "", "", "")
				data_detail += formatStr(6, changeUnits(semi_txs_diff), "", "", "", "")
				data_detail += formatStr(6, changeUnits(semi_yestx_diff), "", "", "", "")
				data_detail += formatStr(5, changeUnits(semi_notx_diff), "", "", "", "")
			}
		}
		title_summit += Colorize(" --semiStatus--- ", green, blue, "", "")
		title_detail += Colorize(" sen cli mas rep|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("   0   0   N   N", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			// fmt.Printf("1 %d 2 %d 3 %d 4 %d 5 %d", second.Rpl_semi_sync_master_net_avg_wait_time, second.Rpl_semi_sync_master_tx_avg_wait_time, second.Rpl_semi_sync_master_no_tx, second.Rpl_semi_sync_master_yes_tx, second.Rpl_semi_sync_master_no_times)
			// 主库当前有几个session在等待备库响应
                        data_detail += formatStr(4, strconv.Itoa(second.Rpl_semi_sync_master_wait_sessions), "", "", "", "y")
			data_detail += formatStr(4, strconv.Itoa(second.Rpl_semi_sync_master_clients), "", "", "", "y")
			data_detail += formatStr(4, second.Rpl_semi_sync_master_status, "", "", "", "y")
			data_detail += formatStr(4, second.Rpl_semi_sync_slave_status, "", "", "", "y")
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//threads ------threads------
	if flag_info["slave"] == true {
		if flag_info["ext"] == true {
			title_summit += Colorize("---------MasterLogPos--------", green, blue, "", "")
			title_detail += Colorize("       Read       Exec  chkRE", green, "", "y", "")
			if count == 0 {
				data_detail += Colorize("    0    0     0", "", "", "", "") + Colorize("|", green, "", "", "")
	                } else {
				checkNum := second.Read_Master_Log_Pos - second.Exec_Master_Log_Pos
				if checkNum < 0 {
					checkNum = second.var_max_binlog_size + second.Read_Master_Log_Pos - second.Exec_Master_Log_Pos
				}
				data_detail += formatStr(11, changeUnits(second.Read_Master_Log_Pos), "", "", "", "")
				data_detail += formatStr(11, changeUnits(second.Exec_Master_Log_Pos), "", "", "", "")
				data_detail += formatStr(7, changeUnits(checkNum), "", "", "", "")
			}
		} else {

		}
		title_summit += Colorize("-----Slave---- ", green, blue, "", "")
		title_detail += Colorize(" SecBM  IO SQL|", green, "", "y", "")
		if count == 0 {
			data_detail += Colorize("     0    N    N", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			if second.Seconds_Behind_Master > 300 {
				data_detail += formatStr(6,strconv.Itoa(second.Seconds_Behind_Master), red, "", "", "")
			} else {
				data_detail += formatStr(6,strconv.Itoa(second.Seconds_Behind_Master), green, "", "", "")
			}
			if second.Slave_IO_Running == "Yes" {
			        data_detail += formatStr(4, second.Slave_IO_Running, green, "", "", "")
			} else {
			        data_detail += formatStr(4, second.Slave_IO_Running, red, "", "", "")
			}
			if second.Slave_SQL_Running == "Yes" {
			        data_detail += formatStr(4, second.Slave_SQL_Running, green, "", "", "")
			} else {
			        data_detail += formatStr(4, second.Slave_SQL_Running, red, "", "", "")
			}
			data_detail += Colorize("|", green, "", "", "")
		}
	}

	//rt
	if flag_info["rt"] == true {
		// 1,000 皮秒 = 1纳秒
		// 1,000,000 皮秒 = 1微秒
		// 1,000,000,000 皮秒 = 1毫秒
		// 1,000,000,000,000 皮秒 = 1秒
		var rt_count, rt_avg, rt_95avg, rt_99avg string
		title_summit += Colorize("--------tcprstat(us)-------- ", green, blue, "", "") + " "
		title_detail += Colorize("  count    Avg  95Avg  99Avg|", green, "", "y", "")
		if count == 0 {
			data_detail = Colorize("  0  0  0  0", "", "", "", "") + Colorize("|", green, "", "", "")
		} else {
			if second.rt_count > 1000 {
				rt_count = formatStr(7, strconv.Itoa(second.rt_count), red, "", "", "")
			} else {
				rt_count = formatStr(7, strconv.Itoa(second.rt_count), "", "", "", "")
			}

			if second.rt_avg/1000 > 60 {
				rt_avg = formatStr(7,strconv.Itoa(second.rt_avg), red, "", "", "")
			} else {
				rt_avg = formatStr(7,strconv.Itoa(second.rt_avg), green, "", "", "")
			}

			if second.rt_a5/1000 > 100 && second.rt_a5 != 0 {
				rt_95avg = formatStr(7, strconv.Itoa(second.rt_a5), red, "", "", "")
			} else {
				rt_95avg = formatStr(7, strconv.Itoa(second.rt_a5), green, "", "", "")
			}

			if second.rt_a9/1000 == 100 && second.rt_a9 != 0 {
				rt_99avg = formatStr(7, strconv.Itoa(second.rt_a9), red, "", "", "")
			} else {
				rt_99avg = formatStr(7, strconv.Itoa(second.rt_a9), green, "", "", "")
			}

			data_detail += rt_count + rt_avg + rt_95avg + rt_99avg + Colorize("|", green, "", "", "")
		}

	}

	if count == 0 {
		fmt.Println(pic)
		fmt.Println(title_summit)
		fmt.Println(title_detail)
		add_log(flag_info, pic)
		add_log(flag_info, title_summit)
		add_log(flag_info, title_detail)
	}
	if count != 0 && count%20 == 0 {
		fmt.Println(title_summit)
		fmt.Println(title_detail)
		add_log(flag_info, title_summit)
		add_log(flag_info, title_detail)
	}
	if count != 0 && count%20 != 0 {
		fmt.Println(data_detail)
		add_log(flag_info, data_detail)

	}

}

func add_log(flag_log map[string]interface{}, info string) {
	var file_name string
	if flag_log["logfile_by_day"].(bool) == true {
		t := time.Now()
		if flag_log["logfile"].(string) != "none" {
			file_name = flag_log["logfile"].(string) + "_" + fmt.Sprintf("%s", t.Format("2006-01-02")) + ".log"
		} else {
			file_name = "/tmp/orzdba" + "_" + fmt.Sprintf("%s", t.Format("2006-01-02")) + ".log"
		}
	} else {
		if flag_log["logfile"].(string) != "none" {
			file_name = flag_log["logfile"].(string)
		} else {
			file_name = "/tmp/orzdba.log"
		}
	}

	lf, err := os.OpenFile(file_name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0606)
	checkErr(err)

	defer lf.Close()

	l := log.New(lf, "", os.O_APPEND)

	l.Printf("%s\n", info)

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	info := GetValue()
	// ss := basic{}
	second := basic{}
	//捕获退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		select {
		case s := <-c:
			fmt.Printf("\n\033[1;4;31m%s:罒灬罒: Buy htmoon !\033[0m\n", s)
			execCommand("killall tcprstat")
			//不写就退不出来了
			os.Exit(1)
			// panic("退出")
		}
	}()

	//nocolor
	if info["nocolor"] == true {
		black  = ""
		red    = ""
		green  = ""
		yellow = ""
		blue   = ""
		purple = ""
		dgreen = ""
		white  = ""

	} else {
		black  = "black"
		red    = "red"
		green  = "green"
		yellow = "yellow"
		blue   = "blue"
		purple = "purple"
		dgreen = "dgreen"
		white  = "white"
	}
	//rt
	if info["rt"] == true {
		go func() {
			var rt_cmd string
			rt_cmd = "tcprstat --no-header -p " + info["port"].(string) + " -t " + info["interval"].(string) + " -n 0 -l `/sbin/ifconfig | grep 'addr:[^ ]\\+' -o | cut -f 2 -d : | xargs echo | sed -e 's/ /,/g'` 1>/tmp/orzdba_tcprstat.log"
			execCommand(rt_cmd)
		}()
	}

	if info["mysql"] == true {
		info["com"] = true
		info["hit"] = true
		info["threads"] = true
		info["bytes"] = true
	}

	if info["innodb"] == true {
		info["innodb_pages"] = true
		info["innodb_data"] = true
		info["innodb_log"] = true
		info["innodb_status"] = true
	}

	if info["sys"] == true {
		info["load"] = true
		info["cpu"] = true
		info["swap"] = true
	}

	if info["lazy"] == true {
		info["time"] = true
		info["load"] = true
		info["cpu"] = true
		info["swap"] = true
		info["com"] = true
		info["hit"] = true
	}

	first := createCommand(info, 0)
	//计算CPU核数
	cpu_core_cmd := "grep processor /proc/cpuinfo | wc -l"
	cpu_string := execCommand(cpu_core_cmd)
	cpu_string = strings.Replace(cpu_string, "\n", "", -1)
	first.cpu_core, _ = strconv.ParseFloat(cpu_string, 64)

	interval, _ := strconv.Atoi(info["interval"].(string))
	if info["count"] == 0 {
		i := 0
		for {
			second = createCommand(info, i)
			second.cpu_core, _ = strconv.ParseFloat(cpu_string, 64)
			gotNumber(info, first, second, i)
			first = second
			time.Sleep(time.Second * time.Duration(interval))
			i++
		}
	} else {
		for i := 0; i <= info["count"].(int); i++ {
			second = createCommand(info, i)
			second.cpu_core, _ = strconv.ParseFloat(cpu_string, 64)
			gotNumber(info, first, second, i)
			first = second
			time.Sleep(time.Second * time.Duration(interval))
		}
		execCommand("killall tcprstat")
		os.Exit(0)
	}
}
