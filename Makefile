release:
	docker build -t "singaravelan21/match_todo_list_gateway_srv" .;
	docker tag "singaravelan21/match_todo_list_gateway_srv" "singaravelan21/match_todo_list_gateway_srv:v2.18.0";
	docker push "singaravelan21/match_todo_list_gateway_srv:v2.18.0"