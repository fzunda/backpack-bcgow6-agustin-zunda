-- Implementar la base de datos en PHPMyAdmin o MySQL Workbench, cargar cinco registros en cada tabla y probar algunas 
-- consultas planteadas en el Ejercicio 1. 

/*create database biblioteca_db;
use biblioteca_db;

-- Created tables 

CREATE TABLE books (
    `id_book` int(10) NOT NULL AUTO_INCREMENT,
    `title` varchar(50) null,
    `leader` varchar(50) null,
    `zone` varchar(50) null,
    PRIMARY KEY (`id_book`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Inserts books
INSERT INTO `biblioteca_db`.`books`
(`title`,`leader`,`zone`) VALUES ('Libro 1','Autor 1','Aventura');
INSERT INTO `biblioteca_db`.`books`
(`title`,`leader`,`zone`) VALUES ('Libro 2','Autor 2','terror');
INSERT INTO `biblioteca_db`.`books`
(`title`,`leader`,`zone`) VALUES ('Libro 3','Autor 3','Romance');
INSERT INTO `biblioteca_db`.`books`
(`title`,`leader`,`zone`) VALUES ('Libro 4','Autor 4','Misterio');
INSERT INTO `biblioteca_db`.`books`
(`title`,`leader`,`zone`) VALUES ('Libro 5','Autor 5','Paranormal');


CREATE TABLE author (
    `id_author` int(10) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) null,
    `nationality` varchar(50) null,
    PRIMARY KEY (`id_author`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Inserts author
INSERT INTO `biblioteca_db`.`author`
(`name`,`nationality`) VALUES ('Autor 1', 'Argentina');
INSERT INTO `biblioteca_db`.`author`
(`name`,`nationality`) VALUES ('Autor 2', 'Argentina');
INSERT INTO `biblioteca_db`.`author`
(`name`,`nationality`) VALUES ('Autor 3', 'Chile');
INSERT INTO `biblioteca_db`.`author`
(`name`,`nationality`) VALUES ('Autor 4', 'Chile');
INSERT INTO `biblioteca_db`.`author`
(`name`,`nationality`) VALUES ('Autor 5', 'Argentina');
INSERT INTO `biblioteca_db`.`author`
(`name`,`nationality`) VALUES ('Autor 6', 'Italiana');

CREATE TABLE student (
    `id_student` int(10) NOT NULL AUTO_INCREMENT,
    `first_name` varchar(50) null,
    `last_name` varchar(50) null,
    `address` varchar(50) null,
    `career` varchar(50) null,
    `age` int(50) null,
    PRIMARY KEY (`id_student`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- inserts students
INSERT INTO `biblioteca_db`.`student`
(`first_name`, `last_name`, `address`, `career`, `age`) VALUES ('Estudiante 1', 'Gomez 1', 'Domicilio 1', 'Sistemas', 22);
INSERT INTO `biblioteca_db`.`student`
(`first_name`, `last_name`, `address`, `career`, `age`) VALUES ('Estudiante 2', 'Gomez 2', 'Domicilio 2', 'Sistemas', 25);
INSERT INTO `biblioteca_db`.`student`
(`first_name`, `last_name`, `address`, `career`, `age`) VALUES ('Estudiante 3', 'Gomez 3', 'Domicilio 3', 'Quimica', 30);
INSERT INTO `biblioteca_db`.`student`
(`first_name`, `last_name`, `address`, `career`, `age`) VALUES ('Estudiante 4', 'Gomez 4', 'Domicilio 4', 'Biologia', 21);
INSERT INTO `biblioteca_db`.`student`
(`first_name`, `last_name`, `address`, `career`, `age`) VALUES ('Estudiante 5', 'Gomez 5', 'Domicilio 5', 'Civil', 35);


CREATE TABLE book_author (
    `id_book_author` int(10) NOT NULL AUTO_INCREMENT,
    `id_book_assigned` int(10) null,
    `id_author_assigned` int(10) null,
    PRIMARY KEY (`id_book_author`),
    KEY `id_book_assigned_foreign` (`id_book_assigned`),
	KEY `id_author_assigned_foreign` (`id_author_assigned`),
	CONSTRAINT `id_book_assigned_foreign` FOREIGN KEY (`id_book_assigned`) REFERENCES `books` (`id_book`),
	CONSTRAINT `id_author_assigned_foreign` FOREIGN KEY (`id_author_assigned`) REFERENCES `author` (`id_author`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- inserts book_author
INSERT INTO `biblioteca_db`.`book_author`
(`id_book_assigned`, `id_author_assigned`) VALUES (149, 153);
INSERT INTO `biblioteca_db`.`book_author`
(`id_book_assigned`, `id_author_assigned`) VALUES (150, 151);
INSERT INTO `biblioteca_db`.`book_author`
(`id_book_assigned`, `id_author_assigned`) VALUES (151, 152);
INSERT INTO `biblioteca_db`.`book_author`
(`id_book_assigned`, `id_author_assigned`) VALUES (152, 149);
INSERT INTO `biblioteca_db`.`book_author`
(`id_book_assigned`, `id_author_assigned`) VALUES (153, 150);

CREATE TABLE loan (
    `id_loan` int(10) NOT NULL AUTO_INCREMENT,
    `loan_date` varchar(50) null,
    `return_date` varchar(50) null,
    `id_student_assigned` int(10) null,
    `id_book_assigned` int(10) null,
    PRIMARY KEY (`id_loan`),
    KEY `id_student_assigned_foreign` (`id_student_assigned`),
	KEY `id_book_assigned_foreign` (`id_book_assigned`),
	CONSTRAINT `id_student_assigned_foreign` FOREIGN KEY (`id_student_assigned`) REFERENCES `student` (`id_student`),
	CONSTRAINT `id_book_assigned` FOREIGN KEY (`id_book_assigned`) REFERENCES `books` (`id_book`)
)ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `biblioteca_db`.`loan`
(`loan_date`, `return_date`, `id_student_assigned`, `id_book_assigned`) VALUES ('12/01/2022', '14/01/2022', 149, 149);
INSERT INTO `biblioteca_db`.`loan`
(`loan_date`, `return_date`, `id_student_assigned`, `id_book_assigned`) VALUES ('18/01/2022', '20/01/2022', 150, 151);
INSERT INTO `biblioteca_db`.`loan`
(`loan_date`, `return_date`, `id_student_assigned`, `id_book_assigned`) VALUES ('20/01/2022', '23/01/2022', 151, 152);
INSERT INTO `biblioteca_db`.`loan`
(`loan_date`, `return_date`, `id_student_assigned`, `id_book_assigned`) VALUES ('12/01/2022', '14/01/2022', 152, 149);
INSERT INTO `biblioteca_db`.`loan`
(`loan_date`, `return_date`, `id_student_assigned`, `id_book_assigned`) VALUES ('12/01/2022', '14/01/2022', 153, 150);

-- Listar los datos de los autores.
*/
-- Listar nombre y edad de los estudiantes
select concat(first_name, ' ', last_name), age from biblioteca_db.student;

