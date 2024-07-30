(ns gitcall.main
  (:gen-class)
  (:use org.httpkit.server)
  (:require [clojure.data.json :as json]))

;; parse GITCALL_PORT
(defn get_port [port]
  (if (not= port nil)
    (Integer/parseInt port)
    ((println "GITCALL_PORT env is required but not set") (System/exit 1))))

;; handle
(defn handle [data]
  (assoc data :hello "Hello world!"))

;; request
(defn request [req]
   (let [msg (json/read (clojure.java.io/reader (:body req) :encoding "UTF-8") :key-fn keyword)]
    (let [jsonrpc (:jsonrpc msg) id (:id msg) params (:params msg)]
      (try
         (let [data (handle params)]
           {:status  200
            :headers {"Content-Type" "application/json"}
            :body (json/write-str {:jsonrpc jsonrpc :id id, :result data} )})
         (catch Exception e
           {:status  200
            :headers {"Content-Type" "application/json"}
            :body (json/write-str {:jsonrpc jsonrpc :id id, :error {:code 1 :message (ex-message e)}})})))))

;; start TCP server   
(defn -main [& args]
  (def port (get_port (System/getenv "GITCALL_PORT"))),
  (println (format "Listening on http://0.0.0.0:%d" port))
  (run-server request {:port port}))