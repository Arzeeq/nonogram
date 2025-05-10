package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"nonogram"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	sc := bufio.NewScanner(file)
	if !sc.Scan() {
		log.Fatalf("nonogram size is not provided")
	}

	n, m, err := parseSize(sc.Text())
	if err != nil {
		log.Fatalf("failed to parse nonogram size: %v", err)
	}

	rows := make(nonogram.FillPattern, n)
	for i := range n {
		if !sc.Scan() {
			log.Fatalf("not enought rows")
		}

		p, err := parsePatternLine(sc.Text())
		if err != nil {
			log.Fatalf("invalid pattern '%s'", sc.Text())
		}

		rows[i] = p
	}

	columns := make(nonogram.FillPattern, m)
	for i := range m {
		if !sc.Scan() {
			log.Fatalf("not enought rows")
		}

		p, err := parsePatternLine(sc.Text())
		if err != nil {
			log.Fatalf("invalid pattern '%s'", sc.Text())
		}

		columns[i] = p
	}

	var s nonogram.Solver
	if err := s.Solve(rows, columns); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(s.StringCaged(5))
	s.SavePNG("solved.png", 10)
}

func parseSize(s string) (int, int, error) {
	sizeStr := strings.Split(s, " ")
	if len(sizeStr) != 2 {
		return 0, 0, errors.New("expected two numbers N and M")
	}

	n, err := strconv.Atoi(sizeStr[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse '%s' as N", sizeStr[0])
	}

	m, err := strconv.Atoi(sizeStr[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse '%s' as M", sizeStr[1])
	}

	return n, m, nil
}

func parsePatternLine(s string) ([]int, error) {
	strs := strings.Split(s, " ")

	res := make([]int, 0)
	for i := range strs {
		x, err := strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}

		res = append(res, x)
	}

	return res, nil
}
