-- 用户
INSERT INTO users (username, hashed_password,full_name, phone, email) VALUES ('adminx', '$2a$10$6vd7N7ujEJgJEMYQduqk1.ijKqBtYeVr6ha6WHO9esJFXSbvJIiuO', 'adminx', '13400000000', 'chenzhenying@88.com');

-- 菜单
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/home', 'Home', 'home/index', '', '首页', '', false, true, true, false, '{home}', 'HomeFilled', 0, 2, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/system', 'System', 'layout/routerView/parent', '', '系统设置', '', false, true, false, false, '{system}', 'Setting', 0, 1, 2);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/system/menu', 'SystemMenu', 'system/menu/index', '', '菜单管理', '', false, true, false, false, '{system:menu}', 'Tickets', 2, 2, 15);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/system/user', 'SystemUser', 'system/user/index', '', '用户管理', '', false, true, false, false, '{system:user}', 'User', 2, 2, 1);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/system/role', 'SystemRole', 'system/role/index', '', '角色管理', '', false, true, false, false, '{system:role}', 'SetUp', 2, 2, 5);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/system/api', 'SystemAPI', 'system/api/index', '', '接口管理', '', false, true, false, false, '{system:api}', 'Promotion', 2, 2, 10);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('/system/role/permission/:id', 'SystemRolePermission', 'system/role/permission', '', '角色权限', '', true, true, false, false, '{system:role:permission}', 'Minus', 2, 2, 99);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '新建用户', '', false, false, false, false, '{system:user:create}', '', 4, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '编辑用户', '', false, false, false, false, '{system:user:edit}', '', 4, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '删除用户', '', false, false, false, false, '{system:user:delete}', '', 4, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '新建角色', '', false, false, false, false, '{system:role:create}', '', 5, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '角色权限', '', false, false, false, false, '{system:role:permission}', '', 5, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '编辑角色', '', false, false, false, false, '{system:role:edit}', '', 5, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '删除角色', '', false, false, false, false, '{system:role:delete}', '', 5, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '新建接口', '', false, false, false, false, '{system:api:create}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '编辑接口', '', false, false, false, false, '{system:api:edit}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '删除接口', '', false, false, false, false, '{system:api:delete}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '接口分组列表', '', false, false, false, false, '{system:api-group:list}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '新建接口分组', '', false, false, false, false, '{system:api-group:create}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '编辑接口分组', '', false, false, false, false, '{system:api-group:edit}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '删除接口分组', '', false, false, false, false, '{system:api-group:delete}', '', 6, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '新建顶级菜单', '', false, false, false, false, '{system:top-menu:create}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '保存菜单表单', '', false, false, false, false, '{system:menu-form:save}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '右键新建菜单', '', false, false, false, false, '{system:menu:create}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '右键删除菜单', '', false, false, false, false, '{system:menu:delete}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '新建页面按钮', '', false, false, false, false, '{system:menu-button:create}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '编辑页面按钮', '', false, false, false, false, '{system:menu-button:edit}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '删除页面按钮', '', false, false, false, false, '{system:menu-button:delete}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '菜单接口管理', '', false, false, false, false, '{system:menu-api:manager}', '', 3, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '保存权限', '', false, false, false, false, '{system:role-permission:save}', '', 7, 3, 0);
INSERT INTO menus (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort) VALUES ('', '', '', '', '接口权限管理', '', false, false, false, false, '{system:role-permission-api:manager}', '', 7, 3, 0);

-- API分组
INSERT INTO api_groups (name, remark) VALUES ('系统管理', '');

-- API接口
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取用户列表', '/sys/user', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取用户详情', '/sys/user/info', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('通过ID获取用户详情', '/sys/user/info/:username', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('新建用户', '/sys/user', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('编辑用户', '/sys/user/:username', 'PUT', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('删除用户', '/sys/user/:username', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取菜单树', '/sys/menu/tree', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('新建菜单', '/sys/menu', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('编辑菜单', '/sys/menu/:id', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('删除菜单', '/sys/menu/single/:id', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('批量删除菜单', '/sys/menu/batch', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取菜单对应的按钮', '/sys/menu/button/:id', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('菜单绑定/解绑接口', '/sys/menu/api/:id', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取菜单绑定的接口', '/sys/menu/api/:menu', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取菜单对应的接口', '/sys/menu/api-list/:menu', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取角色列表', '/sys/role', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('新建角色', '/sys/role', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('编辑角色', '/sys/role/:id', 'PUT', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('删除角色', '/sys/role/single/:id', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('批量删除角色', '/sys/role/batch', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('角色授权/解除菜单权限', '/sys/role/permission/:id', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('角色授权/解除接口权限', '/sys/role/api/:id', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取角色授权的接口权限', '/sys/role/api/:id', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取角色授权的菜单权限', '/sys/role/permission/:id', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('获取接口列表', '/sys/api', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('新建接口', '/sys/api', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('编辑接口', '/sys/api/:id', 'PUT', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('删除接口', '/sys/api/single/:id', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('批量删除接口', '/sys/api/batch', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('接口分组列表', '/sys/api-group', 'GET', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('新建接口分组', '/sys/api-group', 'POST', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('编辑接口分组', '/sys/api-group', 'PUT', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('删除接口分组', '/sys/api-group/single/:id', 'DELETE', 1, '');
INSERT INTO apis (title, url, method, groups, remark) VALUES ('批量删除接口分组', '/sys/api-group/batch', 'DELETE', 1, '');