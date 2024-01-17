# Asosiy imojdan Nginx olish
FROM nginx:latest

# Ilova katalogini workdir sifatida belgilash
WORKDIR /usr/share/nginx/html

# Local JS, CSS, va HTML fayllarini ilova katalogiga kopyalash
COPY . .

# Nginx konfiguratsiya faylini almashtirish
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# 80-chi portni ochish
EXPOSE 80

# Nginxni ishga tushirish
CMD ["nginx", "-g", "daemon off;"]
