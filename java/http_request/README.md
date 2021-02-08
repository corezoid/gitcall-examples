# Corezoid Java Example

The project requires custom build with import dependencies.
To do so specify build script in Corezoit Gitcall node:
```
./gradlew build
```
Gradle compiles source codes and adds all dependencies
to the resulting usercode.jar file. Gitcall will include
the usercode.jar file to java classpass.
