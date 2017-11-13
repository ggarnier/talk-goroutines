## Example: Healthchecks in Java

```java
	class Healthcheck implements Runnable {
		private static int okCount = 0;
		private static int errCount = 0;

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
			...
		}
	}
```

<span class="fragment current-only" data-code-focus="2,3"></span>
<span class="fragment current-only" data-code-focus="5,6"></span>
