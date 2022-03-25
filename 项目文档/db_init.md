## 表结构

### t_project -- 机构

```sql
CREATE TABLE `t_project` (
    `id` bigint(20) unsigned NOT NULL COMMENT 'ID',
    `created` datetime(6) NOT NULL COMMENT '创建时间',
    `updated` datetime(6) NOT NULL COMMENT '变更时间',
    `deleted` datetime(6) COMMENT '作废时间',
    `code` varchar(16) NOT NULL DEFAULT "" COMMENT '编号',
    `name` varchar(32) NOT NULL DEFAULT "" COMMENT '机构名',
    `life_cycle` tinyint(4) NOT NULL COMMENT '生命周期: 1.有效，2.作废',
    PRIMARY KEY (`id`),
    KEY `idx_created` (`created`) USING BTREE,
    KEY `idx_updated` (`updated`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='机构表';
```



### t_agreement -- 任务

```sql
CREATE TABLE `t_task` (
    `id` bigint(20) unsigned NOT NULL COMMENT 'ID',
    `created` datetime(6) NOT NULL COMMENT '创建时间',
    `updated` datetime(6) NOT NULL COMMENT '变更时间',
    `deleted` datetime(6) COMMENT '作废时间',
    `project_id` bigint(20) NOT NULL COMMENT '机构ID',
    `name` varchar(32) NOT NULL DEFAULT "" COMMENT '协议名称',
    `type` tinyint(4) NOT NULL COMMENT '任务类型',
    `life_cycle` tinyint(4) NOT NULL COMMENT '生命周期: 1.有效，2.作废',
    PRIMARY KEY (`id`),
    KEY `idx_created` (`created`) USING BTREE,
    KEY `idx_updated` (`updated`) USING BTREE,
    KEY `idx_project_id` (`project_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务表';
```

### t_thread_group -- 线程组

```sql
CREATE TABLE `t_thread_group` (
    `id` bigint(20) unsigned NOT NULL COMMENT 'ID',
    `created` datetime(6) NOT NULL COMMENT '创建时间',
    `updated` datetime(6) NOT NULL COMMENT '变更时间',
    `deleted` datetime(6) COMMENT '作废时间',
    `name` varchar(32) NOT NULL DEFAULT "" COMMENT '线程组名称',
    `thread_count` int(11) NOT NULL COMMENT '线程数量',
    `ew_node` varchar(8) NOT NULL DEFAULT "" COMMENT 'ew节点',
    `sleep_seconds` int(11) NOT NULL COMMENT '空闲时，休眠时长，单位为秒',
    `life_cycle` tinyint(4) NOT NULL COMMENT '生命周期: 1.有效，2.作废',
    PRIMARY KEY (`id`),
    KEY `idx_created` (`created`) USING BTREE,
    KEY `idx_updated` (`updated`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='线程组表';
```

### t_thread_group_task -- 线程组任务

```sql
CREATE TABLE `t_thread_group_task` (
    `id` bigint(20) unsigned NOT NULL COMMENT 'ID',
    `created` datetime(6) NOT NULL COMMENT '创建时间',
    `updated` datetime(6) NOT NULL COMMENT '变更时间',
    `deleted` datetime(6) COMMENT '作废时间',
    `thread_group_id` bigint(20) NOT NULL COMMENT '线程组ID',
    `status` tinyint(4) NOT NULL COMMENT '任务状态: 1.待执行，2.执行中',
    `task_id`  bigint(20) NOT NULL COMMENT '任务ID',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uix_agreement_id` (`task_id`) USING BTREE,
    KEY `idx_created` (`created`) USING BTREE,
    KEY `idx_updated` (`updated`) USING BTREE,
    KEY `idx_thread_group_id` (`thread_group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='线程组任务表';
```

### t_task_plugin -- 任务插件表

```sql
CREATE TABLE `t_task_plugin` (
    `id` bigint(20) unsigned NOT NULL COMMENT 'ID',
    `created` datetime(6) NOT NULL COMMENT '创建时间',
    `updated` datetime(6) NOT NULL COMMENT '变更时间',
    `deleted` datetime(6) COMMENT '作废时间',
    `task_id` bigint(20) NOT NULL COMMENT '任务ID',
    `service_plugin` int(11) NOT NULL COMMENT '任务类型',
    `next_query_time` datetime(6) NOT NULL COMMENT '下次查询时间',
    `last_query_time` datetime(6) COMMENT '最后一次查询时间',
    `interval` int(11) NOT NULL COMMENT '间隔时间，单位秒',
    `life_cycle` tinyint(4) NOT NULL COMMENT '生命周期: 1.有效，2.作废',
    `params` varchar(256) NOT NULL DEFAULT "" COMMENT '配置参数',
    PRIMARY KEY (`id`),
    KEY `idx_created` (`created`) USING BTREE,
    KEY `idx_updated` (`updated`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务插件表';
```

