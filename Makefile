- kuberbuild:
	kubectl apply -f deployments.yml
	kubectl apply -f services.yml

- kuberstop:
	kubectl delete deployment goapp
	kubectl delete service goapp 