package gosql

import "testing"

func TestMerge_SQL(t *testing.T) {
	t.Run("merge_1", func(t *testing.T) {
		// MERGE INTO customer_account ca
		// USING recent_transactions t
		//  ON t.customer_id = ca.customer_id
		//  WHEN MATCHED THEN
		//  UPDATE SET balance = balance + transaction_value
		//  WHEN NOT MATCHED THEN
		//  INSERT (customer_id, balance)
		//  VALUES (t.customer_id, t.transaction_value);
		m := NewMerge().
			Into("customer_account ca").
			Using("recent_transactions t ON t.customer_id = ca.customer_id")

		m.When().Update().Add("balance = balance + transaction_value")
		m.When().Insert().Columns("customer_id", "balance").Values().Add("t.customer_id", "t.transaction_value")

		t.Log(m.String())
		if m.String() != "MERGE INTO customer_account ca USING recent_transactions t ON t.customer_id = ca.customer_id WHEN MATCHED THEN UPDATE SET balance = balance + transaction_value WHEN NOT MATCHED THEN INSERT (customer_id, balance) VALUES (t.customer_id, t.transaction_value);" {
			t.Fatal("wrong merge_1")
		}
	})

	t.Run("merge_2", func(t *testing.T) {
		// MERGE INTO customer_account ca
		// USING (SELECT customer_id, transaction_value FROM recent_transactions) AS t
		//  ON t.customer_id = ca.customer_id
		//  WHEN MATCHED THEN
		//  UPDATE SET balance = balance + transaction_value
		//  WHEN NOT MATCHED THEN
		//  INSERT (customer_id, balance)
		//  VALUES (t.customer_id, t.transaction_value);

		sub := NewSelect().From("recent_transactions")
		sub.Columns().Add("customer_id", "transaction_value")
		sub.SubQuery = true

		m := NewMerge().
			Into("customer_account ca").
			Using(sub.String() + " AS t ON t.customer_id = ca.customer_id")

		m.When().Update().Add("balance = balance + transaction_value")
		m.When().Insert().Columns("customer_id", "balance").
			Values().Add("t.customer_id", "t.transaction_value")

		t.Log(m)
		if m.String() != "MERGE INTO customer_account ca USING (SELECT customer_id, transaction_value FROM recent_transactions) AS t ON t.customer_id = ca.customer_id WHEN MATCHED THEN UPDATE SET balance = balance + transaction_value WHEN NOT MATCHED THEN INSERT (customer_id, balance) VALUES (t.customer_id, t.transaction_value);" {
			t.Fatal("wrong merge 2")
		}
	})

	t.Run("merge_3", func(t *testing.T) {
		// MERGE INTO wines w
		// USING wine_stock_changes s
		//  ON s.winename = w.winename
		//  WHEN NOT MATCHED AND s.stock_delta > 0 THEN
		//  INSERT VALUES(s.winename, s.stock_delta)
		//  WHEN MATCHED AND w.stock + s.stock_delta > 0 THEN
		//  UPDATE SET stock = w.stock + s.stock_delta
		//  WHEN MATCHED THEN
		//  DELETE;

		m := NewMerge().
			Into("wines w").
			Using("wine_stock_changes s ON s.winename = w.winename")

		insertWhen := m.When()
		insertWhen.Condition().AddExpression("s.stock_delta > 0")
		insertWhen.Insert().Values().Add("s.winename", "s.stock_delta")

		updateWhen := m.When()
		updateWhen.Condition().AddExpression("w.stock + s.stock_delta > 0")
		updateWhen.Update().Add("stock = w.stock + s.stock_delta")

		m.When().Delete()

		t.Log(m.String())

		if m.String() != "MERGE INTO wines w USING wine_stock_changes s ON s.winename = w.winename WHEN NOT MATCHED AND (s.stock_delta > 0) THEN INSERT VALUES (s.winename, s.stock_delta) WHEN MATCHED AND (w.stock + s.stock_delta > 0) THEN UPDATE SET stock = w.stock + s.stock_delta WHEN MATCHED THEN DELETE;" {
			t.Fatal("merge_3")
		}
	})

	t.Run("merge_4", func(t *testing.T) {
		// MERGE INTO station_data_actual sda
		// USING station_data_new sdn
		//  ON sda.station_id = sdn.station_id
		//  WHEN MATCHED THEN
		//  UPDATE SET a = sdn.a, b = sdn.b, updated = DEFAULT
		//  WHEN NOT MATCHED THEN
		//  INSERT (station_id, a, b)
		//  VALUES (sdn.station_id, sdn.a, sdn.b);

		m := NewMerge().
			Into("station_data_actual sda").
			Using("station_data_new sdn ON sda.station_id = sdn.station_id")

		m.When().Update().Add("a = sdn.a", "b = sdn.b", "updated = DEFAULT")
		insertWhen := m.When().Insert()
		insertWhen.Columns("station_id", "a", "b")
		insertWhen.Values().Add("sdn.station_id", "sdn.a", "sdn.b")

		t.Log(m.String())

		if m.String() != "MERGE INTO station_data_actual sda USING station_data_new sdn ON sda.station_id = sdn.station_id WHEN MATCHED THEN UPDATE SET a = sdn.a, b = sdn.b, updated = DEFAULT WHEN NOT MATCHED THEN INSERT (station_id, a, b) VALUES (sdn.station_id, sdn.a, sdn.b);" {
			t.Fatal("merge_4")
		}
	})
}

// goos: darwin
// goarch: amd64
// pkg: github.com/dimonrus/gosql
// cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
// BenchmarkMerge_String
// BenchmarkMerge_String-8   	 1020259	      1504 ns/op	    1584 B/op	      21 allocs/op
func BenchmarkMerge_String(b *testing.B) {
	m := NewMerge().
		Into("station_data_actual sda").
		Using("station_data_new sdn ON sda.station_id = sdn.station_id")

	m.When().Update().Add("a = sdn.a", "b = sdn.b", "updated = DEFAULT")
	insertWhen := m.When().Insert()
	insertWhen.Columns("station_id", "a", "b")
	insertWhen.Values().Add("sdn.station_id", "sdn.a", "sdn.b")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.String()
	}
	b.ReportAllocs()
}
