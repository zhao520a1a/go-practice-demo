package main

//枚举从 1 开始, 除非使用零值是有意义的，如:当零值是理想的默认行为时。
type Operation int
const (
	Add Operation = iota + 1
	Subtract
	Multiply
)
// Add=1, Subtract=2, Multiply=3

type LogOutput int
const (
	LogToStdout LogOutput = iota
	LogToFile
	LogToRemote
)
// LogToStdout=0, LogToFile=1, LogToRemote=2