const WebSocket = require("ws");

const wss = new WebSocket.Server({ port: 3000 });

wss.on("connection", function connection(ws) {
  console.log("✅ Client connected");

  ws.send(
    JSON.stringify({ message: "🟢 Welcome to Saarthi WebSocket Server!" })
  );

  ws.on("message", function incoming(message) {
    console.log("📩 Received:", message);
    ws.send(`Echo: ${message}`);
  });

  ws.on("close", () => console.log("❌ Client disconnected"));
});

console.log("🚀 WebSocket server running on port 3000");
wss.on("error", (error) => {
  console.error("❌ WebSocket error:", error);
});
