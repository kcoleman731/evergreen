CREATE TABLE companies (
    company_id serial NOT NULL,
    name character varying(100) NOT NULL,
    funding character varying(100),
    website character varying(500) NOT NULL,
    created date
);

CREATE TABLE users (
    user_id serial NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(500) NOT NULL,
    email character varying(100),
    department character varying(100),
    position character varying(100),
    created date
);

CREATE TABLE products (
    product_id serial NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    platforms character varying(100),
    teams character varying(100),
    created date
);

CREATE TABLE features (
    feature_id serial NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created date
);

CREATE TABLE platforms (
    platform_id serial NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created date
);

CREATE TABLE teams (
    team_id serial NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created date
);

CREATE TABLE questions (
    qustion_id serial NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created date
);

CREATE TABLE answers (
    answer_id serial NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(500) NOT NULL,
    created date
);
