drop table if exists posts;
drop table if exists threads;
drop table if exists sessions;
drop table if exists users;

CREATE TABLE users (
  id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  uuid VARCHAR(64) NOT NULL UNIQUE COMMENT 'uuid',
  name VARCHAR(255) NOT NULL COMMENT '用户名',
  email VARCHAR(255) NOT NULL COMMENT '邮件',
  password VARCHAR(255) NOT NULL COMMENT '密码',
  created_at TIMESTAMP NOT NULL COMMENT '创建时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB;

CREATE TABLE sessions (
  id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  uuid VARCHAR(64) NOT NULL UNIQUE COMMENT 'uuid',
  email VARCHAR(255) NOT NULL COMMENT '邮件',
  user_id int(11) NOT NULL COMMENT '用户id',
  created_at TIMESTAMP NOT NULL COMMENT '创建时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB;

CREATE TABLE threads (
  id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  uuid VARCHAR(64) NOT NULL UNIQUE COMMENT 'uuid',
  topic text NOT NULL COMMENT '内容文本',
  user_id int(11) NOT NULL COMMENT '用户id',
  created_at TIMESTAMP NOT NULL COMMENT '创建时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB;

CREATE TABLE posts (
  id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  uuid VARCHAR(64) NOT NULL UNIQUE COMMENT 'uuid',
  body text NOT NULL COMMENT '内容文本',
  user_id int(11) NOT NULL COMMENT '用户id',
  thread_id int(11) NOT NULL COMMENT '用户id',
  created_at TIMESTAMP NOT NULL COMMENT '创建时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB;