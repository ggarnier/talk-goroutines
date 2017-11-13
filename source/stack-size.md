# Stack size

<div style="display: flex; justify-content: space-between">
  <div style="width: 600px">
    <h3>Threads</h3>
    <ul>
      <li>fixed (~ 1-8 MB)</li>
      <li>wasteful for small jobs</li>
      <li>not enough for large jobs</li>
    </ul>
  </div>

  <div style="width: 600px">
    <h3>Goroutines</h3>
    <ul>
      <li>dynamic</li>
      <li>starts small (~ 4 kB)</li>
      <li>grows and shrinks when needed</li>
    </ul>
  </div>
</div>
