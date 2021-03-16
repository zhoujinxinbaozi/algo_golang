/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/3/25 1:48 下午
 */

package yaml

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"

	"github.com/creasty/defaults"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `yaml:"name"`
	Age  int `yaml:"age"`
	Year int `default:"-1" yaml:"year"`
}

func TestYaml(t *testing.T) {
	conf := "./user.yaml"
	b, err := ioutil.ReadFile(conf)
	assert.NoError(t, err)
	var user User
	yaml.Unmarshal(b, &user)
	bb, err := json.Marshal(user)
	assert.NoError(t, err)
	fmt.Println(string(bb))

	conf = "./user_copy.yaml"
	b, err = ioutil.ReadFile(conf)
	assert.NoError(t, err)
	var userCopy User
	yaml.Unmarshal(b, &userCopy)
	bb, err = json.Marshal(userCopy)
	assert.NoError(t, err)
	fmt.Println(string(bb))
}

// UnmarshalYAML 实现该方法，在yaml未配置的项 会使用默认值替代 比如year为配置是-1 并不是0
func (s *User) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(s)

	type plain User
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	return nil
}
