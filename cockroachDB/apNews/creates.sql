-- create user if not exists bboyle;

create database if not exists apnews_db;

-- grant all on * to bboyle;

use apnews_db;

drop table if exists article;

CREATE TABLE Article (
    ArticleID int,
    Link string,
    Header string,
    Paragraph string,
    PRIMARY KEY (ArticleID)
);

insert into Article (articleid, link, header, paragraph) values (1, 'https://apnews.com/', 'First Header', 'First Paragraph');

insert into Article (articleid, link, header, paragraph) values (2, 'https://apnews.com/article/somehting_happened', 'Second Header', 'Second Paragraph');

insert into Article (articleid, link, header, paragraph) values (3, 'https://apnews.com/article/someone_died', 'Something happened and we dont know what', 'Guess what? Something happened and now I am testing the length of varchars in SQL because I havent done SQL in 2 years');
