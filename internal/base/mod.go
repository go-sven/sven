package base

// ModulePath return go module path
// 返回layout go.mod 中第一行 module 后面到内容
func ModulePath(filename string) (string, error) {
	return "github.com/go-sven/layout", nil
	//modBytes, err := os.ReadFile(filename)
	/*if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil//*/

}

// ModuleVersion returns module version.
/*func ModuleVersion(path string) (string, error) {
	stdout := &bytes.Buffer{}
	fd := exec.Command("go", "mod", "graph")
	fd.Stdout = stdout
	fd.Stderr = stdout
	if err := fd.Run(); err != nil {
		return "", err
	}
	rd := bufio.NewReader(stdout)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			return "", err
		}
		str := string(line)
		i := strings.Index(str, "@")
		if strings.Contains(str, path+"@") && i != -1 {
			return path + str[i:], nil
		}
	}
}

// SvenMod returns sven mod.
func SvenMod() string {
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}
	if path, err := ModuleVersion("github.com/go-sven/sven/v2"); err == nil {
		// $GOPATH/pkg/mod/github.com/go-sven/sven@v2
		return filepath.Join(cachePath, path)
	}
	// $GOPATH/src/github.com/go-sven/sven
	return filepath.Join(gopath, "src", "github.com", "go-sven", "sven")
}//*/
