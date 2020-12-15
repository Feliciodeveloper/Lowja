Create database lowja;
use lowja;

create table produtos (
id int auto_increment,
nome varchar(100) not null,
descricao varchar(200),
preco float default 0.00,
quantidade int default 0,
primary key (id)
);

create table usuarios (
id int auto_increment,
nome varchar(100) not null,
email varchar(100) not null unique key,
senha varchar(100) not null,
primary key (id)
);