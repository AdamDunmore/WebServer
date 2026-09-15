package main

import (
	"os"
	"log"
)

func checkConfig(){
	var envs = [4]string{	
		"WEBSERVER_PASSWORD",
		"WEBSERVER_IP",
		"WEBSERVER_PORT",
		"WEBSERVER_LIDARR_APIKEY",
	}

	var globals = [4]*string{
		&password,
		&ip,
		&port,
		&lidarr_apikey,
	}

	var defaults = [4]string{
		"",
		"localhost",
		"1913",
		"",
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

