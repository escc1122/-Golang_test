PGDMP         '    
            x            admin    13.1 (Debian 13.1-1.pgdg100+1)    13.0     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16384    admin    DATABASE     Y   CREATE DATABASE admin WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';
    DROP DATABASE admin;
                admin    false                        2615    16385    first_collation    SCHEMA        CREATE SCHEMA first_collation;
    DROP SCHEMA first_collation;
                admin    false                        2615    16386    original    SCHEMA        CREATE SCHEMA original;
    DROP SCHEMA original;
                admin    false            �            1259    16470 +   first_collation_securities_investment_trust    TABLE     \  CREATE TABLE first_collation.first_collation_securities_investment_trust (
    id bigint,
    stock_no character varying NOT NULL,
    stock_name character varying,
    buy_trading_volume bigint DEFAULT 0 NOT NULL,
    sell_trading_volume bigint DEFAULT 0 NOT NULL,
    difference bigint DEFAULT 0 NOT NULL,
    stock_date character(8) NOT NULL
);
 H   DROP TABLE first_collation.first_collation_securities_investment_trust;
       first_collation         heap    admin    false    7            �            1259    16387    first_collation_stock_day    TABLE     (  CREATE TABLE first_collation.first_collation_stock_day (
    stock_date character varying NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    change double precision DEFAULT 0,
    percent_change double precision,
    ma_5 double precision DEFAULT 0,
    ma_10 double precision DEFAULT 0,
    ma_20 double precision DEFAULT 0,
    ma_60 double precision DEFAULT 0,
    stock_no character varying NOT NULL
);
 6   DROP TABLE first_collation.first_collation_stock_day;
       first_collation         heap    admin    false    7            �            1259    16422 $   original_securities_investment_trust    TABLE     �   CREATE TABLE original.original_securities_investment_trust (
    id bigint,
    date character(8) NOT NULL,
    json_data jsonb
);
 :   DROP TABLE original.original_securities_investment_trust;
       original         heap    admin    false    4            �            1259    16398    original_stock_day    TABLE     �   CREATE TABLE original.original_stock_day (
    id bigint,
    stock_no character varying NOT NULL,
    stock_json_data jsonb NOT NULL,
    stock_date character(8) NOT NULL,
    stock_date_year_month character(6)
);
 (   DROP TABLE original.original_stock_day;
       original         heap    admin    false    4                       2606    16485 \   first_collation_securities_investment_trust first_collation_securities_investment_trust_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY first_collation.first_collation_securities_investment_trust
    ADD CONSTRAINT first_collation_securities_investment_trust_pkey PRIMARY KEY (stock_no, stock_date);
 �   ALTER TABLE ONLY first_collation.first_collation_securities_investment_trust DROP CONSTRAINT first_collation_securities_investment_trust_pkey;
       first_collation            admin    false    205    205                       2606    16405 8   first_collation_stock_day first_collation_stock_day_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY first_collation.first_collation_stock_day
    ADD CONSTRAINT first_collation_stock_day_pkey PRIMARY KEY (stock_date, stock_no);
 k   ALTER TABLE ONLY first_collation.first_collation_stock_day DROP CONSTRAINT first_collation_stock_day_pkey;
       first_collation            admin    false    202    202                       2606    16483 N   original_securities_investment_trust original_securities_investment_trust_pkey 
   CONSTRAINT     �   ALTER TABLE ONLY original.original_securities_investment_trust
    ADD CONSTRAINT original_securities_investment_trust_pkey PRIMARY KEY (date);
 z   ALTER TABLE ONLY original.original_securities_investment_trust DROP CONSTRAINT original_securities_investment_trust_pkey;
       original            admin    false    204                       2606    16407 *   original_stock_day original_stock_day_pkey 
   CONSTRAINT     |   ALTER TABLE ONLY original.original_stock_day
    ADD CONSTRAINT original_stock_day_pkey PRIMARY KEY (stock_no, stock_date);
 V   ALTER TABLE ONLY original.original_stock_day DROP CONSTRAINT original_stock_day_pkey;
       original            admin    false    203    203           