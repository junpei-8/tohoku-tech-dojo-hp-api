-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "email" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create index "user_email" to table: "users"
CREATE INDEX "user_email" ON "users" ("email");
-- create "news_list" table
CREATE TABLE "news_list" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "uid" bigint NOT NULL, "title" character varying NOT NULL, "html" text NOT NULL, "markdown" text NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NULL, "creator_id" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "news_list_users_created_news_list" FOREIGN KEY ("creator_id") REFERENCES "users" ("id") ON DELETE CASCADE);
-- create index "news_list_uid_key" to table: "news_list"
CREATE UNIQUE INDEX "news_list_uid_key" ON "news_list" ("uid");
-- create index "news_uid_title_creator_id_creator_id" to table: "news_list"
CREATE INDEX "news_uid_title_creator_id_creator_id" ON "news_list" ("uid", "title", "creator_id", "creator_id");
-- create "tags" table
CREATE TABLE "tags" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "title" character varying NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create index "tags_title_key" to table: "tags"
CREATE UNIQUE INDEX "tags_title_key" ON "tags" ("title");
-- create index "tag_id_title" to table: "tags"
CREATE INDEX "tag_id_title" ON "tags" ("id", "title");
-- create "news_taggings" table
CREATE TABLE "news_taggings" ("news_id" bigint NOT NULL, "tag_id" bigint NOT NULL, PRIMARY KEY ("news_id", "tag_id"), CONSTRAINT "news_taggings_news_list_news" FOREIGN KEY ("news_id") REFERENCES "news_list" ("id") ON DELETE CASCADE, CONSTRAINT "news_taggings_tags_tag" FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE);
-- create index "newstagging_news_id_tag_id_news_id_tag_id" to table: "news_taggings"
CREATE INDEX "newstagging_news_id_tag_id_news_id_tag_id" ON "news_taggings" ("news_id", "tag_id", "news_id", "tag_id");
