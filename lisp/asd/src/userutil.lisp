(defpackage #:userutil
  (:use #:cl)
  (:export :run))
(in-package #:userutil)

(defun run (data)
  (setf (gethash "template" data) 
    (mustache:render* "Hello {{person}}!" '((:person . "World")))
    ) data)