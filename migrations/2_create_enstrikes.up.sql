CREATE TABLE IF NOT EXISTS enstrikes (
    id serial PRIMARY KEY,
    cloud boolean,
    time timestamptz,
    latitude numeric(6,4),
    longitude numeric(6,4),
    signal smallint,
    height smallint,
    sensors smallint,
    	count              int,
	maxPozitiveSignal  int,
	maxNegativeSignal  int,
	pozitiveSignal     int,
	negativeSignal     int,
	cloudTypeRelation   numeric(6,4),
	groundTypeRelation numeric(6,4)
);
