<script>
  let inputJson = $state();
  let indentedJson = $state();

  async function submitJson() {
    const response = await fetch('http://localhost:8080/json/indent', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: new URLSearchParams({ input: inputJson }),
    });

    if (response.ok) {
      indentedJson = await response.text();
    } else {
      indentedJson = 'Error: ' + response.statusText;
    }
  }
</script>

<textarea bind:value={inputJson} rows="5" cols="50"></textarea>
<button onclick={submitJson}>Submit JSON</button>

<pre>{indentedJson}</pre>
