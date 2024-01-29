CREATE TABLE IF NOT EXISTS corrpoints (
    station varchar(150) NOT NULL,
    time timestamptz,
    duration interval,
    raduis numeric(6,2),
    wind_speed numeric(6,2),
    wind_dir varchar(5),
    wind_run numeric(6,2),
    wind_chill numeric(6,2),
    bar numeric(6,2),
    rain numeric(6,2),
    rain_rate numeric(6,2),
    count              integer,
	maxPozitiveSignal  integer,
	maxNegativeSignal  integer,
	pozitiveSignal     integer,
	negativeSignal     integer,
	cloudTypeRelation  numeric(6,4),
	groundTypeRelation numeric(6,4)

);