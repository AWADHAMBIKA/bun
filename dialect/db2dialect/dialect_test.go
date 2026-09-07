package db2dialect

import (
	"testing"

	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/dialect/feature"
)

func TestDialect(t *testing.T) {
	d := New()

	if d.Name() != dialect.DB2 {
		t.Fatalf("expected dialect.DB2, got %v", d.Name())
	}
	if d.IdentQuote() != '"' {
		t.Fatalf("expected identifier quote to be a double quote")
	}
	if !d.Features().Has(feature.OffsetFetch) {
		t.Fatal("expected OffsetFetch feature to be enabled")
	}
	if !d.Features().Has(feature.Identity) {
		t.Fatal("expected Identity feature to be enabled")
	}
	if d.DefaultVarcharLen() != 255 {
		t.Fatalf("expected default varchar length 255, got %d", d.DefaultVarcharLen())
	}
	if d.DefaultSchema() != "" {
		t.Fatalf("expected empty default schema (current authorization ID), got %q", d.DefaultSchema())
	}
}

func TestAppendBool(t *testing.T) {
	d := New()

	if got := string(d.AppendBool(nil, true)); got != "1" {
		t.Fatalf("expected %q, got %q", "1", got)
	}
	if got := string(d.AppendBool(nil, false)); got != "0" {
		t.Fatalf("expected %q, got %q", "0", got)
	}
}
