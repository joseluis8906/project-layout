.PHONY: pb
pb: 
	@./scripts/generatepb.sh
# mtx
.PHONY: mtx-build
mtx-build:
	@cd scripts/mtx && make build

.PHONY: mtx-run
mtx-run:
	@cd scripts/mtx && make run

.PHONY: mtx-debug
mtx-debug:
	@cd scripts/mtx && make debug

# osp
.PHONY: osp-build
osp-build:
	@cd scripts/osp && make build

.PHONY: osp-run
osp-run:
	@cd scripts/osp && make run

.PHONY: osp-debug
osp-debug:
	@cd scripts/osp && make debug

# itau
.PHONY: itau-build
itau-build:
	@cd scripts/itau && make build

.PHONY: itau-run
itau-run:
	@cd scripts/itau && make run

.PHONY: itau-debug
itau-debug:
	@cd scripts/itau && make debug

# citybank
.PHONY: citybank-build
citybank-build:
	@cd scripts/citybank && make build

.PHONY: citybank-run
citybank-run:
	@cd scripts/citybank && make run

.PHONY: citybank-debug
citybank-debug:
	@cd scripts/citybank && make debug

