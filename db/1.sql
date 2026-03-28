-- 插入比赛提交记录用于测试 ranking
-- 比赛ID: 1005, 题目: 1000, 1001, 1002

-- 用户 test 的提交
INSERT INTO submission (pid, username, cid, result, language, memory_cost, time_cost, code, create_at, update_at) VALUES
(1000, 'test', 1005, 'Accepted', 'cpp', 1024, 10, '#include<iostream>', NOW(), NOW()),
(1001, 'test', 1005, 'Wrong Answer', 'cpp', 2048, 50, '#include<iostream>', NOW(), NOW()),
(1001, 'test', 1005, 'Accepted', 'cpp', 2048, 45, '#include<iostream>', NOW(), NOW()),
(1002, 'test', 1005, 'Time Limit Exceeded', 'cpp', 4096, 2000, '#include<iostream>', NOW(), NOW());

-- 用户 test1 的提交
INSERT INTO submission (pid, username, cid, result, language, memory_cost, time_cost, code, create_at, update_at) VALUES
(1000, 'test1', 1005, 'Wrong Answer', 'cpp', 1024, 5, '#include<iostream>', NOW(), NOW()),
(1000, 'test1', 1005, 'Wrong Answer', 'cpp', 1024, 5, '#include<iostream>', NOW(), NOW()),
(1000, 'test1', 1005, 'Accepted', 'cpp', 1024, 8, '#include<iostream>', NOW(), NOW()),
(1001, 'test1', 1005, 'Accepted', 'cpp', 2048, 30, '#include<iostream>', NOW(), NOW());

-- 用户 zrzring 的提交（全对）
INSERT INTO submission (pid, username, cid, result, language, memory_cost, time_cost, code, create_at, update_at) VALUES
(1000, 'zrzring', 1005, 'Accepted', 'cpp', 512, 5, '#include<iostream>', NOW(), NOW()),
(1001, 'zrzring', 1005, 'Accepted', 'cpp', 1024, 20, '#include<iostream>', NOW(), NOW()),
(1002, 'zrzring', 1005, 'Accepted', 'cpp', 2048, 100, '#include<iostream>', NOW(), NOW());

-- 用户 admin 的提交（只做了一题）
INSERT INTO submission (pid, username, cid, result, language, memory_cost, time_cost, code, create_at, update_at) VALUES
(1000, 'admin', 1005, 'Runtime Error', 'cpp', 1024, 100, '#include<iostream>', NOW(), NOW()),
(1000, 'admin', 1005, 'Accepted', 'cpp', 1024, 12, '#include<iostream>', NOW(), NOW());
