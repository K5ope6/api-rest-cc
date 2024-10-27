CREATE DATABASE IF NOT EXISTS mydatabasecc;

USE mydatabasecc;

CREATE TABLE IF NOT EXISTS Alumno (
    idAlumno INT PRIMARY KEY NOT NULL,
    nombre VARCHAR(30) NOT NULL,
    apellidoP VARCHAR(30) NOT NULL,
    apellidoM VARCHAR(30) NOT NULL,
    matricula VARCHAR(30) NOT NULL,
    correo VARCHAR(30) NOT NULL,
    carrera VARCHAR(50) NOT NULL,
    semestre VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS Creditos (
    idCredito INT PRIMARY KEY NOT NULL,
    nombreCredito VARCHAR(30) NOT NULL,
    estado VARCHAR(20) DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS AlumnoCreditos (
    idAlumno INT,
    idCredito INT,
    PRIMARY KEY (idAlumno, idCredito),
    FOREIGN KEY (idAlumno) REFERENCES Alumno(idAlumno),
    FOREIGN KEY (idCredito) REFERENCES Creditos(idCredito)
);

INSERT INTO Alumno (idAlumno, nombre, apellidoP, apellidoM, matricula, correo, carrera, semestre)
VALUES
(1, 'Juan', 'Pérez', 'Gómez', '19240609', 'juan.perez@example.com', 'Ingeniería', 11),
(2, 'Ana', 'Martínez', 'López', '19241234', 'ana.martinez@example.com', 'Derecho', 11),
(3, 'Carlos', 'Ramírez', 'Díaz', '19249443', 'carlos.ramirez@example.com', 'Arquitectura', 11),
(4, 'María', 'García', 'Hernández', '19240589', 'maria.garcia@example.com', 'Medicina', 11),
(5, 'Luis', 'Sánchez', 'Torres', '20249532', 'luis.sanchez@example.com', 'Administración', 10),
(6, 'Sofía', 'Martín', 'Castillo', '20240675', 'sofia.martin@example.com', 'Biología', 9),
(7, 'David', 'Cruz', 'Mendoza', '19240684', 'david.cruz@example.com', 'Química', 9),
(8, 'Isabel', 'Vásquez', 'González', '18248683', 'isabel.vasquez@example.com', 'Física', 11),
(9, 'Fernando', 'Morales', 'Nava', '19240784', 'fernando.morales@example.com', 'Psicología', 10),
(10, 'Elena', 'Ponce', 'Rojas', '20249583', 'elena.ponce@example.com', 'Ingeniería', 10),
(11, 'Arturo', 'Blanco', 'Duarte', '18248211', 'arturo.blanco@example.com', 'Sociología', 12),
(12, 'Valeria', 'Salazar', 'Alvarez', '18240698', 'valeria.salazar@example.com', 'Ciencias Políticas', 12),
(13, 'Diego', 'Córdova', 'Jiménez', '19240967', 'diego.cordova@example.com', 'Historia', 9),
(14, 'Carmen', 'Salinas', 'Rentería', '21249599', 'carmen.salinas@example.com', 'Relaciones Internacionales', 6),
(15, 'Hugo', 'Núñez', 'Ríos', '21249679', 'hugo.nunez@example.com', 'Mercadotecnia', 6);

INSERT INTO Creditos(idCredito, nombreCredito, estado)
values
('1','Actividad extraescolar','APROBADO'),
('2','Actividad extraescolar','NO APROBADO'),
('3','Conferencia','APROBADO'),
('4','Conferencia','NO APROBADO'),
('5','Taller','APROBADO'),
('6','Taller','NO APROBADO'),
('7','Tutoria','APROBADO'),
('8','Tutoria','NO APROBADO'),
('9','Ayuda al tec','APROBADO'),
('10','Ayuda al tec','NO APROBADO');

INSERT INTO AlumnoCreditos(idAlumno, idCredito)
VALUES
('1','1'), ('1','4'),('1','5'), ('1','7'), ('1','9'),
('2','1'), ('2','3'), ('2','5'), ('2','7'), ('2','9'),
('3', '1'), ('3', '3'), ('3', '5'),('3', '7'),('3', '9'),
('4', '2'), ('4', '4'), ('4', '6'), ('4', '8'), ('4', '10'), 
('5', '1'), ('5', '3'), ('5', '5'), ('5', '8'), ('5', '10'), 
('6', '2'), ('6', '4'), ('6', '5'), ('6', '7'), ('6', '9'), 
('7', '1'), ('7', '4'), ('7', '6'), ('7', '7'), ('7', '9'), 
('8', '2'), ('8', '3'), ('8', '5'), ('8', '7'), ('8', '10'), 
('9', '1'), ('9', '3'), ('9', '6'),('9', '7'),('9', '9'),
('10', '2'),('10', '3'),('10', '6'),('10', '8'),('10', '10'), 
('11', '1'),('11', '3'),('11', '6'),('11', '7'),('11', '10'), 
('12', '2'), ('12', '3'), ('12', '5'), ('12', '8'), ('12','10'),
('13', '1'), ('13', '3'), ('13', '6'), ('13', '7'), ('13', '9'), 
('14', '1'), ('14', '3'), ('14', '5'), ('14', '7'), ('14', '10'), 
('15', '1'), ('15', '3'), ('15', '6'), ('15', '8'), ('15', '9'); 