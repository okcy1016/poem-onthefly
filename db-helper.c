#include <stdio.h>
#include <sqlite3.h>
#include <json-c/json.h>
#include "sds.h"

#define POEM_TANG 0
#define POEM_SONGCI 1

struct poem {
  const char* title;
  const char* author;
  const char* paragraphs;
};

int loop_poem_and_insert(json_object*, int, const char*);
int read_poem_2_db(const char*, int, const char*);
int insert_single_poem(sqlite3*, sds, struct poem*);
static int callback(void*, int, char**, char**);
sds build_insert_sql(sds, struct poem*);

static int callback(void *NotUsed, int argc, char **argv, char **azColName){
  int i;
  for(i=0; i<argc; i++){
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : "NULL");
  }
  printf("\n");
  return 0;
}

int read_poem_2_db(const char* jsonpath, int poem_type, const char* dbpath) {
  /* char* is a mutable pointer to a mutable character/string. */

  /* const char* is a mutable pointer to an immutable character/string. */
  /* You cannot change the contents of the location(s) this pointer points to. */
  /* Also, compilers are required to give error messages when you try to do so. */
  /* For the same reason, conversion from const char * to char* is deprecated. */
  
  /* char* const is an immutable pointer (it cannot point to any other location) */
  /* but the contents of location at which it points are mutable. */

  /* const char* const is an immutable pointer to an immutable character/string. */
  
  json_object* root = json_object_from_file(jsonpath);
  int poem_count = loop_poem_and_insert(root, poem_type, dbpath);
  
  json_object_put(root);
  return poem_count;
}

int loop_poem_and_insert(json_object* root, int poem_type, const char* db_path) {
  sqlite3* db;
  int poem_count;
  int para_count;
  struct poem* poem = malloc(sizeof(struct poem));
  sds table; 

  table = sdsempty();
  switch(poem_type) {
  case POEM_TANG:
    table = sdscat(table, "tang");
    break;
  case POEM_SONGCI:
    table = sdscat(table, "songci");
    break;
  default:
    table = sdscat(table, "unknown");
    break;
  }
  
  poem_count = json_object_array_length(root);

  // open db
  int rc = sqlite3_open(db_path, &db);
  if (rc != SQLITE_OK) {
    fprintf(stderr, "can't open database: %s\n", sqlite3_errmsg(db));
    sqlite3_close(db);
    return -1;
  }

  // create `poem` table
  sds create_table_sql = sdscatprintf(sdsempty(), "CREATE TABLE %s (title string, author string, paragraphs string);", table);
  sqlite3_exec(db, create_table_sql, callback, 0, NULL);
  sdsfree(create_table_sql);

  // wrap in one transaction, speed up bulk insert!
  // see https://stackoverflow.com/questions/1711631/improve-insert-per-second-performance-of-sqlite
  sqlite3_exec(db, "BEGIN TRANSACTION", NULL, NULL, NULL);
  
  for(int i=0; i<poem_count; i++) {
    json_object *poem_elem = json_object_array_get_idx(root, i);

    json_object *poem_title_elem;
    if (poem_type == POEM_SONGCI)
      poem_title_elem = json_object_object_get(poem_elem, "rhythmic");
    else
      poem_title_elem = json_object_object_get(poem_elem, "title");
    json_object *poem_author_elem = json_object_object_get(poem_elem, "author");
    sds poem_title = sdsnew(json_object_get_string(poem_title_elem));
    sds poem_author = sdsnew(json_object_get_string(poem_author_elem));
    
    json_object *poem_paragraphs_elem = json_object_object_get(poem_elem, "paragraphs");
    sds poem_paragraphs = sdsempty();
    para_count = json_object_array_length(poem_paragraphs_elem);
    for(int j=0; j<para_count; j++) {
      json_object *para_elem = json_object_array_get_idx(poem_paragraphs_elem, j);
      poem_paragraphs = sdscat(poem_paragraphs, json_object_get_string(para_elem));
      if (j != (para_count - 1))
        poem_paragraphs = sdscat(poem_paragraphs, "\n");
    }
    
    poem->title = poem_title;
    poem->author = poem_author;
    poem->paragraphs = poem_paragraphs;

    insert_single_poem(db, table, poem);
    sdsfree(poem_title);
    sdsfree(poem_author);
    sdsfree(poem_paragraphs);
  }

  // end sqlite transaction
  sqlite3_exec(db, "END TRANSACTION", NULL, NULL, NULL);
  
  free(poem);
  sqlite3_close(db);
  return poem_count;
}

int insert_single_poem(sqlite3* db, sds table, struct poem* poem) {  
  char* zErrMsg = 0;

  sds sql = build_insert_sql(table, poem);
  /* printf("sql: %s\n", sql); */
  int rc = sqlite3_exec(db, sql, callback, 0, &zErrMsg);
  if(rc != SQLITE_OK){
    fprintf(stderr, "sql error: %s\n", zErrMsg);
    sqlite3_free(zErrMsg);
  }
  sdsfree(sql);
  
  return 0;
}

sds build_insert_sql(sds table, struct poem* poem) {
  sds sql = sdsempty();
  sql = sdscatprintf(sql, "INSERT INTO %s VALUES (\"%s\", \"%s\", \"%s\");", table, poem->title, poem->author, poem->paragraphs);
  /* printf("%d\n", sdslen(sql)); */
  return sql;
}
