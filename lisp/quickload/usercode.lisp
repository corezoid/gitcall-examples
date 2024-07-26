(ql:quickload '(:cl-mustache) :silent t)

(defpackage #:usercode
(:use #:cl)
(:export :handle))
(in-package #:usercode)

(defun handle (data)
;; write your code here
(setf (gethash "template" data) 
    (mustache:render* "Hello {{person}}!" '((:person . "World")))
    ) data)