package common

func getReason(Reason ...string) string {
	if len(Reason) > 1 {
		return Reason[0]
	}
	return ""
}
