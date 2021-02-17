(display "building c extensions ...")
(newline)

(define cmd-compile-net-helper "gcc `curl-config --cflags` net-helper.c `curl-config --libs` -shared -o net-helper.so -fPIC")
(define cmd-compile-db-helper "gcc `pkg-config --cflags sqlite3 json-c` db-helper.c `pkg-config --libs sqlite3 json-c` -shared -o db-helper.so -fPIC")

(define cmd-list
  (list cmd-compile-net-helper
        cmd-compile-db-helper))

(for-each (lambda (cmd)
            (display cmd)
            (newline)
            (system cmd))
          cmd-list)
