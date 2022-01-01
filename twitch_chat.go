package twitch_chat

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type Message struct {
	// Username of user sending message
	Username string
	// Actual chat message
	Message string
	// if the user is subscriber or not
	Subscriber string
	// id of user sending message
	UserId string
	// timestamp in milliseconds
	Timestamp string
	// message to which this message is a reply
	ReplyTo string
	// username of user to whom this message is a reply
	ReplyToUser string
	// if the user is mod or not
	Mod string
	// System messages are messages by twitch
	// These include gifted subs and other actions
	SystemMsg string
	// Recipient of sub gift or other action
	MsgParamRecipientUsername string
	// Recipient display name
	MsgParamRecipientDisplayName string
	// Number of months of gifted sub
	MsgParamGiftMonths string
	// Number of months
	MsgParamMonths string
	// msg-id=subgift means gifted sub
	MsgId string
}

const (
	TWITCH_NICK         = "justinfan28165"
	TWITCH_WS_URL       = "irc-ws.chat.twitch.tv"
	PASS                = "SCHMOOPIIE"
	USER_TYPE_SEPERATOR = "PRIVMSG #"
)

var (
	// https://ircv3.net/specs/extensions/message-tags.html
	IRC_MSG_TAGS = []string{"subscriber", "user-id", "user-type", "mod", "tmi-sent-ts", "display-name", "reply-parent-msg-body", "reply-parent-user-login", "system-msg", "msg-param-recipient-user-name", "msg-param-recipient-display-name", "msg-param-gift-months", "msg-param-months", "msg-id"}
	// ws connection
	c *websocket.Conn
)

func sliceContains(s []string, e string, checkString bool) bool {
	for _, a := range s {
		if a == e {
			return true
		}
		if checkString {
			if strings.Contains(a, e) {
				return true
			}
		}
	}
	return false
}

// Returns message from user-type IRC tag
// Example-
//	 :tarik!tarik@tarik.tmi.twitch.tv PRIVMSG #summit1g :@okkiewokkie Happy new year :)
func parseUserType(userType string) string {
	_userType := strings.SplitN(userType, "PRIVMSG", 2)
	if len(_userType) == 2 {
		_userType = strings.SplitN(_userType[1], ":", 2)
		if len(_userType) == 2 {
			return _userType[1]
		}
	}
	return ""
}

// Parse twitch ws response and return Message struct
// Example of ws new chat message response-
//
// @badge-info=;badges=;client-nonce=3b6178753cbb5d5654e546105f1b3714;color=#009EC3;
// display-name=killertrip7;emotes=;first-msg=0;flags=;id=ec1296ae-c5f2-4eec-bf88-74b0088689fb;mod=0;
// room-id=26490481;subscriber=0;tmi-sent-ts=1641025567523;turbo=0;user-id=64035912;
// user-type= :killertrip7!killertrip7@killertrip7.tmi.twitch.tv PRIVMSG #summit1g :!p wing
func ParseTags(tagsB string) (Message, error) {
	combinedTags := string(tagsB)
	combinedTags = strings.ReplaceAll(combinedTags, "\\:", ";")
	combinedTags = strings.ReplaceAll(combinedTags, "\\s", " ")
	tags := strings.Split(combinedTags, ";")

	// Ignore ws messages that are not chat messages
	// bans or communication messages
	if !sliceContains(tags, "display-name", true) {
		return Message{}, fmt.Errorf("igoring ws response as it is not a chat message")
	}

	subscriber := ""
	userType := ""
	mod := ""
	userId := ""
	timestamp := ""
	username := ""
	replyParentMsgBody := ""
	replyParentUserLogin := ""
	systemMsg := ""
	msgParamRecipientUserName := ""
	msgParamRecipientDisplayName := ""
	msgParamGiftMonths := ""
	msgParamMonths := ""
	msgId := ""
	for _, msg := range tags {
		// split on `=`
		if strings.Contains(msg, "=") {
			splitMsg := strings.Split(msg, "=")
			if len(splitMsg) == 2 {
				key := splitMsg[0]
				val := splitMsg[1]
				if sliceContains(IRC_MSG_TAGS, key, false) {
					switch key {
					case "subscriber":
						subscriber = val
					case "user-id":
						userId = val
					case "mod":
						mod = val
					case "tmi-sent-ts":
						timestamp = val
					case "display-name":
						username = val
					case "reply-parent-msg-body":
						replyParentMsgBody = val
					case "reply-parent-user-login":
						replyParentUserLogin = val
					case "user-type":
						userType = val
					case "system-msg":
						systemMsg = val
					case "msg-param-recipient-user-name":
						msgParamRecipientUserName = val
					case "msg-param-recipient-display-name":
						msgParamRecipientDisplayName = val
					case "msg-param-gift-months":
						msgParamGiftMonths = val
					case "msg-param-months":
						msgParamMonths = val
					case "msg-id":
						msgId = val
					}
				}

			}
		}
	}

	// Fetch message from user-type tag
	message := parseUserType(userType)

	return Message{
		Username:                     username,
		Subscriber:                   subscriber,
		UserId:                       userId,
		Message:                      message,
		Timestamp:                    timestamp,
		ReplyTo:                      replyParentMsgBody,
		ReplyToUser:                  replyParentUserLogin,
		Mod:                          mod,
		SystemMsg:                    systemMsg,
		MsgParamRecipientUsername:    msgParamRecipientUserName,
		MsgParamRecipientDisplayName: msgParamRecipientDisplayName,
		MsgParamGiftMonths:           msgParamGiftMonths,
		MsgParamMonths:               msgParamMonths,
		MsgId:                        msgId,
	}, nil
}

// Opens a websocket connection to twitch chat
// Call this function again to reconnect
func InitializeConnection(channel_name string) error {
	u := url.URL{Scheme: "wss", Host: TWITCH_WS_URL}

	_c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	c = _c
	if err != nil {
		return err
	}
	c.WriteMessage(websocket.TextMessage, []byte("CAP REQ :twitch.tv/tags twitch.tv/commands"))
	c.WriteMessage(websocket.TextMessage, []byte("PASS "+PASS))
	c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s%s", "NICK ", TWITCH_NICK)))
	c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("USER %s 8 * :%s", TWITCH_NICK, TWITCH_NICK)))
	c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("JOIN #%s", channel_name)))

	return nil
}

func FetchMessages() (string, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
