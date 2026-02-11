import * as React from "react";
import ReactDOM from "react-dom/client";
import { MantineProvider } from "@mantine/core";
import App from "./App";
import "@mantine/core/styles.css";
import "./index.css";

const reactNamespace = React as unknown as { default?: typeof React };
if (!reactNamespace.default) {
  reactNamespace.default = React;
}


try {
  ReactDOM.createRoot(document.getElementById("root")!).render(
    <React.StrictMode>
      <MantineProvider defaultColorScheme="light">
        <App />
      </MantineProvider>
    </React.StrictMode>
  );
} catch (e) {
  console.error("Fatal UI error:", e);
  document.body.innerHTML = `
    <pre style="padding:16px;color:#fff;background:#000">
    UI failed to start.
    Check DevTools console for details.
    </pre>
  `;
}
