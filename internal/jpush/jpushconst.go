package jpush

// @Project: go-push-api
// @Author: houseme
// @Description:
// @File: jpushconst
// @Version: 1.0.0
// @Date: 2021/6/28 01:27
// @Package builder

// AudienceType .
type AudienceType string

const (
	// Tag .
	Tag = "tag"
	// TagAnd .
	TagAnd = "tag_and"
	// TagNot .
	TagNot = "tag_not"
	// ALIAS .
	ALIAS = "alias"
	// Segment .
	Segment = "segment"
	// Abtest .
	Abtest = "abtest"
	// RegistrationID .
	RegistrationID = "registration_id"
	// File .
	File = "file"
)
