create database platcontApi 

create table pagoPendiente(
	id_pago char(36) PRIMARY KEY,
	n_fact int,
	n_period time default CURRENT_TIMESTAMP() ,
	n_clie varchar(100),
	l_detalle varchar(80),
	s_impo decimal(10, 2)
) CREATE TABLE Users (
	id varchar(36) NOT NULL PRIMARY KEY,
	email varchar(100) DEFAULT '' UNIQUE,
	password varchar(200) DEFAULT ''
);

CREATE TABLE Clients (
	id varchar(36) NOT NULL PRIMARY KEY,
	l_clie varchar(100) NOT NULL,
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

CREATE TABLE ClientProducts (
	id varchar(36) NOT NULL,
	host varchar(100) DEFAULT '',
	users varchar(20) DEFAULT '',
	data_base varchar(20) DEFAULT '',
	password varchar(100) DEFAULT '',
	modulos varchar(50) DEFAULT '',
	multi int DEFAULT 0,-- 0 = single, 1 = multi empresa
	date_facture varchar(10) DEFAULT '',
	uid varchar(36) NOT NULL,
	FOREIGN KEY (uid) REFERENCES Users (id)
);

CREATE TABLE ProductosDetalle (
	id_detail varchar(36) NOT NULL,
	l_detalle varchar(100),
	s_impo decimal(10, 2)
);

CREATE TABLE Facturas (
	id_fact varchar(36) NOT NULL,
	n_period time default CURRENT_TIMESTAMP,
	n_fact int DEFAULT 0,
	n_clie varchar(100) NOT NULL,
	l_detalle varchar(80) NOT NULL,
	s_impo decimal(10, 2)
);