package gosql

import "testing"

func TestIndex_String(t *testing.T) {
	t.Run("title_idx", func(t *testing.T) {
		// CREATE UNIQUE INDEX title_idx ON films (title);
		idx := CreateIndex("films", "title").Name("title_idx").Unique()
		t.Log(idx.String())
		if idx.String() != "CREATE UNIQUE INDEX title_idx ON films (title);" {
			t.Fatal("wrong title_idx")
		}
		idx = CreateIndex().Table("films").Name("title_idx").Unique()
		idx.Expression().Add("title")
		t.Log(idx.String())
		if idx.String() != "CREATE UNIQUE INDEX title_idx ON films (title);" {
			t.Fatal("wrong title_idx")
		}
		idx = CreateIndex("films", "title").Unique().AutoName()
		t.Log(idx.String())
		if idx.String() != "CREATE UNIQUE INDEX films_title_uidx ON films (title);" {
			t.Fatal("wrong films_title_uidx")
		}
	})

	t.Run("include", func(t *testing.T) {
		// CREATE UNIQUE INDEX title_idx ON films (title) INCLUDE (director, rating);
		idx := CreateIndex("films", "title").Name("title_idx").Include("director", "rating").Unique()
		t.Log(idx.String())
		if idx.String() != "CREATE UNIQUE INDEX title_idx ON films (title) INCLUDE (director, rating);" {
			t.Fatal("wrong include")
		}
		idx = CreateIndex("films", "title").AutoName().Include("director", "rating").Unique()
		t.Log(idx.String())
		if idx.String() != "CREATE UNIQUE INDEX films_title_uidx ON films (title) INCLUDE (director, rating);" {
			t.Fatal("wrong include")
		}
	})

	t.Run("with", func(t *testing.T) {
		// CREATE INDEX title_idx ON films (title) WITH (deduplicate_items = off);
		idx := CreateIndex("films", "title").Name("title_idx").With("deduplicate_items = off")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX title_idx ON films (title) WITH (deduplicate_items = off);" {
			t.Fatal("wrong with")
		}
	})

	t.Run("expression", func(t *testing.T) {
		// CREATE INDEX ON films ((lower(title)));
		idx := CreateIndex("films", "(lower(title))")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX ON films ((lower(title)));" {
			t.Fatal("wrong expression")
		}
	})

	t.Run("collate", func(t *testing.T) {
		// CREATE INDEX title_idx_german ON films (title COLLATE "de_DE");
		idx := CreateIndex("films", `title COLLATE "de_DE"`).Name("title_idx_german")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX title_idx_german ON films (title COLLATE \"de_DE\");" {
			t.Fatal("wrong collate")
		}
	})

	t.Run("nulls_first", func(t *testing.T) {
		// CREATE INDEX title_idx_nulls_low ON films (title NULLS FIRST);
		idx := CreateIndex("films", `title NULLS FIRST`).Name("title_idx_nulls_low")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX title_idx_nulls_low ON films (title NULLS FIRST);" {
			t.Fatal("wrong nulls_first")
		}
	})

	t.Run("fillfactor", func(t *testing.T) {
		// CREATE UNIQUE INDEX title_idx ON films (title) WITH (fillfactor = 70);
		idx := CreateIndex("films", "title").Name("title_idx").With("fillfactor = 70").Unique()
		t.Log(idx.String())
		if idx.String() != "CREATE UNIQUE INDEX title_idx ON films (title) WITH (fillfactor = 70);" {
			t.Fatal("wrong fillfactor")
		}
	})

	t.Run("gin", func(t *testing.T) {
		// CREATE INDEX gin_idx ON documents_table USING GIN (locations) WITH (fastupdate = off);
		idx := CreateIndex("documents_table", "locations").Name("gin_idx").With("fastupdate = off").Using("GIN")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX gin_idx ON documents_table USING GIN (locations) WITH (fastupdate = off);" {
			t.Fatal("wrong gin")
		}
	})

	t.Run("tablespace", func(t *testing.T) {
		// CREATE INDEX code_idx ON films (code) TABLESPACE indexspace;
		idx := CreateIndex("films", "code").Name("code_idx").TableSpace("indexspace")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX code_idx ON films (code) TABLESPACE indexspace;" {
			t.Fatal("wrong gin")
		}
	})

	t.Run("gist", func(t *testing.T) {
		// CREATE INDEX pointloc ON points USING gist (box(location,location));
		idx := CreateIndex("points", "box(location,location)").Name("pointloc").Using("gist")
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX pointloc ON points USING gist (box(location,location));" {
			t.Fatal("wrong gin")
		}
	})

	t.Run("concurrently", func(t *testing.T) {
		// CREATE INDEX CONCURRENTLY sales_quantity_index ON sales_table (quantity);
		idx := CreateIndex("sales_table", "quantity").Name("sales_quantity_index").Concurrently()
		t.Log(idx.String())
		if idx.String() != "CREATE INDEX CONCURRENTLY sales_quantity_index ON sales_table (quantity);" {
			t.Fatal("wrong gin")
		}
	})
}
