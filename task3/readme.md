Task 3.  BASH

Преобразовать следующий однострочник в красивый скрипт:
sudo netstat -tunapl | awk '/firefox/ {print $5}' | cut -d: -f1 | sort | uniq -c |
sort | tail -n5 | grep -oP '(\d+\.){3}\d+' | while read IP ; do whois $IP | 
awk -F':' '/^Organization/ {print $2}' ; done
* создайте README.md и опишите, что будет делать ваш скрипт
* скрипт должен принимать PID или имя процесса в качестве аргумента
* количество строк вывода должно регулироваться пользователем
* должна быть возможность видеть другие состояния соединений
* скрипт должен выводить понятные сообщения об ошибках
* скрипт не должен зависеть от привилегий запуска, выдавать предупреждение
** скрипт выводит число соединений к каждой организации 
** скрипт позволяет получать другие данные из вывода `whois`
** скрипт умеет работать с `ss`, вы используете утилиты/built-ins, не вошедшие в 
однострочник    