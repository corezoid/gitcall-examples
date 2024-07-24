# Php extension example

The project requires custom build with docker-php-ext-install usage.
To do so specify build script in Corezoit Gitcall node.
#### Build command:
```
docker-php-ext-install bcmath
```

Or with composer dependencies.
#### Build command:
```
docker-php-ext-install bcmath && composer install --no-dev --prefer-dist --optimize-autoloader --no-interaction
```