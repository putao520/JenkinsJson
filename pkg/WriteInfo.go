package ScriptHelper

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type ProjectConfig struct {
	DockerRepository 	string
	Port				string
}

func WriteScript(name string, version string){
	// 读取本地配置文件
	var conf ProjectConfig
	_, err := toml.DecodeFile("project.toml", &conf)
	if err != nil {
		// handle error
		fmt.Println("project.toml Not Found")
	}

	do(name, version, conf.DockerRepository, conf.Port)
}

func do(name string, version string, dockerRepository string, port string){
	// 更新 docker 文件
	/*
	dockerContent, _ := readDockerfile()
	if dockerContent != nil {
		if _, ok := dockerContent["EXPOSE"]; ok {
			dockerContent["EXPOSE"] = port
		}

		writeDockerfile(dockerContent)
	}
	*/

	// 更新 jenkins 文件
	jenkinsEnv, _ := readJenkinsfile()
	writeJenkinsfile(jenkinsEnv)
	// 更新 helm 文件
}