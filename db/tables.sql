create user bmlist identified by 'bmlist_u01';
-- ER_NOT_SUPPORTED_AUTH_MODE : Client does not support authentication protocol requested by server;
alter user 'bmlist' @'%' identified with mysql_native_password by 'bmlist_u01';
create database bmlist;
grant all privileges on bmlist.* to 'bmlist' @'%';
--
--
--------------------------------------------------------------------------
-- main table for xnote
CREATE TABLE `xnote_xnote` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(60) DEFAULT NULL,
  `content` text,
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL default current_timestamp on update current_timestamp,
  `uid` int(11) NOT NULL default 1,
  `status` smallint(6) NOT NULL default 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
--
---------------------------------------------------------------------------
-- ----- sample data for test usage
insert into
  xnote_xnote(title, content, create_time)
values('This is a test', 'this is content', now())