package util

import (
	"fmt"
	"strings"
)

//go:generate mockery --name=FilterString
type FilterString interface {
	MakeID(id string) (filter string)
	MakeIDFilters(id string) (filters []string)
	MakeNotEqualID(id string) (filter string)
	MakeUserID(userID string) (filter string)
	MakeKeySlug(keySlug string) (filter string)
	MakeUserName(username string) (filter string)
	MakePassword(password string) (filter string)
	MakeEmail(email string) (filter string)
	MakeUserGroup(status string) (filter string)
	MakeStatus(status string) (filter string)
	MakeStatusIn(status ...string) (filter string)
	MakeDeletedAtIsNull() (filter string)
	MakeUserAgent(userAgent string) (filter string)
	MakeAccessToken(accessToken string) (filter string)
	MakeAccessTokenExpiredAtGreaterThan(time int64) (filter string)
	MakeRefreshTokenExpiredAtLessThan(time int64) (filter string)
	MakeRefreshToken(refreshToken string) (filter string)
	MakeUserGroupRoleID(roleID string) (filter string)
	MakeEqualUserType(userType string) (filter string)
	MakeVerified(val bool) (filter string)
	MakeRoleID(roleID string) (filter string)
	MakeActivateTokenString(activateTokenString string) (filter string)
	MakeActivateExpiredAtGreaterThan(time int64) (filter string)
	MakeVerifiedBool(val bool) (filter string)
	MakePermission(key string) (filter string)
	MakeSubMenuIn(subMenu []string) (filter string)
}

type filter struct{}

func NewFilterString() (f FilterString) {
	return &filter{}
}

func (f *filter) MakeID(id string) (filter string) {
	return fmt.Sprintf("id:eq:%s", id)
}

func (f *filter) MakeIDFilters(id string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", id),
	}
}

func (f *filter) MakeNotEqualID(id string) (filter string) {
	return fmt.Sprintf("_id:ne:%s", id)
}

func (f *filter) MakeUserID(userID string) (filter string) {
	return fmt.Sprintf("user_id:eq:%s", userID)
}

func (f *filter) MakeKeySlug(keySlug string) (filter string) {
	return fmt.Sprintf("key_slug:eq:%s", keySlug)
}

func (f *filter) MakeUserName(username string) (filter string) {
	return fmt.Sprintf("username:eq:%s", username)
}

func (f *filter) MakePassword(password string) (filter string) {
	return fmt.Sprintf("password:eq:%s", password)
}

func (f *filter) MakeEmail(email string) (filter string) {
	return fmt.Sprintf("email:eq:%s", email)
}

func (f *filter) MakeUserGroup(user_group string) (filter string) {
	return fmt.Sprintf("user_group:eq:%s", user_group)
}

func (f *filter) MakeStatus(status string) (filter string) {
	return fmt.Sprintf("status:eq:%s", status)
}

func (f *filter) MakeStatusIn(status ...string) (filter string) {
	return fmt.Sprintf("status:in:%s", strings.Join(status, "|"))
}

func (f *filter) MakeDeletedAtIsNull() (filter string) {
	return "deleted_at:isNull"
}

func (f *filter) MakeUserAgent(userAgent string) (filter string) {
	return fmt.Sprintf("user_agent:eq:%s", userAgent)
}

func (f *filter) MakeAccessToken(accessToken string) (filter string) {
	return fmt.Sprintf("access_token:eq:%s", accessToken)
}

func (f *filter) MakeAccessTokenExpiredAtGreaterThan(time int64) (filter string) {
	return fmt.Sprintf("access_token_exp_at:gt:%d", time)
}

func (f *filter) MakeRefreshTokenExpiredAtLessThan(time int64) (filter string) {
	return fmt.Sprintf("refresh_token_exp_at:lt:%d", time)
}

func (f *filter) MakeRefreshToken(refreshToken string) (filter string) {
	return fmt.Sprintf("refresh_token:eq:%s", refreshToken)
}

func (f *filter) MakeUserGroupRoleID(roleID string) (filter string) {
	return fmt.Sprintf("roles:elemMatch:eq:%s", roleID)
}

func (f *filter) MakeEqualUserType(userType string) (filter string) {
	return fmt.Sprintf("user_type:eq:%s", userType)
}

func (f *filter) MakeVerified(val bool) (filter string) {
	return fmt.Sprintf("verified:eq:%t", val)
}

func (f *filter) MakeRoleID(roleID string) (filter string) {
	return fmt.Sprintf("role:eq:%s", roleID)
}

func (f *filter) MakeActivateTokenString(activateToken string) (filter string) {
	return fmt.Sprintf("activate_token:eq:%s", activateToken)
}

func (f *filter) MakeActivateExpiredAtGreaterThan(time int64) (filter string) {
	return fmt.Sprintf("activate_expired_at:gt:%d", time)
}

func (f *filter) MakeVerifiedBool(val bool) (filter string) {
	return fmt.Sprintf("verified:eqBool:%t", val)
}

func (f *filter) MakePermission(key string) (filter string) {
	return fmt.Sprintf("permissions.key_slug:eq:%s", key)
}

func (f *filter) MakeSubMenuIn(subMenu []string) (filter string) {
	return fmt.Sprintf("sub_menus.key_slug:in:%s", strings.Join(subMenu, "|"))
}
