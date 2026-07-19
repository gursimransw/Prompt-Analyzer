package loader

//This package will be responsible for loading the prompt library json file into the detection logic.

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/gursimransw/BearBreach/services/trust-guard/internal/types"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"` //Creating a struct HTTPServer that will store server address information
}

type Config struct {
	Env                  string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	DetectionRuleLibrary string `yaml:"detection_rule_library" detection_rule_library-required:"true" detection_rule_library-default:"config/rules/rules.json"`
	PolicyConfig         string `yaml:"policy_config" policy_config-required:"true" policy_config-default:"config/policy/config.json"`
	HTTPServer           `yaml:"http_server"`
}

//Here we have created a struct called Config , an instance of this struct will be used to store configurations
//from the YAML Config file , so that we can reference all those configuration variables in our program easily.

//Flag Usage - In the struct definition, we are also defining how the configurations from the YAML file will we loaded into the struct instance.

//Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
// yaml:"env" → reads value from env key in the YAML config file
// env:"ENV" → can override value using environment variable ENV
// env-required:"true" → program will fail if no value is provided
// env-default:"production" → uses "production" if nothing is set

// yaml:"detection_rule_library" → reads value from detection_rule_library in config file
// env-required:"true" → must be provided, otherwise program exits

// yaml:"policy_config" → reads value from policy_config in config file
// env-required:"true" → must be provided, otherwise program exits

//HTTPServer `yaml:"http_server"`
// yaml:"http_server" → maps nested config under http_server into this struct

func LoadDetectionRules(fileName string) (*[]types.DetectionRule, error) {

	//This functions simply takes file path as an argument, this is the filepath where the detection rule library is saved.
	//The function returns the pointer of array of DetectionRule struct , this struct is like a structure for a detection rule which contains both detection logic and metadata
	//(Refer internal/types/types.go for DetectionRule struct)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//Reading the file for file contents (Detection Rules JSON File in this case)

	var detectionRules []types.DetectionRule
	//Variable declaration for storing JSON file content

	err = json.Unmarshal(data, &detectionRules)
	if err != nil {
		return nil, err
	}
	//Here we are loading the contents from the "data", which is a string of bytes and we are storing that data in the address location of
	//config variable of type (array of DetectionRule) , that we declared above.

	return &detectionRules, nil //We are returning pointer of detectionRule array here

}

func LoadPolicyConfig(fileName string) (*types.PolicyConfig, error) {

	//Just like above, this function loads the Policy Configuration into the program, this policy configuration
	//consists of both severity configuration and actions configuration. Here we are loading the whole policy into
	//Variable of type PolicyConfig

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//Reading the file for file contents (Prompt Library JSON File in this case)

	var policyConfig types.PolicyConfig
	//Variable declaration for storing JSON file content

	err = json.Unmarshal(data, &policyConfig)
	if err != nil {
		return nil, err
	}
	return &policyConfig, nil //We are returning pointer of policyConfig here

}

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
