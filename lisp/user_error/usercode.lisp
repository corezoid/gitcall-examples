(defpackage #:usercode
(:use #:cl)
(:export :handle))
(in-package #:usercode)

(defun handle (data)
(error "my custom error"))