build:
	@mkdir -p bin
	@cd bin && go build ..

clean:
	@cd bin && rm -f *