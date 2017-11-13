import java.net.HttpURLConnection;
import java.net.URL;

class Healthcheck implements Runnable {
  public static int okCount = 0;
  public static int errCount = 0;
  private String url;

  public Healthcheck(String url) {
    this.url = url;
  }

  public void run() {
    try {
      if (getStatusCode(this.url) == 200) {
        okCount++;
      } else {
        errCount++;
      }
    } catch(Exception e) {
      errCount++;
    }
  }

  private int getStatusCode(String url) throws Exception {
    URL u = new URL(url);
    HttpURLConnection conn = (HttpURLConnection) u.openConnection();
    conn.setRequestMethod("GET");
    return conn.getResponseCode();
  }
}

public class Main {
  public static void main(String args[]) {
    String[] urls = {
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com"
    };
    Thread[] threads = new Thread[urls.length];
    for (int i = 0; i < urls.length; i++) {
      threads[i] = new Thread(new Healthcheck(urls[i]));
      threads[i].start();
    }

    try {
      for (Thread t : threads) {
        t.join();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }

    System.out.println(urls.length);
    System.out.println(Healthcheck.okCount + " ok, " + Healthcheck.errCount + " errors");
  }
}
