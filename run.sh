#!/bin/bash
wk='/root/go/src/weixinsdk'
m_BuildingName=weixinsdk

# 进程pid
#命令帮助
shell_usage(){
 	echo "Usage: $0 (restart|dev|build)"
}

do_restart(){
    process=$(ps -ef | grep $m_BuildingName | grep -v grep | awk '{print $2}')
    if [ "$process" != "" ];then
         echo $(kill -TERM $process)
         echo $(nohup ./$m_BuildingName > $m_BuildingName.log 2>&1 &)
    fi
}

#开发环境 ./run.sh dev
do_dev(){
    process=$(ps -ef | grep $m_BuildingName | grep -v grep | awk '{print $2}')
    if [ "$process" != "" ];then
         echo $(kill -TERM $process)
    fi

    cd ${wk}/src &&go build -o $wk/$m_BuildingName . 
    cd ${wk} && chmod +x $m_BuildingName && ./$m_BuildingName
}

#正式环境
do_build(){
    process=$(ps -ef | grep $m_BuildingName | grep -v grep | awk '{print $2}')
    if [ "$process" != "" ];then
         echo $(kill -TERM $process)
    fi
    cd ${wk}/src &&go build -o $wk/$m_BuildingName . 
    cd ${wk} && chmod +x $m_BuildingName
    echo $(nohup ./$m_BuildingName > $m_BuildingName.log 2>&1 &)
}


#启动函数
main(){
	#switch切换函数
    case $1 in
        dev)
            do_dev
            ;;
        build)
            do_build
            ;;
        restart)
            do_restart
            ;;
		*)
			shell_usage;
			exit 1
	esac
}

#执行函数
main $1 $2