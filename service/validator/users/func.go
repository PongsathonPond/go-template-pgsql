package users

import (
	"context"
	"fmt"
	"regexp"

	"idev-cms-service/domain"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateUser) checkEmailUnique(ctx context.Context, structLV validator.StructLevel, email string) {
	filters := []string{
		fmt.Sprintf("email:eq:%s", email),
		fmt.Sprintf("deleted_at:isNull"),
	}
	if err := c.repo.Read(ctx, filters, &domain.Users{}); err == nil {
		structLV.ReportError(email, "Email", "email", "unique", "")
	}
}

func (c *customValidateUser) checkUpdateEmailUnique(ctx context.Context, structLV validator.StructLevel, id string, email string) {
	filters := []string{
		fmt.Sprintf("_id:ne:%s", id),
		fmt.Sprintf("email:eq:%s", email),
		fmt.Sprintf("deleted_at:isNull"),
	}
	if err := c.repo.Read(ctx, filters, &domain.Users{}); err == nil {
		structLV.ReportError(email, "Email", "email", "unique", "")
	}
}

//func (c *customValidateUser) checkRolesExist(structLV validator.StructLevel, roleID string) {
//	grpcRole, err := c.grpcRoleService.GetRole(&pb.RoleRequest{RoleId: roleID})
//	if err != nil {
//		structLV.ReportError(roleID, "RoleID", "role_id", "not_exist", "")
//	}
//	fmt.Println("grpcRole", grpcRole)
//}

func (c *customValidateUser) checkFormatMobileNumber(structLV validator.StructLevel, mobileNumber string) {
	if mobileNumber != "" {
		re := regexp.MustCompile("^[0-9]{10}$")
		ok := re.MatchString(mobileNumber)
		if !ok {
			structLV.ReportError(mobileNumber, "mobileNumber", "mobileNumber", "match", "")
		}
	}
}

func (c *customValidateUser) checkPasswordLength(structLv validator.StructLevel, password string) {
	if len([]rune(password)) < 8 {
		structLv.ReportError(password, "Password", "Password", "length", "")
	}
}
