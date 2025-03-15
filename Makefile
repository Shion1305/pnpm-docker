dev: dev-build
	docker build -t pnpm-sample -f sample-app/Dockerfile sample-app/next-app

dev-build:
	docker build -t pnpm-image -f images/22/Dockerfile .
