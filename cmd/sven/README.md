###使用方法
####说明
    基于gin框架 
    wire依赖注入
    借鉴kratos

####一、安装

    go install github.com/go-sven/sven/cmd/sven@latest


####二、新建项目

    sven new sven-demo

####三、新增模块

    sven new app/server/user --nomod

####四、运行

    1.修改config配置文件
    2.cd cmd 修改wire.go 假如新增的ProviderSet
    3. 执行  wire
    4.在跟目录执行 go run main.go