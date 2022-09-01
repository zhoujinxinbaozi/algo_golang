package err

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestError(t *testing.T) {
	err := makeErr()
	if err != nil {
		fmt.Printf("err : %v", errors.Wrap(err, "msg"))
	}
}

func makeErr() error {
	return errors.New("tt")
}

type rpcInfo interface {
	From() rpcInfo1
}

type rpcInfo1 interface {
}

type rpcInfoImpl struct {
}

func (p *rpcInfoImpl) From() rpcInfo1 {
	return nil
}

const (
	key = "rpc_info"
)

func TestPanic(t *testing.T) {
	ctx := context.Background()
	var infoImpl *rpcInfoImpl = nil
	var info rpcInfo
	info = infoImpl

	ctx = context.WithValue(ctx, key, info)
	r, ok := ctx.Value(key).(rpcInfo)
	fmt.Println("ok : ", ok)
	fmt.Println("r is nil : ", r == nil)
	fmt.Println("r.value is nil : ", reflect.ValueOf(r).IsNil())
	fmt.Println("success")

	var i interface{}
	fmt.Println(reflect.ValueOf(i).IsNil())
}

type ss struct {
	Name string
}

func TestStruct(t *testing.T) {
	s1 := &ss{}
	s2 := &ss{}
	fmt.Println(*s1 == *s2)
}

func TestMap(t *testing.T) {
	m := make(map[int]bool)
	fmt.Printf("m : %v", m)
	addMap(m)
	fmt.Printf("m : %v", m)
}
func addMap(m map[int]bool) {
	m[1] = false
}

