package event

import (
	"github.com/tencent-connect/botgo/dto"
)

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler

	Guild       GuildEventHandler
	GuildMember GuildMemberEventHandler
	Channel     ChannelEventHandler

	Message             MessageEventHandler
	MessageReaction     MessageReactionEventHandler
	ATMessage           ATMessageEventHandler
	DirectMessage       DirectMessageEventHandler
	MessageAudit        MessageAuditEventHandler
	MessageDelete       MessageDeleteEventHandler
	PublicMessageDelete PublicMessageDeleteEventHandler
	DirectMessageDelete DirectMessageDeleteEventHandler

	Audio AudioEventHandler

	Thread     ThreadEventHandler
	Post       PostEventHandler
	Reply      ReplyEventHandler
	ForumAudit ForumAuditEventHandler

<<<<<<< HEAD
	Interaction    InteractionEventHandler
	GroupAtMessage GroupAtMessageEventHandler
	GroupMessage   GroupMessageEventHandler
	C2CMessage     C2CMessageEventHandler

	GroupAddRobot   GroupAddRobotEventHandle
	GroupDelRobot   GroupDelRobotEventHandle
	GroupMsgReject  GroupMsgRejectEventHandle
	GroupMsgReceive GroupMsgReceiveEventHandle

	FriendAdd     FriendAddEventHandle
	FriendDel     FriendDelEventHandle
	C2CMsgReject  C2CMsgRejectHandle
	C2CMsgReceive C2CMsgReceiveHandle
=======
	Interaction InteractionEventHandler
>>>>>>> fbfd4112b279aa509a885d86af8b0678db55e765
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *dto.WSPayload, data *dto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *dto.WSPayload, message []byte) error

// GuildEventHandler 频道事件handler
type GuildEventHandler func(event *dto.WSPayload, data *dto.WSGuildData) error

// GuildMemberEventHandler 频道成员事件 handler
type GuildMemberEventHandler func(event *dto.WSPayload, data *dto.WSGuildMemberData) error

// ChannelEventHandler 子频道事件 handler
type ChannelEventHandler func(event *dto.WSPayload, data *dto.WSChannelData) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *dto.WSPayload, data *dto.WSMessageData) error

// MessageDeleteEventHandler 消息事件 handler
type MessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSMessageDeleteData) error

// PublicMessageDeleteEventHandler 消息事件 handler
type PublicMessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSPublicMessageDeleteData) error

// DirectMessageDeleteEventHandler 消息事件 handler
type DirectMessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSDirectMessageDeleteData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *dto.WSPayload, data *dto.WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *dto.WSPayload, data *dto.WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *dto.WSPayload, data *dto.WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *dto.WSPayload, data *dto.WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *dto.WSPayload, data *dto.WSMessageAuditData) error

// ThreadEventHandler 论坛主题事件 handler
type ThreadEventHandler func(event *dto.WSPayload, data *dto.WSThreadData) error

// PostEventHandler 论坛回帖事件 handler
type PostEventHandler func(event *dto.WSPayload, data *dto.WSPostData) error

// ReplyEventHandler 论坛帖子回复事件 handler
type ReplyEventHandler func(event *dto.WSPayload, data *dto.WSReplyData) error

// ForumAuditEventHandler 论坛帖子审核事件 handler
type ForumAuditEventHandler func(event *dto.WSPayload, data *dto.WSForumAuditData) error

// InteractionEventHandler 互动事件 handler
type InteractionEventHandler func(event *dto.WSPayload, data *dto.WSInteractionData) error

<<<<<<< HEAD
type GroupAtMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupATMessageData) error

type GroupMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupMessageData) error

type C2CMessageEventHandler func(event *dto.WSPayload, data *dto.WSC2CMessageData) error

type GroupAddRobotEventHandle func(event *dto.WSPayload, data *dto.WSGroupAddRobotData) error

type GroupDelRobotEventHandle func(event *dto.WSPayload, data *dto.WSGroupDelRobotData) error

type GroupMsgRejectEventHandle func(event *dto.WSPayload, data *dto.WSGroupMsgRejectData) error

type GroupMsgReceiveEventHandle func(event *dto.WSPayload, data *dto.WSGroupMsgReceiveData) error

type FriendAddEventHandle func(event *dto.WSPayload, data *dto.WSFriendAddData) error

type FriendDelEventHandle func(event *dto.WSPayload, data *dto.WSFriendDelData) error

type C2CMsgRejectHandle func(event *dto.WSPayload, data *dto.WSFriendMsgRejectData) error

type C2CMsgReceiveHandle func(event *dto.WSPayload, data *dto.WSFriendMsgReveiceData) error

