package main

import "spider/initlize"

func main() {
	initlize.InitMysql()
	initlize.InitEs()
	initlize.InitMq()
}
