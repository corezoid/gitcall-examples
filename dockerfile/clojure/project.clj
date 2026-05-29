(defproject gitcall "1.0.0"
  :description "Gitcall runner"
  :url "http://example.com/"
  :dependencies [[org.clojure/clojure "1.12.5"] 
                 [http-kit "2.8.1"] 
                 [org.clojure/data.json "2.5.2"]]
  :main gitcall.main
  :aot [gitcall.main]
  :profiles {:uberjar {:aot :all}})
