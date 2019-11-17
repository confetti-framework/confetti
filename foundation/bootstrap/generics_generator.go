package bootstrap

import (
	"io/ioutil"
	"lanvard/config"
	"lanvard/foundation"
	"lanvard/interface/decorator"
	"lanvard/support"
	"os"
	"regexp"
	"strings"
)

type GenericGeneratorStruct struct {
	environmentVariables map[string]string
}

func GenericGenerator() decorator.Bootstrap {
	return GenericGeneratorStruct{}
}

func (l GenericGeneratorStruct) Bootstrap(app foundation.Application) foundation.Application {
	// for _, generic := range config.Generics {
	// 	file := generic.Struct.Path()
	// 	content := contentByPath(file)
	//
	// 	content = replaceDocs(content, file)
	// 	content = replacePackage(content, generic.SaveTo)
	// 	content = pushImport(content, generic.Vars)
	// 	content = replaceType(content, generic.Vars)
	//
	// 	ensureDir(filepath.Dir(generic.SaveTo))
	// 	_ = os.Remove(generic.SaveTo)
	// 	updateOrCreate(generic.SaveTo, content)
	// }

	return app
}

func updateOrCreate(file string, content string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}

	_ = f.Close()
}

func contentByPath(file string) string {
	fileObject, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(fileObject)
}

func replaceDocs(content string, originalFile string) string {
	originalRelative := strings.Replace(originalFile, config.App.BasePath, "", 1)
	re := regexp.MustCompile("This is the original file")

	return re.ReplaceAllString(content, "Original file: "+originalRelative)
}

func replacePackage(content string, saveTo string) string {
	packageName := packageNameFromPath(saveTo)

	re := regexp.MustCompile("package.*")
	return re.ReplaceAllString(content, "package "+packageName)
}

func replaceType(content string, vars config.Vars) string {
	for typeName, targetVar := range vars {
		re := regexp.MustCompile("type " + typeName + " .*")
		content = re.ReplaceAllString(content, "type "+typeName+" "+support.Name(targetVar))
	}

	return content
}

func pushImport(content string, vars config.Vars) string {
	for _, importName := range uniqueImports(vars) {
		importRegex := regexp.MustCompile(`import \(`)
		content = importRegex.ReplaceAllString(content, "import (\n    \""+importName+"\"")
	}

	return content
}

func uniqueImports(vars config.Vars) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range vars {
		importName := support.Package(entry)
		if _, value := keys[importName]; !value {
			keys[importName] = true
			list = append(list, importName)
		}
	}

	return list
}

func ensureDir(dirName string) {
	// todo use
	err := os.MkdirAll(dirName, 0755)

	if err != nil && !os.IsExist(err) {
		panic("Directory creation failed with error: " + err.Error())
	}
}

func packageNameFromPath(path string) string {
	parts := strings.Split(path, string(os.PathSeparator))
	if len(parts) < 2 {
		panic("You must provide a path that contains at least one directory and a file.")
	}

	return parts[len(parts)-2]
}
