package core

import (
	"bytes"
	"fmt"
	"gvb/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)
//对于log的封装
type LogFormatter struct {
}

// Formatter实现对LogFormatter的接口
func (t *LogFormatter) Format(entry *logrus.Entry)([]byte ,error) {
//根据不同的level去展示颜色
var levelcolor int
 switch entry.Level{
 case logrus.DebugLevel,logrus.TraceLevel:
	levelcolor=gray
 case logrus.WarnLevel:
	levelcolor=yellow
 case logrus.ErrorLevel,logrus.FatalLevel,logrus.PanicLevel:
	levelcolor=red
 default:
	levelcolor=blue
 }
var b *bytes.Buffer
if entry.Buffer!=nil {
	b=entry.Buffer
}else{
	b=&bytes.Buffer{}
}
//定义日期格式
log:=global.Config.Logger
timestamp:=entry.Time.Format("2006/01/02-15:04:05")
if entry.HasCaller() {
	//自定义文件路径
	funcval:=entry.Caller.Function
	fileval:=fmt.Sprintf("%s:%d",path.Base(entry.Caller.File),entry.Caller.Line)
	fmt.Fprintf(b,"%s %s \x1b[%dm[%s]\x1b[0m %s %s %s\n",log.Prefix,timestamp,levelcolor,entry.Level,fileval,funcval,entry.Message)
}else{
	fmt.Fprintf(b,"%s %s \x1b[%dm[%s]\x1b[0m %s\n",log.Prefix,timestamp,levelcolor,entry.Level,entry.Message)
}
return b.Bytes(),nil
}
//log初始化
func InitLogger() *logrus.Logger {
	mlog:=logrus.New()
	mlog.SetOutput(os.Stdout)//设置输出类型
	mlog.SetReportCaller(true)//开启返回函数名和行号
	mlog.SetFormatter(&LogFormatter{})//设置自己定义的Formatter
	mlog.SetLevel(logrus.DebugLevel)
	return mlog
}
//
func InitDefaultLogger()  {
	//全局log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

}