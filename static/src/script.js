document.getElementById("inputForm").addEventListener("submit", async function (event) {
    event.preventDefault();

    const inputString = document.getElementById("inputString").value;

    const responseDiv = document.getElementById("response");
    responseDiv.innerHTML = "Processing...";

    try {
        const response = await fetch("/process", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ inputString })
        });

        const data = await response.json();

        responseDiv.innerHTML = `Processed Response: ${data.response}`;
    } catch (error) {
        responseDiv.innerHTML = "Error processing the request.";
    }
});
