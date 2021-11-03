package http

import (
	"context"
	"fmt"
	"code.byted.org/im_cloud/lark_bot_sdk/v2"
)

func TestBot(){

	ct1 := v2.CardTemplate1{
		OpenChatId:    "oc_d01dd8fa75cd7068c82e45099fbe7930",
		Title:      "OCI SDK TEST",
		Message:    "Test using sdk [Try this!](https://www.google.com)",
		FootNote:   "this is a footnote",
	}

	ctd := v2.CardTemplateDefault{
		OpenId:     "ou_72427976ac979e3a9fd97c3c422ab7e2",
		Message:    "Default template",
	}
	ctx := context.Background()
	err := v2.BotCard(ctx).SetTemplate(&ct1).BuildAndSend()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = v2.BotCard(ctx).SetTemplate(&ctd).BuildAndSend()
	if err != nil {
		fmt.Println(err.Error())
	}


}