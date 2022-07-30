# gosql

## *Postgres SQL builder*

### Create table (support full [PG14 SQL specification](https://www.postgresql.org/docs/current/sql-createtable.html)) examples

###### Table with named primary key constraint
```
CREATE TABLE films (
    code        char(5) CONSTRAINT firstkey PRIMARY KEY,
    title       varchar(40) NOT NULL,
    did         integer NOT NULL,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute
);

films := gosql.CreateTable("films")
films.AddColumn("code").Type("char(5)").Constraint().Name("firstkey").PrimaryKey()
films.AddColumn("title").Type("varchar(40)").Constraint().NotNull()
films.AddColumn("did").Type("integer").Constraint().NotNull()
films.AddColumn("date_prod").Type("date")
films.AddColumn("kind").Type("varchar(10)")
films.AddColumn("len").Type("interval hour to minute")
```

######  Table with unique named constraint
```
CREATE TABLE films (
    code        char(5),
    title       varchar(40),
    did         integer,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute,
    CONSTRAINT production UNIQUE(date_prod)
);

films := gosql.CreateTable("films")
films.AddColumn("code").Type("char(5)")
films.AddColumn("title").Type("varchar(40)")
films.AddColumn("did").Type("integer")
films.AddColumn("date_prod").Type("date")
films.AddColumn("kind").Type("varchar(10)")
films.AddColumn("len").Type("interval hour to minute")
films.AddConstraint().Name("production").Unique().Columns().Add("date_prod")
```

###### Table with primary key constraint
```
CREATE TABLE distributors (
    did     integer,
    name    varchar(40),
    PRIMARY KEY(did)
);

distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)")
distributors.AddConstraint().PrimaryKey().Columns().Add("did")
```

###### Table with primary key in column definition
```
CREATE TABLE distributors (
    did     integer PRIMARY KEY,
    name    varchar(40)
);

distributors = gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer").Constraint().PrimaryKey()
distributors.AddColumn("name").Type("varchar(40)")
```

###### Table with named constraint not null
```
CREATE TABLE distributors (
    did     integer CONSTRAINT no_null NOT NULL,
    name    varchar(40) NOT NULL
);

distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer").Constraint().Name("no_null").NotNull()
distributors.AddColumn("name").Type("varchar(40)").Constraint().NotNull()
```

###### Table with unique column
```
CREATE TABLE distributors (
    did     integer,
    name    varchar(40) UNIQUE
);

distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)").Constraint().Unique()
```

###### Table with unique constraint with storage parameter
```
CREATE TABLE distributors (
    did     integer,
    name    varchar(40),
    UNIQUE(name) WITH (fillfactor=70)
)
WITH (fillfactor=70);

distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)")
unique := distributors.AddConstraint().Unique()
unique.Columns().Add("name")
unique.IndexParameters().With().Add("fillfactor=70")
distributors.With().Expression().Add("fillfactor=70")
```

###### Table with name constraint primary key on multiple column
```
CREATE TABLE films (
    code        char(5),
    title       varchar(40),
    did         integer,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute,
    CONSTRAINT code_title PRIMARY KEY(code,title)
);

films := gosql.CreateTable("films")
films.AddColumn("code").Type("char(5)")
films.AddColumn("title").Type("varchar(40)")
films.AddColumn("did").Type("integer")
films.AddColumn("date_prod").Type("date")
films.AddColumn("kind").Type("varchar(10)")
films.AddColumn("len").Type("interval hour to minute")
films.AddConstraint().Name("code_title").PrimaryKey().Columns().Add("code", "title")
```

###### Table with check constraint
```
CREATE TABLE distributors (
    did     integer CHECK (did > 100),
    name    varchar(40)
);

distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer").Constraint().Check().Expression().Add("did > 100")
distributors.AddColumn("name").Type("varchar(40)")
```

###### Table with default values in column definition
```
CREATE TABLE distributors (
    name      varchar(40) DEFAULT 'Luso Films',
    did       integer DEFAULT nextval('distributors_serial'),
    modtime   timestamp DEFAULT current_timestamp
);

distributors := gosql.CreateTable("distributors")
distributors.AddColumn("name").Type("varchar(40)").Constraint().Default("'Luso Films'")
distributors.AddColumn("did").Type("integer").Constraint().Default("nextval('distributors_serial')")
distributors.AddColumn("modtime").Type("timestamp").Constraint().Default("current_timestamp")
```

###### Table with tablespace
```
CREATE TABLE cinemas (
    id serial,
    name text,
    location text
) TABLESPACE diskvol1;

cinemas := gosql.CreateTable("cinemas")
cinemas.AddColumn("id").Type("serial")
cinemas.AddColumn("name").Type("text")
cinemas.AddColumn("location").Type("text")
cinemas.TableSpace("diskvol1")
```

###### Table with options and default constraint
```
CREATE TABLE employees OF employee_type (
    PRIMARY KEY (name),
    salary WITH OPTIONS DEFAULT 1000
);

employees := gosql.CreateTable("employees")
employees.OfType().Name("employee_type")
employees.OfType().Columns().AddColumn("name").Constraint().PrimaryKey()
salary := employees.OfType().Columns().AddColumn("salary")
salary.Constraint().Default("1000")
salary.WithOptions()
```

