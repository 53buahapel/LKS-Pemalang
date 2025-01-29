FROM php:apache

WORKDIR /var/www/html

COPY . .

RUN mv flag.txt /flag_$(head -n 10 /dev/urandom | md5sum | cut -d " " -f 1).txt

CMD ["apache2-foreground"]