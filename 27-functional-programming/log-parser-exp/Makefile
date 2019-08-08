SHELL := /bin/bash
LINES = $$(wc -l log.txt | cut -f1 -d' ')
ECHO_LINES = echo -e ">> log.txt has $(LINES) lines"

s: 
	time go run . < log.txt

r: 
	go run . < log.txt

n ?= 18
load: restore
	@echo "enlarging the file with itself, please wait..."	

	@for i in {1..$(n)}; do \
		awk 1 log.txt log.txt > log_.txt; \
		mv log_.txt log.txt; \
		rm -f log_.txt; \
	done
	
	@$(ECHO_LINES)

restore:
	@echo "restoring the file..."
	@git checkout log.txt
	@$(ECHO_LINES)

lines:
	@$(ECHO_LINES)