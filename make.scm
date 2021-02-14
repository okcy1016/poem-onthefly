(display "building c extensions ...")
(newline)

(define cmd-compile-net-helper "gcc `pkg-config --cflags guile-2.2` `curl-config --cflags` net-helper.c `curl-config --libs` -shared -o net-helper.so -fPIC")
(define cmd-compile-db-helper "gcc `pkg-config --cflags guile-2.2 sqlite3` db-helper.c `pkg-config --libs guile-2.2 sqlite3 json-c` -shared -o db-helper.so -fPIC")

(define cmd-list
  (list cmd-compile-net-helper
        cmd-compile-db-helper))

(for-each (lambda (cmd)
            (display cmd)
            (newline)
            (system cmd))
          cmd-list)
