# RabbitMQのログイン方法

    localhostであれば、http://localhost:8080を開いて
    ID/PW: guest/guestでログインできる。
    
    localhost以外だと、上記ではログインできないらしい。
    // ユーザ名sugimoto、パスワードsugimotoで作成
    $ sudo rabbitmqctl add_user sugimoto sugimoto
    // 権限設定
    $ sudo rabbitmqctl set_permissions sugimoto ".*" ".*" ".*"
    // ロール指定
    $ sudo rabbitmqctl set_user_tags sugimoto administrator
    
    みたいな設定でいけるみたいだけど、ここは未検証
