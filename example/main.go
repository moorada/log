package main

import (
	"fmt"

	"github.com/moorada/log"
)

func main() {

	config := log.FormatConfigBasic
	config.Format = "{level:name} {message}"
	err := log.AddOutput("logs_test1.txt", log.INFO, config, true)
	if err != nil {
		panic(err)
	}

	config2 := log.FormatConfigBasic
	config2.Format = "{time} {level:color}{level:name}{reset} {message}"
	err = log.AddOutput("", log.INFO, config2, false)
	if err != nil {
		panic(err)
	}

	config3 := log.FormatConfigBasic
	config3.Format = "{level:name}    {message}"
	err = log.AddOutput("logs_test2.txt", log.ERROR, config3, true)
	if err != nil {
		panic(err)
	}

	//log.CloseOutputs()

	log.Info("Hello world! %s", "ciao1")
	log.Debug("Hello world! %s", "ciao1")
	log.Error("Hello world! %s", "ciao1")
	log.Raw("Hello world! %s", "Raw1")

	_ = log.RemoveOutput("logs_test1.txt")
	log.Info("Hello world! %s", "ciao2")
	log.Debug("Hello world! %s", "ciao2")
	log.Error("Hello world! %s", "ciao2")
	log.Raw("Hello world! %s", "Raw2")

	_ = log.RemoveOutput("logs_test2.txt")
	log.Info("Hello world! %s", "ciao3")
	log.Debug("Hello world! %s", "ciao3")
	log.Error("Hello world! %s", "ciao3")
	log.Raw("Hello world! %s", "Raw3")

	err = log.RemoveOutput("")
	if err != nil {
		fmt.Println(err)
	}
	log.Info("Hello world! %s", "ciao4")
	log.Debug("Hello world! %s", "ciao4")
	log.Error("Hello world! %s", "ciao4")
	log.Raw("Hello world! %s", "Raw4")

}
