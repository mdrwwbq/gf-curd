package src

import (
	"context"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"strings"
)

type Handle struct {
	ModuleName        string          `json:"moduleName"`
	ToolsName         string          `json:"toolsName"`
	OutputPath        string          `json:"outputPath"`
	Cover             bool            `json:"cover"`
	Ctx               context.Context `json:"ctx"`
	curApiPath        string          // 当前api目录
	curI18nPath       string          // 当前i18n目录
	curControllerPath string          // 当前controller目录
	curLogicPath      string          // 当前logic目录
	utilityPath       string          // 当前utility函数目录
}

var (
	defaultOutputPath = "./output"
	defaultCover      = false
)

func NewHandle(ctx context.Context) *Handle {
	return &Handle{}
}
func (h *Handle) Run() {
	if h.ToolsName == "" {
		log.Println("参数-t 不能为空")
		os.Exit(0)
		//log.Fatal("参数-t 不能为空")
	}
	// 初始化模板文件

	// 初始化目录
	h.initDir()
	// 写API文件
	h.WriteApi()
	// write i18n
	h.WriteI18n()
	// write controller
	h.WriteController()
	// write logic
	h.WriteLogic()
	// write utility
	h.WriteUtility()
}

// WriteApi write api
func (h *Handle) WriteApi() {
	apiTemplatePath := "./template/api.template"
	outFile := fmt.Sprintf("%s/%s.go", h.curApiPath, h.getSnakeToolsName())
	moduleName := "v1"
	if h.ModuleName != "" {
		moduleName = h.ModuleName
	}
	h.WriteLastResult(apiTemplatePath, outFile, []string{
		"{{ModuleName}}",
		"{{UpperToolsName}}",
		"{{LowerToolsName}}",
		"{{ToolsName}}",
	}, []string{
		moduleName,
		h.getUpperToolsName(),
		h.getLowerToolsName(),
		h.ToolsName,
	})
}

// WriteI18n write language
func (h *Handle) WriteI18n() {
	i18nTemplatePath := "./template/i18n.template"
	outFile := fmt.Sprintf("%s/%s.toml", h.curI18nPath, h.getSnakeToolsName())
	h.WriteLastResult(i18nTemplatePath, outFile, []string{
		"{{UpperModuleName}}",
		"{{UpperToolsName}}",
	}, []string{
		h.getUpperModuleName(),
		h.getUpperToolsName(),
	})
}

// WriteController write controller
func (h *Handle) WriteController() {
	apiTemplatePath := "./template/controller.template"
	outFile := fmt.Sprintf("%s/%s.go", h.curControllerPath, h.getSnakeToolsName())
	moduleName := "v1"
	packageName := "controller"
	if h.ModuleName != "" && h.ModuleName != moduleName {
		moduleName = h.ModuleName
		packageName = h.ModuleName
	}
	h.WriteLastResult(apiTemplatePath, outFile, []string{
		"{{PACKAGE}}",
		"{{ModuleName}}",
		"{{UpperToolsName}}",
		"{{LowerToolsName}}",
		"{{ToolsName}}",
	}, []string{
		packageName,
		moduleName,
		h.getUpperToolsName(),
		h.getLowerToolsName(),
		h.ToolsName,
	})
}

// WriteLogic write logic
func (h *Handle) WriteLogic() {
	apiTemplatePath := "./template/logic.template"
	outFile := fmt.Sprintf("%s/%s.go", h.curLogicPath, h.getSnakeToolsName())
	moduleName := "v1"
	if h.ModuleName != "" {
		moduleName = h.ModuleName
	}
	h.WriteLastResult(apiTemplatePath, outFile, []string{
		"{{ModuleName}}",
		"{{UpperToolsName}}",
		"{{LowerToolsName}}",
		"{{ToolsName}}",
		"{{UpperModuleName}}",
	}, []string{
		moduleName,
		h.getUpperToolsName(),
		h.getLowerToolsName(),
		h.ToolsName,
		h.getUpperModuleName(),
	})
}

// WriteUtility write utility
func (h *Handle) WriteUtility() {
	// 写入resp通用数据
	apiTemplatePath := "./template/utility.template"
	outFile := fmt.Sprintf("%s/resp.go", h.utilityPath)
	h.WriteLastResult(apiTemplatePath, outFile, []string{}, []string{})

	// 写入语言包
	apiTemplatePath = "./template/language.template"
	outFile = fmt.Sprintf("%s/language.go", h.utilityPath)
	h.WriteLastResult(apiTemplatePath, outFile, []string{}, []string{})

	// 写入通用db函数
	apiTemplatePath = "./template/db.template"
	outFile = fmt.Sprintf("%s/db.go", h.utilityPath)
	h.WriteLastResult(apiTemplatePath, outFile, []string{}, []string{})

	// 写入consts
	apiTemplatePath = "./template/consts.template"
	outFile = fmt.Sprintf("%s/consts.go", h.utilityPath)
	h.WriteLastResult(apiTemplatePath, outFile, []string{}, []string{})

}

