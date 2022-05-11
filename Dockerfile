FROM ubuntu:jammy

RUN DEBIAN_FRONTEND=noninteractive \
	ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime \
	&& apt-get update \
	&& apt-get install -y \
	build-essential \
	curl \
	php \
	python3 \
	python3-pip \
	nodejs \
	ruby \
	&& gem install activesupport \
	&& pip install python-dateutil \
	&& curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y \
	&& curl https://dl.google.com/go/go1.18.2.linux-amd64.tar.gz | tar xz -C /usr/local

ENV PATH="/usr/local/go/bin:/root/.cargo/bin:${PATH}"