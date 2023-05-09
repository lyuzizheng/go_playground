package processor

import (
	"code.byted.org/gopkg/logs"
	diagnoseTool "code.byted.org/im_cloud/diagnose_tool_sdk/core"
	"code.byted.org/gopkg/metrics"
	"code.byted.org/im_cloud/push/thrift_gen/push_strategy"
	"code.byted.org/kite/kitutil"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

const (
	BBPush =  "BackBonePush"
	PNSPush = "PNSPush"
)


var (
	bbPushTags = []metrics.T{
		metrics.Tag("method", BBPush),
		metrics.Tag("empty", "false"),
	}
	pnsPushTags = []metrics.T{
		metrics.Tag("method", PNSPush),
		metrics.Tag("empty", "false"),
	}
)

//func ExportBackBonePush(ctx context.Context, toUserId int64, conshortId int64, convCursor int64, userCursorStart int64, userEndCursor int64,  pushRes *processor.Pusher) {
func ExportBackBonePush(ctx context.Context, pusher *Pusher) {
	//In Case of push failure
	if pusher.PushRes == nil {
		return
	}
	logid, _ := kitutil.GetCtxLogID(ctx)

	record := diagnoseTool.DataExportRecord{
		Userid:      pusher.ToUserId,
		ConvShortId: pusher.MessageBody.ConversationShortId,
		Ext:         make(map[string]string),
		TimeStamp:   time.Now().UnixNano(),
		LogId:       logid,
		Method:      BBPush,
	}

	convInfo := &diagnoseTool.MsgsInfo{
		ConvShortId:     pusher.MessageBody.ConversationShortId,
		ConvStartCursor: pusher.IndexInConversation,
		ConvEndCursor:   pusher.IndexInConversation,
		Num:             1,
	}
	userInfo := &diagnoseTool.UserInfo{
		UserStartCursor: pusher.PreIndexInUser,
		UserEndCursor:   pusher.IndexInUser,
		Num:             1,
	}
	pushResultMap := make(map[string]string)
	for _, pushResult := range pusher.PushRes {
		pushResultMap[strconv.Itoa(int(pushResult.DeviceID))] = pushResult.DevicePlatform.String() + ":" + pushResult.Code.String()
	}

	//Add Ext Fields
	convInfoBytes, _ := json.Marshal(convInfo)
	userInfoBytes, _ := json.Marshal(userInfo)
	pushResultBytes, _ := json.Marshal(pushResultMap)
	record.Ext = make(map[string]string)
	record.Ext["msg_info"] = string(convInfoBytes)
	record.Ext["user_info"] = string(userInfoBytes)
	record.Ext["push_result"] = string(pushResultBytes)

	diagnoseTool.ExportDataRecord(ctx, &record, bbPushTags...)
	logs.CtxDebug(ctx,"[DataExportHelper] Send 1 message to eventbus")
	return

}

func ExportPNSPushRecord(ctx context.Context, pe *push_strategy.PushEvent) {

	logid, _ := kitutil.GetCtxLogID(ctx)

	record := diagnoseTool.DataExportRecord{
		Userid:      pe.ToUser,
		DeviceId:    "",
		ConvShortId: pe.Message.ConversationShortId,
		Ext:         make(map[string]string),
		TimeStamp:   time.Now().UnixNano(),
		LogId:       logid,
		Method:      PNSPush,
	}

	convInfo := &diagnoseTool.MsgsInfo{
		ConvShortId:     pe.Message.ConversationShortId,
		ConvStartCursor: pe.IndexInConversation,
		ConvEndCursor:   pe.IndexInConversation,
		Num:             1,
	}
	userInfo := &diagnoseTool.UserInfo{
		UserStartCursor: pe.PreIndexInUser,
		UserEndCursor:   pe.IndexInUser,
		Num:             1,
	}

	convInfoBytes, _ := json.Marshal(convInfo)
	convInfoStr := string(convInfoBytes)
	userInfoBytes, _ := json.Marshal(userInfo)
	userInfoStr := string(userInfoBytes)
	record.Ext = make(map[string]string)
	record.Ext["msg_info"] = convInfoStr
	record.Ext["user_info"] = userInfoStr


	diagnoseTool.ExportDataRecord(ctx, &record, pnsPushTags...)

	logs.CtxDebug(ctx,"[DataExportHelper] Send 1 message to eventbus")
	return

}