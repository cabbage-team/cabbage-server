/*
 Navicat Premium Data Transfer

 Source Server         : myMachine
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 127.0.0.1:3308
 Source Schema         : cabbage

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 12/12/2023 13:55:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cabbage_comment
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_comment`;
CREATE TABLE `cabbage_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `post_id` bigint(20) NULL DEFAULT NULL,
  `parent` bigint(20) NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `reply` bigint(20) NULL DEFAULT NULL,
  `create_by` bigint(20) NULL DEFAULT NULL,
  `like` bigint(20) NULL DEFAULT NULL,
  `diss` bigint(20) NULL DEFAULT NULL,
  `share` bigint(20) NULL DEFAULT NULL,
  `favorite` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_comment_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_cabbage_comment_post_id`(`post_id`) USING BTREE,
  INDEX `idx_cabbage_comment_parent`(`parent`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_comment_operator
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_comment_operator`;
CREATE TABLE `cabbage_comment_operator`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `comment_id` bigint(20) NULL DEFAULT NULL,
  `op_code` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_comment_operator_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_cabbage_comment_operator_user_id`(`user_id`) USING BTREE,
  INDEX `idx_cabbage_comment_operator_comment_id`(`comment_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_post
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_post`;
CREATE TABLE `cabbage_post`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `author` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `like` bigint(20) NULL DEFAULT NULL,
  `diss` bigint(20) NULL DEFAULT NULL,
  `share` bigint(20) NULL DEFAULT NULL,
  `favorite` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_post_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_post_operator
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_post_operator`;
CREATE TABLE `cabbage_post_operator`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `post_id` bigint(20) NULL DEFAULT NULL,
  `op_code` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_post_operator_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_cabbage_post_operator_user_id`(`user_id`) USING BTREE,
  INDEX `idx_cabbage_post_operator_post_id`(`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_post_tag
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_post_tag`;
CREATE TABLE `cabbage_post_tag`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `post_id` bigint(20) NULL DEFAULT NULL,
  `tag_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_post_tag_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_cabbage_post_tag_post_id`(`post_id`) USING BTREE,
  INDEX `idx_cabbage_post_tag_tag_id`(`tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_tag
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_tag`;
CREATE TABLE `cabbage_tag`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tag_id` binary(16) NOT NULL COMMENT '\'uuid\'',
  `tag_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'tag name\'',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '\'tag description\'',
  `usage_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '\'usage count\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_tag_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_user
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_user`;
CREATE TABLE `cabbage_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'uuid\'',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(320) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `email`(`email`) USING BTREE,
  INDEX `idx_cabbage_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_user_follows
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_user_follows`;
CREATE TABLE `cabbage_user_follows`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uid` bigint(20) NULL DEFAULT NULL,
  `follower` bigint(20) NULL DEFAULT NULL,
  `ship` enum('F','N','D') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'D',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `follow`(`follower`) USING BTREE,
  INDEX `idx_cabbage_user_follows_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_cabbage_user_follows_uid`(`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_user_profile
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_user_profile`;
CREATE TABLE `cabbage_user_profile`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `platform` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `uri` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `only_self` enum('Y','N') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N',
  `only_friend` enum('Y','N') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N',
  `public` enum('Y','N') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Y',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_user_profile_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_cabbage_user_profile_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_user_profile_platform
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_user_profile_platform`;
CREATE TABLE `cabbage_user_profile_platform`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `platform` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_user_profile_platform_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabbage_user_tag
-- ----------------------------
DROP TABLE IF EXISTS `cabbage_user_tag`;
CREATE TABLE `cabbage_user_tag`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` binary(16) NOT NULL COMMENT '\'uuid\'',
  `tag_id` binary(16) NOT NULL COMMENT '\'uuid\'',
  `relationship` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '\'1: 正常状态, 2: follow, 3: hide\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cabbage_user_tag_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
