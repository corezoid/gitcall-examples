:- module(usercode, [
    handle/2
]).
		
handle(Data, Result) :-
  http_client:http_get('https://reqres.in/api/users?page=1', Reply, [content("application/json")]),
  put_dict(reply, Data, Reply, Result).