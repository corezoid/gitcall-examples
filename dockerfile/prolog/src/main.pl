:- use_module(library(http/http_server)).
:- use_module(library(http/http_json)).
:- use_module(library(http/json_convert)).
:- use_module(usercode).

:- use_module(library(http/http_unix_daemon)).
:- initialization http_daemon.
:- http_handler(root(.), request, []).



request(Request) :-
  http_read_json_dict(Request, Json),
  catch(call_usercode(Json), Err, resp_error(Json, Err)).

call_usercode(Json) :-
  usercode:handle(Json.params, Result),
  reply_json_dict(json([jsonrpc=Json.jsonrpc, id=Json.id, result=Result])).

resp_error(Json, Err) :-
  print_message(error, Err),
  format(atom(Error), '~q', [Err]),
  reply_json(json([jsonrpc=Json.jsonrpc, id=Json.id, error=json([code=1, message=Error])])).
