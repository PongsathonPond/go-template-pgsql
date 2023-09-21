package implement

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"idev-cms-service/domain"
	"idev-cms-service/service/roles/inout"
)

func (impl *implementation) CheckPermission(ctx context.Context, input *inout.CheckPermissionInput) (isPermission bool) {
	endpointStripBasePath := strings.Replace(input.EndPoint, fmt.Sprintf("%s/", impl.Config.SwaggerInfoBasePath), "", 1)
	endpointStripPathID := strings.Replace(endpointStripBasePath, "/:id", "", 1)
	endpointStripPathKey := strings.Replace(endpointStripPathID, "/:key", "", 1)
	endpointStripPathKeySlug := strings.Replace(endpointStripPathKey, "/:key_slug", "", 1)
	keySlug := mapMenu[input.Service][endpointStripPathKeySlug]

	role := &domain.Roles{}
	filters := []string{
		impl.FilterString.MakeID(input.RoleID),
		impl.FilterString.MakePermission(keySlug),
		impl.FilterString.MakeStatus("active"),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err := impl.Repo.Read(ctx, filters, role); err != nil {
		impl.Log.Error("CheckPermissionError: ", err)
		return false
	}

	for _, v := range role.Permissions {
		if v.KeySlug == keySlug {
			perm := impl.Perm.ConvertLevelToCRUD(v.Permission)
			switch input.Method {
			case http.MethodPost:
				return perm.CanCreate
			case http.MethodGet:
				return perm.CanRead
			case http.MethodPut:
				return perm.CanUpdate
			case http.MethodDelete:
				return perm.CanDelete
			}
			break
		}
	}

	return false
}
