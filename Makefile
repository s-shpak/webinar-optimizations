.PHONY: clean
clean: clean-strconv-vs-fmt clean-slices-cap clean-bce clean-sync-pool

.PHONY: clean-strconv-vs-fmt
clean-strconv-vs-fmt:
	find cmd/strconv_vs_fmt -type f \( -name "*.test" -o -name "*.out" -o -name "*.s" \) -delete

.PHONY: clean-slices-cap
clean-slices-cap:
	find cmd/slices_cap -type f \( -name "*.test" -o -name "*.out" -o -name "*.s" \) -delete

.PHONY: clean-bce
clean-bce:
	find cmd/bce -type f \( -name "*.test" -o -name "*.out" -o -name "*.s" \) -delete

.PHONY: clean-sync-pool
clean-sync-pool:
	find cmd/sync-pool -type f \( -name "*.test" -o -name "*.out" -o -name "*.s" \) -delete