// checkCover 检查是否覆盖写入
func (h *Handle) checkCover(path string) {
	if h.Cover == false {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("文件“%s”已经存在，如果想覆盖写入请使用 “-c=true”", path))
		}
	}
}

// WriteLastResult 最终写入文件内容
func (h *Handle) WriteLastResult(templateFile, outFile string, oldSlice, newSlice []string) {
	template, err := h.getTemplate(templateFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	curTemplate := h.batchReplaceAll(template, oldSlice, newSlice)
	// 检查是否覆盖写入
	h.checkCover(outFile)
	err = os.WriteFile(outFile, []byte(curTemplate), 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(outFile, "success")
}

// batchReplaceAll 批量替换功能
// 有点像php str_replace 中的str_replace(["A","B"],["C","D"], "Content")
func (h *Handle) batchReplaceAll(s string, old []string, new []string) string {
	result := s
	for index, oldVal := range old {
		newVal := ""
		if index >= 0 && index < len(new) {
			newVal = new[index]
		}
		// 替换字符串
		result = strings.ReplaceAll(result, oldVal, newVal)
	}
	return result
}

// getUpperToolsName get upper camel tools name
func (h *Handle) getUpperToolsName() string {
	return strcase.ToCamel(h.ToolsName)
}

// getLowerToolsName get lower camel tools name
func (h *Handle) getLowerToolsName() string {
	return strcase.ToLowerCamel(h.ToolsName)
}

// getSnakeToolsName get snake tools name
func (h *Handle) getSnakeToolsName() string {
	return strcase.ToSnake(h.ToolsName)
}

// getUpperModuleName get upper camel tools name
func (h *Handle) getUpperModuleName() string {
	return strcase.ToCamel(h.ModuleName)
}

// getLowerModuleName get lower camel tools name
func (h *Handle) getLowerModuleName() string {
	return strcase.ToLowerCamel(h.ModuleName)
}

// getSnakeModuleName get snake tools name
func (h *Handle) getSnakeModuleName() string {
	return strcase.ToSnake(h.ModuleName)
}

// getTemplate 获得指定模板文件内容
func (h *Handle) getTemplate(path string) (string, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		con := GetDefaultTemplateContent(path)
		if len(con) == 0 {
			return "", errors.New(fmt.Sprintf(`模板文件“%s”内容为空，或模板文件不存在`, path))
		}
		return GetDefaultTemplateContent(path), nil
		// return "", errors.New(fmt.Sprintf("“%s”文件不存在", path))
	}
	contentByte, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(contentByte), nil
}

// initDir 初始化目录
func (h *Handle) initDir() {
	curPath := defaultOutputPath
	if h.OutputPath == "" {
		curPath = h.OutputPath
	}
	// 创建根据目录
	h.mkDirHandle(curPath)
	// api/basic
	i18nModelName := h.ModuleName
	moduleName := h.ModuleName
	// api的目录名，是一个特殊需要处理的目录，如果没有h.ModuleName
	if h.ModuleName == "" {
		i18nModelName = "v1"
		moduleName = "/"
	}
	apiDir := fmt.Sprintf("%s/api/%s", curPath, i18nModelName)
	// 创建api目录
	h.mkDirHandle(apiDir)
	h.curApiPath = apiDir

	// i18n目录
	i18nDir := fmt.Sprintf("%s/i18n/zh-CN/%s", curPath, h.ModuleName)
	// 创建i18n目录
	h.mkDirHandle(i18nDir)
	h.curI18nPath = i18nDir

	// controller目录
	controllerDir := fmt.Sprintf("%s/internal/controller/%s", curPath, moduleName)
	// 创建controller目录
	h.mkDirHandle(controllerDir)
	h.curControllerPath = controllerDir

	// logic目录
	logicDir := fmt.Sprintf("%s/internal/logic/%s", curPath, h.ToolsName)
	// 创建logic目录
	h.mkDirHandle(logicDir)
	h.curLogicPath = logicDir

	// utility目录
	utilityDir := fmt.Sprintf("%s/utility/", curPath)
	// 创建utility目录
	h.mkDirHandle(utilityDir)
	h.utilityPath = utilityDir
}

// mkDirHandle 创建目录
func (h *Handle) mkDirHandle(path string) {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
