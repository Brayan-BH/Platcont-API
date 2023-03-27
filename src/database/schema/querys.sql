CREATE TABLE Users (
    id_user varchar(36) NOT NULL PRIMARY KEY,
    email varchar(100) NOT NULL,
    password varchar(200) DEFAULT '',
    password_admin varchar(200) DEFAULT '',
    constraint email unique (email)
);

CREATE TABLE Clients (
    id_clie varchar(36) NOT NULL PRIMARY KEY,
    n_docu varchar(11) NOT NULL,
    l_orga varchar(150) DEFAULT '',
    l_dire varchar(200) DEFAULT '',
    n_celu varchar(25) DEFAULT '',
    l_emai varchar(150) DEFAULT '',
    n_repr varchar(8) DEFAULT '',
    l_repr varchar(70) DEFAULT '',
    FOREIGN KEY (id_clie) REFERENCES Users (id_user),
    constraint n_docu unique (n_docu)
);

CREATE TABLE ClientProducts (
    id_clipd varchar(36) NOT NULL PRIMARY KEY,
    host varchar(100) DEFAULT '',
    users varchar(20) DEFAULT '',
    data_base varchar(20) DEFAULT '',
    password varchar(100) DEFAULT '',
    modulos varchar(50) DEFAULT '',
    -- 1 := general
    -- 2 := caja
    -- 4 := contabilidad
    -- 5 := Financiera
    -- 6 := Recursos Humanos
    -- 7 := Stock
    -- 8 := Punto de Venta
    -- 9 := Activo Fijo
    -- 10 := Cuenta Corriente
    -- 11 := Costos
    -- 12 := Taxi
    -- 13 := Transporte
    multi int DEFAULT 0,
    -- 0 = single, 1 = multi empresa
    date_facture varchar(10) DEFAULT '',
    id_clie varchar(36) NOT NULL,
    FOREIGN KEY (id_clie) REFERENCES Clients (id_clie)
);

CREATE TABLE Facturas (
    id_fact varchar(36) NOT NULL PRIMARY KEY,
    n_docu varchar(11) NOT NULL,
    years int not null,
    months int not null,
    c_comp varchar (2) NOT NULL,
    n_seri varchar(4) NOT NULL,
    n_com varchar(10) NOT NULL,
    f_venc varchar (10) NOT NULL,
    f_comp varchar (10) NOT NULL,
    f_pago varchar (10) NOT NULL,
    s_impo float8 NOT NULL,
    s_igv float8 NOT NULL,
    s_desc float8 NOT NULL,
    s_tota float8 NOT NULL,
    l_obse varchar(100) NOT NULL,
    k_stad int DEFAULT 0,
    id_clipd varchar(36) not null,
    foreign key (id_clipd) references ClientProducts (id_clipd)
);

CREATE TABLE ProductosDetalle (
    id_pddt varchar(36) not null primary key,
    id_clipd varchar(36) NOT NULl,
    l_deta varchar(100) NOT NULL,
    s_impo float8,
    months int not null,
    years int not null,
    id_fact varchar(36) NOT NULL,
    FOREIGN KEY (id_clipd) REFERENCES ClientProducts (id_clipd)
    -- FOREIGN KEY (id_fact) REFERENCES Facturas (id_fact)

);

CREATE TABLE Facturas_Detalle (
    id_fact varchar(36) NOT null primary Key,
    n_item int not null,
    c_prod varchar(8) NOT NULL,
    s_impo float8 NOT NULL,
    s_igv float8 NOT NULL,
    s_desc float8 NOT NULL,
    s_tota float8 NOT NULL,
    l_peri varchar(7) not null,
    id_pddt varchar(36) not null,
    foreign key (id_fact) references Facturas (id_fact),
    foreign key (id_pddt) references ProductosDetalle (id_pddt),
    constraint id_fact_n_item unique (id_fact, n_item)
);