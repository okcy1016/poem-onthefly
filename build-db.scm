(define poem-archive-uri "https://github.com/chinese-poetry/chinese-poetry/archive/master.zip")
(define tmp-dir "./tmp")
(define poem-archive-filepath "./tmp/chinese-poetry-master.zip")
(define poem-db-path "./poems.db")

(define get-poet-tang-filename-list
  (lambda (dir)
    (use-modules (ice-9 ftw))
    (let ((filenames-in-dir (scandir dir)))
      (filter (lambda (item)
                (string-prefix? "poet.tang" item))
              filenames-in-dir))))

(define get-poet-tang-filepath-list
  (lambda (dir)
    (let ((poet-tang-filename-list (get-poet-tang-filename-list dir)))
      (map
       (lambda (dir filename) (string-append dir filename))
       (make-list (length poet-tang-filename-list) dir)
       poet-tang-filename-list))))

;; create tmp dir if not exists
(if (not (access? tmp-dir F_OK))
    (mkdir tmp-dir))

;; download poem if not exists
(if (not (access? poem-archive-filepath F_OK))
    (begin
      (display "first time use, downloading poems from GitHub ...")
      (newline)
      (load-extension "./net-helper" "init_poem_net_helper")
      (poem-download-archive poem-archive-uri poem-archive-filepath)
      ;; extract archive
      (display "extracting archive ...")
      (newline)
      ;; which os?
      (if (string=? (vector-ref (uname) 0) "Linux")
          (system (string-append
                   "unzip"
                   " "
                   "-q"
                   " "
                   poem-archive-filepath
                   " "
                   "-d"
                   " "
                   tmp-dir))
          (begin
            (display "win32 or macos are not supported currently.\n")
            (quit)))))

;;;; insert poems to db
;; delete db first
(if (access? poem-db-path F_OK)
    (delete-file poem-db-path))

(load-extension "./db-helper" "init_poem_db_helper")
(use-modules (ice-9 format))
(for-each
 (lambda (elem)
   (display
    (format
     #f
     "~d poems inserted from ~a\n"
     (poem-json2db elem poem-db-path)
     elem)))
 (get-poet-tang-filepath-list "./tmp/chinese-poetry-master/json/"))

(display (format #f "generated ~a\n" poem-db-path))
