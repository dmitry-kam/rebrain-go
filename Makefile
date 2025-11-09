sources:
	find src/ | sed -e "s/[^-][^\/]*\// |/g" -e "s/|\([^ ]\)/|-\1/"