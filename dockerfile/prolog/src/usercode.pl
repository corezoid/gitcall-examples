:- module(usercode, [handle/2]).

handle(Data, Result) :-
    put_dict(prolog, Data, "Hello, world!", Result).