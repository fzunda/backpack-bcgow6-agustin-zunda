/*
create database empresa_internet_db;
use empresa_internet_db;

-- Created tables 

CREATE TABLE clients (
    `id_client` int(10) NOT NULL AUTO_INCREMENT,
    `first_name` varchar(50) null,
    `last_name` varchar(50) null,
    `date_birth` varchar(50) null,
    `province` varchar(50) null,
    `city` varchar(50) null,
    PRIMARY KEY (`id_client`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Inserts clients
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Ricardo','Flores', '01/05/1990', 'Cordoba', 'Cordoba');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Florencia','Dominguez', '21/05/1990', 'Cordoba', 'Cordoba');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('hernan','Garcia', '01/08/1987', 'Rosario', 'Rosario');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Ricardo','Garcia', '01/06/1995', 'Rosario', 'Rosario');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Esteban','Flores', '01/05/1980', 'Buenos Aires', 'Buenos Aires');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Ricardo','Garcia', '01/05/1985', 'Buenos Aires', 'Buenos Aires');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Jose','Vasquez', '01/07/1975', 'Santa fe', 'Santa fe');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Ruben','Gimenez', '01/06/1990', 'Santa fe', 'Santa fe');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('German','Paredes', '01/10/1980', 'Santa Cruz', 'Caleta');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Florencia','Paredes', '01/11/1971', 'Santa Cruz', 'Caleta');
INSERT INTO `empresa_internet_db`.`clients`
(`first_name`, `last_name`, `date_birth`,`province`, `city`) VALUES ('Ricardo','Vasquez', '01/05/1992', 'Cordoba', 'Cordoba');


CREATE TABLE pack (
    `id_pack` int(10) NOT NULL AUTO_INCREMENT,
    `navigation_speed` varchar(50) null,
    `price` double null,
    PRIMARY KEY (`id_pack`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Inserts pack
INSERT INTO `empresa_internet_db`.`pack`
(`navigation_speed`, `price`) VALUES ('100 Megas', 5200.0);
INSERT INTO `empresa_internet_db`.`pack`
(`navigation_speed`, `price`) VALUES ('300 Megas', 10000.0);
INSERT INTO `empresa_internet_db`.`pack`
(`navigation_speed`, `price`) VALUES ('50 Megas', 2700.0);
INSERT INTO `empresa_internet_db`.`pack`
(`navigation_speed`, `price`) VALUES ('3 Megas', 1500.0);
INSERT INTO `empresa_internet_db`.`pack`
(`navigation_speed`, `price`) VALUES ('1 Mega', 1000.0);



CREATE TABLE plans (
    `id_plan` int(10) NOT NULL AUTO_INCREMENT,
    `id_client` int(10) null,
    `id_pack` int(10) null,
    `discount` int(10) null,
    `final_price` double null,
    PRIMARY KEY (`id_plan`),
    KEY `id_client_foreign` (`id_client`),
	KEY `id_pack_foreign` (`id_pack`),
	CONSTRAINT `id_client_foreign` FOREIGN KEY (`id_client`) REFERENCES `clients` (`id_client`),
	CONSTRAINT `id_pack_foreign` FOREIGN KEY (`id_pack`) REFERENCES `pack` (`id_pack`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- inserts students
INSERT INTO `empresa_internet_db`.`plans`
(`id_client`,`id_pack`,`discount`,`final_price`) VALUES (149, 149, 200, 1500);
INSERT INTO `empresa_internet_db`.`plans`
(`id_client`,`id_pack`,`discount`,`final_price`) VALUES (150, 151, 300, 2000);
INSERT INTO `empresa_internet_db`.`plans`
(`id_client`,`id_pack`,`discount`,`final_price`) VALUES (159, 149, 200, 1500);
INSERT INTO `empresa_internet_db`.`plans`
(`id_client`,`id_pack`,`discount`,`final_price`) VALUES (153, 150, 500, 1000);
INSERT INTO `empresa_internet_db`.`plans`
(`id_client`,`id_pack`,`discount`,`final_price`) VALUES (154, 149, 200, 1500);
*/
-- Consultas
select * from empresa_internet_db.clients;

select * from empresa_internet_db.pack;

select * from empresa_internet_db.plans;

select pl.id_pack, pl.discount, pl.final_price, c.first_name, c.last_name from empresa_internet_db.plans pl
inner join empresa_internet_db.clients c on c.id_client = pl.id_client;

select pl.id_pack, pl.discount, pl.final_price, c.first_name, c.last_name from empresa_internet_db.plans pl
inner join empresa_internet_db.clients c on c.id_client = pl.id_client
where c.id_client = 149;

select pl.id_pack, pl.discount, pl.final_price, c.first_name, c.last_name, pk.* 
from empresa_internet_db.plans pl
inner join empresa_internet_db.clients c on c.id_client = pl.id_client
inner join empresa_internet_db.pack pk on pk.id_pack = pl.id_pack
order by c.first_name;

select pl.id_pack, pl.discount, pl.final_price, c.first_name, c.last_name, pk.navigation_speed, pk.price 
from empresa_internet_db.plans pl
inner join empresa_internet_db.clients c on c.id_client = pl.id_client
inner join empresa_internet_db.pack pk on pk.id_pack = pl.id_pack
where pk.id_pack in (149)
order by c.first_name;

select count(*) as 'clientes en BsAS' from empresa_internet_db.clients
where city like '%Buenos Aires%';

select min(final_price) as 'Plan mas barato' from empresa_internet_db.plans;

select avg(discount) as 'Descuento promedio en planes' from empresa_internet_db.plans;