func TestMsgReq(t *testing.T) { // putMsgReq result :
	//str := "eyJUeXBlIjoxMSwiQ2hhbm5lbElEIjo3MTAxOTkxMDEyMDQ5OTkzNzI5LCJjaGFubmVsX3R5cGUiOjEsIk93bmVySUQiOjY4MjgzOTg4MjE1ODg5MjY0NjUsIkNvbnRlbnQiOnsidGV4dCI6IiIsImltYWdlX2tleSI6IiIsInRpdGxlIjoiIiwiYXR0YWNobWVudHMiOm51bGwsImlzX25vdGlmaWVkIjpmYWxzZSwibWVzc2FnZV90ZW1wbGF0ZSI6IiIsInRlbXBsYXRlX2tleV92YWx1ZSI6bnVsbCwic3lzdGVtX2NvbnRlbnRfdHlwZSI6MCwiZmlsZV9rZXkiOiIiLCJmaWxlX21pbWUiOiIiLCJmaWxlX3NpemUiOjAsImZpbGVfbmFtZSI6IiIsImF1ZGlvX2tleSI6IiIsImF1ZGlvX2R1cmF0aW9uIjowLCJhdWRpb19zaXplIjowLCJzaGFyZV9jaGF0X3Rva2VuIjoiIiwic2hhcmVfY2hhdF9pZCI6MCwidG9rZW5fZXhwaXJlX3RpbWUiOjAsIndpZHRoIjowLCJoZWlnaHQiOjAsIm1lcmdlX2NvbnRlbnQiOnsibWVzc2FnZXNfYnl0ZXMiOiJEd0FCREFBQUFBSUtBQUZpc0VYVEVVQ0FBZ2dBQWdBQUFBUUtBQU1BQUFBQVlyQkYwd29BQkY3RFdJekJnUUFCREFBRkN3QUJBQUFBQzNSbGMzUWc1cldMNksrVkN3QUNBQUFBQUFzQUF3QUFBQUFPQUFRTEFBQUFBQUlBQlFBTEFBWUFBQUFBRFFBSEN3c0FBQUFBQ0FBUEFBQUFBQXNBQ0FBQUFBQUxBQWtBQUFBQUNnQUtBQUFBQUFBQUFBQUxBQXNBQUFBQUN3QU1BQUFBQUFnQURRQUFBQUFLQUE0QUFBQUFBQUFBQUFzQUVBQUFBQUFLQUJFQUFBQUFBQUFBQUFvQUVnQUFBQUFBQUFBQUNBQVRBQUFBQUFnQUZBQUFBQUFQQUJrTUFBQUFBQXdBR3c4QUFRc0FBQUFCQUFBQUF6RXRNUXNBQWdBQUFBQUxBQU1BQUFBZUNod0tBekV0TVJJVkNBRWFFUW9MZEdWemRDRG10WXZvcjVVU0FCZ0JEd0FFQ3dBQUFBQVBBQVVMQUFBQUFBOEFCZ3NBQUFBQUR3QUhDd0FBQUFBQ0FBZ0JBQThBTlF3QUFBQUFBQXdBQmdJQUFnQUlBQW9BQUFBSEFBb0FDV0tQVnpLUFFjQUJDQUFLQUFBQUFRc0FDd0FBQUNReVpqUXlPREJoT0MxalpUZG1MVFJsWmpNdFltRmhOaTFtTWprMll6a3hOREpqTW1FS0FBMEFBQUFBWXJCRjB3SUFEZ0VLQUJCaXNFWFRFVUNBQWdvQUVXS3dSZE1SUUlBQ0NnQVNZckJGMHhGQWdBSUNBQlFBRHdBVkNnQUFBQUFLQUJkaWoxY3lqMEhBQVFnQUdBQUFBQUVLQUJ4aWoxY3lrRU9BQWdvQUhRQUFBWUdBa01DWkNnQWVBQUFCZ1lDUXdKa1BBQ0FLQUFBQUFBc0FJUUFBQUFKNmFBZ0FJd0FBQUFFSUFDUUFBQUFDQ2dBbVhzTllqTUdCQUFFS0FDcGdUdnRsQXNFQUhBQUtBQUZpc0VYVDlBQ0FBUWdBQWdBQUFBUUtBQU1BQUFBQVlyQkYxQW9BQkY3RFdJekJnUUFCREFBRkN3QUJBQUFBTm5SbGMzUWc1cmFJNW9HdlBHRjBJSFZ6WlhKZmFXUTlJbnQ3Wnk1QlZGVlRSVkl4ZlgwaVBrRGxyb3ZtbUl6cG00UThMMkYwUGdzQUFnQUFBQUFMQUFNQUFBQUFEZ0FFQ3dBQUFBQUNBQVVBQ3dBR0FBQUFBQTBBQndzTEFBQUFBQWdBRHdBQUFBQUxBQWdBQUFBQUN3QUpBQUFBQUFvQUNnQUFBQUFBQUFBQUN3QUxBQUFBQUFzQURBQUFBQUFJQUEwQUFBQUFDZ0FPQUFBQUFBQUFBQUFMQUJBQUFBQUFDZ0FSQUFBQUFBQUFBQUFLQUJJQUFBQUFBQUFBQUFnQUV3QUFBQUFJQUJRQUFBQUFEd0FaREFBQUFBQU1BQnNQQUFFTEFBQUFBZ0FBQUFNeExURUFBQUFETVMweUN3QUNBQUFBQUFzQUF3QUFBRVlLSEFvRE1TMHhFaFVJQVJvUkNndDBaWE4wSU9hMmlPYUJyeElBR0FBS0pnb0RNUzB5RWg4SUJSb2JDZzE3ZTJjdVFWUlZVMFZTTVgxOUVncEE1YTZMNXBpTTZadUVEd0FFQ3dBQUFBQVBBQVVMQUFBQUFBOEFCZ3NBQUFBQUR3QUhDd0FBQUFBQ0FBZ0JBQThBTlF3QUFBQUFBQXdBQmdJQUFnQUlBQW9BQUFBSEFBb0FDV0tQVnpLUFFjQUJDQUFLQUFBQUFnc0FDd0FBQUNSbVlqbGxNekl5TXkxak9UZ3pMVFEwWmpZdE9XSm1aaTAyTldOaVpqWXhOemswTjJFS0FBMEFBQUFBWXJCRjFBSUFEZ0VLQUJCaXNFWFQ5QUNBQVFvQUVXS3dSZFAwQUlBQkNnQVNZckJGMC9RQWdBRUNBQlFBRHdBVkNnQUFBQUFLQUJkaWoxY3lqMEhBQVFnQUdBQUFBQUVLQUJ4aWoxY3lrRU9BQWdvQUhRQUFBWUdBa01RbENnQWVBQUFCZ1lDUXhDVVBBQ0FLQUFBQUFBc0FJUUFBQUFKNmFBZ0FJd0FBQUFFSUFDUUFBQUFDQ2dBbVhzTllqTUdCQUFFS0FDcGdUdnRsQXNFQUhBQUEiLCJvcmlnaW5fY2hhdF90eXBlIjoyLCJncm91cF9jaGF0X25hbWUiOiLmma7pgJrnvqTvvIxnLlVTRVIxIOS4uue+pOS4uyIsImNoYXR0ZXJzIjp7IjY4MjgzOTg4MjE1ODg5MjY0NjUiOnsiaWQiOjY4MjgzOTg4MjE1ODg5MjY0NjUsIm5hbWUiOiJnLlVTRVIyIiwiYXZhdGFyX3Rvc19rZXkiOiJkYWZjMDAwZjg1ZmVmMzg4MWVkYSIsInR5cGUiOjEsImkxOG5fbmFtZSI6eyJlbl91cyI6ImcuVVNFUjIiLCJqYV9qcCI6IiIsInpoX2NuIjoiIn0sImNoYXRfdXNlcl90eXBlIjoxfX0sImlzX3NpbmdsZV9lbmNvZGVkIjp0cnVlLCJyZWFjdGlvbl9zbmFwc2hvdHMiOlt7Im1lc3NhZ2VfaWQiOjcxMTEyNjA1ODQ0NDc4Njg5MzAsInJlYWN0aW9ucyI6W119LHsibWVzc2FnZV9pZCI6NzExMTI2MDU4ODI1MjEwMjY1NywicmVhY3Rpb25zIjpbXX1dLCJvcmlnaW5fY2hhdF9pZCI6NzEwMTk5MTAxMjA0OTk5MzcyOSwiaTE4bl90aHJlYWRfdGl0bGUiOm51bGx9LCJ1cmxfcHJldmlld19oYW5ncG9pbnRzIjpudWxsfSwiQ2lkIjoiNTAyMTNiYTgtMmJkNy00MzljLWJiNWYtNWFiOTVjZmE5NThlIiwiTmVlZFB1c2giOnRydWUsIkhhc0Rlc2t0b3BCYWRnZSI6dHJ1ZSwiTmVlZFVwZGF0ZUZlZWQiOnRydWUsIkNoYXRNb2RlIjoxLCJDcmVhdGVPcHRpb24iOnt9LCJCYXNlIjpudWxsfQ=="
	str := "eyJUeXBlIjoxMSwiQ2hhbm5lbElEIjo3MDM2NzMwNzA0MDk2NDExNjUwLCJjaGFubmVsX3R5cGUiOjEsIk93bmVySUQiOjcwMzY1NTYyNjQwNjY5OTAwODIsIkNvbnRlbnQiOnsidGV4dCI6IiIsImltYWdlX2tleSI6IiIsInRpdGxlIjoiIiwiYXR0YWNobWVudHMiOm51bGwsImlzX25vdGlmaWVkIjpmYWxzZSwibWVzc2FnZV90ZW1wbGF0ZSI6IiIsInRlbXBsYXRlX2tleV92YWx1ZSI6bnVsbCwic3lzdGVtX2NvbnRlbnRfdHlwZSI6MCwiZmlsZV9rZXkiOiIiLCJmaWxlX21pbWUiOiIiLCJmaWxlX3NpemUiOjAsImZpbGVfbmFtZSI6IiIsImF1ZGlvX2tleSI6IiIsImF1ZGlvX2R1cmF0aW9uIjowLCJhdWRpb19zaXplIjowLCJzaGFyZV9jaGF0X3Rva2VuIjoiIiwic2hhcmVfY2hhdF9pZCI6MCwidG9rZW5fZXhwaXJlX3RpbWUiOjAsIndpZHRoIjowLCJoZWlnaHQiOjAsIm1lcmdlX2NvbnRlbnQiOnsibWVzc2FnZXNfYnl0ZXMiOiJEd0FCREFBQUFBRUtBQUZpc1Nia3ZFSUFCQWdBQWdBQUFBUUtBQU1BQUFBQVlyRW01QW9BQkdJRnlLem53TUFCREFBRkN3QUJBQUFBU09pK20raUxwdVM3aXVXa3FlYWNpZWVwdXVXR2plV2tqZWVPc09TNGkrV1FwKys4ak9Xd3NlaTlyT1dTc2VTN3JPUy9xZWVhaE9hMmlPYUJyK1dRcHlCYjZMQ2k2TENpWFFzQUFnQUFBQUFMQUFNQUFBQUFEZ0FFQ3dBQUFBQUNBQVVBQ3dBR0FBQUFBQTBBQndzTEFBQUFBQWdBRHdBQUFBQUxBQWdBQUFBQUN3QUpBQUFBQUFvQUNnQUFBQUFBQUFBQUN3QUxBQUFBQUFzQURBQUFBQUFJQUEwQUFBQUFDZ0FPQUFBQUFBQUFBQUFMQUJBQUFBQUFDZ0FSQUFBQUFBQUFBQUFLQUJJQUFBQUFBQUFBQUFnQUV3QUFBQUFJQUJRQUFBQUFEd0FaREFBQUFBQU1BQnNQQUFFTEFBQUFBZ0FBQUFFekFBQUFBVFFMQUFJQUFBQUFDd0FEQUFBQWJRcExDZ0V6RWtZSUFScENDa0RvdnB2b2k2Ymt1NHJscEtubW5Jbm5xYnJsaG8zbHBJM25qckRrdUl2bGtLZnZ2SXpsc0xIb3ZhemxrckhrdTZ6a3Y2bm5tb1RtdG9qbWdhL2xrS2NnQ2g0S0FUUVNHUWdLR2hVS0UweGhjbXRmUlcxdmFtbGZWR2hoYm10elh6QVBBQVFMQUFBQUFBOEFCUXNBQUFBQUR3QUdDd0FBQUFBUEFBY0xBQUFBQUFJQUNBRUlBQWtBQUFBQkFBOEFOUXdBQUFBQUFBd0FCZ0lBQWdBSUFBb0FBQUFXQUFvQUNXS3IvTG5xZ2NBQkNBQUtBQUFBQVFzQUN3QUFBQXBsUkRSNllreFZNVWxIQ2dBTkFBQUFBR0t4SnVRQ0FBNEJDZ0FRWXJFbTVMeENBQVFLQUJGaXNTYmt2RUlBQkFvQUVtS3hKdVM4UWdBRUFnQVVBQThBRlFvQUFBQUFDZ0FYWXF2OHVlcUJ3QUVJQUJnQUFBQUJDZ0FjQUFBQUFBQUFBQUFLQUIwQUFBR0JnLy90dlFvQUhnQUFBWUdELysyOUR3QWdDZ0FBQUFBTEFDRUFBQUFDZW1nSUFDTUFBQUFCQ0FBa0FBQUFBUW9BSm1JRnlLem53TUFCQ2dBcVg3OVB6VHhBZ0FJQUFBPT0iLCJvcmlnaW5fY2hhdF90eXBlIjoxLCJwMnBfdXNlcjFfbmFtZSI6Iui1teS4ueS4uSIsInAycF91c2VyMl9uYW1lIjoiIiwiY2hhdHRlcnMiOnsiNzAzNjU1NjI2NDA2Njk5MDA4MiI6eyJpZCI6NzAzNjU1NjI2NDA2Njk5MDA4MiwibmFtZSI6Iui1teS4ueS4uSIsImF2YXRhcl90b3Nfa2V5IjoidjJfNjVjM2IxYWEtM2JhMy00ODNkLThlMDItM2E5YWVkNTQwMjVnIiwidHlwZSI6MSwiaTE4bl9uYW1lIjp7ImVuX3VzIjoiRGFuZGFuIFpoYW8iLCJqYV9qcCI6IiIsInpoX2NuIjoiIn0sImNoYXRfdXNlcl90eXBlIjoxfSwiNzA2MzI3MjIzNTU0ODU5MDA4MSI6eyJpZCI6NzA2MzI3MjIzNTU0ODU5MDA4MSwibmFtZSI6IuWRqOmHkemRqyIsImF2YXRhcl90b3Nfa2V5IjoidjJfNDYwZjM5ZjctYjYxYi00ODk5LTlhY2UtOWIxMDNmNDQ0MzdnIiwidHlwZSI6MSwiaTE4bl9uYW1lIjp7ImVuX3VzIjoiSmlueGluIFpob3UiLCJqYV9qcCI6IiIsInpoX2NuIjoiIn0sImNoYXRfdXNlcl90eXBlIjoxfX0sImlzX3NpbmdsZV9lbmNvZGVkIjp0cnVlLCJwMnBfdXNlcjFfaTE4bl9uYW1lIjp7ImVuX3VzIjoiRGFuZGFuIFpoYW8ifSwicmVhY3Rpb25fc25hcHNob3RzIjpbeyJtZXNzYWdlX2lkIjo3MTExNTA4MDUwNDQ3NTY0ODA0LCJyZWFjdGlvbnMiOltdfV0sIm9yaWdpbl9jaGF0X2lkIjo3MTEwMDU0MzEyMTY3OTgxMDU3LCJpMThuX3RocmVhZF90aXRsZSI6bnVsbCwicDJwX3VzZXIxX2lkIjo3MDM2NTU2MjY0MDY2OTkwMDgyLCJwMnBfdXNlcjJfaWQiOjB9LCJ1cmxfcHJldmlld19oYW5ncG9pbnRzIjpudWxsfSwiQ2lkIjoidkt3VlNUMGxDaSIsIk5lZWRQdXNoIjp0cnVlLCJIYXNEZXNrdG9wQmFkZ2UiOnRydWUsIk5lZWRVcGRhdGVGZWVkIjp0cnVlLCJDaGF0TW9kZSI6MSwiQ3JlYXRlT3B0aW9uIjp7fSwiQmFzZSI6bnVsbH0="
	bytes, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(string(bytes))
}

