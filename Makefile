- kuberbuild:
	kubectl apply -f deployments.yml
	kubectl apply -f services.yml

- kuberautoscale:
	kubectl autoscale deployment goapp --cpu-percent=5 --min=1 --max=10

- kuberstop:
	kubectl delete deployment goapp
	kubectl delete service goapp 
