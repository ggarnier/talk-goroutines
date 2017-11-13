class Fibonacci implements Runnable {
  private String name;

  public Fibonacci(String name) {
    this.name = name;
  }

  public void run() {
    for (int i = 0; i < 100; i++) {
      fib(20);
      System.out.println(this.name);
    }
  }

  private int fib(int n) {
    if (n <= 1) {
      return 1;
    }
    return fib(n-1) + fib(n-2);
  }
}

public class Main {
  public static void main(String args[]) {
    int num = 10;
    Thread[] threads = new Thread[num];
    String spacing = "";
    for (int i = 0; i < num; i++) {
      threads[i] = new Thread(new Fibonacci(spacing + i));
      threads[i].start();
      spacing += "\t";
    }

    try {
      for (Thread t : threads) {
        t.join();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
