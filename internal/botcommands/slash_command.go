package botcommands

import "regexp"

var (
	userRegex *regexp.Regexp = regexp.MustCompile(`<@(?P<userid>\w+)(?:\|(?P<username>[\w._-]+))?>`)
)

type UserInfo struct {
	UserID   string
	UserName string
}

func ParseSlackTzCommand(text string) []UserInfo {
	var users = make([]UserInfo, 0)
	for _, userMatch := range userRegex.FindAllStringSubmatch(text, -1) {
		users = append(users, UserInfo{
			UserID:   userMatch[1],
			UserName: userMatch[2],
		})
	}

	return users
}
