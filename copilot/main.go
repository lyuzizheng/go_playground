package main

import (
	"code.byted.org/gopkg/jsonx"
	"code.byted.org/gopkg/logs"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	var riddleQn *RiddleQns
	ctx := context.Background()
	copliotResp, err := GetCopilotReps(ctx)
	if err != nil {
		riddleQn = GetHardCodeRiddleQns()
	} else {
		riddleQn, err = ParseRespToRiddleQuestion(ctx, copliotResp)
		if err != nil {
			riddleQn = GetHardCodeRiddleQns()
		}
	}

	logs.CtxInfo(ctx, jsonx.ToString(riddleQn))

}

func GetHardCodeRiddleQns() *RiddleQns {
	return nil
}

type CopilotResp struct {
	Result       string        `json:"result"`
	Sources      []interface{} `json:"sources"`
	Quota        int           `json:"quota"`
	DefaultQuota int           `json:"default_quota"`
	PackageQuota int           `json:"package_quota"`
}

type RiddleQns struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int32    `json:"correct_answer"`
	Reason        string   `json:"reason"`
}

func GetCopilotReps(ctx context.Context) (*CopilotResp, error) {

	proxyURL := fmt.Sprintf("http://%s@%s", "tiktoksocial_copilothubforhackathon_kongzy:LIJEDMytwsL6X7hm", "id1634.http-sg-idc-idc-sg-flow.forward-proxy.byted.org:8080")
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(proxyURL)
			},
		},
	}
	data := url.Values{
		"api_key": {"d9aebd30-e7d4-4595-ba91-e7dc56f9ccfb"},
		"query":   {"Give a me a funny riddle with one correct answer and two wrong answers, shuffle the answer order and explain the reason,"},
	}
	req, err := http.NewRequest("POST", "https://api.copilothub.ai/openapi/v1/query", strings.NewReader(data.Encode()))
	if err != nil {
		logs.CtxError(ctx, "new req error", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		logs.CtxError(ctx, "call copilot error", err)
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.CtxError(ctx, "read body error", err)
		return nil, err
	}

	respStruct := &CopilotResp{}
	logs.CtxInfo(ctx, string(respBody))

	err = jsonx.Unmarshal(respBody, respStruct)
	if err != nil {
		logs.CtxError(ctx, "unmarshall err", err)
		return nil, err
	}
	return respStruct, nil

}

func ParseRespToRiddleQuestion(ctx context.Context, copilotResp *CopilotResp) (*RiddleQns, error) {
	myStr := copilotResp.Result
	myStrArr := strings.Split(myStr, "\n\n")
	riddleQn := RiddleQns{}
	fmt.Println(len(myStrArr))
	if len(myStrArr) == 4 {
		fmt.Println("length is 4")
		riddleQn.Question = myStrArr[1]
		riddleQn.Reason = myStrArr[3]
	}

	answers := strings.Split(myStrArr[2], "\n")
	if len(answers) == 3 {
		fmt.Println("length is 3")
		riddleQn.Answers = answers
	}
	reasonCopy := riddleQn.Reason
	after, found := strings.CutPrefix(reasonCopy, "The correct answer is option ")
	if found && len(after) > 0 {
		correct := after[0:1]
		switch strings.ToLower(correct) {
		case "a":
			riddleQn.CorrectAnswer = 0
		case "b":
			riddleQn.CorrectAnswer = 1
		case "c":
			riddleQn.CorrectAnswer = 2
		default:
			riddleQn.CorrectAnswer = -1
		}
	}
	fmt.Println(jsonx.ToString(riddleQn))
	return nil, nil
}
