## Example: Healthchecks in Java

```java
	public class Main {
		public static void main(String args[]) {
			String[] urls = {"url1", "url2", ...};
			Thread[] threads = new Thread[urls.length];
			for (int i = 0; i < urls.length; i++) {
				threads[i] = new Thread(new Healthcheck(urls[i]));
				threads[i].start();
			}

			try {
				for (Thread t : threads) {
					t.join();
				}
			} catch (Exception e) {}

			System.out.println(Healthcheck.getOkCount() + " ok, " +
				Healthcheck.getErrCount() + " errors");
		}
	}
```

<span class="fragment current-only" data-code-focus="3"></span>
<span class="fragment current-only" data-code-focus="6,7,12"></span>
<span class="fragment current-only" data-code-focus="16,17"></span>
