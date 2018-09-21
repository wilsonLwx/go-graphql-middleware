package logger

type FileLogger struct{
	level int
	logPath string
	logName string
}

func NewFileLogger(level int, logPath, logName string) (LogInterface){
	
}