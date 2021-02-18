(define poem-archive-uri "https://github.com/chinese-poetry/chinese-poetry/archive/master.zip")
(define tmp-dir "./tmp")
(define poem-archive-filepath "./tmp/chinese-poetry-master.zip")
(define poem-db-path "./poems.db")
(define poem-tang-json-dir (string-append
                            tmp-dir
                            "/chinese-poetry-master/json"))
(define poem-songci-json-dir (string-append
                              tmp-dir
                              "/chinese-poetry-master/ci"))
(define poem-archive-dir "./tmp/chinese-poetry-master")

(define string-prefix?
  (lambda (x y)
    (let ([n (string-length x)])
      (and (fx<= n (string-length y))
           (let prefix? ([i 0])
             (or (fx= i n)
                 (and (char=? (string-ref x i) (string-ref y i))
                      (prefix? (fx+ i 1)))))))))

(define get-poet-filename-list
  (lambda (poem-type)
    (let ([filenames-in-dir (directory-list
                             (case poem-type
                               ["tang" poem-tang-json-dir]
                               ["songci" poem-songci-json-dir]))]
          [filename-prefix (case poem-type
                             ["tang" "poet.tang"]
                             ["songci" "ci.song"])])
      (filter (lambda (item)
                (string-prefix? filename-prefix item))
              filenames-in-dir))))

(define get-poet-filepath-list
  (lambda (poem-type)
    (let ([poet-filename-list (get-poet-filename-list poem-type)]
          [dir (case poem-type
                 ["tang" poem-tang-json-dir]
                 ["songci" poem-songci-json-dir])])
      (map
       (lambda (dir filename) (string-append dir "/" filename))
       (make-list (length poet-filename-list) dir)
       poet-filename-list))))

(define create-tmp-dir
  (lambda ()
    ;; create tmp dir if not exists
    (if (not (file-exists? tmp-dir))
        (mkdir tmp-dir))))

(define os-type
  (lambda ()
    (case (machine-type)
      [(i3le ti3le a6le ta6le) "linux"]
      [(i3osx ti3osx a6osx ta6osx) "macos"]
      [(i3nt ti3nt a6nt ta6nt) "windows"]
      [else "unknown"])))

;; load c dynamic library
(case (os-type)
  ["windows" (begin
               (load-shared-object "./net-helper.dll")
               (load-shared-object "./db-helper.dll"))]
  ["linux" (begin
             (load-shared-object "./net-helper.so")
             (load-shared-object "./db-helper.so"))])

(define poem-download-archive
  (foreign-procedure "download_poem_archive" (string int string) int))

(define poem-json2db
  (foreign-procedure "read_poem_2_db" (string int string) int))

(define poem-get-type-int
  (lambda (poem-type)
    (case poem-type
      ["tang" 0]
      ["songci" 1])))

(define download-poem-archive
  (lambda ()
    ;; download poem if not exists
    (if (not (file-exists? poem-archive-filepath))
        (begin
          (display "first time use, downloading poems from GitHub ...")
          (newline)
          (poem-download-archive poem-archive-uri poem-archive-filepath)))))

(define extract-poem-archive
  (lambda ()
    (if (not (file-exists? poem-archive-dir))
        (begin
          (display "extracting archive ...")
          (newline)
          (if (string=? (os-type) "linux")
              (system (format #f "unzip -q ~a -d ~a" poem-archive-filepath tmp-dir))
              (begin
                (display "win32 or macos are not supported currently.\n")
                (exit)))))))

(define generate-db
  (lambda ()
    ;; delete old db first
    (if (file-exists? poem-db-path)
        (delete-file poem-db-path))

    ;; insert poems to db    
    ;; poem tang
    (for-each
     (lambda (elem)
       (display
        (format
         "~d poems inserted from ~a\n"
         (poem-json2db elem (poem-get-type-int "tang") poem-db-path)
         elem)))
     (get-poet-filepath-list "tang"))

    ;; poem songci
    (for-each
     (lambda (elem)
       (display
        (format
         #f
         "~d poems inserted from ~a\n"
         (poem-json2db elem (poem-get-type-int "songci") poem-db-path)
         elem)))
     (get-poet-filepath-list "songci"))
    
    (display (format "generated ~a\n" poem-db-path))))

(define build-db
  (lambda ()
    (create-tmp-dir)
    (download-poem-archive)
    (extract-poem-archive)
    (generate-db)))

(build-db)
