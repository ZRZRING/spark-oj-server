package service

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

type judgeTestCase struct {
	InputPath  string
	OutputPath string
}

// CollectTestCases 加载测试用例路径
func CollectTestCases(problemId int) ([]*judgeTestCase, error) {
	uploadPath := g.Cfg().MustGet(gctx.New(), "upload.path.testcases").String()
	problemDir := filepath.Join(uploadPath, gconv.String(problemId))
	if !gfile.Exists(problemDir) {
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "测试用例不存在")
	}

	inputFiles, err := gfile.ScanDirFile(problemDir, "*.in", false)
	if err != nil {
		return nil, err
	}
	if len(inputFiles) == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "未找到输入测试用例")
	}
	sort.Strings(inputFiles)

	testCases := make([]*judgeTestCase, 0, len(inputFiles))
	for _, inputFile := range inputFiles {
		outputFile := strings.TrimSuffix(inputFile, ".in") + ".out"
		if !gfile.Exists(outputFile) {
			return nil, gerror.NewCode(gcode.CodeInvalidRequest, "测试用例缺少输出文件")
		}

		testCases = append(testCases, &judgeTestCase{
			InputPath:  inputFile,
			OutputPath: outputFile,
		})
	}

	return testCases, nil
}
