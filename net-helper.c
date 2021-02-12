// borrowed from libcurl url2file.c example file

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h> 
#include <curl/curl.h>
#include <libguile.h>
 
static size_t write_data(void *ptr, size_t size, size_t nmemb, void *stream)
{
  size_t written = fwrite(ptr, size, nmemb, (FILE *)stream);
  return written;
}
 
int download_peom_archive(const char *uri, const char *filepath)
{
  CURL *curl_handle;
  const char *pagefilename = filepath;
  FILE *pagefile;
 
  if(uri == NULL) {
    printf("err NULL ptr of uri\n");
    return 1;
  }
 
  curl_global_init(CURL_GLOBAL_ALL);
 
  /* init the curl session */ 
  curl_handle = curl_easy_init();
 
  /* set URL to get here */ 
  curl_easy_setopt(curl_handle, CURLOPT_URL, uri);
 
  /* Switch on full protocol/debug output while testing */ 
  curl_easy_setopt(curl_handle, CURLOPT_VERBOSE, 0L);
 
  /* disable progress meter, set to 0L to enable it */ 
  curl_easy_setopt(curl_handle, CURLOPT_NOPROGRESS, 0L);

  // follow redirection
  curl_easy_setopt(curl_handle, CURLOPT_FOLLOWLOCATION, 1L);
 
  /* send all data to this function  */ 
  curl_easy_setopt(curl_handle, CURLOPT_WRITEFUNCTION, write_data);
 
  /* open the file */ 
  pagefile = fopen(pagefilename, "wb");
  if(pagefile) {
 
    /* write the page body to this file handle */ 
    curl_easy_setopt(curl_handle, CURLOPT_WRITEDATA, pagefile);
 
    /* get it! */ 
    curl_easy_perform(curl_handle);
 
    /* close the header file */ 
    fclose(pagefile);
  }
 
  /* cleanup curl stuff */ 
  curl_easy_cleanup(curl_handle);
 
  curl_global_cleanup();
 
  return 0;
}

SCM download_peom_archive_wrapper(SCM uri, SCM filepath) {
  download_peom_archive(scm_to_locale_string(uri), scm_to_locale_string(filepath));
  return scm_from_int(0);
}

void init_poem_net_helper() {
  scm_c_define_gsubr("poem-download-archive", 2, 0, 0, download_peom_archive_wrapper);
}