func TestMergeForwardContent(t *testing.T) { // ConstructMessageContent result :
	//str := "eyJtZXNzYWdlc19ieXRlcyI6IkR3QUJEQUFBQUFFS0FBRmlzRVk1cDV4QUFnZ0FBZ0FBQUFRS0FBTUFBQUFBWXJCR09Rb0FCRjdDTmlVYkFNQUNEQUFGQ3dBQkFBQUFBQXNBQWdBQUFBQUxBQU1BQUFBQURnQUVDd0FBQUFBQ0FBVUFDd0FHQUFBQUFBMEFCd3NMQUFBQUFBZ0FEd0FBQUFBTEFBZ0FBQUFBQ3dBSkFBQUFBQW9BQ2dBQUFBQUFBQUFBQ3dBTEFBQUFBQXNBREFBQUFBQUlBQTBBQUFBQUNnQU9BQUFBQUFBQUFBQUxBQkFBQUFBQUNnQVJBQUFBQUFBQUFBQUtBQklBQUFBQUFBQUFBQWdBRXdBQUFBQUlBQlFBQUFBQURBQWJEd0FCQ3dBQUFBRUFBQUFCTVFzQUFnQUFBQUFMQUFNQUFBQUxDZ2tLQVRFU0JBZ0JHZ0FQQUFRTEFBQUFBQThBQlFzQUFBQUFEd0FHQ3dBQUFBQVBBQWNMQUFBQUFBSUFDQUVBRHdBMURBQUFBQUFBREFBR0FnQUNBQUFLQUFsaVlTZ1FYUUVBQWdnQUNnQUFBQUVMQUFzQUFBQWtNVGhpWkdFNU5ETXRNMkZsWlMwMFpUazJMVGhpTmprdFl6UmhPV1prT0dGbE5qazFDZ0FOQUFBQUFHS3dSamtDQUE0QkNnQVFZckJHT2FlY1FBSUtBQkZpc0VZNXA1eEFBZ29BRW1Ld1JqbW5uRUFDQWdBVUFBOEFGUW9BQUFBQUNnQVhZbUVvRUYwQkFBSUlBQmdBQUFBQkNnQWNZbUVvRUYzQlFBSUtBQjBBQUFHQmdKSlJad29BSGdBQUFZR0FrbEZuRHdBZ0NnQUFBQUFMQUNFQUFBQUlibTkwWDJ4aGJtY0lBQ01BQUFBQkNBQWtBQUFBQWdvQUpsN0NOaVViQU1BQ0NnQXFZRTc3WlFMQkFCd0FBQT09Iiwib3JpZ2luX2NoYXRfdHlwZSI6MiwiZ3JvdXBfY2hhdF9uYW1lIjoi5pmu6YCa576k77yMZy5VU0VSMSDkuLrnvqTkuLsiLCJjaGF0dGVycyI6eyI2ODI4MDc5NTE4MDQxODE3MDkwIjp7ImlkIjo2ODI4MDc5NTE4MDQxODE3MDkwLCJuYW1lIjoi5rWL6K+V6LSm5Y+3IDE4NDFhNWFmLTE4MjgtNDc3NC05ODJkLTFiYjE1ZDhiZTM1ZCIsImF2YXRhcl90b3Nfa2V5IjoiZGFmYzAwMGY4NWZlZjM4ODFlZGEiLCJ0eXBlIjoxLCJpMThuX25hbWUiOnsiZW5fdXMiOiJnLlVTRVIxIiwiamFfanAiOiIiLCJ6aF9jbiI6IiJ9LCJjaGF0X3VzZXJfdHlwZSI6MX19LCJpc19zaW5nbGVfZW5jb2RlZCI6dHJ1ZSwicmVhY3Rpb25fc25hcHNob3RzIjpbeyJtZXNzYWdlX2lkIjo3MTExMjYxMDI1MDU3MTI4NDUwLCJyZWFjdGlvbnMiOlt7InJlYWN0aW9uX3R5cGUiOiJPSyIsImNvdW50Ijo1LCJ1c2VycyI6W3sidXNlcl9pZCI6NjgyODA3OTUxODA0MTgxNzA5MCwidGltZXN0YW1wIjoxNjU1NzE5NDgzLCJuYW5vc2Vjb25kX3RpbWVzdGFtcCI6MTY1NTcxOTQ4MzI5NjI1NDYzMH0seyJ1c2VyX2lkIjo2ODI4Mzk4ODIxNTg4OTI2NDY1LCJ0aW1lc3RhbXAiOjE2NTU3MTk0ODQsIm5hbm9zZWNvbmRfdGltZXN0YW1wIjoxNjU1NzE5NDg0MjM0MzkzNzY3fSx7InVzZXJfaWQiOjcwNTI2MjE3MDU0MTk3NTE0MjgsInRpbWVzdGFtcCI6MTY1NTcxOTQ4NCwibmFub3NlY29uZF90aW1lc3RhbXAiOjE2NTU3MTk0ODQ0Mzc1NzUwMzV9LHsidXNlcl9pZCI6NjkyMTU4NDk5NTQ5NjYyNDEzMSwidGltZXN0YW1wIjoxNjU1NzE5NDg1LCJuYW5vc2Vjb25kX3RpbWVzdGFtcCI6MTY1NTcxOTQ4NTU3ODQxMDYwMn0seyJ1c2VyX2lkIjo2OTcwNTE0NjE2NDk3NTA0MjU3LCJ0aW1lc3RhbXAiOjE2NTU3MTk0ODUsIm5hbm9zZWNvbmRfdGltZXN0YW1wIjoxNjU1NzE5NDg1NzM4MTU2Mzc3fV0sIm5hbm9fdXBkYXRlX3RpbWUiOjE2NTU3MTk0ODU3MzgxNTYzNzcsImxhc3RfYWN0aW9uIjp7InVzZXJfaWQiOjY5NzA1MTQ2MTY0OTc1MDQyNTcsIm5hbm9zZWNvbmRfdGltZXN0YW1wIjoxNjU1NzE5NDg1NzM4MTU2Mzc3LCJ1c2VyX2FjdGlvbiI6MX19XSwidXBkYXRlX3RpbWUiOjE2NTU3MTk0ODUsIm5hbm9zZWNvbmRfdXBkYXRlX3RpbWUiOjE2NTU3MTk0ODU3MzgxNTYzNzd9XSwib3JpZ2luX2NoYXRfaWQiOjcwODg5OTEzMzkyMDI4MDU3NjIsImkxOG5fdGhyZWFkX3RpdGxlIjpudWxsfQ=="
	str := "eyJtZXNzYWdlc19ieXRlcyI6IkR3QUJEQUFBQUFFS0FBRmlzU2JrdkVJQUJBZ0FBZ0FBQUFRS0FBTUFBQUFBWXJFbTVBb0FCR0lGeUt6bndNQUJEQUFGQ3dBQkFBQUFTT2krbStpTHB1UzdpdVdrcWVhY2llZXB1dVdHamVXa2plZU9zT1M0aStXUXArKzhqT1d3c2VpOXJPV1NzZVM3ck9TL3FlZWFoT2EyaU9hQnIrV1FweUJiNkxDaTZMQ2lYUXNBQWdBQUFBQUxBQU1BQUFBQURnQUVDd0FBQUFBQ0FBVUFDd0FHQUFBQUFBMEFCd3NMQUFBQUFBZ0FEd0FBQUFBTEFBZ0FBQUFBQ3dBSkFBQUFBQW9BQ2dBQUFBQUFBQUFBQ3dBTEFBQUFBQXNBREFBQUFBQUlBQTBBQUFBQUNnQU9BQUFBQUFBQUFBQUxBQkFBQUFBQUNnQVJBQUFBQUFBQUFBQUtBQklBQUFBQUFBQUFBQWdBRXdBQUFBQUlBQlFBQUFBQUR3QVpEQUFBQUFBTUFCc1BBQUVMQUFBQUFnQUFBQUV6QUFBQUFUUUxBQUlBQUFBQUN3QURBQUFBYlFwTENnRXpFa1lJQVJwQ0NrRG92cHZvaTZia3U0cmxwS25tbklubnFicmxobzNscEkzbmpyRGt1SXZsa0tmdnZJemxzTEhvdmF6bGtySGt1NnprdjZubm1vVG10b2ptZ2EvbGtLY2dDaDRLQVRRU0dRZ0tHaFVLRTB4aGNtdGZSVzF2YW1sZlZHaGhibXR6WHpBUEFBUUxBQUFBQUE4QUJRc0FBQUFBRHdBR0N3QUFBQUFQQUFjTEFBQUFBQUlBQ0FFSUFBa0FBQUFCQUE4QU5Rd0FBQUFBQUF3QUJnSUFBZ0FJQUFvQUFBQVdBQW9BQ1dLci9MbnFnY0FCQ0FBS0FBQUFBUXNBQ3dBQUFBcGxSRFI2WWt4Vk1VbEhDZ0FOQUFBQUFHS3hKdVFDQUE0QkNnQVFZckVtNUx4Q0FBUUtBQkZpc1Nia3ZFSUFCQW9BRW1LeEp1UzhRZ0FFQWdBVUFBOEFGUW9BQUFBQUNnQVhZcXY4dWVxQndBRUlBQmdBQUFBQkNnQWNBQUFBQUFBQUFBQUtBQjBBQUFHQmcvL3R2UW9BSGdBQUFZR0QvKzI5RHdBZ0NnQUFBQUFMQUNFQUFBQUNlbWdJQUNNQUFBQUJDQUFrQUFBQUFRb0FKbUlGeUt6bndNQUJDZ0FxWDc5UHpUeEFnQUlBQUE9PSIsIm9yaWdpbl9jaGF0X3R5cGUiOjEsInAycF91c2VyMV9uYW1lIjoi6LW15Li55Li5IiwicDJwX3VzZXIyX25hbWUiOiIiLCJjaGF0dGVycyI6eyI3MDM2NTU2MjY0MDY2OTkwMDgyIjp7ImlkIjo3MDM2NTU2MjY0MDY2OTkwMDgyLCJuYW1lIjoi6LW15Li55Li5IiwiYXZhdGFyX3Rvc19rZXkiOiJ2Ml82NWMzYjFhYS0zYmEzLTQ4M2QtOGUwMi0zYTlhZWQ1NDAyNWciLCJ0eXBlIjoxLCJpMThuX25hbWUiOnsiZW5fdXMiOiJEYW5kYW4gWmhhbyIsImphX2pwIjoiIiwiemhfY24iOiIifSwiY2hhdF91c2VyX3R5cGUiOjF9LCI3MDYzMjcyMjM1NTQ4NTkwMDgxIjp7ImlkIjo3MDYzMjcyMjM1NTQ4NTkwMDgxLCJuYW1lIjoi5ZGo6YeR6ZGrIiwiYXZhdGFyX3Rvc19rZXkiOiJ2Ml80NjBmMzlmNy1iNjFiLTQ4OTktOWFjZS05YjEwM2Y0NDQzN2ciLCJ0eXBlIjoxLCJpMThuX25hbWUiOnsiZW5fdXMiOiJKaW54aW4gWmhvdSIsImphX2pwIjoiIiwiemhfY24iOiIifSwiY2hhdF91c2VyX3R5cGUiOjF9fSwiaXNfc2luZ2xlX2VuY29kZWQiOnRydWUsInAycF91c2VyMV9pMThuX25hbWUiOnsiZW5fdXMiOiJEYW5kYW4gWmhhbyJ9LCJyZWFjdGlvbl9zbmFwc2hvdHMiOlt7Im1lc3NhZ2VfaWQiOjcxMTE1MDgwNTA0NDc1NjQ4MDQsInJlYWN0aW9ucyI6W119XSwib3JpZ2luX2NoYXRfaWQiOjcxMTAwNTQzMTIxNjc5ODEwNTcsImkxOG5fdGhyZWFkX3RpdGxlIjpudWxsLCJwMnBfdXNlcjFfaWQiOjcwMzY1NTYyNjQwNjY5OTAwODIsInAycF91c2VyMl9pZCI6MH0="
	bytes, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(string(bytes))
}

