(ns usercode.usercode
    (:require [clj-http.client :as client]))

(defn handle [data]
    (assoc data :reply (:body (client/get "https://reqres.in/api/users?page=1"))))