### Install
    go get github.com/go-sven/sven/contrib/conf/viper
### Use

    filePath := "path/conf.yaml"
	svenViper, err := viper.NewSvenViper(filePath)
	if err != nil {
		fmt.Println(err)
	}
	c := new(AppConfig)
	err = svenViper.Parse(c)
	if err != nil {
		fmt.Println("err:", err)
	}
    fmt.Println(c.Database.Connection)