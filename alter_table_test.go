package gosql

import (
	"testing"
)

func TestAlter(t *testing.T) {
	t.Run("add_column", func(t *testing.T) {
		alter := AlterTable("dictionary").IfExists()
		alter.Action().Add().Column("name", "TEXT").Constraint().NotNull()
		alter.Action().Add().Column("created_at", "TIMESTAMP WITH TIME ZONE").Constraint().NotNull().Default("localtimestamp")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE IF EXISTS dictionary ADD COLUMN name TEXT NOT NULL, ADD COLUMN created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT localtimestamp;" {
			t.Fatal("wrong add_column")
		}
	})
	t.Run("rename_column", func(t *testing.T) {
		alter := AlterTable("dictionary").IfExists().RenameColumn("name", "full_name")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE IF EXISTS dictionary RENAME COLUMN name TO full_name;" {
			t.Fatal("wrong rename_column")
		}
	})
	t.Run("attach_partition", func(t *testing.T) {
		alter := AlterTable("dictionary").IfExists()
		partition := alter.AttachPartition("my_partition")
		partition.From().Add("'2016-07-01'")
		partition.To().Add("'2016-08-01'")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE IF EXISTS dictionary ATTACH PARTITION my_partition FOR VALUES FROM ('2016-07-01') TO ('2016-08-01');" {
			t.Fatal("wrong attach_partition")
		}
	})
	t.Run("add_varchar_column", func(t *testing.T) {
		// ALTER TABLE distributors ADD COLUMN address varchar(30);
		alter := AlterTable("distributors")
		alter.Action().Add().Column("address", "varchar(30)")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD COLUMN address varchar(30);" {
			t.Fatal("wrong add_varchar_column")
		}
	})
	t.Run("add_column_with_constraint", func(t *testing.T) {
		// ALTER TABLE measurements ADD COLUMN mtime timestamp with time zone DEFAULT now();
		alter := AlterTable("measurements")
		alter.Action().Add().Column("mtime", "timestamp with time zone").Constraint().Default("now()")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE measurements ADD COLUMN mtime timestamp with time zone DEFAULT now();" {
			t.Fatal("wrong add_column_with_constraint")
		}
	})
	t.Run("alter_column", func(t *testing.T) {
		// ALTER TABLE transactions
		//    ALTER COLUMN status SET default 'current';
		alter := AlterTable("transactions")
		alter.Action().AlterColumn("status").Set().Default("'current'")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE transactions ALTER COLUMN status SET DEFAULT 'current';" {
			t.Fatal("wrong alter_column")
		}
	})
	t.Run("add_column_and_alter_column", func(t *testing.T) {
		// ALTER TABLE transactions
		//    ADD COLUMN status varchar(30) DEFAULT 'old',
		//    ALTER COLUMN status SET default 'current';
		alter := AlterTable("transactions")
		alter.Action().Add().Column("status", "varchar(30)").Constraint().Default("'old'")
		alter.Action().AlterColumn("status").Set().Default("'current'")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE transactions ADD COLUMN status varchar(30) DEFAULT 'old', ALTER COLUMN status SET DEFAULT 'current';" {
			t.Fatal("wrong add_column_and_alter_column")
		}
	})
	t.Run("drop_column", func(t *testing.T) {
		// ALTER TABLE distributors DROP COLUMN address RESTRICT;
		alter := AlterTable("distributors")
		alter.Action().Drop().Column("address").Restrict()
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors DROP COLUMN address RESTRICT;" {
			t.Fatal("wrong drop_column")
		}
	})
	t.Run("change_column_types", func(t *testing.T) {
		// ALTER TABLE distributors
		//    ALTER COLUMN address SET DATA TYPE varchar(80),
		//    ALTER COLUMN name SET DATA TYPE varchar(100);
		alter := AlterTable("distributors")
		alter.Action().AlterColumn("address").Set().DataType("varchar(80)")
		alter.Action().AlterColumn("name").Set().DataType("varchar(100)")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ALTER COLUMN address SET DATA TYPE varchar(80), ALTER COLUMN name SET DATA TYPE varchar(100);" {
			t.Fatal("wrong change_column_types")
		}
	})
	t.Run("change_column_types_using", func(t *testing.T) {
		// ALTER TABLE foo
		//    ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone
		//    USING
		//    timestamp with time zone 'epoch' + foo_timestamp * interval '1 second';
		alter := AlterTable("foo")
		alter.Action().AlterColumn("foo_timestamp").Set().DataType("timestamp with time zone").
			Using("timestamp with time zone 'epoch' + foo_timestamp * interval '1 second'")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE foo ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone USING timestamp with time zone 'epoch' + foo_timestamp * interval '1 second';" {
			t.Fatal("wrong change_column_types_using")
		}
	})
	t.Run("change_column_types_using_with_default", func(t *testing.T) {
		// ALTER TABLE foo
		//    ALTER COLUMN foo_timestamp DROP DEFAULT,
		//    ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone
		//        USING
		//        timestamp with time zone 'epoch' + foo_timestamp * interval '1 second',
		//    ALTER COLUMN foo_timestamp SET DEFAULT now();
		alter := AlterTable("foo")
		alter.Action().AlterColumn("foo_timestamp").Drop().Default()
		alter.Action().AlterColumn("foo_timestamp").Set().DataType("timestamp with time zone").
			Using("timestamp with time zone 'epoch' + foo_timestamp * interval '1 second'")
		alter.Action().AlterColumn("foo_timestamp").Set().Default("now()")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE foo ALTER COLUMN foo_timestamp DROP DEFAULT, ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone USING timestamp with time zone 'epoch' + foo_timestamp * interval '1 second', ALTER COLUMN foo_timestamp SET DEFAULT now();" {
			t.Fatal("wrong change_column_types_using_with_default")
		}
	})
	t.Run("rename_existing_column", func(t *testing.T) {
		// ALTER TABLE distributors RENAME COLUMN address TO city;
		alter := AlterTable("distributors").RenameColumn("address", "city")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors RENAME COLUMN address TO city;" {
			t.Fatal("wrong rename_existing_column")
		}
	})
	t.Run("rename_table", func(t *testing.T) {
		// ALTER TABLE distributors RENAME TO suppliers;
		alter := AlterTable("distributors").Rename("suppliers")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors RENAME TO suppliers;" {
			t.Fatal("wrong rename_table")
		}
	})
	t.Run("rename_constraint", func(t *testing.T) {
		// ALTER TABLE distributors RENAME CONSTRAINT zipchk TO zip_check;
		alter := AlterTable("distributors").RenameConstraint("zipchk", "zip_check")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors RENAME CONSTRAINT zipchk TO zip_check;" {
			t.Fatal("wrong rename_constraint")
		}
	})
	t.Run("set_column_not_null", func(t *testing.T) {
		// ALTER TABLE distributors ALTER COLUMN street SET NOT NULL;
		alter := AlterTable("distributors")
		alter.Action().AlterColumn("street").Set().NotNull()
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ALTER COLUMN street SET NOT NULL;" {
			t.Fatal("wrong set_column_not_null")
		}
	})
	t.Run("drop_column_not_null", func(t *testing.T) {
		// ALTER TABLE distributors ALTER COLUMN street DROP NOT NULL;
		alter := AlterTable("distributors")
		alter.Action().AlterColumn("street").Drop().NotNull()
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ALTER COLUMN street DROP NOT NULL;" {
			t.Fatal("wrong drop_column_not_null")
		}
	})
	t.Run("add_constraint", func(t *testing.T) {
		// ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5);
		alter := AlterTable("distributors")
		alter.Action().Add().TableConstraint().Name("zipchk").Check().AddExpression("char_length(zipcode) = 5")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5);" {
			t.Fatal("wrong add_constraint")
		}
	})
	t.Run("add_constraint_no_inherit", func(t *testing.T) {
		// ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5) NO INHERIT;
		alter := AlterTable("distributors")
		alter.Action().Add().TableConstraint().Name("zipchk").NoInherit().Check().AddExpression("char_length(zipcode) = 5")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5) NO INHERIT;" {
			t.Fatal("wrong add_constraint")
		}
	})
	t.Run("drop_constraint", func(t *testing.T) {
		// ALTER TABLE distributors DROP CONSTRAINT zipchk;
		alter := AlterTable("distributors")
		alter.Action().Drop().Constraint("zipchk")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors DROP CONSTRAINT zipchk;" {
			t.Fatal("wrong drop_constraint")
		}
	})
	t.Run("drop_constraint_only", func(t *testing.T) {
		// ALTER TABLE ONLY distributors DROP CONSTRAINT zipchk;
		alter := AlterTable("distributors").Only()
		alter.Action().Drop().Constraint("zipchk")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE ONLY distributors DROP CONSTRAINT zipchk;" {
			t.Fatal("wrong drop_constraint_only")
		}
	})
	t.Run("add_constraint_foreign", func(t *testing.T) {
		// ALTER TABLE distributors ADD CONSTRAINT distfk FOREIGN KEY (address) REFERENCES addresses (address);
		alter := AlterTable("distributors")
		alter.Action().Add().TableConstraint().Name("distfk").
			ForeignKey().Column("address").References().RefTable("addresses").Column("address")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD CONSTRAINT distfk FOREIGN KEY (address) REFERENCES addresses (address);" {
			t.Fatal("wrong add_constraint_foreign")
		}
	})
	t.Run("add_constraint_foreign_and_validate", func(t *testing.T) {
		// ALTER TABLE distributors ADD CONSTRAINT distfk FOREIGN KEY (address) REFERENCES addresses (address) NOT VALID;
		alter := AlterTable("distributors")
		alter.Action().Add().NotValid().TableConstraint().Name("distfk").
			ForeignKey().Column("address").References().RefTable("addresses").Column("address")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD CONSTRAINT distfk FOREIGN KEY (address) REFERENCES addresses (address) NOT VALID;" {
			t.Fatal("wrong add_constraint_foreign_and_validate")
		}
		// ALTER TABLE distributors VALIDATE CONSTRAINT distfk;
		alter = AlterTable("distributors")
		alter.Action().ValidateConstraint("distfk")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors VALIDATE CONSTRAINT distfk;" {
			t.Fatal("wrong add_constraint_foreign_and_validate")
		}
	})
	t.Run("add_multicolumn_unique_constraint", func(t *testing.T) {
		// ALTER TABLE distributors ADD CONSTRAINT dist_id_zipcode_key UNIQUE (dist_id, zipcode);
		alter := AlterTable("distributors")
		alter.Action().Add().TableConstraint().Name("dist_id_zipcode_key").Unique().Column("dist_id", "zipcode")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD CONSTRAINT dist_id_zipcode_key UNIQUE (dist_id, zipcode);" {
			t.Fatal("wrong add_multicolumn_unique_constraint")
		}
	})
	t.Run("add_primary_key", func(t *testing.T) {
		// ALTER TABLE distributors ADD PRIMARY KEY (dist_id);
		alter := AlterTable("distributors")
		alter.Action().Add().TableConstraint().PrimaryKey().Column("dist_id")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors ADD PRIMARY KEY (dist_id);" {
			t.Fatal("wrong add_primary_key")
		}
	})
	t.Run("set_table_space", func(t *testing.T) {
		// ALTER TABLE distributors SET TABLESPACE fasttablespace;
		alter := AlterTable("distributors").SetTableSpace("fasttablespace")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors SET TABLESPACE fasttablespace;" {
			t.Fatal("wrong set_table_space")
		}
	})
	t.Run("set_schema", func(t *testing.T) {
		// ALTER TABLE myschema.distributors SET SCHEMA yourschema;
		alter := AlterTable("myschema.distributors").SetSchema("yourschema")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE myschema.distributors SET SCHEMA yourschema;" {
			t.Fatal("wrong set_schema")
		}
	})
	t.Run("recreate_primary_key", func(t *testing.T) {
		// CREATE UNIQUE INDEX CONCURRENTLY dist_id_temp_idx ON distributors (dist_id);
		unique := CreateIndex("distributors", "dist_id").Name("dist_id_temp_idx").Concurrently().Unique()
		t.Log(unique.String())
		if unique.String() != "CREATE UNIQUE INDEX CONCURRENTLY dist_id_temp_idx ON distributors (dist_id);" {
			t.Fatal("wrong recreate_primary_key index")
		}
		// ALTER TABLE distributors DROP CONSTRAINT distributors_pkey,
		//                         ADD CONSTRAINT distributors_pkey PRIMARY KEY USING INDEX dist_id_temp_idx;
		alter := AlterTable("distributors")
		alter.Action().Drop().Constraint("distributors_pkey")
		alter.Action().Add().TableConstraintUsingIndex().Name("distributors_pkey").PrimaryKey().Using("dist_id_temp_idx")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE distributors DROP CONSTRAINT distributors_pkey, ADD CONSTRAINT distributors_pkey PRIMARY KEY USING INDEX dist_id_temp_idx;" {
			t.Fatal("wrong recreate_primary_key")
		}
	})
	t.Run("attach_partition", func(t *testing.T) {
		// ALTER TABLE measurement
		//    ATTACH PARTITION measurement_y2016m07 FOR VALUES FROM ('2016-07-01') TO ('2016-08-01');
		alter := AlterTable("measurement")
		bound := alter.AttachPartition("measurement_y2016m07")
		bound.From().Add("'2016-07-01'")
		bound.To().Add("'2016-08-01'")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE measurement ATTACH PARTITION measurement_y2016m07 FOR VALUES FROM ('2016-07-01') TO ('2016-08-01');" {
			t.Fatal("wrong attach_partition")
		}
	})
	t.Run("attach_partition_to_a_list", func(t *testing.T) {
		// ALTER TABLE cities
		//    ATTACH PARTITION cities_ab FOR VALUES IN ('a', 'b');
		alter := AlterTable("cities")
		alter.AttachPartition("cities_ab").In().Add("'a'", "'b'")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE cities ATTACH PARTITION cities_ab FOR VALUES IN ('a', 'b');" {
			t.Fatal("wrong attach_partition_to_a_list")
		}
	})
	t.Run("attach_partition_for_values_with", func(t *testing.T) {
		// ALTER TABLE orders
		//    ATTACH PARTITION orders_p4 FOR VALUES WITH (MODULUS 4, REMAINDER 3);
		alter := AlterTable("orders")
		alter.AttachPartition("orders_p4").With().Add("MODULUS 4", "REMAINDER 3")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE orders ATTACH PARTITION orders_p4 FOR VALUES WITH (MODULUS 4, REMAINDER 3);" {
			t.Fatal("wrong attach_partition_for_values_with")
		}
	})
	t.Run("attach_partition_default", func(t *testing.T) {
		// ALTER TABLE cities
		//    ATTACH PARTITION cities_partdef DEFAULT;
		alter := AlterTable("cities")
		alter.AttachDefaultPartition("cities_partdef")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE cities ATTACH PARTITION cities_partdef DEFAULT;" {
			t.Fatal("wrong attach_partition_default")
		}
	})
	t.Run("detach_partition", func(t *testing.T) {
		// ALTER TABLE measurement
		//    DETACH PARTITION measurement_y2015m12;
		alter := AlterTable("measurement")
		alter.DetachPartition("measurement_y2015m12")
		t.Log(alter.String())
		if alter.String() != "ALTER TABLE measurement DETACH PARTITION measurement_y2015m12;" {
			t.Fatal("wrong detach_partition")
		}
	})
}
