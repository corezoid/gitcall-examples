(ns usercode.usercode
    (:require [clj-uuid :as uuid]))

(defn handle [data]
    (assoc data :uuid (uuid/v1)))