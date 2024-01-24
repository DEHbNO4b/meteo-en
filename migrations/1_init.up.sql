CREATE TABLE IF NOT EXISTS stations (
    id serial,
    name varchar(150) NOT NULL UNIQUE,
    latitude numeric(6,4),
    longitude numeric(6,4)
);

CREATE TABLE IF NOT EXISTS meteodata (
    station varchar(150) NOT NULL,
    time timestamptz NOT NULL,
    temp_out numeric(6,2),
    wind_speed numeric(6,2),
    wind_dir varchar(5),
    wind_run numeric(6,2),
    wind_chill numeric(6,2),
    bar numeric(6,2),
    rain numeric(6,2),
    rain_rate numeric(6,2)

);
