.PHONY: pb
pb: 
	@./scripts/generatepb.sh

# pos
.PHONY: pos-build
	@cd scripts/pos && make build

.PHONY: pos-run
pos-run:
	@cd scripts/pos && make run

.PHONY: pos-debug
pos-debug:
	@cd scripts/pos && make debug 

# banking
.PHONY: banking-build
banking-build:
	@cd scripts/banking && make build

.PHONY: banking-run
banking-run:
	@cd scripts/banking && make run

.PHONY: banking-debug
banking-debug:
	@cd scripts/banking && make debug 

