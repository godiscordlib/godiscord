package utils

func GetReason(Reason ...string) string {
	if len(Reason) > 1 {
		return Reason[0]
	}
	return ""
}
