clean:
	rm -rf ./VENV

venv:
	test -d VENV || python3 -m venv VENV

deps: venv
	source ./VENV/bin/activate && \
	pip install --upgrade pip && \
	pip install -r requirements-dev.txt

setup: venv deps

test: deps
	source ./VENV/bin/activate && \
	python -m pytest -vvvv --showlocals tests