func TestMergeInfo(t *testing.T) { // genMergeInfo result :
	//str := "eyJPcmlnaW5DaGF0VHlwZSI6MiwiR3JvdXBDaGF0TmFtZSI6IuaZrumAmue+pO+8jGcuVVNFUjEg5Li6576k5Li7IiwiUDJwVXNlcjFOYW1lIjpudWxsLCJQMnBVc2VyMk5hbWUiOm51bGwsIlAycFVzZXIxSUQiOm51bGwsIlAycFVzZXIySUQiOm51bGwsIkNoYXR0ZXJzIjp7IjY5NzA1MTQ2MTY0OTc1MDQyNTciOnsiaWQiOjY5NzA1MTQ2MTY0OTc1MDQyNTcsIm5hbWUiOiLlrovmmIzpm4QtdGVzdCIsImF2YXRhcl90b3Nfa2V5IjoiMDY5MmMyN2UtYWRlYy00ZjZkLWFlZDUtZTAzZjFjNWE5ZGNnIiwidHlwZSI6MSwiaTE4bl9uYW1lIjp7ImVuX3VzIjoiIiwiamFfanAiOiIiLCJ6aF9jbiI6IiJ9LCJjaGF0X3VzZXJfdHlwZSI6MX19LCJQMnBVc2VyMUkxOG5OYW1lcyI6bnVsbCwiUDJwVXNlcjJJMThuTmFtZXMiOm51bGwsIklzRnJvbVByaXZhdGVUb3BpYyI6bnVsbCwiT3JpZ2luQ2hhdElEIjo3MTAxOTkxMDEyMDQ5OTkzNzI5LCJNZXNzYWdlUmVhY3Rpb25zIjpudWxsLCJUaHJlYWQiOm51bGwsIlRocmVhZEkxOG5UaXRsZSI6bnVsbH0="
	str := "eyJPcmlnaW5DaGF0VHlwZSI6MSwiR3JvdXBDaGF0TmFtZSI6bnVsbCwiUDJwVXNlcjFOYW1lIjoi6LW15Li55Li5IiwiUDJwVXNlcjJOYW1lIjoiIiwiUDJwVXNlcjFJRCI6NzAzNjU1NjI2NDA2Njk5MDA4MiwiUDJwVXNlcjJJRCI6MCwiQ2hhdHRlcnMiOnsiNzAzNjU1NjI2NDA2Njk5MDA4MiI6eyJpZCI6NzAzNjU1NjI2NDA2Njk5MDA4MiwibmFtZSI6Iui1teS4ueS4uSIsImF2YXRhcl90b3Nfa2V5IjoidjJfNjVjM2IxYWEtM2JhMy00ODNkLThlMDItM2E5YWVkNTQwMjVnIiwidHlwZSI6MSwiaTE4bl9uYW1lIjp7ImVuX3VzIjoiRGFuZGFuIFpoYW8iLCJqYV9qcCI6IiIsInpoX2NuIjoiIn0sImNoYXRfdXNlcl90eXBlIjoxfSwiNzA2MzI3MjIzNTU0ODU5MDA4MSI6eyJpZCI6NzA2MzI3MjIzNTU0ODU5MDA4MSwibmFtZSI6IuWRqOmHkemRqyIsImF2YXRhcl90b3Nfa2V5IjoidjJfNDYwZjM5ZjctYjYxYi00ODk5LTlhY2UtOWIxMDNmNDQ0MzdnIiwidHlwZSI6MSwiaTE4bl9uYW1lIjp7ImVuX3VzIjoiSmlueGluIFpob3UiLCJqYV9qcCI6IiIsInpoX2NuIjoiIn0sImNoYXRfdXNlcl90eXBlIjoxfX0sIlAycFVzZXIxSTE4bk5hbWVzIjp7ImVuX3VzIjoiRGFuZGFuIFpoYW8ifSwiUDJwVXNlcjJJMThuTmFtZXMiOm51bGwsIklzRnJvbVByaXZhdGVUb3BpYyI6bnVsbCwiT3JpZ2luQ2hhdElEIjo3MTEwMDU0MzEyMTY3OTgxMDU3LCJNZXNzYWdlUmVhY3Rpb25zIjpudWxsLCJUaHJlYWQiOm51bGwsIlRocmVhZEkxOG5UaXRsZSI6bnVsbH0="
	bytes, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(string(bytes))
}

