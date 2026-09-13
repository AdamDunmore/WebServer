package main

import (
	"os"
	"log"
)

func checkConfig(){
	var envs = [3]string{	
		"WEBSERVER_PASSWORD",
		"WEBSERVER_IP",
		"WEBSERVER_PORT",
	}

	var globals = [3]*string{
		&password,
		&ip,
		&port,
	}

	// ip := "100.99.196.79"	
	var defaults = [3]string{
		"",
		"localhost",
		"1913",
	}

	for i, env := range envs {
		env_str := os.Getenv(env)
		if env_str == "" {
			log.Println("WARNING: environment variable " + envs[i] + " is not set")
			*globals[i] = defaults[i]
		} else {
			*globals[i] = env_str
		}
	}	
}

