package main

import (
	"os"
	"log"
)

func checkConfig(){
	const LEN = 3;

	var envs = [LEN]string{	
		"WEBSERVER_PASSWORD",
		"WEBSERVER_IP",
		"WEBSERVER_PORT",
	}

	var globals = [LEN]*string{
		&password,
		&ip,
		&port,
	}

	var defaults = [LEN]string{
		"",
		"localhost",
		"1913",
	}

	for i, env := range envs {
		env_str := os.Getenv(env)
		if env_str == "" {
			log.Println("WARNING: environment variable " + envs[i] + " is not set. Defaulting to " + defaults[i])
			*globals[i] = defaults[i]
		} else {
			*globals[i] = env_str
		}
	}	
}

