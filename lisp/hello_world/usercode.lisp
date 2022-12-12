(defpackage #:usercode
  (:use #:cl)
  (:export :handle))
(in-package #:usercode)

(defun handle (data)
  ;; write your code here
  (setf (gethash 'hello data) 'Hello-world) data)
