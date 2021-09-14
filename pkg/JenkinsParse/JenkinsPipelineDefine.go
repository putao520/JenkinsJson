package JenkinsParse

import (
	"ScriptHelper/pkg/JenkinsParse/JenkinsDirective"
	"github.com/pkg/errors"
	"github.com/putao520/ourjson"
	"strings"
)

// 替换 root.xxx.xxx 定义到预定义模型
func preBuildJson(root *ourjson.JsonObject, j *ourjson.JsonObject, value string) error {
	lvArr := strings.Split(value, ".")
	switch lvArr[0] {
	case "root":
		temp := root
		for _, v := range lvArr[1:] {
			temp = temp.GetJsonObject(v)
		}
		j = temp
		j.Replace(temp)
	}
	return nil
}

// 预处理定义
func preBuildDefine(root *ourjson.JsonObject, j *ourjson.JsonObject) error {
	for _, v := range j.Values(){
		switch v.Data().(type){
		case ourjson.JsonObject:
			preBuildDefine(root, v.JsonObject())
		case string:
			str, err := v.String()
			if err == nil {
				props, pErr := JenkinsDirective.New(str)
				if pErr == nil {
					switch props.Name {
					case "json":
						preBuildJson(root, j, props.Value)
					}
				}
			}

		}
	}
	return nil
}

func buildJenkinsDefine() (*ourjson.JsonObject, error){
	pipelineModel := `{
  "pipeline": {
    "agent": {
      "any": "",
      "none": "",
      "kubernetes": {
        "defaultContainer": "",
        "yaml": "",
        "yamlFile": ""
      },
      "docker": {
        "reuseNode": "bool",
        "image": "",
        "label": "",
        "args": "",
        "registryUrl": "",
        "registryCredentialsId": "",
        "customWorkspace": ""
      },
      "dockerfile": {
        "reuseNode": "bool",
        "filename": "",
        "dir": "",
        "label": "",
        "additionalBuildArgs": "",
        "args": "",
        "customWorkspace": ""
      },
      "node": {
        "label": "",
        "customWorkspace": ""
      },
      "label": ""
    },
    "stages": {
      "stage": {
        "when": {
          "whenQualifiers": "",
          "whenPredicate": "",
          "whenCondition": ""
        },
        "failFast": "bool",
        "options": {
          "*": ""
        },
        "agent": "#json:root.agent#",
        "environment": "=",
        "input": {
          "message": "",
          "submitter": "",
          "id": ""
        },
        "tools": {
          "IDENT": ""
        },
        "steps": {
          "*": ""
        },
        "parallel": {
          "*": ""
        },
        "stages": "self"
      }
    },
    "environment": "=",
    "options": {
      "*": ""
    },
    "parameters": {
      "*": ""
    },
    "post": {
      "always": {
        "*": ""
      },
      "changed": {
        "*": ""
      },
      "fixed": {
        "*": ""
      },
      "regression": {
        "*": ""
      },
      "aborted": {
        "*": ""
      },
      "failure": {
        "*": ""
      },
      "success": {
        "*": ""
      },
      "unstable": {
        "*": ""
      },
      "unsuccessful": {
        "*": ""
      },
      "cleanup": {
        "*": ""
      }
    },
    "tools": {
      "IDENT": ""
    },
    "triggers": {
      "*": ""
    }
  }
}`
	model, err := ourjson.ParseObject(pipelineModel)
	if err != nil {
		return nil, errors.Errorf("define model error")
	}
	return model, nil
}
