package JenkinsParse

import (
	"ScriptHelper/pkg/JenkinsParse/JenkinsDirective"
	"github.com/putao520/ourjson"
	"io/ioutil"
	"strings"
)

// 扫描block
func scanBlock(model *ourjson.JsonObject, lines []string, directives *ourjson.JsonObject) error{
	_directive := ""
	for _,_line := range lines {
		line := strings.TrimLeft(_line, " ")
		if line == "}" {
			_directive = ""
			continue
		}
		directive, content, ok := scanDirective(line, directives)
		if ok {
			// 上个指令结束
			if len(_directive) > 0 {
				content, err := model.GetString(_directive)
				if err == nil {
					_model := ourjson.New()
					if scanBlock(_model, strings.Split(content, "\n"), directives.GetJsonObject(_directive) ) == nil {
						model.Put(_directive, _model)
					}
				}
			}
			// 进入下个指令解析
			_directive = directive
			model.Put(directive, content)
		} else {
			if len(_directive) > 0 {
				// 补充内容
				_content,err := model.GetString(_directive)
				if err != nil {
					continue
				}
				model.Put(_directive, _content + _line)
			}
		}
	}
	return nil
}

func scanDirective(line string, directives *ourjson.JsonObject) (string, string, bool){
	if len(line) == 0 {
		return "", line, false
	}
	sep := " "
	for directive, _struct := range directives.Values(){
		// 处理 sep 更新
		vStr, vErr := _struct.String()
		if vErr == nil {
			props, pErr := JenkinsDirective.New(vStr)
			if pErr == nil {
				switch props.Name {
				case "split":
					sep = props.Value
				}
			}
		}
		segmentArr := strings.Split(line, sep)
		if len(segmentArr) <= 1 {
			continue
		}
		// 找到符合的指令
		if strings.HasPrefix( strings.ToLower(segmentArr[0]), directive ) {
			return directive, strings.Join(segmentArr[1:], " "), true
		}
	}
	return "", line, false
}

// FromFile ---------------------------------------------------------
func FromFile(path string) (*ourjson.JsonObject, error){
	jf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return from(string(jf))
}

func from(data string) (*ourjson.JsonObject, error) {
	modelDefine, err := buildJenkinsDefine()
	if err != nil {
		return nil, err
	}
	model := ourjson.New()
	lines := strings.Split(data, "\n")

	scanErr := scanBlock(model, lines, modelDefine)
	if scanErr != nil {
		return nil, scanErr
	}

	return model, nil
}