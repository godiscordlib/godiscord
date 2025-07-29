package utils

func RoleString(id string) string {
	return "<@&" + id + ">"
}

func UserString(id string) string {
	return "<@" + id + ">"
}

func ChannelString(id string) string {
	return "<#" + id + ">"
}
