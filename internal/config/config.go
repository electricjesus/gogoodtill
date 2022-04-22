package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

type Kv map[string]interface{}

type cfg struct {
	values Kv
}

func (c *cfg) get(k string) interface{} {
	return c.values[k]
}
func (c *cfg) set(k string, v interface{}) {
	c.values[k] = v
}

var defaultConfig *cfg

func init() {
	defaultConfig = new(cfg)
	if err := Load(); err != nil {
		panic(err)
	}
}

const (
	defaultConfigFile = "config.yaml"
	defaultConfigPath = "."
)

type KVPair struct {
	Key, Value string
}

func Load() error {
	return load(afero.NewOsFs())
}

func load(_ afero.Fs) error {
	var values Kv
	filePath := path.Join(defaultConfigPath, defaultConfigFile)
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, &values); err != nil {
		return err
	}
	defaultConfig = &cfg{values}
	return nil
}

func Save() error {
	return save(afero.NewOsFs())
}

func save(_ afero.Fs) error {
	b, err := yaml.Marshal(defaultConfig.values)
	if err != nil {
		return err
	}
	filePath := path.Join(defaultConfigPath, defaultConfigFile)
	if err := os.WriteFile(filePath, b, 0644); err != nil {
		return nil
	}
	return nil
}

func Update(kvs Kv) {
	for k, v := range kvs {
		defaultConfig.set(k, v)
	}
}

func Get(k string) interface{} {
	return defaultConfig.get(k)
}

func GetS(k string) (string, bool) {
	vv := defaultConfig.get(k)
	v, ok := vv.(string)
	return v, ok
}

func Set(k string, v interface{}) {
	defaultConfig.set(k, v)
}
