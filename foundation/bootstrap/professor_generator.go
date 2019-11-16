package bootstrap

import (
	"fmt"
	"io/ioutil"
	"lanvard/config"
	"lanvard/foundation"
	"lanvard/interface/decorator"
	"lanvard/support"
	"regexp"
)

type ProfessorGeneratorStruct struct {
	environmentVariables map[string]string
}

func ProfessorGenerator() decorator.Bootstrap {
	return ProfessorGeneratorStruct{}
}

func (l ProfessorGeneratorStruct) Bootstrap(app foundation.Application) foundation.Application {
	for _, professor := range config.Professors {
		file := professor.Target.Path()
		content := contentByPath(file)

		replaceVars(content, professor.Vars)
		fmt.Println(content)
	}

	return app
}

func contentByPath(file string) string {
	fileObject, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}

	return string(fileObject)
}

func replaceVars(content string, vars config.Vars) {

	content = pushImport(vars, content)
	content = replaceType(vars, content)

	fmt.Println(content)
}

func replaceType(vars config.Vars, content string) string {
	for typeName, targetVar := range vars {
		re := regexp.MustCompile("type " + typeName + " .*")
		content = re.ReplaceAllString(content, "type "+typeName+" "+support.Name(targetVar))
	}
	return content
}

func pushImport(vars config.Vars, content string) string {
	for _, importName := range uniqueImports(vars) {
		importRegex := regexp.MustCompile(`import \(`)
		content = importRegex.ReplaceAllString(content, "import (\n    \""+importName+"\"")
	}
	return content
}

func uniqueImports(intSlice config.Vars) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range intSlice {
		importName := support.Package(entry)
		if _, value := keys[importName]; !value {
			keys[importName] = true
			list = append(list, importName)
		}
	}
	return list
}
