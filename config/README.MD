### Install
    go get github.com/go-sven/sven/config/viper
### Use

    filePath := "path/conf.yaml"
	conf, err := viper.NewViperConfig(filePath)
	if err != nil {
		fmt.Println(err)
	}
	c := new(AppConfig)
	err = conf.Parse(c)
	if err != nil {
		fmt.Println("err:", err)
	}
    fmt.Println(c.Database.Connection)