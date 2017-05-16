package csv

import (
	"encoding/csv"
	"github.com/yaronsumel/geo-coding-test/src/place"
	"io"
	"os"
	"strconv"
)

// CSV Handler implements Data Interface
type Handler struct {
	file      *os.File
	csvReader *csv.Reader
}

// NewFileHandler - return new FileHandler or error
// open the requested file and create the csv
// create csv reader to help us decode it easily (and well tested)
func NewCsvHandler(path string) (*Handler, error) {
	// open requested file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// create new csv reader
	cr := csv.NewReader(f)
	// skip csv header
	cr.Read()
	return &Handler{
		file:      f,
		csvReader: cr,
	}, nil
}

// Next - Implements the Next(p *Place) error function
// Read next line onto p and return error at EOF
// Usage:
// p := Place{}
// for f.Next(&p){
//	  do something with p
// }
// do something with eof
func (h *Handler) Next(p *place.Place) bool {
	// keep reading till error comes
	for {
		// read line
		line, err := h.csvReader.Read()
		// return EOF
		// continue on any other error
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		// parseLine into p
		// on error keep going
		if err := h.parseLine(line, p); err != nil {
			continue
		}
		// return true - will keep looping out next
		return true
	}
	return false
}

// Close implements Close()error
// close all readers and resources
func (h *Handler) Close() error {
	// close os.File
	return h.file.Close()
}

// parseLine - populate p based on value positions in line
// pos 0 - PlaceID
// pos 1 - Lat
// pos 2 - Lon
// return error if something went wrong
func (h *Handler) parseLine(line []string, p *place.Place) error {
	for k, v := range line {
		switch k {
		case 0:
			x, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			p.Id = x
		case 1:
			x, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			p.Lat = x
		case 2:
			x, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			p.Lon = x
		}
	}
	return nil
}
