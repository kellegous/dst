run: env/bin/pip3
	env/bin/python3 main.py

env/bin/pip3:
	python3 -m venv env
	env/bin/pip3 install -r requirements.txt

clean:
	rm -rf env
