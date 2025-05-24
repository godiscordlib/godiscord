package enums

type ActivityFlags int

const (
	AF_Instance    ActivityFlags = 1 << iota // 1 << 0 = 1
	AF_Join                                  // 1 << 1 = 2
	AF_Spectate                              // 1 << 2 = 4
	AF_JoinRequest                           // 1 << 3 = 8
	AF_Sync                                  // 1 << 4 = 16
)
