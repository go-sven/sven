package base

import "testing"

func TestPath(t *testing.T) {
	gitUrl := "https://gitee.com/gosven/sven-layout.git"
	t.Log(svenHome())                                 // ~/.sven
	t.Log(svenHomeWithDir("repo/"))                   // ~/.sven/repo
	t.Log(repoDir(gitUrl))                            // gitee.com/gosven
	t.Log(svenHomeWithDir("repo/" + repoDir(gitUrl))) // ~/.sven/repo/gitee.com/gosven

}
