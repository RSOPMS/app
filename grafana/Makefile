.PHONY: start
start:
	@docker run -d --name=grafana -p 3000:3000 grafana/grafana

.PHONY: stop
stop:
	@docker stop grafana
	@docker rm grafana