=======
>>>>>>> fbfd4112b279aa509a885d86af8b0678db55e765
// RegisterHandlers 注册事件回调，并返回 intent 用于 websocket 的鉴权
func RegisterHandlers(handlers ...interface{}) dto.Intent {
	var i dto.Intent
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case AudioEventHandler:
			DefaultHandlers.Audio = handle
			i = i | dto.EventToIntent(
				dto.EventAudioStart, dto.EventAudioFinish,
				dto.EventAudioOnMic, dto.EventAudioOffMic,
			)
		case InteractionEventHandler:
			DefaultHandlers.Interaction = handle
			i = i | dto.EventToIntent(dto.EventInteractionCreate)
<<<<<<< HEAD
		case GroupAtMessageEventHandler:
			DefaultHandlers.GroupAtMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupATMessageCreate)
		case GroupMessageEventHandler:
			DefaultHandlers.GroupMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupMessageCreate)
		case C2CMessageEventHandler:
			DefaultHandlers.C2CMessage = handle
			i = i | dto.EventToIntent(dto.EventC2CMessageCreate)
		case GroupAddRobotEventHandle:
			DefaultHandlers.GroupAddRobot = handle
			i = i | dto.EventToIntent(dto.EventGroupAddRobbot)
		case GroupDelRobotEventHandle:
			DefaultHandlers.GroupDelRobot = handle
			i = i | dto.EventToIntent(dto.EventGroupDelRobbot)
		case GroupMsgRejectEventHandle:
			DefaultHandlers.GroupMsgReject = handle
			i = i | dto.EventToIntent(dto.EventGroupMsgReject)
		case GroupMsgReceiveEventHandle:
			DefaultHandlers.GroupMsgReceive = handle
			i = i | dto.EventToIntent(dto.EventGroupMsgReceive)
		case FriendAddEventHandle:
			DefaultHandlers.FriendAdd = handle
			i = i | dto.EventToIntent(dto.EventFriendAdd)
		case FriendDelEventHandle:
			DefaultHandlers.FriendDel = handle
			i = i | dto.EventToIntent(dto.EventFriendDel)
		case C2CMsgRejectHandle:
			DefaultHandlers.C2CMsgReject = handle
			i = i | dto.EventToIntent(dto.EventC2CMsgReject)
		case C2CMsgReceiveHandle:
			DefaultHandlers.C2CMsgReceive = handle
			i = i | dto.EventToIntent(dto.EventC2CMsgReceive)
=======
>>>>>>> fbfd4112b279aa509a885d86af8b0678db55e765
		default:
		}
	}
	i = i | registerRelationHandlers(i, handlers...)
	i = i | registerMessageHandlers(i, handlers...)
	i = i | registerForumHandlers(i, handlers...)

	return i
}

func registerForumHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case ThreadEventHandler:
			DefaultHandlers.Thread = handle
			i = i | dto.EventToIntent(
				dto.EventForumThreadCreate, dto.EventForumThreadUpdate, dto.EventForumThreadDelete,
			)
		case PostEventHandler:
			DefaultHandlers.Post = handle
			i = i | dto.EventToIntent(dto.EventForumPostCreate, dto.EventForumPostDelete)
		case ReplyEventHandler:
			DefaultHandlers.Reply = handle
			i = i | dto.EventToIntent(dto.EventForumReplyCreate, dto.EventForumReplyDelete)
		case ForumAuditEventHandler:
			DefaultHandlers.ForumAudit = handle
			i = i | dto.EventToIntent(dto.EventForumAuditResult)
		default:
		}
	}
	return i
}

// registerRelationHandlers 注册频道关系链相关handlers
func registerRelationHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case GuildEventHandler:
			DefaultHandlers.Guild = handle
			i = i | dto.EventToIntent(dto.EventGuildCreate, dto.EventGuildDelete, dto.EventGuildUpdate)
		case GuildMemberEventHandler:
			DefaultHandlers.GuildMember = handle
			i = i | dto.EventToIntent(dto.EventGuildMemberAdd, dto.EventGuildMemberRemove, dto.EventGuildMemberUpdate)
		case ChannelEventHandler:
			DefaultHandlers.Channel = handle
			i = i | dto.EventToIntent(dto.EventChannelCreate, dto.EventChannelDelete, dto.EventChannelUpdate)
		default:
		}
	}
	return i
}

// registerMessageHandlers 注册消息相关的 handler
func registerMessageHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case MessageEventHandler:
			DefaultHandlers.Message = handle
			i = i | dto.EventToIntent(dto.EventMessageCreate)
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handle
			i = i | dto.EventToIntent(dto.EventAtMessageCreate)
		case DirectMessageEventHandler:
			DefaultHandlers.DirectMessage = handle
			i = i | dto.EventToIntent(dto.EventDirectMessageCreate)
		case MessageDeleteEventHandler:
			DefaultHandlers.MessageDelete = handle
			i = i | dto.EventToIntent(dto.EventMessageDelete)
		case PublicMessageDeleteEventHandler:
			DefaultHandlers.PublicMessageDelete = handle
			i = i | dto.EventToIntent(dto.EventPublicMessageDelete)
		case DirectMessageDeleteEventHandler:
			DefaultHandlers.DirectMessageDelete = handle
			i = i | dto.EventToIntent(dto.EventDirectMessageDelete)
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
			i = i | dto.EventToIntent(dto.EventMessageReactionAdd, dto.EventMessageReactionRemove)
		case MessageAuditEventHandler:
			DefaultHandlers.MessageAudit = handle
			i = i | dto.EventToIntent(dto.EventMessageAuditPass, dto.EventMessageAuditReject)
		default:
		}
	}
	return i
}
