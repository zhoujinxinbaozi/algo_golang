/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/12/11 3:42 下午
 */

package basic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/aviddiviner/go-murmur"
)

func TestExp(t *testing.T) {
	lines := readFile("/Users/zhoujinxin/Desktop/common_kv")
	if len(lines) == 0 {
		fmt.Println("hashArr is 0")
		return
	}

	hashArr := hashArr(lines)
	lines = readFile("/Users/zhoujinxin/Desktop/cuid")
	cuid2hashMap := converMap(lines)
	judge(hashArr, cuid2hashMap)
}

func judge(hashArr []string, cuid2hashMap map[string]string) {
	for cuid, hash := range cuid2hashMap {
		for _, v := range hashArr {
			if v != hash {
				continue
			}
			fmt.Println(cuid)
		}
	}
}

func ComputeBalanceHash(uniqKey string) uint64 {
	return murmur.MurmurHash64A([]byte(uniqKey), 0) % 994009
}

// key cuid
// val hash(cuid)
func converMap(lines []string) map[string]string {
	var result map[string]string
	result = make(map[string]string, len(lines))
	for _, v := range lines {
		result[v] = strconv.Itoa(int(ComputeBalanceHash(v)))
	}
	return result
}

// 过滤sid
func hashArr(lines []string) []string{
	var result []string
	for _, v := range lines {
		if strings.Index(v, "7034_13901") != -1 {
			result = append(result, strings.Split(v,"\t")[0])
		}
	}
	return result
}
// 按行读取文件
func readFile(filePath string) []string {
	var result []string
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		result = append(result, string(a))
	}
	return result
}