-- ¿Qué estudiantes pertenecen a la carrera informática?
select id_student, concat(first_name, ' ', last_name) as 'name' 
from biblioteca_db.student
where career like 'Sistemas';

-- ¿Qué autores son de nacionalidad francesa o italiana?
select name from biblioteca_db.author
where nationality like 'Francesa' or nationality like 'Italiana';

-- ¿Qué libros no son del área de internet?
select title from biblioteca_db.books
where zone not like 'Internet';

-- Listar los libros de la editorial Salamandra.
select * from biblioteca_db.books
where leader like 'Salamandra';

-- Listar los datos de los estudiantes cuya edad es mayor al promedio.
select * from biblioteca_db.student
where age > (select avg(age) from biblioteca_db.student);

-- Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
select first_name, last_name from biblioteca_db.student
where last_name like 'G%';

-- Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
select a.name from biblioteca_db.books b
inner join biblioteca_db.book_author ba on b.id_book = ba.id_book_assigned
inner join biblioteca_db.author a on a.id_author = ba.id_author_assigned
where b.title like 'El Universo: Guía de viaje'; 

-- ¿Qué libros se prestaron al lector “Filippo Galli”?
select b.* from biblioteca_db.books b
inner join biblioteca_db.loan l on l.id_book_assigned = b.id_book
inner join biblioteca_db.student s on s.id_student = l.id_student_assigned
where s.first_name like 'Filippo' and s.last_name like 'Galli'; 

-- Listar el nombre del estudiante de menor edad.
select first_name, last_name from biblioteca_db.student
where age = (select min(age) from biblioteca_db.student);

-- Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
select * from biblioteca_db.student
where id_student in (select id_loan from biblioteca_db.loan);

-- Listar los libros que pertenecen a la autora J.K. Rowling.
select b.* from biblioteca_db.books b
inner join biblioteca_db.book_author ba on ba.id_author_assigned = b.id_book
inner join biblioteca_db.author a on a.id_author = ba.id_author_assigned
where a.name like 'J.K. Rowling';

-- Listar títulos de los libros que debían devolverse el 16/07/2021.
select b.* from biblioteca_db.books b
inner join biblioteca_db.loan l on l.id_book_assigned = b.id_book
where l.return_date like '16/07/2021';
