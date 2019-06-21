/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 110003
 Source Host           : localhost:5432
 Source Catalog        : blog
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110003
 File Encoding         : 65001

 Date: 21/06/2019 10:10:15
*/


-- ----------------------------
-- Table structure for blog
-- ----------------------------
DROP TABLE IF EXISTS "public"."blog";
CREATE TABLE "public"."blog" (
  "id" int4 NOT NULL DEFAULT nextval('blog_id_seq'::regclass),
  "uuid" varchar(64) COLLATE "pg_catalog"."default",
  "useruuid" varchar(64) COLLATE "pg_catalog"."default",
  "title" varchar(64) COLLATE "pg_catalog"."default",
  "info" text COLLATE "pg_catalog"."default",
  "build_time" timestamp(6),
  "status" varchar(16) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."blog" OWNER TO "liudong";

-- ----------------------------
-- Uniques structure for table blog
-- ----------------------------
ALTER TABLE "public"."blog" ADD CONSTRAINT "blog_uuid_key" UNIQUE ("uuid");

-- ----------------------------
-- Primary Key structure for table blog
-- ----------------------------
ALTER TABLE "public"."blog" ADD CONSTRAINT "blog_pkey" PRIMARY KEY ("id");
