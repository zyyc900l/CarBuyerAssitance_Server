-- 插入礼品测试数据
INSERT INTO `gift` (`gift_name`, `required_points`, `stock_quantity`, `cover_image_url`, `is_online`) VALUES
-- 低积分礼品 (100-500积分)
('定制汽车钥匙扣', 1, 50, 'https://img.alicdn.com/imgextra/i4/3589751681/O1CN01vIaMRu1OHwfAuLy71_!!0-item_pic.jpg_q50.jpg', 1),
('车载手机支架', 150, 30, 'https://img.alicdn.com/imgextra/i3/3296263975/O1CN01R2RmF21fEbG7z6R75_!!3296263975.jpg_q50.jpg', 1),
('汽车品牌贴纸套装', 200, 100, 'https://img.alicdn.com/imgextra/i4/2210371521619/O1CN01ryGLxi1NpY0Pf9aqk_!!0-item_pic.jpg_q50.jpg', 1),
('车载充电器', 300, 25, 'https://img.alicdn.com/imgextra/i1/2612928539/O1CN01XU8nWR2Cwv3C1ieOC_!!2612928539.jpg_q50.jpg', 1),
('汽车香薰片', 250, 80, 'https://img.alicdn.com/imgextra/i1/2201465910053/O1CN01jeiyN51CGK11qTdnB_!!2201465910053.jpg_q50.jpg', 1),
('擦车毛巾', 120, 60, 'https://img.alicdn.com/imgextra/i2/1963082586/O1CN01eZ0DQ91UyRJiOTyy9_!!4611686018427380570-0-item_pic.jpg_q50.jpg', 1),

-- 中等积分礼品 (500-2000积分)
('品牌保温杯', 500, 40, 'https://gw.alicdn.com/bao/uploaded/i1/1714128138/O1CN01xMhwfy29zGEZ2YKLl_!!1714128138.jpg', 1),
('汽车遮阳挡', 600, 35, 'https://img.alicdn.com/imgextra/i3/2091512131/O1CN01chjKOz1Rc2walqtho_!!2091512131.jpg_q50.jpg', 1),
('车载收纳箱', 800, 20, 'https://img.alicdn.com/imgextra/i4/2096020116/O1CN01wp8Bvt1CjB1f8eKI3_!!4611686018427386516-0-item_pic.jpg_q50.jpg', 1),
('真皮驾驶证套', 450, 55, 'https://img.alicdn.com/imgextra/i1/2206952262958/O1CN01xaVgGV1Xioaru6cOR_!!4611686018427383086-0-item_pic.jpg_q50.jpg', 1),
('应急救生锤', 700, 15, 'https://gw.alicdn.com/bao/uploaded/i1/752922641/O1CN01EWMJZU1VNcrD6WVgn_!!752922641.jpg', 1),
('车载吸尘器', 1500, 10, 'https://gw.alicdn.com/imgextra/O1CN01sNUy391Dy7a8NqxSs_!!735670284-0-picasso.jpg_q50.jpg', 1),
('汽车座套', 1200, 8, 'https://gw.alicdn.com/imgextra/O1CN01220Kw92FPtJ9qO2lL_!!2204173188873.jpg_q50.jpg', 1),

-- 高积分礼品 (2000-5000积分)
('行车记录仪', 2500, 12, 'https://gw.alicdn.com/bao/uploaded/i4/2455154774/O1CN01sRatzY1l8Xtgh9y3U_!!2455154774.jpg', 1),
('车载空气净化器', 1800, 18, 'https://gw.alicdn.com/bao/uploaded/i3/2103754304/O1CN01aYiM9G1hfHhKofLq0_!!2103754304.jpg', 1),
('汽车模型(1:18)', 3000, 5, 'https://img.alicdn.com/imgextra/i4/354063145/O1CN012a4vbf1Z6Sf4G9wXh_!!354063145.jpg_q50.jpg', 1),
('品牌冲锋衣', 2200, 7, 'https://img.alicdn.com/imgextra/i4/3035493001/O1CN01kveR5W1Y2VgqI6Etj_!!3035493001.jpg_q50.jpg', 1),
('车载冰箱', 4000, 3, 'https://img.alicdn.com/imgextra/i3/2216811456848/O1CN01jQANbo20SRCQckNwt_!!2216811456848.jpg_q50.jpg', 1),
('全车贴膜券', 3500, 20, 'https://img.alicdn.com/imgextra/i4/596685497/O1CN01fvhkPy1qTg2xX29FY_!!596685497.jpg_q50.jpg', 1),

-- 限量/高价礼品 (5000+积分)
('品牌轮胎', 6000, 4, 'https://img.alicdn.com/imgextra/i2/2208354892100/O1CN010mqsNI1RNqLuZmWjj_!!2208354892100.jpg_q50.jpg', 1),
('汽车保养套餐', 8000, 15, 'https://img.alicdn.com/imgextra/i4/1718241991/O1CN01ocboV61QZvL13kdVE_!!1718241991.png_q50.jpg', 1),
('车载音响升级', 10000, 2, 'https://gw.alicdn.com/imgextra/O1CN01wHGN0p1h9DscSuXWN_!!3256574234.jpg_q50.jpg', 1),

-- 下架商品示例
('旧款导航仪', 1500, 0, 'https://gw.alicdn.com/imgextra/O1CN01LaxDuF1igwhRjCZyz_!!2212246924443.jpg_q50.jpg', 0),
(' discontinued 车模', 2000, 0, 'https://img.alicdn.com/imgextra/i3/2454171253/O1CN01KIJEiO1L7vMWZuUSH_!!4611686018427383413-0-item_pic.jpg_q50.jpg', 0);


INSERT INTO `user` (
    `user_id`,
    `username`,
    `password`,
    `phone`,
    `budget_min`,
    `budget_max`,
    `preferred_type`,
    `preferred_brand`,
    `status`,
    `address`
) VALUES
      ('admin', '林先生', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '1231242142', 10000.00, 100000.00, 'SUV', '比亚迪、吉利', 2, '福州大学旗山校区'),
      ('user001', '张三', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138001', 80000.00, 150000.00, '轿车', '丰田、本田', 1, '北京市朝阳区'),
      ('user002', '李四', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138002', 50000.00, 100000.00, 'SUV', '吉利、长安', 1, '上海市浦东新区'),
      ('user003', '王五', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138003', 120000.00, 250000.00, 'MPV', '别克、大众', 1, '广州市天河区'),
      ('user004', '赵六', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138004', 150000.00, 300000.00, '新能源', '特斯拉、比亚迪', 1, '深圳市南山区'),
      ('user005', '钱七', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138005', 60000.00, 120000.00, '轿车', '日产、现代', 1, '杭州市西湖区'),
      ('user006', '孙八', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138006', 200000.00, 500000.00, '豪华车', '奔驰、宝马', 1, '成都市武侯区'),
      ('user007', '周九', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138007', 30000.00, 80000.00, '小型车', '五菱、奇瑞', 1, '武汉市江汉区'),
      ('user008', '吴十', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138008', 100000.00, 200000.00, 'SUV', '长城、荣威', 1, '南京市鼓楼区'),
      ('manager', '陈经理', '$2a$10$2Hw1OZ88zuXg3K8vthq10eK7VjUbrvPeetMOk92vuhpyhxZ7JTzxK', '13800138009', 0.00, 0.00, '', '', 2, '福州市仓山区');