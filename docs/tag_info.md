# 标签管理

## Models

### `Tag` 标签信息
Name|Type|Doc
:---|:---|:--
`TagID`|`int`|标签ID，非负整型，指定此参数时new_id无效
`TagName`|`string`|标签名称，长度限制为32个字（汉字或英文字符），标签不可与其他标签重名

### `TagList` 标签列表
Name|Type|Doc
:---|:---|:--
`TagList`|`[]Tag`|标签列表

### `TagUser` 标签成员
Name|Type|Doc
:---|:---|:--
`UserID`|`string`|成员UserID，在管理端可获取到
`Name`|`string`|成员名称，此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有存量第三方应用不再返回，后续第三方仅通讯录应用可获取，第三方页面需要通过通讯录展示组件来展示名字

### `TagMember` 标签成员信息
Name|Type|Doc
:---|:---|:--
`TagID`|`int`|标签ID
`UserList`|`[]TagUser`|标签成员列表
`PartyList`|`[]int`|标签所在的部门列表
