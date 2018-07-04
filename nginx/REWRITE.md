条件匹配

```nginx
server {
    listen       80;
    server_name  www.example.org  example.org;
    if ($http_host = example.org) {
        rewrite  (.*)  http://www.example.org$1;
    }
    ...
}
    
server {
    listen       80;
    server_name  example.org;
    return       301 http://www.example.org$request_uri;
}    

rewrite      ^ http://www.example.org$request_uri?;

location / {
    root       /var/www/myapp.com/current/public;

    try_files  /system/maintenance.html
               $uri  $uri/index.html $uri.html
               @mongrel;
}

location @mongrel {
    proxy_pass  http://mongrel;
}
```

break;

中断顺序执行