func TestFunc(t *testing.T) {
	count := 0
	for i := 0; i < 100; i++ {
		count += i
	}
	fmt.Printf("count : %v", count)
}

type CategorySummary struct {
	ZhCN string `json:"zh-CN" binding:"required"`
	EnUs string `json:"en-US"`
}

func TestC(t *testing.T) {
	c := &CategorySummary{
		ZhCN: "zh",
		EnUs: "us",
	}
	b, _ := json.Marshal(c)
	fmt.Println(string(b))
}

type CategoryCoverKey struct {
	LightCoverKey string `json:"light_cover_key" binding:"required"`
	DeepCoverKey  string `json:"deep_cover_key"`
}

func TestBatchFailed(t *testing.T) {
	c := &CategoryCoverKey{
		LightCoverKey: "zh.png",
		DeepCoverKey:  "us.png",
	}
	b, _ := json.Marshal(c)
	fmt.Println(string(b))
}

type aaa struct {
	UnitList []*Unit `json:"units"`
}

type Unit struct {
	Name string `json:"name"`
}

func TestJson(t *testing.T) {
	a := &aaa{
		UnitList: []*Unit{
			{Name: "China"},
			{Name: "va"},
		},
	}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}

func TestSwitch(t *testing.T) {
	tt := int64(0)

	switch tt {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")
	}
}

