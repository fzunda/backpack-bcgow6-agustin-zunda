
-- I create the database and the tables
CREATE DATABASE empresa_db;
USE empresa_db;

CREATE TABLE Employees (
    `cod_emp` varchar(50) NOT NULL,
    `fist_name` varchar(50) null,
    `last_name` varchar(50) null,
    `position` varchar(50) null,
    `release_date` varchar(50) null,
    `salary` double null,
    `commission` double null,
    `dpto_number` varchar(50),
    PRIMARY KEY (`cod_emp`),
    KEY dpto_number_foreign (`dpto_number`),
    CONSTRAINT `dpto_number_foreign` FOREIGN KEY (`dpto_number`) REFERENCES department (`depto_nro`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE department (
    `depto_nro` varchar(50) NOT NULL,
    `depto_name` varchar(50) null,
    `depto_town` varchar(50) null,
    PRIMARY KEY (`depto_nro`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- I create the records in the tables
use empresa_db;

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0001','César', 'Piñero', 'Vendedor', '12/05/2018', 80000.0, 15000.0, 'D-000-4');

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0002','Yosep', 'Kowaleski', 'Analista', '14/07/2015', 140000.0, 0.0, 'D-000-2');

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0003','Mariela', 'Barrios', 'Director', '05/06/2014', 185000.0, 0.0, 'D-000-3');

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0004','Jonathan', 'Aguilera', 'Vendedor', '03/06/2015', 85000.0, 10000.0, 'D-000-4');

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0005','Daniel', 'Brezezicki', 'Vendedor', '03/03/2018', 83000.0, 10000.0, 'D-000-4');

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0006','Mito', 'Barchuk', 'Presidente', '05/06/2014', 190000.0, 0.0, 'D-000-3');

INSERT INTO `empresa_db`.`Employees` (`cod_emp`, `fist_name`, `last_name`, `position`, `release_date`, `salary`, `commission`, `dpto_number`)
VALUES ('E-0007','Emilio', 'Galarza', 'Desarrollador', '02/08/2014', 60000.0, 0.0, 'D-000-1');


INSERT INTO `empresa_db`.`department` (`depto_nro`, `depto_name`, `depto_town`) VALUES ('D-000-1','Software', 'Los Tigres');
INSERT INTO `empresa_db`.`department` (`depto_nro`, `depto_name`, `depto_town`) VALUES ('D-000-2','Sistemas', 'Guadalupe');
INSERT INTO `empresa_db`.`department` (`depto_nro`, `depto_name`, `depto_town`) VALUES ('D-000-3','Contabilidad', 'La Roca');
INSERT INTO `empresa_db`.`department` (`depto_nro`, `depto_name`, `depto_town`) VALUES ('D-000-4','Ventas', 'Plata');

-- Se requiere obtener las siguientes consultas:

-- Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select e.fist_name, e.last_name, e.position, d.depto_town from  empresa_db.Employees as e 
inner join empresa_db.department as d on e.dpto_number = d.depto_nro;

-- Visualizar los departamentos con más de cinco empleados.
select dpto_number, count(cod_emp) from empresa_db.Employees 
group by dpto_number
having count(cod_emp) > 5;

-- Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select e.fist_name, e.salary, d.depto_name from empresa_db.Employees as e
inner join empresa_db.department as d on e.dpto_number = d.depto_nro
where e.position = (select e2.position from empresa_db.Employees as e2
					where e2.fist_name like 'Mito' and e2.last_name like 'Barchuk');
                    
-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
select * from empresa_db.employees
where dpto_number = (select depto_nro from empresa_db.department
					 where depto_name like 'Contabilidad')
order by fist_name;

-- Mostrar el nombre del empleado que tiene el salario más bajo.
select concat(fist_name, " ", last_name) as name_employee from empresa_db.employees
where salary = (select Min(e2.salary) from empresa_db.employees as e2);

-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
select e.* from empresa_db.employees as e
where e.salary = (select max(e2.salary) from empresa_db.employees as e2
				  inner join empresa_db.department as d on e2.dpto_number = d.depto_nro
				  where d.depto_name like 'Ventas');






















