package runcommand

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Target struct {
	Name       string `yaml:"name"`
	SecretId   string `yaml:"secretid"`
	SecretKey  string `yaml:"secretkey"`
	BucketName string `yaml:"bucketname"`
	Location   string `yaml:"location"`
}

type Config struct {
	CurrentTarget *Target            `yaml:"current_target"`
	Targets       map[string]*Target `yaml:"targets"`
}

func Update(target, secretid, secretkey, bucketname, location string) error {
	var config *Config
	var err error

	config, err = Load()
	//出错则代表读取文件失败，需要新创建
	if err != nil {
		config = &Config{
			&Target{},
			make(map[string]*Target),
		}
	}
	currentTarget := &Target{
		Name:       target,
		SecretId:   secretid,
		SecretKey:  secretkey,
		BucketName: bucketname,
		Location:   location,
	}

	config.CurrentTarget = currentTarget
	config.Targets[target] = currentTarget

	configContents, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filepath.Join(os.Getenv("HOME"), ".copy-pastarc"), configContents, 0666)
	return nil
}
func Load() (*Config, error) {
	var config *Config

	byteContent, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".copyy"))
	if err != nil {
		return nil, fmt.Errorf("Unable to load the targets, please check if ~/.copyy exists %s", err.Error())
	}
	err = yaml.Unmarshal(byteContent, &config)
	if err != nil {
		return nil, fmt.Errorf("Parsing failed %s", err.Error())
	}
	return config, err
}
