import java.net.HttpURLConnection;
import java.net.URL;

class Healthcheck implements Runnable {
  private static int okCount = 0;
  private static int errCount = 0;
  private String url;

  public Healthcheck(String url) {
    this.url = url;
  }

  public static synchronized void addOk() { Healthcheck.okCount++; }
  public static synchronized void addErr() { Healthcheck.errCount++; }
  public static int getOkCount() { return Healthcheck.okCount; }
  public static int getErrCount() { return Healthcheck.errCount; }

  public void run() {
    try {
      if (getStatusCode(this.url) == 200) {
        Healthcheck.addOk();
      } else {
        Healthcheck.addErr();
      }
    } catch(Exception e) {
      Healthcheck.addErr();
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
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
      "https://www.oracleimg.com/us/assets/metrics/ora_docs.js",
      "http://yahoo.com",
      "https://www.yahoo.com",
      "http://unknownurl.com",
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
    System.out.println(Healthcheck.getOkCount() + " ok, " + Healthcheck.getErrCount() + " errors");
  }
}
