package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rlog "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

//tabFormatter tab数据格式化
type tabFormatter struct {
	log.TextFormatter
}

// Format 自定义日志输出格式
func (c *tabFormatter) Format(entry *log.Entry) ([]byte, error) {
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s:%d", fileName, frame.Line)
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s] %s", entry.Time.Format(c.TimestampFormat), // 输出日志时间
		strings.ToUpper(entry.Level.String())))
	if entry.HasCaller() {
		b.WriteString(fmt.Sprintf("[%s]", prettyCaller(entry.Caller))) // 输出日志所在文件，行数位置
	}
	b.WriteString(fmt.Sprintf(" %s\n", entry.Message)) // 输出日志内容
	return b.Bytes(), nil
}

// InitLog 初始化日志
func InitLog() {
	appConf := GetGlobalConf()
	// 设置日志级别
	level, err := log.ParseLevel(appConf.Log.Level)
	if err != nil {
		panic("log level parse err:" + err.Error())
	}
	// 设置日志格式为json格式
	log.SetFormatter(&tabFormatter{
		log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		}})
	//log.SetReportCaller(false)
	log.SetLevel(level)

	switch appConf.Log.LogPattern {
	case "stdout":
		log.SetOutput(os.Stdout)
		setGinLog(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
		setGinLog(os.Stderr)
	case "file":
		logger, err := rlog.New(appConf.Log.LogPath+".%Y%m%d%H%M",
			rlog.WithLinkName(appConf.Log.LogPath),
			rlog.WithRotationCount(appConf.Log.SaveDays),
			rlog.WithRotationTime(time.Hour*24),
		)
		if err != nil {
			panic("log conf err " + err.Error())
		}
		log.SetOutput(logger)
		setGinLog(logger)
	default:
		panic("log conf err, check log_pattern in config.yaml")
	}
	log.Infof("api Conf %#v", appConf)
}

func setGinLog(out io.Writer) {
	gin.DefaultWriter = out
	gin.DefaultErrorWriter = out
}
