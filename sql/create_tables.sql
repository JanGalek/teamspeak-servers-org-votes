create table voters
(
  id       int auto_increment
    primary key,
  username varchar(180)    not null,
  votes    int default '0' not null,
  constraint voters_username_uindex
  unique (username)
);

create table votes
(
  id       int auto_increment
    primary key,
  username varchar(180)    not null,
  votes    int default '0' not null,
  month    varchar(8)      null,
  constraint votes_month_username_uindex
  unique (month, username)
);