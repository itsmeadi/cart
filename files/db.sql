create database NTUC;

create table NTUC.cart(
    id int(10) primary key auto_increment not null,
    user_id int(10) not null,
    status int(2) not null,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

create table NTUC.cart_items (
    id int(10) primary key auto_increment not null,
    cart_id int(10),
    product_id int(10) not null,
    qty int(4) not null,
    status int(2) not null,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    unique(product_id, cart_id),
    foreign key (cart_id) references NTUC.cart(id)
);

create table NTUC.user(
    id int(10) primary key auto_increment not null,
    sub varchar(50) unique not null
);