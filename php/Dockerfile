# Use the official PHP 8.5 FPM image (replace with cli/apache if preferred)
FROM php:8.5-fpm

# Install system dependencies and MySQL extension
RUN apt-get update && apt-get install -y \
        default-mysql-client \
    && docker-php-ext-install pdo_mysql mysqli \
    && rm -rf /var/lib/apt/lists/*

# (Optional) install common PHP extensions
# RUN docker-php-ext-install intl bcmath

WORKDIR /var/www/html

# Copy application source (optional)
# COPY . .

EXPOSE 9000
ENTRYPOINT ["php", "./index.php"]