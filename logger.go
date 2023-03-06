package gosms

import "log"

var DefaultName = "GoTool"

var _ ILogger = (*Logger)(nil)

type Logger struct {
	Log *log.Logger
}

func (l Logger) I() ILogger {
	return &l
}
func (l *Logger) Error(format string, v ...any) {
	l.output("ERROR "+DefaultName+" "+format, v...)
}
func (l *Logger) Warn(format string, v ...any) {
	l.output("WARN "+DefaultName+" "+format, v...)
}

func (l *Logger) Debug(format string, v ...any) {
	l.output("DEBUG "+DefaultName+" "+format, v...)
}
func (l *Logger) output(format string, v ...any) {
	if len(v) == 0 {
		l.Log.Print(format)
		return
	}
	l.Log.Printf(format, v...)
}
