:- module(usercode, [
    handle/2
]).
:- use_module(library(matrix)).
        
handle(Data, Result) :-
    determinant([[2,-1,0],[-1,2,-1],[0,-1,2]],M),
    put_dict(matrix, Data, M, Result).