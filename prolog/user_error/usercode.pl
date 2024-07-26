:- module(usercode, [
  handle/2
]).

handle(Data, Result) :-
  throw("my custom error").