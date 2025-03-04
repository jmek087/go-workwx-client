package workwx

import (
	"fmt"
	"strconv"
)

func reshapeDeptInfo(
	ids []int64,
	orders []uint32,
	leaderStatuses []int,
) ([]UserDeptInfo, error) {
	if len(ids) != len(orders) {
		return nil, fmt.Errorf(
			"server API breakage: len(DeptIDs) (%d) != len(DeptOrder) (%d)",
			len(ids),
			len(orders),
		)
	}
	// sometimes leaderStatuses could be empty, but if not, the arrays should
	// have equal length
	if len(ids) != len(leaderStatuses) && len(leaderStatuses) != 0 {
		return nil, fmt.Errorf(
			"server API breakage: len(DeptIDs) (%d) != len(IsLeaderInDept) (%d)",
			len(ids),
			len(leaderStatuses),
		)
	}

	result := make([]UserDeptInfo, len(ids))
	for i := range ids {
		result[i].DeptID = ids[i]
		result[i].Order = orders[i]
		if i < len(leaderStatuses) {
			// apparently leaderStatuses could sometimes be empty, don't set
			// anybody as leader in that case
			// see https://github.com/jmek087/go-workwx-client/pull/78
			result[i].IsLeader = leaderStatuses[i] != 0
		}
	}

	return result, nil
}

func userGenderFromGenderStr(x string) (UserGender, error) {
	if x == "" {
		return UserGenderUnspecified, nil
	}
	n, err := strconv.Atoi(x)
	if err != nil {
		return UserGenderUnspecified, fmt.Errorf("gender string parse failed: %+v", err)
	}

	return UserGender(n), nil
}

func (x UserDetail) intoUserInfo() (UserInfo, error) {
	deptInfo, err := reshapeDeptInfo(x.DeptIDs, x.DeptOrder, x.IsLeaderInDept)
	if err != nil {
		return UserInfo{}, err
	}

	gender, err := userGenderFromGenderStr(x.Gender)
	if err != nil {
		return UserInfo{}, err
	}

	// 转换扩展属性
	var extAttrs *UserExtAttrs
	if x.ExtAttr != nil {
		attrs := make([]UserExtAttr, len(x.ExtAttr.Attrs))
		for i, attr := range x.ExtAttr.Attrs {
			attrs[i] = UserExtAttr{
				Type: attr.Type,
				Name: attr.Name,
			}
			if attr.Text != nil {
				attrs[i].Text = &UserExtAttrText{
					Value: attr.Text.Value,
				}
			}
			if attr.Web != nil {
				attrs[i].Web = &UserExtAttrWeb{
					URL:   attr.Web.URL,
					Title: attr.Web.Title,
				}
			}
		}
		extAttrs = &UserExtAttrs{
			Attrs: attrs,
		}
	}

	return UserInfo{
		UserID:         x.UserID,
		Name:           x.Name,
		Position:       x.Position,
		Departments:    deptInfo,
		Mobile:         x.Mobile,
		Gender:         gender,
		Email:          x.Email,
		AvatarURL:      x.AvatarURL,
		Telephone:      x.Telephone,
		IsEnabled:      x.IsEnabled != 0,
		Alias:          x.Alias,
		Status:         UserStatus(x.Status),
		QRCodeURL:      x.QRCodeURL,
		MainDepartment: x.MainDepartment,
		DirectLeader:   x.DirectLeader,
		ExtAttr:        extAttrs,
	}, nil
}
