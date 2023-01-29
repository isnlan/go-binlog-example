create table mydb.user
(
  id int auto_increment primary key,
  name varchar(40) null,
  status enum("active","deleted") DEFAULT "active",
  created timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
)
  engine=InnoDB;

INSERT Into mydb.user (`id`,`name`) VALUE (1,"Jack");
UPDATE mydb.user SET name="Jonh" WHERE id=1;
DELETE FROM mydb.user WHERE id=1;