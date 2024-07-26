(ql:quickload '(:dexador) :silent t)

(defpackage #:usercode
(:use #:cl)
(:export :handle))
(in-package #:usercode)

(defun handle (data)
    (setf (gethash "reply" data) (dex:get "https://reqres.in/api/users?page=1'")) data)