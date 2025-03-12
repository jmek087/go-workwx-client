package workwx

// CreateTag 创建标签
func (c *WorkwxApp) CreateTag(tagName string) (int, error) {
	resp, err := c.execTagCreate(reqTagCreate{
		TagName: tagName,
	})
	if err != nil {
		return 0, err
	}
	return resp.TagID, nil
}

// UpdateTag 更新标签名字
func (c *WorkwxApp) UpdateTag(tagID int, tagName string) error {
	_, err := c.execTagUpdate(reqTagUpdate{
		TagID:   tagID,
		TagName: tagName,
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag 删除标签
func (c *WorkwxApp) DeleteTag(tagID int) error {
	_, err := c.execTagDelete(reqTagDelete{
		TagID: tagID,
	})
	if err != nil {
		return err
	}
	return nil
}

// GetTagUsers 获取标签成员
func (c *WorkwxApp) GetTagUsers(tagID int) ([]TagUser, error) {
	resp, err := c.execTagListUsers(reqTagListUsers{
		TagID: tagID,
	})
	if err != nil {
		return nil, err
	}
	return resp.UserList, nil
}

// AddTagUsers 增加标签成员
func (c *WorkwxApp) AddTagUsers(tagID int, userList []string, partyList []int) (string, error) {
	resp, err := c.execTagAddUsers(reqTagAddUsers{
		TagID:     tagID,
		UserList:  userList,
		PartyList: partyList,
	})
	if err != nil {
		return "", err
	}
	return resp.InvalidList, nil
}

// DeleteTagUsers 删除标签成员
func (c *WorkwxApp) DeleteTagUsers(tagID int, userList []string, partyList []int) (string, error) {
	resp, err := c.execTagDeleteUsers(reqTagDeleteUsers{
		TagID:     tagID,
		UserList:  userList,
		PartyList: partyList,
	})
	if err != nil {
		return "", err
	}
	return resp.InvalidList, nil
}

// ListTags 获取标签列表
func (c *WorkwxApp) ListTags() ([]Tag, error) {
	resp, err := c.execTagList(reqTagList{})
	if err != nil {
		return nil, err
	}
	return resp.TagList, nil
}

// AddTagToUsers 批量给用户打标签
func (c *WorkwxApp) AddTagToUsers(tagID int, userIDs []string) (string, error) {
	return c.AddTagUsers(tagID, userIDs, nil)
}

// AddTagToDepts 批量给部门打标签
func (c *WorkwxApp) AddTagToDepts(tagID int, deptIDs []int) (string, error) {
	return c.AddTagUsers(tagID, nil, deptIDs)
}

// HasTag 检查用户是否有某个标签
func (c *WorkwxApp) HasTag(userID string, tagID int) (bool, error) {
	users, err := c.GetTagUsers(tagID)
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.UserID == userID {
			return true, nil
		}
	}
	return false, nil
}
