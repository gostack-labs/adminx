-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-05-19T12:15:01.535Z

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL DEFAULT '',
  "phone" varchar NOT NULL DEFAULT '',
  "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "is_disable" boolean NOT NULL DEFAULT false,
  "key" varchar UNIQUE NOT NULL,
  "sort" int NOT NULL DEFAULT 1,
  "remark" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "menus" (
  "id" bigserial PRIMARY KEY,
  "parent" bigint NOT NULL,
  "title" varchar NOT NULL,
  "path" varchar,
  "name" varchar NOT NULL,
  "component" varchar,
  "redirect" varchar,
  "hyperlink" varchar,
  "is_hide" boolean NOT NULL DEFAULT false,
  "is_keep_alive" boolean NOT NULL DEFAULT true,
  "is_affix" boolean NOT NULL DEFAULT false,
  "is_iframe" boolean NOT NULL DEFAULT false,
  "auth" text[] NOT NULL,
  "icon" varchar,
  "type" int NOT NULL DEFAULT 1,
  "sort" int NOT NULL DEFAULT 1,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "role_menus" (
  "id" bigserial PRIMARY KEY,
  "role" bigint NOT NULL,
  "menu" bigint NOT NULL,
  "type" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "api_groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "remark" text,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "apis" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "url" varchar NOT NULL,
  "method" varchar NOT NULL,
  "groups" bigint NOT NULL,
  "remark" text,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "menu_apis" (
  "id" bigserial PRIMARY KEY,
  "menu" bigint NOT NULL,
  "api" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "role_menus" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id");

ALTER TABLE "role_menus" ADD FOREIGN KEY ("menu") REFERENCES "menus" ("id");

ALTER TABLE "apis" ADD FOREIGN KEY ("groups") REFERENCES "api_groups" ("id");

ALTER TABLE "menu_apis" ADD FOREIGN KEY ("menu") REFERENCES "menus" ("id");

ALTER TABLE "menu_apis" ADD FOREIGN KEY ("api") REFERENCES "apis" ("id");

CREATE INDEX ON "role_menus" ("role");

CREATE INDEX ON "role_menus" ("menu");

CREATE INDEX ON "role_menus" ("role", "menu");

COMMENT ON COLUMN "users"."username" IS '主键，用户名';

COMMENT ON COLUMN "users"."hashed_password" IS '加密后密码';

COMMENT ON COLUMN "users"."full_name" IS '全名';

COMMENT ON COLUMN "users"."email" IS '邮箱';

COMMENT ON COLUMN "users"."phone" IS '手机号';

COMMENT ON COLUMN "users"."password_change_at" IS '修改密码时间';

COMMENT ON COLUMN "sessions"."username" IS '用户名，关联Users表username字段';

COMMENT ON COLUMN "sessions"."refresh_token" IS '刷新密钥';

COMMENT ON COLUMN "sessions"."user_agent" IS '用户代理';

COMMENT ON COLUMN "sessions"."client_ip" IS 'ip';

COMMENT ON COLUMN "sessions"."is_blocked" IS '是否屏蔽';

COMMENT ON COLUMN "sessions"."expires_at" IS '过期时间';

COMMENT ON TABLE "roles" IS '角色表';

COMMENT ON COLUMN "roles"."name" IS '名称';

COMMENT ON COLUMN "roles"."is_disable" IS '是否禁用';

COMMENT ON COLUMN "roles"."key" IS '标识';

COMMENT ON COLUMN "roles"."sort" IS '排序';

COMMENT ON COLUMN "roles"."remark" IS '备注';

COMMENT ON TABLE "menus" IS '菜单表';

COMMENT ON COLUMN "menus"."parent" IS '父级';

COMMENT ON COLUMN "menus"."title" IS '标题';

COMMENT ON COLUMN "menus"."path" IS '路径';

COMMENT ON COLUMN "menus"."name" IS '路由名称';

COMMENT ON COLUMN "menus"."component" IS '组件路径';

COMMENT ON COLUMN "menus"."redirect" IS '跳转路径';

COMMENT ON COLUMN "menus"."hyperlink" IS '超链接';

COMMENT ON COLUMN "menus"."is_hide" IS '是否隐藏';

COMMENT ON COLUMN "menus"."is_keep_alive" IS '是否缓存组件状态';

COMMENT ON COLUMN "menus"."is_affix" IS '是否固定在标签栏';

COMMENT ON COLUMN "menus"."is_iframe" IS '是否内嵌窗口';

COMMENT ON COLUMN "menus"."auth" IS '权限粒子';

COMMENT ON COLUMN "menus"."icon" IS '图标';

COMMENT ON COLUMN "menus"."type" IS '类型：1 目录，2 菜单，3 按钮';

COMMENT ON COLUMN "menus"."sort" IS '顺序';

COMMENT ON TABLE "role_menus" IS '角色菜单关联表';

COMMENT ON COLUMN "role_menus"."role" IS '角色';

COMMENT ON COLUMN "role_menus"."menu" IS '菜单';

COMMENT ON COLUMN "role_menus"."type" IS '类型：1 菜单，2 按钮';

COMMENT ON TABLE "api_groups" IS '接口组表';

COMMENT ON COLUMN "api_groups"."name" IS '名称';

COMMENT ON COLUMN "api_groups"."remark" IS '备注';

COMMENT ON TABLE "apis" IS '接口表';

COMMENT ON COLUMN "apis"."title" IS '标题';

COMMENT ON COLUMN "apis"."url" IS '接口地址';

COMMENT ON COLUMN "apis"."method" IS '请求方式';

COMMENT ON COLUMN "apis"."groups" IS '分组';

COMMENT ON COLUMN "apis"."remark" IS '备注';

COMMENT ON TABLE "menu_apis" IS '菜单接口关联表';

COMMENT ON COLUMN "menu_apis"."menu" IS '菜单';

COMMENT ON COLUMN "menu_apis"."api" IS '接口';
