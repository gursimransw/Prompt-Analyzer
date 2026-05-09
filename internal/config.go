package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"` //Creating a struct HTTPServer that will store server address information
}

type Config struct {
	Env           string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	PromptLibrary string `yaml:"prompt_library" prompt_library-required:"true" prompt_library-default:"config/prompts/patterns.json"`
	HTTPServer    `yaml:"http_server"`
}

//Here we have created a struct called Config , an instance of this struct will be used to store configurations
//from the YAML Config file , so that we can reference all those configuration variables in our program easily.

//Flag Usage - In the struct definition, we are also defining how the configurations from the YAML file will we loaded into the struct instance.

//Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
// yaml:"env" → reads value from env key in the YAML config file
// env:"ENV" → can override value using environment variable ENV
// env-required:"true" → program will fail if no value is provided
// env-default:"production" → uses "production" if nothing is set

// yaml:"prompt_library" → reads value from prompt_library in config file
// env-required:"true" → must be provided, otherwise program exits

//HTTPServer `yaml:"http_server"`
// yaml:"http_server" → maps nested config under http_server into this struct

func MustLoad() *Config { //This function will be responsible for opening, reading and loading config from the config file to struct.
	var configPath string //Setting a variable for configpath

	configPath = os.Getenv("CONFIG_PATH") //This will try to get the config path from environment variables

	//This step will check if we got config path from env variables or not, if not we try alternate methods to get that
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file") //This tries to retrive config file from CLI Args (flags) like - main.go -config=config.yaml
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set") //Program crashes if it is still empty
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist %s", configPath)
	} //Program also crashes if path is give but it is unable to find the file , i.e file did not exist in the path

	var cfg Config //Creating an instance of config , so that we can loadin variables from config file into this.

	err := cleanenv.ReadConfig(configPath, &cfg) //Reading config file and loading it into address location of cfg struct
	if err != nil {

		log.Fatalf("Unable to read config file - %s", err.Error())

	}

	return &cfg //Return cfg pointer
}
