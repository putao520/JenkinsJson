package ScriptHelper

import (
	"ScriptHelper/pkg/JenkinsParse"
	"github.com/putao520/ourjson"
)

func readJenkinsfile() (*ourjson.JsonObject, error) {
	r, _ := JenkinsParse.FromFile("./test_file/Jenkinsfile")

	return r, nil
}

func writeJenkinsfile(jenkinsEnvs *ourjson.JsonObject) {
}