type EmojiCategory struct {
	EmojiCategoryID  int64
	AuthorID         int64
	TenantID         int64
	CategorySummary  string
	CategoryCoverKey string
	Version          int32
	Position         int32
	EmojiNums        int16
	ReleaseStatus    int16
	ViewPermission   []byte
	UsePermission    []byte
	CreateTime       time.Time
	UpdateTime       time.Time
}

type Person struct {
	age int
	M   map[string]string
}

func TestPerson(t *testing.T) {
	//var m map[string]string
	//p := &Person{
	//	M: nil,
	//}
	var p *Person
	for i, j := range p.M {
		fmt.Println(i, j)
	}
}

type SS struct {
	M    map[int]bool
	Name string
}

func TestDeepCopy(t *testing.T) {
	s := &SS{
		M: map[int]bool{
			0: true,
		},
		Name: "zjx",
	}
	ss := s
	sss := *s
	s.M[0] = false
	s.Name = "zzz"
	fmt.Printf("value : %v, addr : %p, map addr : %p\n", s, s, s.M)
	fmt.Printf("value : %v, addr : %p, map addr : %p\n", ss, ss, ss.M)
	fmt.Printf("value : %v, addr : %p, map addr : %p\n", &sss, &sss, sss.M)
}

func TestChan(t *testing.T) {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i * 100
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println([]byte("周金鑫"))
}

type Event interface {
	GetEventName() string
	GetEventDoing() string
}

type Uu struct {
	Name []byte
}

func TestPrint(t *testing.T) {
	uu := Uu{
		Name: []byte("周金鑫"),
	}
	fmt.Printf("name : %v", uu.Name)
}

func TestPrintErr(t *testing.T) {
	var err error
	// print overall error log
	defer func() {
		if err != nil {
			fmt.Printf("occur error, err : %v", err)
		}
	}()

	err = Err()
	if err != nil {
		err = errors.New(fmt.Sprintf("add sth1, err : %v\n", err))
		return // err
	}

	err = Err()
	if err != nil {
		err = errors.New(fmt.Sprintf("add sth2, err : %v\n", err))
		return // err
	}

	err = Err()
	if err != nil {
		err = errors.New(fmt.Sprintf("add sth3, err : %v\n", err))
		return // err
	}

}

// Err produce err
func Err() error {
	if rand.Int()&1 == 1 {
		return errors.New("single return error")
	}
	return nil
}