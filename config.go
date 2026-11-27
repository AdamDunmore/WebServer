package main

import (
	"os"
	"log"
)

func checkConfig() []*string {
	const LEN = 4;

	var envs = [LEN]string{	
		"WEBSERVER_PASSWORD",
		"WEBSERVER_IP",
		"WEBSERVER_PORT",
		"WEBSERVER_MUSIC_DIR",
	}

	var globals = [LEN]*string{
		&password,
		&ip,
		&port,
		&music_dir,
	}

	var defaults = [LEN]string{
		"",
		"localhost",
		"1913",
		os.Getenv("HOME") + "/Music",
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

	return globals[:];
}

