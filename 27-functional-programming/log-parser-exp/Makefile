SHELL := /bin/bash
LINES = echo -e ">> log.txt has $$(wc -l log.txt | cut -f1 -d' ') lines"

ifeq ($(n),)
	n := 18
endif

s:
	time go run . < log.txt

r:
	go run . < log.txt

# make n=18
load: restore
	@echo "enlarging the file with itself, please wait..."
	@for i in {1..$(n)}; do awk 1 log.txt log.txt > log_.txt; mv log_.txt log.txt; rm -f log_.txt; done
	@$(LINES)

restore:
	@$(LINES)
	git checkout log.txt
	
lines:
	@$(LINES)