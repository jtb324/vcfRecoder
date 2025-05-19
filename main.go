package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func handle_err(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func recode_genotype(genotype_call string) string {
	var return_call string

	splitCall := strings.Split(genotype_call, ":")

	genotype := splitCall[0]

	switch genotype {

	case "0/0":
		return_call = "0"

	case "./.":
		return_call = "."
	default:
		depth_scores := strings.Split(splitCall[2], ",")

		max_value := 0
		max_value_indx := 0

		for indx, score := range depth_scores {
			convertedScore, err := strconv.Atoi(score)
			handle_err(err, fmt.Sprintf("err converting genotype score %s to an integer", score))
			if convertedScore > max_value {
				max_value_indx = indx
			}
		}
		return_call = strconv.Itoa(max_value_indx)
	}

	return return_call
}

func process_line(line string) string {
	var s strings.Builder

	splitStr := strings.Split(strings.TrimSpace(line), "\t")

	prefix := splitStr[0:9]
	s.WriteString(strings.Join(prefix, "\t"))

	for _, call := range splitStr[9:] {
		s.WriteString("\t")
		recodedCall := recode_genotype(call)
		s.WriteString(recodedCall)
	}

	s.WriteString("\n")
	return s.String()
}

func main() {
	cmd_args := os.Args

	input_file := cmd_args[1]
	output_file := cmd_args[2]

	// opening the input
	filehandle, err := os.Open(input_file)

	handle_err(err, fmt.Sprintf("There was an error reading in the file: %s", input_file))

	defer filehandle.Close()

	gz_handle, err := gzip.NewReader(filehandle)

	handle_err(err, fmt.Sprintf("There was an issue while uncompressing the file: %s", input_file))

	defer gz_handle.Close()

	bc := bufio.NewReader(gz_handle)

	// Now we ned to create an output file to write to
	output_filehandle, err := os.Create(output_file)

	handle_err(err, fmt.Sprintf("There was an issue while opening the otuput file: %s", output_file))

	defer output_filehandle.Close()

	gzipWriter := gzip.NewWriter(output_filehandle)

	defer gzipWriter.Close()

	for {
		line, err := bc.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading in line")
		}
		if !strings.Contains(line, "#") {
			processedStr := process_line(line)
			fmt.Println(processedStr)
			gzipWriter.Write([]byte(processedStr))
		} else {
			gzipWriter.Write([]byte(line))
		}
	}
	gzipWriter.Flush()
}
