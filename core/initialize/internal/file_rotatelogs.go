package internal

import (
	"github.com/SliverFlow/ksmall/core/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *fileRotatelogs) GetWriteSyncer(level string, c config.Zap) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(c.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(c.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if c.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
