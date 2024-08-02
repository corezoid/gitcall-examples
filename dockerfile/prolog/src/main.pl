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
  get_time(Time),
  format(user_output, '[req] time=~f id=~s~n', [float(Time), Json.id]),
  usercode:handle(Json.params, Result),
  get_time(Time2),
  format(user_output, '[res] time=~f id=~s~n', [float(Time2), Json.id]),
  reply_json_dict(json([jsonrpc=Json.jsonrpc, id=Json.id, result=Result])).

resp_error(Json, Err) :-
  print_message(error, Err),
  get_time(Time),
  format(user_output, '[res] time=~f id=~s error=~s~n', [float(Time), Json.id, Error]),
  reply_json(json([jsonrpc=Json.jsonrpc, id=Json.id, error=json([code=1, message=Error])])).
