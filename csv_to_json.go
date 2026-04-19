package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func infer_type(value string) any {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil
	}

	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}

	if value == "true" || value == "false" {
		b, _ := strconv.ParseBool(value)
		return b
	}

	return value

}

func csv_to_json(csv_file, json_file string) error {
	file, err := os.Open(csv_file)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	out, err := os.Create(json_file)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := json.NewEncoder(out)

	header, err := reader.Read()
	if err != nil {
		return err
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		jsonl_map := make(map[string]any)

		for i, col := range header {
			jsonl_map[col] = infer_type(line[i])
		}

		if err := writer.Encode(jsonl_map); err != nil {
			return err
		}

	}

	return nil
}

func main() {

	args := os.Args

	for _, csv_file := range args[1:] {

		json_file := ""
		if strings.HasSuffix(csv_file, ".csv") {
			json_file = strings.TrimSuffix(csv_file, ".csv") + ".jsonl"
		} else {
			fmt.Println("not a .csv file")
			fmt.Println("Exiting")
			os.Exit(0)
		}

		err := csv_to_json(csv_file, json_file)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
