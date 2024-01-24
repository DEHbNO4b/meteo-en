CREATE TABLE IF NOT EXISTS stations (
    id serial,
    name varchar(150) NOT NULL UNIQUE,
    latitude numeric(6,4),
    longitude numeric(6,4)
);
-- CREATE INDEX if NOT EXISTS idx_login ON users(login);

CREATE TABLE IF NOT EXISTS meteodata (
    station varchar(150) NOT NULL,
    time timestamptz NOT NULL,
    temp_out numeric(6,2),
    wind_speed numeric(6,2)
    wind_dir varchar(5),
    wind_run numeric(6,2),
    wind_chill numeric(6,2),
    bar numeric(6,2),
    rain numeric(6,2),
    rain_rate numeric(6,2)

);
-- CREATE TABLE IF NOT EXISTS text_data (
--     user_id integer NOT NULL,
--     text text NOT NULL,
--     meta text
-- );
-- CREATE TABLE IF NOT EXISTS binary_data (
--     user_id integer NOT NULL,
--     data bytea NOT NULL,
--     meta text
-- );
-- CREATE TABLE IF NOT EXISTS card_data (
--     user_id integer NOT NULL,
--     card_id varchar(16) NOT NULL,
--     pass varchar(3) NOT NULL,
--     date varchar(10) NOT NULL,
--     meta text
-- );