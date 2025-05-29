package types

type DefaultForumSortOrder int

const (
	DFSO_LatestActivity DefaultForumSortOrder = iota
	DFSO_CreationDate
)