###### Table with excluding definition
```
CREATE TABLE circles (
    c circle,
    EXCLUDE USING gist (c WITH &&)
);

circles := CreateTable("circles")
circles.AddColumn("c").Type("circle")
exclude := circles.AddConstraint().Exclude().Using("gist")
exclude.ExcludeElement().Expression().Add("c")
exclude.With().Add("&&")
```

###### Table with named check constraint with multiple condition
```
CREATE TABLE distributors (
    did     integer,
    name    varchar(40),
    CONSTRAINT con1 CHECK (did > 100 AND name <> '')
);
distributors := gosql.CreateTable("distributors")
distributors.AddColumn("did").Type("integer")
distributors.AddColumn("name").Type("varchar(40)")
distributors.AddConstraint().Name("con1").
    Check().
    AddExpression("did > 100").
    AddExpression("name <> ''")
```

###### Table with partition by range and clause
```
CREATE TABLE measurement_year_month (
    logdate         date not null,
    peaktemp        int,
    unitsales       int
) PARTITION BY RANGE (EXTRACT(YEAR FROM logdate), EXTRACT(MONTH FROM logdate));

measurement = gosql.CreateTable("measurement_year_month")
measurement.AddColumn("logdate").Type("date").Constraint().NotNull()
measurement.AddColumn("peaktemp").Type("int")
measurement.AddColumn("unitsales").Type("int")
measurement.Partition().By(gosql.PartitionByRange).Clause("EXTRACT(YEAR FROM logdate)", "EXTRACT(MONTH FROM logdate)")
```

###### Table with partition by hash
```
CREATE TABLE orders (
    order_id     bigint not null,
    cust_id      bigint not null,
    status       text
) PARTITION BY HASH (order_id);

orders := gosql.CreateTable("orders")
orders.AddColumn("order_id").Type("bigint").Constraint().NotNull()
orders.AddColumn("cust_id").Type("bigint").Constraint().NotNull()
orders.AddColumn("status").Type("text")
orders.Partition().By(gosql.PartitionByHash).Clause("order_id")
```

###### Table with partition for values
```
CREATE TABLE measurement_y2016m07
    PARTITION OF measurement (
    unitsales DEFAULT 0
) FOR VALUES FROM ('2016-07-01') TO ('2016-08-01');

measurement := gosql.CreateTable("measurement_y2016m07")
measurement.OfPartition().Parent("measurement")
measurement.OfPartition().Columns().AddColumn("unitsales").Constraint().Default("0")
measurement.OfPartition().Values().From().Add("'2016-07-01'")
measurement.OfPartition().Values().To().Add("'2016-08-01'")
```

###### Table with partition for values with constant MINVALUE
```
CREATE TABLE measurement_ym_older
    PARTITION OF measurement_year_month
    FOR VALUES FROM (MINVALUE, MINVALUE) TO (2016, 11);
    
measurement = gosql.CreateTable("measurement_ym_older")
measurement.OfPartition().Parent("measurement_year_month")
measurement.OfPartition().Values().From().Add(PartitionBoundFromMin, PartitionBoundFromMin)
measurement.OfPartition().Values().To().Add("2016", "11")
```

###### Table with partition for values by range
```
CREATE TABLE cities_ab
    PARTITION OF cities (
    CONSTRAINT city_id_nonzero CHECK (city_id != 0)
) FOR VALUES IN ('a', 'b') PARTITION BY RANGE (population);

cities = gosql.CreateTable("cities_ab")
cities.OfPartition().Parent("cities")
cities.OfPartition().Columns().AddConstraint().Name("city_id_nonzero").Check().AddExpression("city_id != 0")
cities.OfPartition().Values().In().Add("'a'", "'b'")
cities.Partition().By(PartitionByRange).Clause("population")
```

###### Table with partition for values from to
```
CREATE TABLE cities_ab_10000_to_100000
    PARTITION OF cities_ab FOR VALUES FROM (10000) TO (100000);
    
citiesAb := gosql.CreateTable("cities_ab_10000_to_100000")
citiesAb.OfPartition().Parent("cities_ab").Values().From().Add("10000")
citiesAb.OfPartition().Parent("cities_ab").Values().To().Add("100000")
```

###### Table with default partition
```
CREATE TABLE cities_partdef
    PARTITION OF cities DEFAULT;
    
citiesPartdef := gosql.CreateTable("cities_partdef")
citiesPartdef.OfPartition().Parent("cities")
```

###### Table with primary key generated by default
```
CREATE TABLE distributors (
    did    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name   varchar(40) NOT NULL CHECK (name <> '')
);

distributors := gosql.CreateTable("distributors")
did := distributors.AddColumn("did").Type("integer")
did.Constraint().PrimaryKey()
did.Constraint().Generated().SetDetail(GeneratedByDefault)
name := distributors.AddColumn("name").Type("varchar(40)")
name.Constraint().NotNull()
name.Constraint().Check().Expression().Add("name <> ''")
```