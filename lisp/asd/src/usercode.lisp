(defpackage #:usercode
   (:use #:cl #:userutil)
  (:export :handle))
(in-package #:usercode)


(defun handle (data)
  ;; write your code here
  (userutil:run data))