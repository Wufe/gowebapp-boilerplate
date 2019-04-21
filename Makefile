PACKAGE_FIND?=boilerplateprj
PACKAGE_REPLACE?=boilerplateprj

clean:
	./project/script/clean.sh $(PACKAGE_FIND) $(PACKAGE_REPLACE)

install:
	cp .env.example .env
	yarn