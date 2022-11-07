use movies_db;

/*
Se propone realizar las siguientes consultas a la base de datos movies_db.sql trabajada en la primera clase.
Importar el archivo movies_db.sql desde PHPMyAdmin o MySQL Workbench y resolver las siguientes consultas:*/

-- Mostrar el título y el nombre del género de todas las series.
select s.title, g.name from movies_db.series s
inner join genres g on g.id = s.genre_id
order by s.title;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select e.title, a.first_name, a.last_name from episodes e
inner join actor_episode ae on ae.episode_id = e.id
inner join actors a on a.id = ae.actor_id
order by e.title;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select s.title, count(se.id) from seasons se
inner join series s on s.id = se.serie_id
group by s.title;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select g.name, count(m.genre_id) as 'Cantidad Peliculas' from genres g
inner join movies m on m.genre_id = g.id
group by g.name
having count(m.genre_id) >= 3
order by g.name;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select distinct a.first_name, a.last_name from movies_db.movies m
inner join movies_db.actor_movie am on am.movie_id = m.id
inner join movies_db.actors a on a.id = am.actor_id
where title like '%La Guerra de las galaxias:%';

