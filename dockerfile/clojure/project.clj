(defproject gitcall "1.0.0"
  :description "Gitcall runner"
  :url "http://example.com/"
  :dependencies [[org.clojure/clojure "1.11.3"] 
                 [http-kit "2.3.0"] 
                 [org.clojure/data.json "2.5.0"]]
  :main gitcall.main
  :aot [gitcall.main]
  :profiles {:uberjar {:aot :all}})
