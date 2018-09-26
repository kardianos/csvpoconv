package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	o := flag.String("o", "out.csv", "output file name")
	flag.Parse()

	if len(*o) == 0 {
		log.Fatal("missing output (o) parameter")
	}

	if len(flag.Args()) == 0 {
		log.Fatal("missing input files after exec")
	}

	of, err := os.OpenFile(*o, os.O_CREATE|os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		log.Fatalf("failed to open %q: %v", *o, err)
	}
	defer of.Close()

	w := csv.NewWriter(of)
	w.UseCRLF = true

	for _, rawFilename := range flag.Args() {
		filenames, err := filepath.Glob(rawFilename)
		if err != nil {
			log.Fatalf("failed to glob %q: %v", rawFilename, err)
		}
		for _, filename := range filenames {
			err = readFile(w, filename)
			if err != nil {
				log.Fatalf("failed to read file %q: %v", filename, err)
			}
		}
	}
}

func readFile(w *csv.Writer, filename string) error {
	rf, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer rf.Close()
	return readFileReader(w, rf)
}

func readFileReader(w *csv.Writer, rf io.Reader) error {
	r := csv.NewReader(rf)
	r.FieldsPerRecord = -1
	r.TrimLeadingSpace = true
	r.TrailingComma = true

	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				w.Flush()
				return w.Error()
			}
			return err
		}
		if len(row) < 7 {
			continue
		}

		date := row[0]
		po := row[1]
		loc := row[2]
		proc := row[3]

		itemNumb := row[4]
		qty := row[5]
		price := row[6]

		if date == "Date" {
			continue
		}
		if strings.HasPrefix(date, "Week") {
			continue
		}
		w.Write([]string{"H", po, proc, date, loc, loc})
		w.Write([]string{"R", itemNumb, qty, strings.TrimLeft(price, "$ ")})
	}
}
