window.onload = function () {
  document.getElementById("keyForm").addEventListener("submit", (e) => {
    e.preventDefault();
    const key = e.target.elements.key.value;

    crypto.subtle
      .digest("SHA-256", new TextEncoder().encode(key))
      .then((hash) => {
        const hashArray = Array.from(new Uint8Array(hash));
        const hashHex = hashArray
          .map((b) => b.toString(16).padStart(2, "0"))
          .join("");
        console.log(hashHex);
        e.target.elements.key.value = hashHex;
        e.target.submit();
      });
  });
};
