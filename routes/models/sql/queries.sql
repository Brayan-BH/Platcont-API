create database platcontApi 

-- create table login (
--     id char(36) PRIMARY KEY,
--     email varchar(50),
--     pass varchar(20),
-- )

-- create table modulo(
--     id_module char(36) PRIMARY KEY,
--     n_module varchar(20),
--     m_price decimal(10,2),
--     m_descri varchar(500),
-- )

-- create table contacto(
--     n_clie varchar(100),
--     l_email varchar(50),
--     n_tele varchar(15),
--     c_descri varchar(500)   
-- )

-- create table cliente(
--     id char(36) PRIMARY KEY,
--     n_clie varchar(100),
--     n_docu varchar(11),
--     l_dire varchar(40),
--     n_tele varchar(11),
--     r_social varchar(100),
--     l_email varchar(50)
-- )
-- create table clienteProducto(
--     id
-- )
-- create table productosDetalle(

-- )

-- create table facturas(

-- )

create table pagoPendiente(
    id_pago char(36) PRIMARY KEY,
    n_fact int,
    n_period datetime default getdate(),
    n_clie varchar(100),
    l_detalle varchar(80),
    s_impo decimal(10,2)
)

CREATE TABLE Users (
	id varchar(36) NOT NULL PRIMARY KEY,
	email varchar(100) DEFAULT '' UNIQUE,
	password varchar(200) DEFAULT ''
);


CREATE TABLE Clients (
	id varchar(36) NOT NULL PRIMARY KEY,
	n_docu varchar(11) NOT NULL UNIQUE,
	l_orga varchar(150) DEfAULT '',
	l_dire varchar(200) DEfAULT '',
	n_celu varchar(25) DEfAULT '',
	l_emai varchar(150) DEfAULT '',
	n_repr varchar(8) DEfAULT '',
	l_repr varchar(70) DEfAULT '',
	uid varchar(36) NOT NULL,
	FOREIGN KEY (uid) REFERENCES Users (id)
	
);


CREATE TABLE ClientProducts
(
	id varchar(36) NOT NULL,
	host varchar(100) DEFAULT '',
	users varchar(20) DEFAULT '',
	data_base varchar(20) DEFAULT '',
	password varchar(100) DEFAULT '',
	modulos varchar(50) DEFAULT '',
	multi int DEFAULT 0, -- 0 = single, 1 = multi empresa
	date_facture varchar(10) DEFAULT '',
	uid varchar(36) NOT NULL,
	FOREIGN KEY (uid) REFERENCES Users (id)
);

CREATE TABLE ProductosDetalle
(
    

);