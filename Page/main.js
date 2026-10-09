const form = document.getElementById("f");
const input = document.getElementById("url");
const out = document.getElementById("out");
            
form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const res = await fetch("/shorten", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ 
      	url: input.value
      })
    });
  
	out.textContent = await res.text();
});
