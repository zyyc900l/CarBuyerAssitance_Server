/*
 Navicat Premium Dump SQL

 Source Server         : db
 Source Server Type    : MySQL
 Source Server Version : 90200 (9.2.0)
 Source Host           : localhost:3306
 Source Schema         : car_consult_system

 Target Server Type    : MySQL
 Target Server Version : 90200 (9.2.0)
 File Encoding         : 65001

 Date: 11/11/2025 12:06:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for consultation
-- ----------------------------
DROP TABLE IF EXISTS `consultation`;
CREATE TABLE `consultation`  (
                                 `consult_id` int NOT NULL AUTO_INCREMENT COMMENT '自增咨询ID',
                                 `user_id` varchar(50) NOT NULL COMMENT '用户ID',
                                 `budget_range` varchar(50) NOT NULL COMMENT '预算范围',
                                 `preferred_type` varchar(20) NOT NULL COMMENT '偏好车型',
                                 `use_case` varchar(30) NOT NULL COMMENT '主要使用场景',
                                 `fuel_type` varchar(20) NOT NULL COMMENT '燃料类型偏好',
                                 `brand_preference` varchar(50) NULL DEFAULT '' COMMENT '品牌偏好',
                                 `consult_content` text NULL COMMENT '用户补充咨询内容',
                                 `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                 PRIMARY KEY (`consult_id`) USING BTREE,
                                 INDEX `idx_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '咨询记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for consult_result
-- ----------------------------
DROP TABLE IF EXISTS `consult_result`;
CREATE TABLE `consult_result`  (
                                   `result_id` int NOT NULL AUTO_INCREMENT COMMENT '自增结果ID',
                                   `consult_id` int NOT NULL COMMENT '关联咨询ID',
                                   `analysis` text NOT NULL COMMENT '分析内容',
                                   `proposal` text NOT NULL COMMENT '总的购车建议',
                                   `recommend_cars` json NOT NULL COMMENT '推荐车型列表（JSON格式存储Car数组）',
                                   `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   PRIMARY KEY (`result_id`) USING BTREE,
                                   INDEX `idx_consult_id`(`consult_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '咨询结果表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for gift
-- ----------------------------
DROP TABLE IF EXISTS `gift`;
CREATE TABLE `gift`  (
                         `gift_id` bigint NOT NULL AUTO_INCREMENT COMMENT '礼品唯一ID',
                         `gift_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '礼品名称',
                         `required_points` int NOT NULL COMMENT '兑换所需积分',
                         `stock_quantity` int NOT NULL DEFAULT 0 COMMENT '库存数量',
                         `cover_image_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '礼品封面图片URL',
                         `is_online` tinyint NOT NULL DEFAULT 1 COMMENT '是否上架（1-上架，0-下架）',
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         PRIMARY KEY (`gift_id`) USING BTREE,
                         INDEX `idx_points`(`required_points` ASC) USING BTREE,
                         INDEX `idx_status`(`is_online` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '汽车周边礼品表（用于积分兑换）' ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `gift_exchange`;
CREATE TABLE `gift_exchange`  (
                                  `exchange_id` int NOT NULL AUTO_INCREMENT COMMENT '自增兑换ID',
                                  `user_id` varchar(50) NOT NULL COMMENT '用户ID',
                                  `gift_name` varchar(50) NOT NULL COMMENT '礼品名称（汽车周边）',
                                  `need_points` int NOT NULL COMMENT '所需积分',
                                  `exchange_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '兑换时间',
                                  `status` tinyint NULL DEFAULT 0 COMMENT '兑换状态（0-待发货/1-已完成）',
                                  PRIMARY KEY (`exchange_id`) USING BTREE,
                                  INDEX `idx_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '礼品兑换表' ROW_FORMAT = Dynamic;
-- ----------------------------
-- Table structure for llm_config
-- ----------------------------
DROP TABLE IF EXISTS `llm_config`;
CREATE TABLE `llm_config`  (
                               `llm_id` int NOT NULL AUTO_INCREMENT COMMENT '自增LLMID',
                               `llm_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'LLM名称',
                               `api_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '加密存储的接口密钥',
                               `api_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '接口请求地址',
                               `status` tinyint NULL DEFAULT 1 COMMENT '接口状态（0-禁用/1-正常）',
                               `weight` tinyint NULL DEFAULT 5 COMMENT '负载均衡权重（1-10）',
                               PRIMARY KEY (`llm_id`) USING BTREE,
                               UNIQUE INDEX `idx_llm_name`(`llm_name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'LLM配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for points
-- ----------------------------
DROP TABLE IF EXISTS `points`;
CREATE TABLE `points`  (
                           `point_id` int NOT NULL AUTO_INCREMENT COMMENT '自增积分记录ID',
                           `user_id` varchar(50) NOT NULL COMMENT '用户ID',
                           `points` int NOT NULL COMMENT '变动积分（正增负减）',
                           `reason` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '积分变动原因（完成咨询/反馈等）',
                           `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '变动时间',
                           `update_time` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                           PRIMARY KEY (`point_id`) USING BTREE,
                           INDEX `idx_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '积分表' ROW_FORMAT = Dynamic;
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `user_id` varchar(50) NOT NULL COMMENT '用户ID',
                         `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '登录用户名',
                         `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '加密存储的密码',
                         `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手机号（注册/登录验证）',
                         `budget_min` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '最低购车预算（元）',
                         `budget_max` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '最高购车预算（元）',
                         `preferred_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '偏好车型（SUV/轿车/MPV等）',
                         `preferred_brand` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '偏好品牌（多个用逗号分隔）',
                         `status` tinyint NULL DEFAULT 1 COMMENT '账号状态（-1-禁用/1-正常/2-管理员）',
                         `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '奖品收货地址',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                         PRIMARY KEY (`user_id`) USING BTREE,
                         UNIQUE INDEX `idx_phone`(`phone` ASC) USING BTREE,
                         INDEX `idx_preferred_type`(`preferred_type` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;