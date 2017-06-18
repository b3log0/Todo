package main

import (
	"path/filepath"
	"io/ioutil"
	"strings"
)

//参数的filename不包含路径
func getFilePathName(filename string) string {
	return filepath.Join(current_dir,filename)
}

//参数的filename包含路径
func editDoingFunc(doingFunc func(string,[]string) error,params []string) error {
	files, _ := ioutil.ReadDir(current_dir)
	for _,value := range files{
		if strings.HasSuffix(value.Name(),doing_suffix) {
			doingFunc(getFilePathName(value.Name()),params)
		}
	}
	return nil
}

func accessGithub(){
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "2612b3f35a40a34b7ecd3506ff37158af93a519c"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err!=nil {
		fmt.Println("error")
	}
	for _,value := range repos{
		fmt.Println(value)
	}
}